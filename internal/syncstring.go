package tmux_status_bar

import "sync"

type SyncString struct {
	str string
	mut sync.Mutex
}

func (s *SyncString) Set(str string) {
	s.mut.Lock()
	s.str = str
	s.mut.Unlock()
}

func (s *SyncString) Get() string {
	s.mut.Lock()
	str := s.str
	s.mut.Unlock()
	return str
}
