package sip

import (
	"context"
	
	"github.com/davecgh/go-spew/spew"
	
	"github.com/oarkflow/gossip/pkg/gb28181"
	"github.com/oarkflow/gossip/pkg/server"
)

func (d *SipHandler) Catalog(ctx context.Context, client server.Client, catalog *gb28181.Catalog) error {
	spew.Dump(catalog)
	// client, err := d.GetClient("34020000001110000002")
	// if err != nil {
	// 	return nil
	// }
	
	// _, err = d.gb28181.GetDeviceInfo(client, catalog.DeviceID)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, item := range catalog.Item {
	// 	_, err = d.gb28181.GetDeviceInfo(client, item.DeviceID)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	
	return nil
}
