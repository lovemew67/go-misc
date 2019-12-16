package test

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func Test_SutMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockMock(ctrl)

	// Asserts that the first and only call to Bar() is passed 99.
	// Anything else will fail.
	m.EXPECT().Bar(gomock.Eq(99)).Return(101)

	SutMock(m)
}

func Test_StubMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockStub(ctrl)

	// Does not make any assertions. Returns 101 when Bar is invoked with 99.
	m.EXPECT().Bar(gomock.Eq(99)).Return(101).AnyTimes()

	// Does not make any assertions. Returns 103 when Bar is invoked with 101.
	m.EXPECT().Bar(gomock.Eq(101)).Return(103).AnyTimes()

	SutStub(m)
}
