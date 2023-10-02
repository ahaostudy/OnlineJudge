// Code generated by Kitex v0.7.2. DO NOT EDIT.

package contestservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	contest "main/kitex_gen/contest"
)

func serviceInfo() *kitex.ServiceInfo {
	return contestServiceServiceInfo
}

var contestServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ContestService"
	handlerType := (*contest.ContestService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetContest":     kitex.NewMethodInfo(getContestHandler, newGetContestArgs, newGetContestResult, false),
		"GetContestList": kitex.NewMethodInfo(getContestListHandler, newGetContestListArgs, newGetContestListResult, false),
		"CreateContest":  kitex.NewMethodInfo(createContestHandler, newCreateContestArgs, newCreateContestResult, false),
		"DeleteContest":  kitex.NewMethodInfo(deleteContestHandler, newDeleteContestArgs, newDeleteContestResult, false),
		"UpdateContest":  kitex.NewMethodInfo(updateContestHandler, newUpdateContestArgs, newUpdateContestResult, false),
		"Register":       kitex.NewMethodInfo(registerHandler, newRegisterArgs, newRegisterResult, false),
		"UnRegister":     kitex.NewMethodInfo(unRegisterHandler, newUnRegisterArgs, newUnRegisterResult, false),
		"IsRegister":     kitex.NewMethodInfo(isRegisterHandler, newIsRegisterArgs, newIsRegisterResult, false),
		"IsAccessible":   kitex.NewMethodInfo(isAccessibleHandler, newIsAccessibleArgs, newIsAccessibleResult, false),
		"ContestRank":    kitex.NewMethodInfo(contestRankHandler, newContestRankArgs, newContestRankResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "contest",
		"ServiceFilePath": ``,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.7.2",
		Extra:           extra,
	}
	return svcInfo
}

func getContestHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(contest.GetContestRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(contest.ContestService).GetContest(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetContestArgs:
		success, err := handler.(contest.ContestService).GetContest(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetContestResult)
		realResult.Success = success
	}
	return nil
}
func newGetContestArgs() interface{} {
	return &GetContestArgs{}
}

func newGetContestResult() interface{} {
	return &GetContestResult{}
}

type GetContestArgs struct {
	Req *contest.GetContestRequest
}

