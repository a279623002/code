package chain

import "fmt"

// 前台
type reception struct {
	next department
}

func (r *reception) execute() {
	fmt.Println("reception registering patient")
	r.next.execute()
}

func (r *reception) setNext(next department) {
	r.next = next
}
