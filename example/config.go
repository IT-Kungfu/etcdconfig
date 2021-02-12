package main

type Config struct {
	ServerHost string `etcd:"/configs/server/{{SERVER_INSTANCE}}/host,watcher" default:"0.0.0.0"`
	ServerPort int    `etcd:"/configs/server/{{SERVER_INSTANCE}}/port,watcher" default:"8000"`

	ServerLogLevel string `etcd:"/configs/server/{{SERVER_INSTANCE}}/log_level,watcher" default:"debug"`
}
