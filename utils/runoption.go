package utils

import "fmt"

var EtcdEndPoint string

// InitRunOption init run config
func InitRunOption(etcd string) {
	EtcdEndPoint = etcd
}

func Test() {
	fmt.Println("asdfafd")
}
