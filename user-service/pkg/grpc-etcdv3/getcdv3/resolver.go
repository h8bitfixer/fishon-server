package getcdv3

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

// ResolveEtcd resolves the endpoints of a service from etcd.
func ResolveEtcd(endpoints []string, serviceName string) ([]string, error) {
	// Create a new etcd client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create etcd client: %v", err)
	}
	defer cli.Close()

	// Get the list of registered endpoints for the service
	key := fmt.Sprintf("/services/%s", serviceName)
	resp, err := cli.Get(context.TODO(), key, clientv3.WithPrefix())
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoints for service: %v", err)
	}

	// Extract the endpoint values from the response
	endpoints = make([]string, 0, len(resp.Kvs))
	for _, kv := range resp.Kvs {
		endpoints = append(endpoints, string(kv.Value))
	}

	return endpoints, nil
}
