package transport

import (
	"context"
)

type Transport interface {
	Send(ctx context.Context, req interface{}) (rsp interface{}, err error)
}
