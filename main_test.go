package main

import (
	"fmt"
	"testing"
	"unique"
)

type testStruct2 struct {
	c1 string
	c2 string
	c3 testStruct3
}
type testStruct3 struct {
	c31 string
	c32 string
}
type testStructData struct {
	a int
	b string
	c testStruct2
}

func BenchmarkMake1(b *testing.B) {
	c1 := testStructData{
		a: 12,
		b: "eee",
		c: testStruct2{
			c1: "aaa",
			c2: "bbb",
			c3: testStruct3{
				c31: "ccc",
				c32: "ddd",
			},
		},
	}
	for i := 0; i < b.N; i++ {
		compareTestStructData(c1, c1)
	}
}
func BenchmarkMake2(b *testing.B) {
	c1 := testStructData{
		a: 12,
		b: "eee",
		c: testStruct2{
			c1: "aaa",
			c2: "bbb",
			c3: testStruct3{
				c31: "ccc",
				c32: "ddd",
			},
		},
	}
	u1 := unique.Make(c1)
	u2 := unique.Make(c1)
	for i := 0; i < b.N; i++ {
		compareTestStructData1(u1, u2)
	}

}

func compareTestStructData1(u1 unique.Handle[testStructData], u2 unique.Handle[testStructData]) bool {
	return u1 == u2
}
func compareTestStructData(c1, c2 testStructData) bool {
	return c1 == c2
}

func TestMap(t *testing.T) {
	a := make(map[int]struct{})
	a[1] = struct{}{}
	fmt.Println(a[1])
	fmt.Println(a[2])

}
