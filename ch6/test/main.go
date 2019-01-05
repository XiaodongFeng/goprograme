package main

import (
	"bytes"
	"fmt"
)

type TestReceive struct {
	Name string
	Age  int
}

type TestReceive1 struct {
	Name string
	Age  int
}

type TestR struct {
	TestReceive
	TestReceive1
	s string
}

func (t *TestReceive1) TestPointer(name string) {
	t.Name = name
}

func (t *TestReceive) TestPointer(name string) {
	t.Name = name
}

func (t *TestReceive) TestPointer1(name string) {
	test := TestReceive{Name: name, Age: 16}
	*t = test
}

func (t *TestReceive) TestPointer2(name string) {
	test := TestReceive{Name: name, Age: 16}
	t = &test
	test1 := TestReceive{Name: name, Age: 17}
	*t = test1
}

func (t TestReceive) TestVar(name string) {
	t.Name = name
}

func (t TestReceive) TestVar1(name string) {
	test := TestReceive{Name: name, Age: 16}
	t = test
}

func (t TestReceive) TestVarP(name string) {
	(&t).Name = name
}

type IntSet struct {
	words []uint64
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	t := TestReceive{}
	t.TestPointer("XiaodongPointer")
	fmt.Println(t)
	t.TestVar("XiaodongVar")
	fmt.Println(t)
	t.TestVarP("XiaodongVarP")
	fmt.Println(t)
	t.TestPointer1("XiaodongPointer1")
	fmt.Println(t)
	t.TestPointer2("XiaodongPointer2")
	fmt.Println(t)
	t.TestVar1("XiaodongVar1")
	fmt.Println(t)
	t1 := TestReceive1{}

	tTest := TestR{t, t1, "test"}
	tTest.TestReceive1.TestPointer("tTest.TestPointer")

	intSet := IntSet{words: []uint64{1, 1, 3}}
	fmt.Println(intSet.String())
}
