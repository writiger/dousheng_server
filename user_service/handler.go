package main

import (
	"context"
	"dousheng_server/user_service/kitex_gen/kitex_gen"
	"fmt"
)

// UserCenterImpl implements the last service interface defined in the IDL.
type UserCenterImpl struct{}

// Ping implements the UserCenterImpl interface.
func (s *UserCenterImpl) Ping(ctx context.Context, req *kitex_gen.Request) (resp *kitex_gen.Response, err error) {
	// TODO: Your code here...
	resp = new(kitex_gen.Response)
	fmt.Println("Get Ping Message:", req.Ping)
	resp.Pong = "Hello Client"
	return
}
