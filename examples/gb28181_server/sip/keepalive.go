package sip

import (
	"context"
	
	"github.com/oarkflow/gossip/pkg/gb28181"
	"github.com/oarkflow/gossip/pkg/server"
)

func (d *SipHandler) Keepalive(ctx context.Context, client server.Client, msg *gb28181.Keepalive) (*server.Response, error) {
	return server.NewResponse(200, "Success"), nil
}
