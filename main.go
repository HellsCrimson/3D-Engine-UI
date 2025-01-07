package main

import (
	"log"

	pb "3d-engine-ui/grpc"

	g "github.com/AllenDang/giu"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client pb.EngineClient

	objects []*pb.Object
)

func loop() {

	g.SingleWindow().Layout(
		g.Label("Testing"),
		g.Column(
			g.Button("Get Objects").OnClick(getObjects),
			g.Custom(func() {
				for _, object := range objects {
					g.BulletText(object.Name).Build()
				}
			}),
		),
	)
}

func main() {
	opts := []grpc.DialOption{}
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client = pb.NewEngineClient(conn)

	wnd := g.NewMasterWindow("3D Engine GUI", 400, 200, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
