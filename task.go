package cartel

type Task interface {
	Execute() interface{}
}
