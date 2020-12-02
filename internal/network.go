package tmux_status_bar

import (
	"bytes"
	"fmt"
	"log"
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/mackerelio/go-osstat/network"
)

var network_interval_sec uint64 = 5

type Network struct {
	Str  SyncString
	prev map[string]network.Stats
}

func (n *Network) Serve() {
	ticker := time.NewTicker(time.Duration(network_interval_sec) * time.Second)
	defer ticker.Stop()
	for ; true; <-ticker.C {
		n.update()
	}
}

func (n *Network) update() {
	nets, err := network.Get()

	if err != nil {
		log.Println("network: ", err.Error())
		n.Str.Set("error")
		return
	}

	var buf bytes.Buffer

	if n.prev != nil {
		for _, net := range nets {
			if prev, ok := n.prev[net.Name]; ok {
				if prev.RxBytes != net.RxBytes || prev.TxBytes != net.TxBytes {
					buf.WriteByte('[')
					buf.WriteString(net.Name)
					buf.WriteByte(':')
					buf.WriteString(fmt.Sprintf(" U% 7s/s", humanize.Bytes((net.TxBytes-prev.TxBytes)/network_interval_sec)))
					buf.WriteString(fmt.Sprintf(" D% 7s/s", humanize.Bytes((net.RxBytes-prev.RxBytes)/network_interval_sec)))
					buf.WriteByte(']')
				}
			}
		}
	}

	n.prev = make(map[string]network.Stats)
	for _, net := range nets {
		n.prev[net.Name] = net
	}

	n.Str.Set(buf.String())
}
