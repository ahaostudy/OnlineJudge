package handle

import rpcTestcase "main/api/testcase"

type TestcaseServer struct {
	rpcTestcase.UnimplementedTestcaseServiceServer
}
