package composite

import "fmt"

type Folder struct {
	name  string
	child []Node
}

func (f *Folder) Add(child Node) {
	f.child = append(f.child, child)
}

func (f *Folder) Print(space int) {
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(f.name)

	space++
	for _, v := range f.child {
		v.Print(space)
	}
}
