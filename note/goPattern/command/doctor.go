package command

import "fmt"

type doctor struct {
}

func (d *doctor) treatEye() {
	fmt.Println("treat eye")
}

func (d *doctor) treatNose() {
	fmt.Println("treat nose")
}
