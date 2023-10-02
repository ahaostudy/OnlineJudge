package test

import (
	"fmt"
	"testing"

	"main/kitex_gen/problem"
)

func TestGetTestcase(t *testing.T) {
	res, err := ProblemCli.GetTestcase(Ctx, &problem.GetTestcaseRequest{
		ID: 10,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("res.GetTestcase(): %#v\n", res.GetTestcase())
}
