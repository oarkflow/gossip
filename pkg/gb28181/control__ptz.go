package gb28181

import (
	"time"
	
	"github.com/oarkflow/gossip/pkg/server"
)

func (g *GB28181) PTZControl(client server.Client, deviceID string, ptzCMD string) error {
	sn := time.Now().Unix()
	_, err := g.SendMessage(client, &Control{
		CmdType:  CmdType__DeviceControl,
		SN:       sn,
		DeviceID: deviceID,
		ControlCommand: ControlCommand{
			PTZCmd: ptzCMD,
			Info: ControlCommandInfo{
				ControlPriority: 1,
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}