func (p *GetContestArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(contest.GetContestRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetContestArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetContestArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetContestArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetContestArgs) Unmarshal(in []byte) error {
	msg := new(contest.GetContestRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetContestArgs_Req_DEFAULT *contest.GetContestRequest

func (p *GetContestArgs) GetReq() *contest.GetContestRequest {
	if !p.IsSetReq() {
		return GetContestArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetContestArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetContestArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetContestResult struct {
	Success *contest.GetContestResponse
}

var GetContestResult_Success_DEFAULT *contest.GetContestResponse

func (p *GetContestResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(contest.GetContestResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetContestResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetContestResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetContestResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetContestResult) Unmarshal(in []byte) error {
	msg := new(contest.GetContestResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetContestResult) GetSuccess() *contest.GetContestResponse {
	if !p.IsSetSuccess() {
		return GetContestResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetContestResult) SetSuccess(x interface{}) {
	p.Success = x.(*contest.GetContestResponse)
}

func (p *GetContestResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetContestResult) GetResult() interface{} {
	return p.Success
}

func getContestListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(contest.GetContestListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(contest.ContestService).GetContestList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetContestListArgs:
		success, err := handler.(contest.ContestService).GetContestList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetContestListResult)
		realResult.Success = success
	}
	return nil
}
func newGetContestListArgs() interface{} {
	return &GetContestListArgs{}
}

func newGetContestListResult() interface{} {
	return &GetContestListResult{}
}

type GetContestListArgs struct {
	Req *contest.GetContestListRequest
}

func (p *GetContestListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(contest.GetContestListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetContestListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetContestListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetContestListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetContestListArgs) Unmarshal(in []byte) error {
	msg := new(contest.GetContestListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetContestListArgs_Req_DEFAULT *contest.GetContestListRequest

func (p *GetContestListArgs) GetReq() *contest.GetContestListRequest {
	if !p.IsSetReq() {
		return GetContestListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetContestListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetContestListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetContestListResult struct {
	Success *contest.GetContestListResponse
}

var GetContestListResult_Success_DEFAULT *contest.GetContestListResponse

func (p *GetContestListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(contest.GetContestListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetContestListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetContestListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetContestListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetContestListResult) Unmarshal(in []byte) error {
	msg := new(contest.GetContestListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetContestListResult) GetSuccess() *contest.GetContestListResponse {
	if !p.IsSetSuccess() {
		return GetContestListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetContestListResult) SetSuccess(x interface{}) {
	p.Success = x.(*contest.GetContestListResponse)
}

func (p *GetContestListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetContestListResult) GetResult() interface{} {
	return p.Success
}

func createContestHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(contest.CreateContestRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(contest.ContestService).CreateContest(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateContestArgs:
		success, err := handler.(contest.ContestService).CreateContest(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateContestResult)
		realResult.Success = success
	}
	return nil
}
func newCreateContestArgs() interface{} {
	return &CreateContestArgs{}
}

func newCreateContestResult() interface{} {
	return &CreateContestResult{}
}

type CreateContestArgs struct {
	Req *contest.CreateContestRequest
}

func (p *CreateContestArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(contest.CreateContestRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateContestArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateContestArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateContestArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateContestArgs) Unmarshal(in []byte) error {
	msg := new(contest.CreateContestRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateContestArgs_Req_DEFAULT *contest.CreateContestRequest

func (p *CreateContestArgs) GetReq() *contest.CreateContestRequest {
	if !p.IsSetReq() {
		return CreateContestArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateContestArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateContestArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateContestResult struct {
	Success *contest.CreateContestResponse
}

var CreateContestResult_Success_DEFAULT *contest.CreateContestResponse

func (p *CreateContestResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(contest.CreateContestResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateContestResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateContestResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateContestResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateContestResult) Unmarshal(in []byte) error {
	msg := new(contest.CreateContestResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateContestResult) GetSuccess() *contest.CreateContestResponse {
	if !p.IsSetSuccess() {
		return CreateContestResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateContestResult) SetSuccess(x interface{}) {
	p.Success = x.(*contest.CreateContestResponse)
}

func (p *CreateContestResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateContestResult) GetResult() interface{} {
	return p.Success
}

func deleteContestHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(contest.DeleteContestRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(contest.ContestService).DeleteContest(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeleteContestArgs:
		success, err := handler.(contest.ContestService).DeleteContest(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteContestResult)
		realResult.Success = success
	}
	return nil
}
func newDeleteContestArgs() interface{} {
	return &DeleteContestArgs{}
}

func newDeleteContestResult() interface{} {
	return &DeleteContestResult{}
}

type DeleteContestArgs struct {
	Req *contest.DeleteContestRequest
}

func (p *DeleteContestArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(contest.DeleteContestRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteContestArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteContestArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteContestArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteContestArgs) Unmarshal(in []byte) error {
	msg := new(contest.DeleteContestRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteContestArgs_Req_DEFAULT *contest.DeleteContestRequest

func (p *DeleteContestArgs) GetReq() *contest.DeleteContestRequest {
	if !p.IsSetReq() {
		return DeleteContestArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteContestArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteContestArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteContestResult struct {
	Success *contest.DeleteContestResponse
}

var DeleteContestResult_Success_DEFAULT *contest.DeleteContestResponse

func (p *DeleteContestResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(contest.DeleteContestResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteContestResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteContestResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteContestResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteContestResult) Unmarshal(in []byte) error {
	msg := new(contest.DeleteContestResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteContestResult) GetSuccess() *contest.DeleteContestResponse {
	if !p.IsSetSuccess() {
		return DeleteContestResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteContestResult) SetSuccess(x interface{}) {
	p.Success = x.(*contest.DeleteContestResponse)
}

func (p *DeleteContestResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteContestResult) GetResult() interface{} {
	return p.Success
}

func updateContestHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(contest.UpdateContestRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(contest.ContestService).UpdateContest(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UpdateContestArgs:
		success, err := handler.(contest.ContestService).UpdateContest(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateContestResult)
		realResult.Success = success
	}
	return nil
}
func newUpdateContestArgs() interface{} {
	return &UpdateContestArgs{}
}

func newUpdateContestResult() interface{} {
	return &UpdateContestResult{}
}

type UpdateContestArgs struct {
	Req *contest.UpdateContestRequest
}

func (p *UpdateContestArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(contest.UpdateContestRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateContestArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateContestArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateContestArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateContestArgs) Unmarshal(in []byte) error {
	msg := new(contest.UpdateContestRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateContestArgs_Req_DEFAULT *contest.UpdateContestRequest

func (p *UpdateContestArgs) GetReq() *contest.UpdateContestRequest {
	if !p.IsSetReq() {
		return UpdateContestArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateContestArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateContestArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateContestResult struct {
	Success *contest.UpdateContestResponse
}

var UpdateContestResult_Success_DEFAULT *contest.UpdateContestResponse

func (p *UpdateContestResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(contest.UpdateContestResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateContestResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateContestResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateContestResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateContestResult) Unmarshal(in []byte) error {
	msg := new(contest.UpdateContestResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateContestResult) GetSuccess() *contest.UpdateContestResponse {
	if !p.IsSetSuccess() {
		return UpdateContestResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateContestResult) SetSuccess(x interface{}) {
	p.Success = x.(*contest.UpdateContestResponse)
}

func (p *UpdateContestResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateContestResult) GetResult() interface{} {
	return p.Success
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(contest.RegisterRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(contest.ContestService).Register(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RegisterArgs:
		success, err := handler.(contest.ContestService).Register(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RegisterResult)
		realResult.Success = success
	}
	return nil
}
func newRegisterArgs() interface{} {
	return &RegisterArgs{}
}

func newRegisterResult() interface{} {
	return &RegisterResult{}
}

type RegisterArgs struct {
	Req *contest.RegisterRequest
}

func (p *RegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(contest.RegisterRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RegisterArgs) Unmarshal(in []byte) error {
	msg := new(contest.RegisterRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RegisterArgs_Req_DEFAULT *contest.RegisterRequest

func (p *RegisterArgs) GetReq() *contest.RegisterRequest {
	if !p.IsSetReq() {
		return RegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RegisterArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RegisterResult struct {
	Success *contest.RegisterResponse
}

var RegisterResult_Success_DEFAULT *contest.RegisterResponse

func (p *RegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(contest.RegisterResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RegisterResult) Unmarshal(in []byte) error {
	msg := new(contest.RegisterResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RegisterResult) GetSuccess() *contest.RegisterResponse {
	if !p.IsSetSuccess() {
		return RegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*contest.RegisterResponse)
}

func (p *RegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RegisterResult) GetResult() interface{} {
	return p.Success
}

func unRegisterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(contest.UnRegisterRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(contest.ContestService).UnRegister(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UnRegisterArgs:
		success, err := handler.(contest.ContestService).UnRegister(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UnRegisterResult)
		realResult.Success = success
	}
	return nil
}
func newUnRegisterArgs() interface{} {
	return &UnRegisterArgs{}
}

func newUnRegisterResult() interface{} {
	return &UnRegisterResult{}
}

type UnRegisterArgs struct {
	Req *contest.UnRegisterRequest
}

func (p *UnRegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(contest.UnRegisterRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UnRegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UnRegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UnRegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UnRegisterArgs) Unmarshal(in []byte) error {
	msg := new(contest.UnRegisterRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UnRegisterArgs_Req_DEFAULT *contest.UnRegisterRequest

func (p *UnRegisterArgs) GetReq() *contest.UnRegisterRequest {
	if !p.IsSetReq() {
		return UnRegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UnRegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UnRegisterArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UnRegisterResult struct {
	Success *contest.UnRegisterResponse
}

var UnRegisterResult_Success_DEFAULT *contest.UnRegisterResponse

func (p *UnRegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(contest.UnRegisterResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UnRegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UnRegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UnRegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UnRegisterResult) Unmarshal(in []byte) error {
	msg := new(contest.UnRegisterResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UnRegisterResult) GetSuccess() *contest.UnRegisterResponse {
	if !p.IsSetSuccess() {
		return UnRegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UnRegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*contest.UnRegisterResponse)
}

func (p *UnRegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UnRegisterResult) GetResult() interface{} {
	return p.Success
}

func isRegisterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(contest.IsRegisterRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(contest.ContestService).IsRegister(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *IsRegisterArgs:
		success, err := handler.(contest.ContestService).IsRegister(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*IsRegisterResult)
		realResult.Success = success
	}
	return nil
}
func newIsRegisterArgs() interface{} {
	return &IsRegisterArgs{}
}

func newIsRegisterResult() interface{} {
	return &IsRegisterResult{}
}

type IsRegisterArgs struct {
	Req *contest.IsRegisterRequest
}

func (p *IsRegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(contest.IsRegisterRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *IsRegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *IsRegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *IsRegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *IsRegisterArgs) Unmarshal(in []byte) error {
	msg := new(contest.IsRegisterRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var IsRegisterArgs_Req_DEFAULT *contest.IsRegisterRequest

func (p *IsRegisterArgs) GetReq() *contest.IsRegisterRequest {
	if !p.IsSetReq() {
		return IsRegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *IsRegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *IsRegisterArgs) GetFirstArgument() interface{} {
	return p.Req
}

type IsRegisterResult struct {
	Success *contest.IsRegisterResponse
}

var IsRegisterResult_Success_DEFAULT *contest.IsRegisterResponse

func (p *IsRegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(contest.IsRegisterResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *IsRegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *IsRegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *IsRegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *IsRegisterResult) Unmarshal(in []byte) error {
	msg := new(contest.IsRegisterResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *IsRegisterResult) GetSuccess() *contest.IsRegisterResponse {
	if !p.IsSetSuccess() {
		return IsRegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *IsRegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*contest.IsRegisterResponse)
}

func (p *IsRegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *IsRegisterResult) GetResult() interface{} {
	return p.Success
}

func isAccessibleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(contest.IsAccessibleRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(contest.ContestService).IsAccessible(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *IsAccessibleArgs:
		success, err := handler.(contest.ContestService).IsAccessible(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*IsAccessibleResult)
		realResult.Success = success
	}
	return nil
}
func newIsAccessibleArgs() interface{} {
	return &IsAccessibleArgs{}
}

func newIsAccessibleResult() interface{} {
	return &IsAccessibleResult{}
}

type IsAccessibleArgs struct {
	Req *contest.IsAccessibleRequest
}

func (p *IsAccessibleArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(contest.IsAccessibleRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *IsAccessibleArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *IsAccessibleArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *IsAccessibleArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *IsAccessibleArgs) Unmarshal(in []byte) error {
	msg := new(contest.IsAccessibleRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var IsAccessibleArgs_Req_DEFAULT *contest.IsAccessibleRequest

func (p *IsAccessibleArgs) GetReq() *contest.IsAccessibleRequest {
	if !p.IsSetReq() {
		return IsAccessibleArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *IsAccessibleArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *IsAccessibleArgs) GetFirstArgument() interface{} {
	return p.Req
}

type IsAccessibleResult struct {
	Success *contest.IsAccessibleResponse
}

var IsAccessibleResult_Success_DEFAULT *contest.IsAccessibleResponse

func (p *IsAccessibleResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(contest.IsAccessibleResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *IsAccessibleResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *IsAccessibleResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *IsAccessibleResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *IsAccessibleResult) Unmarshal(in []byte) error {
	msg := new(contest.IsAccessibleResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *IsAccessibleResult) GetSuccess() *contest.IsAccessibleResponse {
	if !p.IsSetSuccess() {
		return IsAccessibleResult_Success_DEFAULT
	}
	return p.Success
}

func (p *IsAccessibleResult) SetSuccess(x interface{}) {
	p.Success = x.(*contest.IsAccessibleResponse)
}

func (p *IsAccessibleResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *IsAccessibleResult) GetResult() interface{} {
	return p.Success
}

func contestRankHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(contest.ContestRankRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(contest.ContestService).ContestRank(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ContestRankArgs:
		success, err := handler.(contest.ContestService).ContestRank(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ContestRankResult)
		realResult.Success = success
	}
	return nil
}
func newContestRankArgs() interface{} {
	return &ContestRankArgs{}
}

func newContestRankResult() interface{} {
	return &ContestRankResult{}
}

type ContestRankArgs struct {
	Req *contest.ContestRankRequest
}

func (p *ContestRankArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(contest.ContestRankRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ContestRankArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ContestRankArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ContestRankArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ContestRankArgs) Unmarshal(in []byte) error {
	msg := new(contest.ContestRankRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ContestRankArgs_Req_DEFAULT *contest.ContestRankRequest

func (p *ContestRankArgs) GetReq() *contest.ContestRankRequest {
	if !p.IsSetReq() {
		return ContestRankArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContestRankArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ContestRankArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ContestRankResult struct {
	Success *contest.ContestRankResponse
}

var ContestRankResult_Success_DEFAULT *contest.ContestRankResponse

func (p *ContestRankResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(contest.ContestRankResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ContestRankResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ContestRankResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ContestRankResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ContestRankResult) Unmarshal(in []byte) error {
	msg := new(contest.ContestRankResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContestRankResult) GetSuccess() *contest.ContestRankResponse {
	if !p.IsSetSuccess() {
		return ContestRankResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContestRankResult) SetSuccess(x interface{}) {
	p.Success = x.(*contest.ContestRankResponse)
}

func (p *ContestRankResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContestRankResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetContest(ctx context.Context, Req *contest.GetContestRequest) (r *contest.GetContestResponse, err error) {
	var _args GetContestArgs
	_args.Req = Req
	var _result GetContestResult
	if err = p.c.Call(ctx, "GetContest", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetContestList(ctx context.Context, Req *contest.GetContestListRequest) (r *contest.GetContestListResponse, err error) {
	var _args GetContestListArgs
	_args.Req = Req
	var _result GetContestListResult
	if err = p.c.Call(ctx, "GetContestList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateContest(ctx context.Context, Req *contest.CreateContestRequest) (r *contest.CreateContestResponse, err error) {
	var _args CreateContestArgs
	_args.Req = Req
	var _result CreateContestResult
	if err = p.c.Call(ctx, "CreateContest", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteContest(ctx context.Context, Req *contest.DeleteContestRequest) (r *contest.DeleteContestResponse, err error) {
	var _args DeleteContestArgs
	_args.Req = Req
	var _result DeleteContestResult
	if err = p.c.Call(ctx, "DeleteContest", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateContest(ctx context.Context, Req *contest.UpdateContestRequest) (r *contest.UpdateContestResponse, err error) {
	var _args UpdateContestArgs
	_args.Req = Req
	var _result UpdateContestResult
	if err = p.c.Call(ctx, "UpdateContest", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Register(ctx context.Context, Req *contest.RegisterRequest) (r *contest.RegisterResponse, err error) {
	var _args RegisterArgs
	_args.Req = Req
	var _result RegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UnRegister(ctx context.Context, Req *contest.UnRegisterRequest) (r *contest.UnRegisterResponse, err error) {
	var _args UnRegisterArgs
	_args.Req = Req
	var _result UnRegisterResult
	if err = p.c.Call(ctx, "UnRegister", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsRegister(ctx context.Context, Req *contest.IsRegisterRequest) (r *contest.IsRegisterResponse, err error) {
	var _args IsRegisterArgs
	_args.Req = Req
	var _result IsRegisterResult
	if err = p.c.Call(ctx, "IsRegister", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsAccessible(ctx context.Context, Req *contest.IsAccessibleRequest) (r *contest.IsAccessibleResponse, err error) {
	var _args IsAccessibleArgs
	_args.Req = Req
	var _result IsAccessibleResult
	if err = p.c.Call(ctx, "IsAccessible", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContestRank(ctx context.Context, Req *contest.ContestRankRequest) (r *contest.ContestRankResponse, err error) {
	var _args ContestRankArgs
	_args.Req = Req
	var _result ContestRankResult
	if err = p.c.Call(ctx, "ContestRank", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}