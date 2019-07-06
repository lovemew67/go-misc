package structure

import (
	"fmt"
)

type SubInfo struct {
	IsoCode string `json:"isoCode"`
	Name    string `json:"Name"`
}

func (s *SubInfo) String() string {
	return fmt.Sprintf("%+v", *s)
}

type SubInfoSet map[string]SubInfo

func NewSubInfoSet() *SubInfoSet {
	return &SubInfoSet{}
}

func (s *SubInfoSet) Len() int {
	return len(*s)
}

func (s *SubInfoSet) Add(sub SubInfo) {
	(*s)[sub.IsoCode] = sub
}

func (s *SubInfoSet) Contains(sub SubInfo) bool {
	_, ok := (*s)[sub.IsoCode]
	return ok
}

func (s *SubInfoSet) Delete(sub SubInfo) {
	delete(*s, sub.IsoCode)
}
