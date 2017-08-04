package digitalocean

import (
	"fmt"
	"github.com/kris-nova/charlie/network"
	"github.com/kris-nova/kubicorn/apis/cluster"
	"github.com/kris-nova/kubicorn/logger"
	"github.com/kris-nova/kubicorn/profiles"
	"github.com/kris-nova/kubicorn/test"
	"os"
	"testing"
	"time"
)

var testCluster *cluster.Cluster

func TestMain(m *testing.M) {
	logger.TestMode = true
	logger.Level = 4
	var err error
	//func() {
	//	for {
	//		if r := recover(); r != nil {
	//			logger.Critical("Panic: %v", r)
	//		}
	//	}
	//}()
	test.InitRsaTravis()
	testCluster = profiles.NewSimpleDigitalOceanCluster("ubuntu-test")
	testCluster, err = test.Create(testCluster)
	if err != nil {
		fmt.Printf("Unable to create digital ocean test cluster: %v\n", err)
		os.Exit(1)
	}
	status := m.Run()
	exitCode := 0
	if status != 0 {
		fmt.Printf("-----------------------------------------------------------------------\n")
		fmt.Printf("[FAILURE]\n")
		fmt.Printf("-----------------------------------------------------------------------\n")
		exitCode = 1
	}
	_, err = test.Delete(testCluster)
	if err != nil {
		exitCode = 99
		fmt.Println("Failure cleaning up cluster! Abandoned resources!")
	}
	os.Exit(exitCode)
}

const (
	ApiSleepSeconds   = 5
	ApiSocketAttempts = 40
)

func TestApiListen(t *testing.T) {
	success := false
	for i := 0; i < ApiSocketAttempts; i++ {
		_, err := network.AssertTcpSocketAcceptsConnection(fmt.Sprintf("%s:%s", testCluster.KubernetesApi.Endpoint, testCluster.KubernetesApi.Port), "opening a new socket connection against the Kubernetes API")
		if err != nil {
			logger.Info("Attempting to open a socket to the Kubernetes API: %v...\n", err)
			time.Sleep(time.Duration(ApiSleepSeconds) * time.Second)
			continue
		}
		success = true
	}
	if !success {
		t.Fatalf("Unable to connect to Kubernetes API")
	}
}
