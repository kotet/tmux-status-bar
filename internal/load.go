package tmux_status_bar

import (
	"fmt"
	"log"
	"time"

	"github.com/mackerelio/go-osstat/loadavg"
)

type Load struct {
	Str SyncString
}

func (l *Load) Serve() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for ; true; <-ticker.C {
		l.update()
	}
}

func (l *Load) update() {
	la, err := loadavg.Get()

	if err != nil {
		log.Println("loadavg: ", err.Error())
		l.Str.Set("error")
		return
	}

	l.Str.Set(fmt.Sprintf("LA:%.1f", la.Loadavg1))
}
