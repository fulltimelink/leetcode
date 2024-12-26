package prototype

import "fmt"

type Folder struct {
	name     string
	children []INode
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + " " + f.name)
	for _, child := range f.children {
		child.print(indentation + indentation)
	}
}

func (f *Folder) clone() INode {
	cf := &Folder{
		name: f.name + "_clone",
	}
	tempChildren := make([]INode, 0, len(f.children))
	for _, child := range f.children {
		tempChildren = append(tempChildren, child.clone())
	}
	cf.children = tempChildren
	return cf
}
