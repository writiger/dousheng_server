package main

import (
	"context"
	"dousheng_server/snowflake_service/kitex_gen"
	"dousheng_server/snowflake_service/snowfalke"
)

// SnowflakeImpl implements the last service interface defined in the IDL.
type SnowflakeImpl struct{}

// NewID implements the SnowflakeImpl interface.
func (s *SnowflakeImpl) NewID(ctx context.Context, req *kitex_gen.NewIDRequest) (resp *kitex_gen.NewIDResponse, err error) {
	// TODO: Your code here...
	resp = new(kitex_gen.NewIDResponse)
	resp.ID = snowfalke.NewUUID()
	return
}
