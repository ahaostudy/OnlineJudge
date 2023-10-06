package judge

import (
	"fmt"
	"strings"

	"main/common/status"
	"main/services/judge/dal/model"
	"main/services/judge/pkg/code"
	"main/services/judge/pkg/errs"
)

type TestcaseRunOutput struct {
	idx    int
	result code.Result
}

func Judge_(problem *model.Problem, codePath string, langID int) (code.Result, error) {
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
			o.result.SetStatus(status.StatusServerFailed)
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
	result.SetStatus(status.StatusAccepted)
	for i := 0; i < len(problem.Testcases); i++ {
		o := <-out
		fmt.Printf("out_%d: %#v\n", o.idx+1, o.result)
		// 保留第一个错误样例
		if result.Status != status.StatusAccepted {
			continue
		}
		if o.result.Status != status.StatusAccepted {
			result.SetStatus(o.result.Status)
			continue
		}

		result.Time = max(result.Time, o.result.Time)
		result.Memory = max(result.Memory, o.result.Memory)

		// 如果没有其它错误，判断是否wa
		output, ok := problem.Testcases[o.idx].GetOutput()
		if !ok {
			result.SetStatus(status.StatusServerFailed)
			continue
		}
		// TODO: 添加SPJ（特殊判题）
		if !CmpOutput(output, o.result.Output) {
			result.SetStatus(status.StatusWrongAnswer)
			continue
		}
	}

	// 服务器错误
	if result.Status == status.StatusServerFailed {
		return result, errs.ErrJudgeFailed
	}

	return result, nil
}

func Judge(problem *model.Problem, codePath string, langID int) (code.Result, error) {
	// 编译代码
	c := code.NewCodeLimit(codePath, langID, problem.MaxTime, problem.MaxMemory)
	if res, ok := c.Build(); !ok {
		return res, nil
	}
	defer c.Destroy()

	// 执行每个样例
	result := code.Result{}
	result.SetStatus(status.StatusAccepted)
	for i, testcase := range problem.Testcases {
		// 获取样例输入
		input, ok := testcase.GetLocalInput()
		if !ok {
			result.SetStatus(status.StatusServerFailed)
			return result, nil
		}

		// 运行代码
		res, err := c.Run(input)
		if err != nil {
			result.SetStatus(status.StatusServerFailed)
			return result, nil
		}

		result.Message, result.Error = res.Message, res.Error
		if res.Status != status.StatusAccepted {
			result.SetStatus(res.Status)
			return result, nil
		}

		fmt.Printf("out_%d: %#v\n", i+1, res)

		result.Time = max(result.Time, res.Time)
		result.Memory = max(result.Memory, res.Memory)

		// 如果没有其它错误，判断是否wa
		output, ok := problem.Testcases[i].GetOutput()
		if !ok {
			result.SetStatus(status.StatusServerFailed)
			return result, nil
		}
		// TODO: 添加SPJ（特殊判题）
		if !CmpOutput(output, res.Output) {
			result.SetStatus(status.StatusWrongAnswer)
			return result, nil
		}
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

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
