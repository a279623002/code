package chain

func Run() {
	d := &doctor{}

	r := &reception{}
	r.setNext(d)

	r.execute()
}
