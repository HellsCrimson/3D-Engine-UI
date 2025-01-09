package main

import (
	"context"
	"log"
	"time"

	pb "3d-engine-ui/grpc"
)

func getObjects() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objs, err := client.GetObjects(ctx, nil)
	if err != nil {
		log.Fatalf("client.GetFeature failed: %v", err)
	}

	objects = objs.Objects
}

func updateObject(obj *pb.Object) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := client.UpdateObject(ctx, obj)
	if err != nil {
		log.Fatalf("client.UpdateObject failed: %v", err)
	}
}
