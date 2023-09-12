package code

// Status 评测的响应状态码

const (
	StatusAccepted int = 10

	StatusCompileError        int = 20
	StatusRuntimeError        int = 21
	StatusWrongAnswer         int = 22
	StatusTimeLimitExceeded   int = 23
	StatusMemoryLimitExceeded int = 24
	StatusOutputLimitExceeded int = 25

	StatusServerFailed int = 30
)

var msg = map[int]string{
	StatusAccepted: "Accepted",

	StatusCompileError:        "Compile Error",
	StatusRuntimeError:        "Runtime Error",
	StatusWrongAnswer:         "Wrong Answer",
	StatusTimeLimitExceeded:   "Time Limit Exceeded",
	StatusMemoryLimitExceeded: "Memory Limit Exceeded",
	StatusOutputLimitExceeded: "Output Limit Exceeded",

	StatusServerFailed: "server failed",
}

func StatusMsg(status int) string {
	return msg[status]
}
