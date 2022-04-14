package data_structures

type queue struct {
	values []interface{}
}

type Queue interface {
	Enqueue(val interface{})
	Dequeue() interface{}
	Size() int
}

func NewQueue() Queue {
	return &queue{}
}

func (q *queue) Size() int {
	return len(q.values)
}

func (q *queue) Enqueue(val interface{}) {
	q.values = append(q.values, val)
}

func (q *queue) Dequeue() interface{} {
	var val interface{}
	if len(q.values) > 0 {
		val = q.values[0]
		q.values = q.values[1:]
	}
	return val
}
