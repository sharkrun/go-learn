package worker

import "golang.org/x/crypto/ssh"

// Host basis host information
type Host struct {
	HostName       string `json:"hostname"`
	HostIP         string `json:"hostip"`
	UserName       string `json:"username"`
	UserPwd        string `json:"userpwd,omitempty"`
	Prikey         string `json:"prikey,omitempty"`
	PrikeyPwd      string `json:"prikeypwd,omitempty"`
	HostSSHPort    string `json:"hostsshport,omitempty"`
	HostSSHNetwork string `json:"hostsshnetwork,omitempty"`
}

// Master master of kubernetes
type Master struct {
	Host
	CreateTime     int64  `json:"createtime"`
	ClusterName    string `json:"clustername"`
	Registry       string `json:"registry,omitempty"`
	ServiceNetwork string `json:"servicenetwork"`
	PodNetwork     string `json:"podnetwork"`
	KubeServiceIP  string `json:"kubeserviceip"`
	DNSServiceIP   string `json:"dnsserviceip"`
	CaCerts        string `json:"cacerts"`
	APIServerCerts string `json:"apiservercerts"`
	APIServerKey   string `json:"apiserverkey"`
	JoinToken      string `json:"jointoken"`
	ErrorMsg       string `json:"errormsg"`
	Status         string `json:"status"`
	sshClient      *ssh.Client
}

// Node node of kubernetes
type Node struct {
	Host
	CreateTime     int64  `json:"createtime"`
	CaCerts        string `json:"cacerts"`
	ClusterName    string `json:"clustername"`
	Registry       string `json:"registry,omitempty"`
	MasterIP       string `json:"masterip"`
	APISerKubeCert string `json:"apiservercerts"`
	APISerKubeKey  string `json:"apiserverkey"`
	JoinToken      string `json:"jointoken"`
	Status         string `json:"status"`
	ErrorMsg       string `json:"errormsg"`
	sshClient      *ssh.Client
}

type Cluster struct {
	Name        string   `json:"name"`
	Status      string   `json:"status"`
	Masters     []Master `json:"masters"`
	Nodes       []Node   `json:"nodes"`
	ErrorMsg    string   `json:"errormsg"`
	CluInfo     `json:"info"`
}

// CluInfo cluster info
type CluInfo struct {
	CreateTime     int64  `json:"createtime"`
	BaseMasters    int    `json:"basemasters"`
	Vip            string `json:"vip"`
	JoinToken      string `json:"jointoken"`
	CaCerts        string `json:"cacerts"`
	CaKey          string `json:"cakey"`
	APIServerCerts string `json:"apiservercerts"`
	APIServerKey   string `json:"apiserverkey"`
	APISerKubeCert string `json:"apiserkubecert"`
	APISerKubeKey  string `json:"apiserkubekey"`
	FrontProxyCert string `json:"frontproxycacert"`
	FrontProxyKey  string `json:"frontproxycakey"`
	FrPxyCliCert   string `json:"frontproxyclientcert"`
	FrPxyCliKey    string `json:"frontproxyclientkey"`
	SaPub          string `json:"sapub"`
	SaKey          string `json:"sakey"`
	AdminContext   string `json:"admincontext"` // 经过base64编码
}

