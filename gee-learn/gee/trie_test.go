package gee

import (
	"testing"
)

func TestNode(t *testing.T) {
		rootnode :=&node{}
		rootnode.insert("/", []string{},0)
		assertDeepEqual(t, rootnode.pattern, "/")
		rootnode.insert("/hello", []string{"hello"},0)
		assertDeepEqual(t, rootnode.children[0].part, "hello")
		rootnode.insert("/world/hello", []string{"world","hello"}, 0)

		assertDeepEqual(t, rootnode.children[1].part,"world")
		rootnode.insert("/hello/:name", []string{"hello",":name"}, 0)
		node:=rootnode.search([]string{"hello","zhang"},0)
		assertDeepEqual(t, node.pattern,"/hello/:name")
}


