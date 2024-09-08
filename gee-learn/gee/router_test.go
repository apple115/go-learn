package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assert", nil)
	return r
}

func assertDeepEqual(t *testing.T, a, b interface{}) {
	t.Helper()
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("expect %v, but got %v", a, b)
	}
}

func TestParsePatter(t *testing.T) {
	cases := []struct {
		pattern string
		parts   []string
	}{
		{"/p/:name", []string{"p", ":name"}},
		{"/p/*name", []string{"p", "*name"}},
		{"/p/*", []string{"p", "*"}},
	}
	for _, c := range cases {
		parts := parsePattern(c.pattern)
		assertDeepEqual(t, parts, c.parts)
	}
}
func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/geektutu")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "geektutu" {
		t.Fatal("name should be equal to 'geektutu'")
	}
	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])
}

