package main

import (
	"Base/43_RPC02/services"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn,err := grpc.Dial(":8081",grpc.WithInsecure())
	if err!=nil{
		fmt.Println(err)
		log.Fatal(err)
	}
	defer conn.Close()

	prodClient := services.NewProdServiceClient(conn)
	prodRes,err:=prodClient.GetProdStock(context.Background(),
		&services.ProdRequest {ProdId:"12"},
	)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(prodRes.ProdStock)
}