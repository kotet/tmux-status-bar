package main

import (
	"bytes"
	"log"
	"net"
	"os"
	"os/signal"

	tmux "github.com/kotet/tmux-status-bar/internal"
)

var socket_file string = "/tmp/tmux-status-bar.sock"

func main() {
	listener, err := net.Listen("unix", socket_file)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer end()
	defer_chan := make(chan os.Signal, 1)
	signal.Notify(defer_chan, os.Interrupt)
	go func() {
		for signal := range defer_chan {
			close(defer_chan)
			log.Println(signal.String())
			end()
			os.Exit(130)
		}
	}()

	var clock tmux.Clock
	go clock.Serve()
	var battery tmux.Battery
	go battery.Serve()
	var memory tmux.Memory
	go memory.Serve()
	var cpu tmux.CPU
	go cpu.Serve()
	var load tmux.Load
	go load.Serve()
	var net tmux.Network
	go net.Serve()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var buf bytes.Buffer

		buf.WriteString(net.Str.Get())
		buf.WriteByte(' ')
		buf.WriteString(load.Str.Get())
		buf.WriteByte(' ')
		buf.WriteString(cpu.Str.Get())
		buf.WriteByte(' ')
		buf.WriteString(memory.Str.Get())
		buf.WriteByte(' ')
		buf.WriteString(battery.Str.Get())
		buf.WriteByte(' ')
		buf.WriteString(clock.Str.Get())

		conn.Write([]byte(buf.String()))
		conn.Close()
	}
}

func end() {
	err := os.Remove(socket_file)
	if err != nil {
		log.Fatal(err.Error())
	}
}
