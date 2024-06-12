package sip

import (
	"context"
	
	"github.com/davecgh/go-spew/spew"
	
	"github.com/oarkflow/gossip/pkg/gb28181"
	"github.com/oarkflow/gossip/pkg/server"
)

func (d *SipHandler) Broadcast(ctx context.Context, client server.Client, bl *gb28181.BroadcastResponse) {
	spew.Dump(bl)
}

func (d *SipHandler) StartBroadcast(client server.Client, sourceID string, targetID string) (int64, error) {
	return d.gb28181.StartBroadcast(client, sourceID, targetID)
}
