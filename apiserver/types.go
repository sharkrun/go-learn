package apiserver

import (
	"ufleet/cluster/worker"
)

// ClusterResource recourse of cluster
type ClusterResource struct {
	Clusters map[string]*worker.Cluster
}

// Error error type for return
type Error struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}