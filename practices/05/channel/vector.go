package channel

import (
	"sync"
)

type IVector interface {
	Len() int
	GetAt(int) float64
	SetAt(int, float64) 
}

type Vector struct {
	sync.RWMutex
	vec []float64
}

func New(args ...float64) IVector {
	v := new(Vector)
	v.vec = make([]float64, len(args))
	for i, e := range args {
		v.SetAt(i, e)
	}
	return v
}

func WithSize(size int) *Vector {
	v := Vector{
		vec: make([]float64, size),
	}

	for i := 0; i < size; i++ {
		v.SetAt(i, 0.0)
	}

	return &v
}

func (v *Vector) Len() int {
	return len(v.vec)
}

func (v *Vector) GetAt(i int) float64 {
	if i < 0 || i > v.Len() {
		panic("out of range")
	}
	return v.vec[i]
}

func (v *Vector) SetAt(i int, data float64) {
	if i < 0 || i > v.Len() {
		panic("out of range")
	}
	v.Lock()
	v.vec[i] = data
	v.Unlock()
}

func Apply(v1 IVector, v2 IVector, f func(float64, float64) float64) IVector {
	_len := v1.Len()
	if _len != v2.Len() {
		panic("unequal vector length")
	}
	out := WithSize(_len)
	var wg sync.WaitGroup
	for i := 0; i < _len; i++ {
		wg.Add(1)
		go func(v1 IVector, v2 IVector, f func(float64, float64) float64, i int) {
			defer wg.Done()
			out.SetAt(i, f(v1.GetAt(i), v2.GetAt(i)))
		}(v1, v2, f, i)
	}
	wg.Wait()
	return out
}
