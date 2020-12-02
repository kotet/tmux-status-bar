package tmux_status_bar

import (
	"fmt"
	"log"
	"time"

	"github.com/mackerelio/go-osstat/memory"
)

type Memory struct {
	Str SyncString
}

func (m *Memory) Serve() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for ; true; <-ticker.C {
		m.update()
	}
}

func (m *Memory) update() {
	mem, err := memory.Get()

	if err != nil {
		log.Println("memory: ", err.Error())
		m.Str.Set("error")
		return
	}

	memory_usage := float64(mem.Used) / float64(mem.Total) * 100
	swap_usage := float64(mem.SwapUsed) / float64(mem.SwapTotal) * 100

	m.Str.Set(fmt.Sprintf("Mem:%.0f%% Swp:%.0f%%", memory_usage, swap_usage))
}
