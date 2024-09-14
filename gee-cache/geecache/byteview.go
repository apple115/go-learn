package geecache

type ByteView struct {
	b []byte
}
// Len ...
func (v ByteView) Len() int {
	return len(v.b)
}

// BytesSlice ...
func (v ByteView) BytesSlice() []byte {
	return cloneBytes(v.b)
}

// String...
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
