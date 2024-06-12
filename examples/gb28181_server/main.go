package main

import (
	"context"
	"flag"
	"fmt"
	
	"github.com/oarkflow/gossip/examples/gb28181_server/sip"
	"github.com/oarkflow/gossip/pkg/server"
	"github.com/oarkflow/gossip/pkg/utils"
)

func main() {
	ip := flag.String("ip", utils.LocalIp(), "Listening Address")
	port := flag.Uint64("port", 5060, "Listening Port")
	flag.Parse()
	ctx := context.Background()
	handler := sip.NewSipHandler("99920000002000000000", "99999999")
	server := server.NewServer(true, handler)
	handler.SetServer(server)
	
	err := server.SIPListen(ctx, "0.0.0.0", *ip, uint16(*port), "udp", "tcp")
	if err != nil {
		fmt.Println(err)
	}
}
