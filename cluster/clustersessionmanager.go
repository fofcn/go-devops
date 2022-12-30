package cluster

import (
	"errors"
	"io"
	"taskmanager/sshclient"

	"golang.org/x/crypto/ssh"
)

type ClusterSessionManager struct {
	/*
		集群与节点信息对应
		集群->节点map
		节点map: 节点Id->节点Session
	*/
	clusterNodeSessionTable map[string]map[string]ClusterNodeSession

	clusterManager *ClusterManager
}

type ClusterNodeSession struct {
	Client  *ssh.Client
	Session *ssh.Session
	NodeId  string
}

func (csm *ClusterSessionManager) Init(clusterManager *ClusterManager) error {
	csm.clusterNodeSessionTable = make(map[string]map[string]ClusterNodeSession)
	csm.clusterManager = clusterManager
	return nil
}

func (csm *ClusterSessionManager) Shutdown() error {
	for _, clusterNodeSessionTable := range csm.clusterNodeSessionTable {
		for _, nodeSession := range clusterNodeSessionTable {
			if nodeSession.Client != nil {
				sshclient.CloseClient(nodeSession.Client)
			}
		}
	}
	return nil
}

/*
创建集群下所有的node会话
*/
func (csm *ClusterSessionManager) CreateSession(cluster Cluster) error {
	if _, ok := csm.clusterNodeSessionTable[cluster.Name]; ok {
		return nil
	}

	// 创建map
	var nodeSessonMap map[string]ClusterNodeSession = make(map[string]ClusterNodeSession)
	csm.clusterNodeSessionTable[cluster.Name] = nodeSessonMap
	for _, node := range cluster.Node {
		session, err := csm.createSingleSession(node)
		if err != nil {
			continue
		}
		nodeSessonMap[node.Id] = *session
	}

	return nil
}

/*
获取集群下的节点会话
*/
func (csm *ClusterSessionManager) GetSession(clusterName string, nodeId string) (*ClusterNodeSession, error) {
	// 从集群会话中获取节点会话
	// 如果不存在则创建
	// 如果存在则测试连通性
	// 如果不连通重新创建会话
	// 如果连通则直接返回
	if _, ok := csm.clusterNodeSessionTable[clusterName][nodeId]; !ok {
		// create
		node, err := csm.clusterManager.GetNodeById(clusterName, nodeId)
		if err != nil {
			return nil, err
		}

		err = csm.CreateClusterSession(clusterName, *node)
		if err != nil {
			return nil, err
		}
	}

	if _, exists := csm.clusterNodeSessionTable[clusterName][nodeId]; !exists {
		return nil, errors.New("Node Session not exists, " + nodeId)
	}

	session := csm.clusterNodeSessionTable[clusterName][nodeId]
	newSession, err := session.Client.NewSession()
	if err != nil {
		return nil, err
	}
	session.Session = newSession
	return &session, nil
}

/*
创建集群下单个节点会话并添加到集群中
*/
func (csm *ClusterSessionManager) CreateClusterSession(clusterName string, node ClusterNode) error {
	var nodeSessonMap map[string]ClusterNodeSession = csm.clusterNodeSessionTable[clusterName]
	if nodeSessonMap == nil {
		// 创建map
		nodeSessonMap = make(map[string]ClusterNodeSession)
		csm.clusterNodeSessionTable[clusterName] = nodeSessonMap
	}

	session, err := csm.createSingleSession(node)
	if err != nil {
		return nil
	}
	nodeSessonMap[node.Name] = *session

	return nil
}

func (csm *ClusterSessionManager) RunCmd(cluster string, nodeId string, cmd string, stdout io.Writer, stderr io.Writer) error {
	session, err := csm.GetSession(cluster, nodeId)
	if err != nil {
		return err
	}

	session.Session.Stdout = stdout
	session.Session.Stderr = stderr
	err = session.Session.Run(cmd)
	defer session.Session.Close()
	return err
}

func (csm *ClusterSessionManager) createSingleSession(node ClusterNode) (*ClusterNodeSession, error) {
	var sshConnInfo sshclient.SSHConnectionInfo
	sshConnInfo.Address = node.Connection.Address
	sshConnInfo.Proto = node.Connection.Proto
	sshConnInfo.Username = node.Connection.Username
	sshConnInfo.Password = node.Connection.Password
	sshConnInfo.PublicKey = node.Connection.PublicKey
	sshSession, err := sshclient.CreateNewSesion(sshConnInfo)
	if err != nil {
		return nil, err
	}

	var clusterNodeSession ClusterNodeSession
	clusterNodeSession.Client = sshSession.Client
	clusterNodeSession.Session = sshSession.Session
	clusterNodeSession.NodeId = node.Id

	return &clusterNodeSession, nil
}
