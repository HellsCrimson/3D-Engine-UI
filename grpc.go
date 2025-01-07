package main

import (
	"context"
	"log"
	"time"
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
