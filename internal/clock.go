package tmux_status_bar

import "time"

type Clock struct {
	Str SyncString
}

func (c *Clock) Serve() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for ; true; <-ticker.C {
		c.update()
	}
}

func (c *Clock) update() {
	t := time.Now()
	s := t.Format("01/02(Mon)15:04:05")
	c.Str.Set(s)
}
