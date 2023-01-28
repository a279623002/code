package chain

type department interface {
	execute()
	setNext(department)
}
