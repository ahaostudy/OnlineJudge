package test

import (
	"fmt"
	"main/internal/common"
	"testing"
)

type A struct {
	AA int
	BB bool
}

type B struct {
	AA int64
	BB bool
}

type In struct {
	A int16
	B int32
	C int64
	D string
	E bool
	F bool
	G bool
	S []*A
}

type Out struct {
	A int
	B int64
	C int
	D string
	E bool
	F string
	H bool
	S []B
}

func TestBuild(t *testing.T) {
	in := In{A: 1, B: 2, C: 3, D: "hello", E: true, F: false, G: true, S: []*A{
		{
			AA: 1,
			BB: false,
		},
		{
			AA: 1,
			BB: false,
		},
		{
			AA: 1,
			BB: false,
		},
	}}

	out := new(Out)
	builder := new(common.Builder)
	builder.Build(in, out)
	out.S = make([]B, len(in.S))
	for i := range in.S {
		builder.Build(in.S[i], &out.S[i])
	}
	if builder.Error() != nil {
		panic(builder.Error())
	}

	fmt.Printf("out: %#v\n", out)
}
