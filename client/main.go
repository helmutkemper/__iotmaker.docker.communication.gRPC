package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	pb "github.com/helmutkemper/iotmaker.docker.communication.gRPC"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDockerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	pr, _ := nat.NewPort("tcp", "27017")
	pl := nat.PortMap{
		pr: []nat.PortBinding{
			{
				HostIP:   "",
				HostPort: "27017",
			},
		},
	}
	p := pb.ConvertGoNatPortMapToProtoTypeNatPortMapPortMap(pl)

	vl := []mount.Mount{
		{
			Source: "C:\\static\\db",
			Target: "/data/db",
			Type:   mount.Type("bind"),
		},
	}
	v := pb.ConvertGoMountMountToProtoMountVolume(vl)

	r, err := c.ContainerCreate(
		ctx,
		&pb.ContainerCreateRequest{
			ImageName:          "mongo:latest",
			ContainerName:      "delete",
			RestartPolicy:      pb.TypeRestartPolicy_KRestartPolicyUnlessStopped,
			PortExposedList:    p,
			MountVolumes:       v,
			ContainerNetworkId: "",
		},
	)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("ID: %s\nSuccess: %v\nError: %v\n", r.GetContainerID(), r.GetError())
}
