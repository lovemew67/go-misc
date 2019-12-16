package test

type Mock interface {
	Bar(x int) int
}

func SutMock(m Mock) {
	m.Bar(99)
}

type Stub interface {
	Bar(x int) int
}

func SutStub(s Stub) {
	s.Bar(99)
}
