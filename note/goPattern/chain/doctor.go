package chain

import "fmt"

// 医生
type doctor struct {
	next department
}

func (d *doctor) execute() {
	fmt.Println("Doctor checking patient")
	//d.next.execute()
}

func (d *doctor) setNext(next department) {
	d.next = next
}
