package serverman

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	defaultServerMan = NewServerMan()
)

type ServerMan struct {
	servers []Server
}

type Server interface {
	Serve() error
	Stop() error
}

func NewServerMan() *ServerMan {
	return &ServerMan{
		servers: make([]Server, 0, 1),
	}
}

func (m *ServerMan) RegisterServer(server Server) {
	m.servers = append(m.servers, server)
}

func (m *ServerMan) Start() (err error) {
	wg := sync.WaitGroup{}
	done := make(chan error)
	errChan := make(chan error)

	go func() {
		if err = handlerSystemSignal(); err != nil {
			errChan <- err
		}
	}()

	wg.Add(len(m.servers))
	for _, svr := range m.servers {
		go func(s Server) {
			defer wg.Done()
			if err = s.Serve(); err != nil {
				errChan <- err
			}
		}(svr)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case err = <-errChan:
		fmt.Println(err)
	case <-done:
	}

	return
}

func (m *ServerMan) Stop() (err error) {
	for _, svr := range m.servers {
		if err = svr.Stop(); err != nil {
			return
		}
	}
	return
}

func RegisterServer(server Server) {
	defaultServerMan.RegisterServer(server)
}

func Start() error {
	return defaultServerMan.Start()
}

func Stop() error {
	return defaultServerMan.Stop()
}

func handlerSystemSignal() error {
	sChan := make(chan os.Signal)
	for {
		signal.Notify(sChan, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		sig := <-sChan
		switch sig {
		case os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			return Stop()
		}
	}
}
