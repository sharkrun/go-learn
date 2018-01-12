package config

var (
	SERVER_PORT int =9090
	ETCDHOST string = "192.168.4.111"
	ETCDPORT int = 32379
	LOGDIR="/home/lear/gocode/src/ufleet/cluster/log"
	GLogLevel=3 // 3 --info

	// GEtcdListenPort etcd listen port
	GEtcdListenPort = 12379
)