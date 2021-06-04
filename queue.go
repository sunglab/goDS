package queue

type queue struct {
	f, b, len int
	q         []interface{}
}

func New(i int) *queue {
	return new(queue).newQueue(i)
}

func (_q *queue) newQueue(i int) *queue {
	_q.f, _q.b, _q.len = 0, 0, 0
	_q.q = make([]interface{}, i)
	return _q
}

func (_q *queue) Push(i interface{}) {
	if _q.q == nil {
		_q.q = make([]interface{}, 1)
	}

	if _q.f > _q.len {
		// all elements moved to the front
		copy(_q.q, _q.q[_q.f:_q.b])
		_q.f = 0
		_q.b = _q.len
	}

	if !_q.full() {
		_q.q[_q.b] = i
	} else {
		_q.q = append(_q.q, i)
	}
	_q.len++
	_q.b++
}

func (_q *queue) Pop() (ret interface{}) {
	if _q.empty() {
		panic("queue is empty")
	}

	ret = _q.q[_q.f]
	_q.q[_q.f] = nil
	_q.f++
	_q.len--
	if _q.len == 0 {
		_q.f, _q.b = 0, 0
	}
	return
}

func (_q *queue) empty() bool {
	return _q.len == 0
}

func (_q *queue) full() bool {
	return _q.b == len(_q.q)
}

func (_q queue) Front() (ret interface{}) {
	ret = _q.q[_q.f]
	return
}

func (_q queue) Back() (ret interface{}) {
	ret = _q.q[_q.b]
	return
}

func (_q queue) Size() int {
	return _q.len
}

func (_q *queue) Get() *queue {
	return &_q.q
}
