package main

import (
	pb "github.com/Nivas-Mekala/go-grpc-example/proto"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}
