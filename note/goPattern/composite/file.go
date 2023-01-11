package composite

import "fmt"

type File struct {
	name string
}

func (f *File) Add(child Node) {

}

func (f *File) Print(space int) {
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(f.name)
}
