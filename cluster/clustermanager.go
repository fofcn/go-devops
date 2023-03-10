package cluster

import (
	"errors"
	"taskmanager/args"
	"taskmanager/config"

	"github.com/google/uuid"
)

type ClusterManager struct {
	clusterTable map[string]Cluster
}

func (ClusterManager *ClusterManager) Init(obj interface{}) error {
	configfile := obj.(args.ApplicationArgs)

	ClusterManager.clusterTable = make(map[string]Cluster)
	// load cluster
	clusterList, err := config.ConfigManager(clusterConfig{}).Load(configfile.GetCluster())
	if err != nil {
		return err
	}

	// convert list to map
	for _, cluster := range clusterList.([]Cluster) {
		for _, node := range cluster.Node {
			uuid := uuid.New()
			node.Id = uuid.String()
		}
		ClusterManager.clusterTable[cluster.Name] = cluster
	}

	return nil
}

func (ClusterManager *ClusterManager) Start() error {
	return nil
}

func (ClusterManager *ClusterManager) Shutdown() error {
	return nil
}

func (ClusterManager *ClusterManager) GetNodeById(clusterName string, nodeId string) (*ClusterNode, error) {
	if _, node := ClusterManager.clusterTable[clusterName]; node {
		c := ClusterManager.clusterTable[clusterName]
		for _, exists := range c.Node {
			if exists.Name == nodeId {
				return &exists, nil
			}

		}
	}
	return nil, errors.New("Node " + nodeId + " is not existing.")
}

func (ClusterManager *ClusterManager) GetClusterNodeList(clusterName string) ([]ClusterNode, error) {
	if _, node := ClusterManager.clusterTable[clusterName]; node {
		return ClusterManager.clusterTable[clusterName].Node, nil
	}

	return nil, errors.New("Cluster " + clusterName + " does not exists.")
}
