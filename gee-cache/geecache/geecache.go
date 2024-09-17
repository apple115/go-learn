package geecache

import (
	"fmt"
	"log"
	"sync"
)

// 接收 key --> 检查是否被缓存 -----> 返回缓存值 ⑴
//                 |  否                         是
//                 |-----> 是否应当从远程节点获取 -----> 与远程节点交互 --> 返回缓存值 ⑵
//                             |  否
//                             |-----> 调用`回调函数`，获取值并添加到缓存 --> 返回缓存值 ⑶

// (2)
// 使用一致性哈希选择节点        是                                    是
//    |-----> 是否是远程节点 -----> HTTP 客户端访问远程节点 --> 成功？-----> 服务端返回返回值
//                    |  否                                    ↓  否
//                    |----------------------------> 回退到本地节点处理。

// Getter 用于从数据源获取数据 给用户自定义
type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

// Get ...
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

type Group struct {
	name      string
	getter    Getter
	mainCache cache
	peers     PeerPicker
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group) // 全局变量
)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter ")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:   name,
		getter: getter,
		mainCache: cache{
			cachBytes: cacheBytes,
		},
	}
	groups[name] = g
	return g
}

func GetGroup(name string) *Group {
	mu.RLock()
	defer mu.RUnlock()
	return groups[name]
}

// get ...
func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	if v, ok := g.mainCache.get(key); ok {
		log.Printf("[GeeCache] hit")
		return v, nil
	}
	// 缓存不存在,则从数据源获取
	return g.load(key)
}

// RegisterPeers 注册节点选择器
func (g *Group) RegisterPeers(peers PeerPicker) {
	if g.peers != nil {
		panic("RegisterPeerPicker called more than once")
	}
	g.peers = peers
}

// load ...
func (g *Group) load(key string) (value ByteView, err error) {
	if g.peers != nil {
		if peer, ok := g.peers.PickPeer(key); ok {
			if value, err = g.getFromPeer(peer, key); err == nil {
				return value, nil
			}
			log.Println("[GeeCache] Failed to get from peer", err)
		}
	}
	return g.getLocally(key)
}

// getFromPeer ...从远程节点通过key获取数据
func (g *Group) getFromPeer(peer PeerGetter, key string) (ByteView, error) {
	bytes, err := peer.Get(g.name, key)
	if err != nil {
		return ByteView{}, err
	}

	return ByteView{b: bytes}, nil
}

// getLocally ...
func (g *Group) getLocally(key string) (ByteView, error) {
	//从数据源获取
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{b: cloneBytes(bytes)}
	// 缓存到本地
	g.populateCache(key, value)
	return value, nil
}

// populateCache  添加缓存
func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
