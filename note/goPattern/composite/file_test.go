package composite

import "testing"

func TestFile_Print(t *testing.T) {
	dir := &Folder{name: "D:"}
	f1 := &File{name: "file.txt"}
	f2 := &File{name: "file.png"}

	dir.Add(f1)
	dir.Add(f2)

	dir1 := &Folder{name: "work"}
	f3 := &File{name: "file.go"}
	f4 := &File{name: "file.php"}

	dir.Add(dir1)
	dir1.Add(f3)
	dir1.Add(f4)

	dir.Print(0)
}
