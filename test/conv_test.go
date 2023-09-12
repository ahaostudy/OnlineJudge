package test

import (
	"fmt"
	"main/internal/common"
	"testing"
)

type In struct {
	A int16
	B int32
	C int64
	D string
	E bool
	F bool
	G bool
}

type Out struct {
	A int
	B int64
	C int
	D string
	E bool
	F string
	H string
}

func TestBuild(t *testing.T) {
	in := In{ A: 1, B: 2, C: 3, D: "hello", E: true, F: false, G: true }

	out := new(Out)
	common.Build(in, out)

	fmt.Printf("out: %#v\n", out)
}
