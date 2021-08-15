package main

import (
	"Base/43_RPC01/services"
	"net"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
    services.RegisterProdServiceServer(rpcServer,new(services.ProdService))

    lis,_ := net.Listen("tcp",":8081")

    rpcServer.Serve(lis)
}