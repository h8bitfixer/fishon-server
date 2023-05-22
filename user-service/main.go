package main

import (
	"github.com/gin-gonic/gin"
	routesV1 "user-service/api/v1"
)

func main() {

	//config := clientv3.Config{
	//	Endpoints: []string{"http://localhost:2379"}, // Replace with your etcd server addresses
	//}
	//client, err := clientv3.New(config)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer client.Close()
	//
	//key := fmt.Sprintf("%s:///%s/", "fishOn", "UserAuth")
	////"fishOn:///UserAuth/"
	////_, err = client.Put(context.Background(), key, "0.0.0.0:10021")
	////if err != nil {
	////	log.Fatal(err)
	////}
	////
	////log.Println("Key-value pair stored successfully.")
	//
	//resp, err := client.Get(context.Background(), key)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if len(resp.Kvs) > 0 {
	//	value := resp.Kvs[0].Value
	//	log.Printf("Value: %s\n", value)
	//} else {
	//	log.Println("Key not found")
	//}

	r := gin.Default()
	routesV1.SetupRoutes(r)
	// Start the server
	err := r.Run(":10011")
	if err != nil {

	}
}
