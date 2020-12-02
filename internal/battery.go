package tmux_status_bar

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/distatus/battery"
)

type Battery struct {
	Str SyncString
}

func (b *Battery) Serve() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for ; true; <-ticker.C {
		b.update()
	}
}

func (b *Battery) update() {
	bat, err := battery.Get(0)
	if err != nil {
		log.Println("battery: ", err.Error())
		b.Str.Set("error")
		return
	}
	curr := bat.Current

	var buf bytes.Buffer
	buf.WriteString("Bat:")
	buf.WriteString(fmt.Sprintf("%.2fWh", curr/1000))
	buf.WriteByte('(')
	buf.WriteByte(bat.State.String()[0])
	buf.WriteByte(')')
	b.Str.Set(buf.String())
}
