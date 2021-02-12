package main

import (
	"github.com/IT-Kungfu/etcdconfig"
	"github.com/IT-Kungfu/logger"
	"strings"
	"time"
)

const (
	RestartTimeout = 3
)

type Server struct {
	etcd      *etcdconfig.ETCDConfig
	log       *logger.Logger
	restartCh chan struct{}
}

func NewServer(etcd *etcdconfig.ETCDConfig, log *logger.Logger) *Server {
	s := &Server{
		etcd:      etcd,
		log:       log,
		restartCh: make(chan struct{}, 1),
	}

	s.start()
	s.restartTimer()

	return s
}

func (s *Server) start() {

}

func (s *Server) Stop() {

}

func (s *Server) restartTimer() {
	go func() {
		var isChange bool
		for {
			select {
			case <-s.restartCh:
				isChange = true
			case <-time.After(RestartTimeout * time.Second):
				if isChange {
					isChange = false
					s.Stop()
					s.start()
				}
			}
		}
	}()
}

func (s *Server) ETCDValueChanged(key string, value []byte, cfg interface{}) {
	if strings.HasPrefix(key, "/configs/server/") {
		s.log.Infof("server config changed: %s %s", key, value)
		s.restartCh <- struct{}{}
	}
}
