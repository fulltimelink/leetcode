package prototype

import "fmt"

type File struct {
	name string
}

func (f *File) print(indentaion string) {
	fmt.Println(indentaion + " " + f.name)
}

func (f *File) clone() INode {
	return &File{f.name + "_clone"}
}
