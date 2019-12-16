package test

import (
	"bytes"
	"log"
	"net/http"
	"testing"

	gomock "github.com/golang/mock/gomock"
	gock "gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/assert"
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

func Test_Do(t *testing.T) {
	client := &http.Client{}
	defer gock.Off() // Flush pending mocks after test execution
	gock.InterceptClient(client)
	defer gock.RestoreClient(client)

	gock.New("http://test.com").
		Post("/test").
		Reply(200).
		JSON(map[string]string{
			"id": "123",
		})

	req, err := http.NewRequest(http.MethodPost, "http://test.com/test", bytes.NewBuffer([]byte(`abcdefg`)))
	assert.NoError(t, err)
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func BenchmarkClient(b *testing.B) {
	client := &http.Client{}
	defer gock.Off() // Flush pending mocks after test execution
	gock.InterceptClient(client)
	defer gock.RestoreClient(client)

	gock.New("http://test.com").
		Get("/test").
		Reply(200).
		JSON(map[string]string{
			"id": "123",
		})

	req, _ := http.NewRequest(http.MethodGet, "http://test.com/test", nil)
	count := 0
	errCount := 0
	for i := 0; i < b.N; i++ {
		_, err := client.Do(req)
		if err != nil {
			errCount++
		}
		count++
	}
	log.Println("count:", count)
	log.Println("error rate:", float64(errCount)/float64(count))
}
