package judge

import (
	"fmt"
	"main/internal/data/model"
	"main/internal/service/judge/pkg/code"
	"main/internal/service/judge/pkg/errs"
	"strings"
)

type TestcaseRunOutput struct {
	idx    int
	result code.Result
}

func Judge(problem *model.Problem, codePath string, langID int) (code.Result, error) {
	// 编译代码
	c := code.NewCodeLimit(codePath, langID, problem.MaxTime, problem.MaxMemory)
	if res, ok := c.Build(); !ok {
		return res, nil
	}
	defer c.Destroy()

	// 并发执行每个样例
	out := make(chan TestcaseRunOutput)
	for i, testcase := range problem.Testcases {
		go func(i int, testcase *model.Testcase) {
			// 初始化执行结果，使用defer确保每个样例都有结果写入到管道
			o := TestcaseRunOutput{idx: i, result: code.Result{}}
			o.result.SetStatus(code.StatusServerFailed)
			defer func() { out <- o }()

			// 获取样例输入
			input, ok := testcase.GetLocalInput()
			if !ok {
				return
			}

			// 运行代码
			var err error
			o.result, err = c.Run(input)
			if err != nil {
				fmt.Println(err.Error())
			}
		}(i, testcase)
	}

	// 读取结果
	result := code.Result{}
	result.SetStatus(code.StatusAccepted)
	for i := 0; i < len(problem.Testcases); i++ {
		o := <-out
		fmt.Printf("out_%d: %#v\n", o.idx+1, o.result)
		// 保留第一个错误样例
		if result.Status != code.StatusAccepted {
			continue
		}
		if o.result.Status != code.StatusAccepted {
			result.SetStatus(o.result.Status)
			continue
		}

		result.Time += o.result.Time
		result.Memory += o.result.Memory

		// 如果没有其它错误，判断是否wa
		output, ok := problem.Testcases[o.idx].GetOutput()
		if !ok {
			result.SetStatus(code.StatusServerFailed)
			continue
		}
		// TODO: 添加SPJ（特殊判题）
		if !CmpOutput(output, o.result.Output) {
			result.SetStatus(code.StatusWrongAnswer)
			continue
		}
	}

	// 服务器错误
	if result.Status == code.StatusServerFailed {
		return result, errs.ErrJudgeFailed
	}

	return result, nil
}

func CmpOutput(a, b string) bool {
	// 去除末尾回车
	a, b = strings.TrimRight(a, "\n"), strings.TrimRight(b, "\n")

	// 拆分每行
	linesA, linesB := strings.Split(a, "\n"), strings.Split(b, "\n")

	// 如果行数不同则结果不同
	if len(linesA) != len(linesB) {
		return false
	}

	// 遍历每一行，去除行末空格后判断每行是否相同
	for i := range linesA {
		la, lb := strings.TrimRight(linesA[i], " "), strings.TrimRight(linesB[i], " ")
		if la != lb {
			return false
		}
	}

	return true
}
