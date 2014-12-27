package cartel

type Task interface {
	Execute() OutputValue
}

type OutputValue interface {
	Value() interface{}
}
