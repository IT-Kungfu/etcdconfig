package main

import (
	"github.com/IT-Kungfu/etcdconfig"
	"github.com/IT-Kungfu/logger"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var (
	etcd = &etcdconfig.ETCDConfig{}
	cfg  = &Config{}
	log  *logger.Logger
)

func init() {
	var err error
	etcd, err = etcdconfig.GetConfig(cfg)
	if err != nil {
		panic(err)
	}

	if log, err = logger.New(&logger.Config{
		LogLevel:     cfg.ServerLogLevel,
		SentryDSN:    "",
		LogstashAddr: "",
		ServiceName:  "server",
		InstanceName: "dev",
	}); err != nil {
		panic(err)
	}
	etcd.AddObserver(log)

	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	srv := NewServer(etcd, log)
	etcd.AddObserver(srv)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	srv.Stop()
}
