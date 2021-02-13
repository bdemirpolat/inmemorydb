package model

const (
	SET = iota
	GET
)

type Message struct {
	Type  int
	Key   string
	Value string
}

type Response struct {
	Status bool
	Data   string
}