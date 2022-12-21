package cluster

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"taskmanager/zlog"

	"gopkg.in/yaml.v2"
)

type clusterConfig struct {
}

func (clusterConfig) Load(configDir string) (interface{}, error) {
	// list config direcotry file
	files, err := listDirectory(configDir)
	if err != nil {
		return nil, err
	}

	var clusterList []Cluster
	for _, file := range files {
		buf, err := ioutil.ReadFile(file)
		if err != nil {
			zlog.Logger.Fatalf("Read env setup yaml file error, %v", err)
			return nil, err
		}

		var clusterInfo Cluster
		err = yaml.Unmarshal([]byte(buf), &clusterInfo)
		if err != nil {
			zlog.Logger.Fatalf("Unmarshal yaml file error, %v", err)
			return nil, err
		}

		clusterList = append(clusterList, clusterInfo)
	}

	return clusterList, nil
}

func listDirectory(configDir string) ([]string, error) {
	var files []string
	err := filepath.Walk(configDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
