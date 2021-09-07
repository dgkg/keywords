package main

import (
	"errors"
	. "math"
)

var test int = 1

func main() {
	//
}

type MyInt int32

func (mi MyInt) Add(i int) (MyInt, error) {
	res := int64(mi) + int64(i)
	if res > MaxInt32 {
		return -1, errors.New("out of range")
	}
	return MyInt(res), nil
}

func (mi MyInt) Sub(i int) (MyInt, error) {
	res := int64(mi) - int64(i)
	if res < MinInt32 {
		return -1, errors.New("out of range")
	}
	return MyInt(res), nil
}

func (mi MyInt) Multiply(i int) (MyInt, error) {
	res := int64(mi) * int64(i)
	if res > MaxInt32 {
		return -1, errors.New("out of range")
	}
	return MyInt(res), nil
}

func (mi MyInt) Divide(i int) (MyInt, error) {
	if i == 0 {
		return -1, errors.New("try to divide by 0")
	}
	return mi / MyInt(i), nil
}
