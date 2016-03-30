package fnal

type Traversable interface {
	current() interface{}
	next()
	rewine()
	valid() bool
}

func Map(seq interface{}, fn func(interface{}) interface{}) interface{} {
	return nil
}

func Filter(seq interface{}, fn func(interface{}) bool) interface{} {
	return nil
}

func Foldl(seq interface{}, fn func(interface{}, interface{}) interface{}) interface{} {

}
