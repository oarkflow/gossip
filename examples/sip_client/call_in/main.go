package main

import (
	"context"
	"flag"
	"fmt"
	"time"
	
	"github.com/oarkflow/gossip/pkg/client"
	"github.com/oarkflow/gossip/pkg/dialog"
	"github.com/oarkflow/gossip/pkg/sdp"
	"github.com/oarkflow/gossip/pkg/utils"
)

var callsdp = ""

func main() {
	localIP := utils.LocalIp()
	protocol := flag.String("protocol", "udp", "protocol:[udp , tcp],default=udp")
	localAddr := flag.String("local-addr", fmt.Sprintf("%s:35060", localIP), "SIP Local listening address")
	serverAddr := flag.String("server-addr", "172.20.50.121:5060", "SIP Server address")
	
	flag.Parse()
	client, err := client.NewClient(context.Background(), "snails", "snail_in", "abc", *localAddr, nil)
	if err != nil {
		panic(err)
	}
	
	err = client.Registrar(*serverAddr, *protocol)
	if err != nil {
		panic(err)
	}
	
	str := `v=0
o=- 3868331676 3868331676 IN IP4 %s
s=gosip (MacOSX)
t=0 0
m=audio 50006 RTP/AVP 8 0 101
c=IN IP4 %s
a=rtcp:50007
a=rtpmap:8 PCMA/8000
a=rtpmap:0 PCMU/8000
a=rtpmap:101 telephone-event/8000
m=video 50006 RTP/AVP 96
c=IN IP4 %s
a=rtcp:50009
a=rtpmap:96 VP8/90000
`
	
	sd, err := sdp.ParseSDP([]byte(fmt.Sprintf(str, localIP, localIP, localIP)))
	if err != nil {
		panic(err)
	}
	callsdp = sd.Marshal()
	for {
		select {
		case dl := <-client.Receive():
			go doDialog(dl)
		}
	}
}

func doDialog(dl dialog.Dialog) {
	go func() {
		for {
			select {
			case <-dl.Context().Done():
				return
			case state := <-dl.State():
				fmt.Println("state", state.State(), state.Reason())
			}
		}
	}()
	
	fmt.Println("\n\n==============")
	fmt.Println("Receive a call", dl.From().User(), "--->", dl.To().User())
	fmt.Println("==============\n\n")
	dl.Answer(callsdp)
	time.Sleep(5 * time.Second)
	dl.Bye()
}
