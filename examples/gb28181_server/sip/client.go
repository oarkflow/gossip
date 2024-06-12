package sip

import (
	"context"
	"fmt"
	"time"
	
	"github.com/oarkflow/gossip/pkg/dialog"
	"github.com/oarkflow/gossip/pkg/message"
)

type Client struct {
	server   *SipHandler
	user     string
	protocol string
	address  string
	auth     bool
}

func (c *Client) SetTransport(protocol string, address string) {
	if c.protocol != protocol || c.address != address {
		c.address = address
		c.protocol = protocol
		c.auth = false
	}
}

func (c *Client) Transport() (protocol string, address string) {
	return c.protocol, c.address
}

func (c *Client) User() string {
	return c.user
}

func (c *Client) Password() string {
	return "12345678"
}

func (c *Client) SetAuth(auth bool) error {
	c.auth = auth
	if auth {
		go func() {
			time.Sleep(2 * time.Second)
			c.server.gb28181.GetCatalog(c)
			time.Sleep(5 * time.Second)
			c.server.gb28181.GetPresetQuery(c, "34020000001320000001")
			
			// deviceIDs := []string{c.user, "71020001001320000001"}
			// time.Sleep(1 * time.Second)
			// for _, deviceID := range deviceIDs {
			// 	// c.server.gb28181.GetDeviceInfo(c, deviceID)
			// 	// c.server.gb28181.GetDeviceStatus(c, deviceID)
			// 	// c.server.gb28181.GetDeviceConfig(c, deviceID)
			// }
			
			sdp1 := `v=0
o=71020001001320000001 0 0 IN IP4 172.20.30.61
s=Play
c=IN IP4 172.20.30.61
t=0 0
m=video 35000 RTP/AVP 96 97 98
a=recvonly
a=rtpmap:96 PS/90000
a=rtpmap:97 MPEG4/90000
a=rtpmap:98 H264/90000
y=0200010001
m=audio 35002 RTP/AVP 111 0 8
a=rtpmap:111 opus/48000/2
a=rtpmap:0 PCMU/8000
a=rtpmap:8 PCMA/8000
a=mid:audio
a=sendrecv
`
			deviceID := c.user
			// deviceID = "71020001001320000001"
			var (
				dl  dialog.Dialog
				err error
			)
			
			dl, err = c.server.gb28181.Invite(context.Background(), c, deviceID, sdp1, func(msg message.Message) {
			})
			
			if err != nil {
				panic(err)
			}
			if dl != nil {
				for {
					select {
					case <-dl.Context().Done():
						return
					case state := <-dl.State():
						fmt.Println("Receive status updates---------", state)
						if state.State() == dialog.Accepted {
							fmt.Println("The other party has answered the call:", string(dl.SDP()))
							// time.Sleep(20 * time.Second)
							// dl.Bye()
						}
						if state.State() == dialog.Error {
							fmt.Println("mistake:", state.Reason())
						}
					}
				}
			}
			
			// 预制点位调试
			// all := []ptz.PTZ_Type{ptz.Right, ptz.Left, ptz.Left, ptz.Up, ptz.Down, ptz.LeftUp, ptz.LeftDown, ptz.RightUp, ptz.RightDown, ptz.Stop}
			// for _, a := range all {
			// 	fmt.Println("方位调整", string(a))
			// 	c.server.gb28181.PTZControl(c, deviceID, ptz.PTZCmd(a, 50, 0))
			// 	time.Sleep(5 * time.Second)
			// }
			
			time.Sleep(3 * time.Second)
			c.server.gb28181.GetPresetQuery(c, deviceID)
			// fmt.Println("调用预制点位")
			// c.server.gb28181.PTZControl(c, deviceID, ptz.PTZCmd(ptz.CalPos, 0, 1))
		}()
	}
	return nil
}

func (c *Client) IsAuth() bool {
	return c.auth
}

func (c *Client) Logout() error {
	fmt.Println("User logout-----")
	c.auth = false
	return nil
}
