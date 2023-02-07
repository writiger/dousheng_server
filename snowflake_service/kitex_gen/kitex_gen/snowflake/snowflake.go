// Code generated by Kitex v0.4.4. DO NOT EDIT.

package snowflake

import (
	"context"
	kitex_gen "dousheng_server/snowflake_service/kitex_gen/kitex_gen"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return snowflakeServiceInfo
}

var snowflakeServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Snowflake"
	handlerType := (*kitex_gen.Snowflake)(nil)
	methods := map[string]kitex.MethodInfo{
		"NewID": kitex.NewMethodInfo(newIDHandler, newNewIDArgs, newNewIDResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "snowflake",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func newIDHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_gen.NewIDRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.Snowflake).NewID(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *NewIDArgs:
		success, err := handler.(kitex_gen.Snowflake).NewID(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*NewIDResult)
		realResult.Success = success
	}
	return nil
}
func newNewIDArgs() interface{} {
	return &NewIDArgs{}
}

func newNewIDResult() interface{} {
	return &NewIDResult{}
}

type NewIDArgs struct {
	Req *kitex_gen.NewIDRequest
}

func (p *NewIDArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(kitex_gen.NewIDRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *NewIDArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *NewIDArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *NewIDArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in NewIDArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *NewIDArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.NewIDRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var NewIDArgs_Req_DEFAULT *kitex_gen.NewIDRequest

func (p *NewIDArgs) GetReq() *kitex_gen.NewIDRequest {
	if !p.IsSetReq() {
		return NewIDArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *NewIDArgs) IsSetReq() bool {
	return p.Req != nil
}

type NewIDResult struct {
	Success *kitex_gen.NewIDResponse
}

var NewIDResult_Success_DEFAULT *kitex_gen.NewIDResponse

func (p *NewIDResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(kitex_gen.NewIDResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *NewIDResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *NewIDResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *NewIDResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in NewIDResult")
	}
	return proto.Marshal(p.Success)
}

func (p *NewIDResult) Unmarshal(in []byte) error {
	msg := new(kitex_gen.NewIDResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *NewIDResult) GetSuccess() *kitex_gen.NewIDResponse {
	if !p.IsSetSuccess() {
		return NewIDResult_Success_DEFAULT
	}
	return p.Success
}

func (p *NewIDResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_gen.NewIDResponse)
}

func (p *NewIDResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) NewID(ctx context.Context, Req *kitex_gen.NewIDRequest) (r *kitex_gen.NewIDResponse, err error) {
	var _args NewIDArgs
	_args.Req = Req
	var _result NewIDResult
	if err = p.c.Call(ctx, "NewID", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}