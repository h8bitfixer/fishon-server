package getcdv3

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

// RegisterEtcd registers the service with etcd.
func RegisterEtcd(endpoints []string, serviceName string, serviceAddress string, ttl int64) error {
	// Create a new etcd client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return fmt.Errorf("failed to create etcd client: %v", err)
	}
	defer cli.Close()

	// Create a lease with a given TTL
	resp, err := cli.Grant(context.TODO(), ttl)
	if err != nil {
		return fmt.Errorf("failed to create lease: %v", err)
	}

	// Register the service with the lease ID
	key := fmt.Sprintf("/services/%s", serviceName)
	value := serviceAddress
	_, err = cli.Put(context.TODO(), key, value, clientv3.WithLease(resp.ID))
	if err != nil {
		return fmt.Errorf("failed to register service: %v", err)
	}

	// Keep the lease alive in the background
	keepAliveCh, err := cli.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		return fmt.Errorf("failed to start lease keep-alive: %v", err)
	}

	// Start a goroutine to monitor lease keep-alive updates
	go func() {
		for {
			select {
			case _, ok := <-keepAliveCh:
				if !ok {
					log.Printf("Lease keep-alive channel closed")
					return
				}
			}
		}
	}()

	log.Printf("Service '%s' registered with etcd", serviceName)
	return nil
}
