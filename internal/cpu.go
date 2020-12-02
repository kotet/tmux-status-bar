package tmux_status_bar

import (
	"fmt"
	"log"
	"time"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

type CPU struct {
	Str SyncString
}

func (c *CPU) Serve() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for ; true; <-ticker.C {
		c.update()
	}
}

func (c *CPU) update() {
	cpu, err := linuxproc.ReadCPUInfo("/proc/cpuinfo")

	if err != nil {
		log.Println("cpu: ", err.Error())
		c.Str.Set("error")
		return
	}
	c.Str.Set(fmt.Sprintf("Freq:%.1fGhz", cpu.Processors[0].MHz/1024))
}
