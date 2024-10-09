package main

import (
	pb "github.com/Nivas-Mekala/go-grpc-example/coffeeshop_protos"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}
