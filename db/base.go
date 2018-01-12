package db

import (
	"fmt"
	"strings"
	"log"
	"ufleet/cluster/etcd"
	"ufleet/cluster/config"
)

var dEtcd *etcd.LEtcd
var dEndpoints []string


func InitDB(specailEndpoint ...string) error {
	dEtcd = new(etcd.LEtcd)
	if dEtcd == nil {
		log.Printf("can't new etcd, maybe not enough memory in host machine")
		return fmt.Errorf("can't new etcd, maybe not enough memory in host machine")
	}

	var endpoints []string
	if len(specailEndpoint) > 0 {
		endpoints = specailEndpoint
		dEndpoints = specailEndpoint
	} else {
		if dEndpoints != nil {
			endpoints = dEndpoints
		} else {
			endpoints = []string{fmt.Sprintf("http://%s:%d", config.ETCDHOST, config.ETCDPORT)}
			dEndpoints = endpoints
		}
	}
	err := dEtcd.InitByEndpoints(endpoints)
	if err != nil {
		log.Printf("can't init etcd by etcd endpoint(%v): %s", endpoints, err.Error())
		return fmt.Errorf("can't init etcd by etcd endpoint(%v): %s", endpoints, err.Error())
	}

	return nil
}


// TODO 考虑在Ha模式下，可能会出现Ufleet主机异常，切换主机的情况，故可能需要重连Etcd节点
// 重新封装　Get当出现某些错误时，重新连接

func Get(key string) (map[string]string, error){
	if dEtcd == nil {
		if err := InitDB(); err != nil {
			return nil, err
		}
	}

	result, err := dEtcd.Get(key)
	if err != nil && strings.Contains(err.Error(), "reset") {
		// 服务端reset，意味之前的连接已无效，需要重新创建连接
		ierr := InitDB()
		if ierr != nil {
			return nil, fmt.Errorf("because of %s, connect etcd server again but failed: %s", err.Error(), ierr.Error())
		}

		result, err = dEtcd.Get(key)
	}

	return result, nil
}