package exec

// Result 响应结果
type Result struct {
	CpuTime  int64 `json:"cpu_time"`
	RealTime int64 `json:"real_time"`
	Memory   int64 `json:"memory"`
	Signal   int   `json:"signal"`
	ExitCode int   `json:"exit_code"`
	Error    int   `json:"error"`
	Result   int   `json:"result"`
}

// result
const (
	RESULT_SUCCESS                  int = 0
	RESULT_CPU_TIME_LIMIT_EXCEEDED  int = 1
	RESULT_REAL_TIME_LIMIT_EXCEEDED int = 2
	RESULT_MEMORY_LIMIT_EXCEEDED    int = 3
	RESULT_RUNTIME_ERROR            int = 4
	RESULT_SYSTEM_ERROR             int = 5
)

// error
const (
	ERROR_SUCCESS             int = 0
	ERROR_INVALID_CONFIG      int = -1
	ERROR_FORK_FAILED         int = -2
	ERROR_PTHREAD_FAILED      int = -3
	ERROR_WAIT_FAILED         int = -4
	ERROR_ROOT_REQUIRED       int = -5
	ERROR_LOAD_SECCOMP_FAILED int = -6
	ERROR_SETRLIMIT_FAILED    int = -7
	ERROR_DUP2_FAILED         int = -8
	ERROR_SETUID_FAILED       int = -9
	ERROR_EXECVE_FAILED       int = -10
)
