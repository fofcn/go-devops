package cluster

/*
Node 连接信息
*/
type ConnectionInfo struct {
	Address   string `yaml:"address"`
	Proto     string `yaml:"proto"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	PublicKey string `yaml:"publicKey"`
}

/*
Node 信息
*/
type ClusterNode struct {
	Name        string `yaml:"name"`
	Id          string
	Description string         `yaml:"description"`
	Connection  ConnectionInfo `yaml:"connection"`
}

/*
 集群信息
*/
type Cluster struct {
	Name string        `yaml:"name"`
	Node []ClusterNode `yaml:"node"`
}
