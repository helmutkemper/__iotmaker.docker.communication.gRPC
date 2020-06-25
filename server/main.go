package main

import (
	"context"
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	pb "github.com/helmutkemper/iotmaker.docker.communication.gRPC"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedDockerServer
}

func (s *server) ContainerCreateAndChangeExposedPort(
	ctx context.Context,
	r *pb.ContainerCreateAndChangeExposedPortRequest,
) (replay *pb.ContainerCreateReply, err error) {

	var ID string
	var currentPortList, changePortList []nat.Port
	volumesList := pb.ConvertProtoMountVolumeToGoMountMount(r.MountVolumes)
	err, currentPortList = pb.ConvertArrProtoPortToGoArrNatPort(r.CurrentPort)
	err, changePortList = pb.ConvertArrProtoPortToGoArrNatPort(r.ChangeToPort)
	err, ID = dockerSys.ContainerCreateAndChangeExposedPort(
		r.ImageName,
		r.ContainerName,
		iotmakerDocker.RestartPolicy(r.RestartPolicy),
		volumesList,
		nil,
		currentPortList,
		changePortList,
	)

	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	return &pb.ContainerCreateReply{
			Error: &pb.TypeError{
				Success: err == nil,
				Message: errorMessage,
				Trace:   nil,
			},
			ContainerID: ID,
		},
		nil
}

func (s *server) ContainerCreateChangeExposedPortAndStart(
	ctx context.Context,
	r *pb.ContainerCreateAndChangeExposedPortRequest,
) (replay *pb.ContainerCreateReply, err error) {

	var ID string
	var currentPortList, changePortList []nat.Port
	volumesList := pb.ConvertProtoMountVolumeToGoMountMount(r.MountVolumes)
	err, currentPortList = pb.ConvertArrProtoPortToGoArrNatPort(r.CurrentPort)
	err, changePortList = pb.ConvertArrProtoPortToGoArrNatPort(r.ChangeToPort)
	err, ID = dockerSys.ContainerCreateChangeExposedPortAndStart(
		r.ImageName,
		r.ContainerName,
		iotmakerDocker.RestartPolicy(r.RestartPolicy),
		volumesList,
		nil,
		currentPortList,
		changePortList,
	)

	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	return &pb.ContainerCreateReply{
			Error: &pb.TypeError{
				Success: err == nil,
				Message: errorMessage,
				Trace:   nil,
			},
			ContainerID: ID,
		},
		nil
}

func (s *server) ContainerCreate(
	ctx context.Context,
	r *pb.ContainerCreateRequest,
) (replay *pb.ContainerCreateReply, err error) {

	var exposedPortList nat.PortMap
	var ID string

	err, exposedPortList = pb.ConvertProtoTypeNatPortMapPortMapToGoNatPortMap(r.PortExposedList)
	volumesList := pb.ConvertProtoMountVolumeToGoMountMount(r.MountVolumes)
	err, ID = dockerSys.ContainerCreate(
		r.ImageName,
		r.ContainerName,
		iotmakerDocker.RestartPolicy(r.RestartPolicy),
		exposedPortList,
		volumesList,
		nil,
	)

	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	return &pb.ContainerCreateReply{
			Error: &pb.TypeError{
				Success: err == nil,
				Message: errorMessage,
				Trace:   nil,
			},
			ContainerID: ID,
		},
		nil
}

func (s *server) ContainerCreateAndStart(
	ctx context.Context,
	r *pb.ContainerCreateRequest,
) (replay *pb.ContainerCreateReply, err error) {

	var exposedPortList nat.PortMap
	var ID string

	err, exposedPortList = pb.ConvertProtoTypeNatPortMapPortMapToGoNatPortMap(r.PortExposedList)
	volumesList := pb.ConvertProtoMountVolumeToGoMountMount(r.MountVolumes)
	err, ID = dockerSys.ContainerCreateAndStart(
		r.ImageName,
		r.ContainerName,
		iotmakerDocker.RestartPolicy(r.RestartPolicy),
		exposedPortList,
		volumesList,
		nil,
	)

	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	return &pb.ContainerCreateReply{
			Error: &pb.TypeError{
				Success: err == nil,
				Message: errorMessage,
				Trace:   nil,
			},
			ContainerID: ID,
		},
		nil
}

func (s *server) ContainerCreateAndExposePortsAutomatically(
	ctx context.Context,
	r *pb.ContainerCreateAndExposePortsAutomaticallyRequest,
) (replay *pb.ContainerCreateReply, err error) {

	var ID string
	volumesList := pb.ConvertProtoMountVolumeToGoMountMount(r.MountVolumes)
	err, ID = dockerSys.ContainerCreateAndExposePortsAutomatically(
		r.ImageName,
		r.ContainerName,
		iotmakerDocker.RestartPolicy(r.RestartPolicy),
		volumesList,
		nil,
	)

	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	return &pb.ContainerCreateReply{
			Error: &pb.TypeError{
				Success: err == nil,
				Message: errorMessage,
				Trace:   nil,
			},
			ContainerID: ID,
		},
		nil
}

func (s *server) ContainerCreateExposePortsAutomaticallyAndStart(
	ctx context.Context,
	r *pb.ContainerCreateAndExposePortsAutomaticallyRequest,
) (replay *pb.ContainerCreateReply, err error) {

	var ID string
	volumesList := pb.ConvertProtoMountVolumeToGoMountMount(r.MountVolumes)
	err, ID = dockerSys.ContainerCreateExposePortsAutomaticallyAndStart(
		r.ImageName,
		r.ContainerName,
		iotmakerDocker.RestartPolicy(r.RestartPolicy),
		volumesList,
		nil,
	)

	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	return &pb.ContainerCreateReply{
			Error: &pb.TypeError{
				Success: err == nil,
				Message: errorMessage,
				Trace:   nil,
			},
			ContainerID: ID,
		},
		nil
}

func (s *server) ContainerCreateWithoutExposePorts(
	ctx context.Context,
	r *pb.ContainerCreateWithoutExposePortsRequest,
) (replay *pb.ContainerCreateReply, err error) {

	var ID string
	volumesList := pb.ConvertProtoMountVolumeToGoMountMount(r.MountVolumes)
	err, ID = dockerSys.ContainerCreateWithoutExposePorts(
		r.ImageName,
		r.ContainerName,
		iotmakerDocker.RestartPolicy(r.RestartPolicy),
		volumesList,
		nil,
	)

	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	return &pb.ContainerCreateReply{
			Error: &pb.TypeError{
				Success: err == nil,
				Message: errorMessage,
				Trace:   nil,
			},
			ContainerID: ID,
		},
		nil
}

func (s *server) ContainerCreateWithoutExposePortsAndStart(
	ctx context.Context,
	r *pb.ContainerCreateWithoutExposePortsRequest,
) (replay *pb.ContainerCreateReply, err error) {

	var ID string
	volumesList := pb.ConvertProtoMountVolumeToGoMountMount(r.MountVolumes)
	err, ID = dockerSys.ContainerCreateWithoutExposePortsAndStart(
		r.ImageName,
		r.ContainerName,
		iotmakerDocker.RestartPolicy(r.RestartPolicy),
		volumesList,
		nil,
	)

	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	return &pb.ContainerCreateReply{
			Error: &pb.TypeError{
				Success: err == nil,
				Message: errorMessage,
				Trace:   nil,
			},
			ContainerID: ID,
		},
		nil
}

func (s *server) ContainerFindIdByName(
	ctx context.Context,
	r *pb.RequestByNameRequest,
) (replay *pb.RequestByNameReply, err error) {

	var ID string
	err, ID = dockerSys.ContainerFindIdByName(r.Name)

	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	return &pb.RequestByNameReply{
			Error: &pb.TypeError{
				Success: err == nil,
				Message: errorMessage,
				Trace:   nil,
			},
			Id: ID,
		},
		nil
}

func (s server) ContainerFindIdByNameContains(
	ctx context.Context,
	r *pb.RequestByNameRequest,
) (replay *pb.RequestByNameReply, err error) {

	var ID string
	err, ID = dockerSys.ContainerFindIdByNameContains(r.Name)

	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	return &pb.RequestByNameReply{
			Error: &pb.TypeError{
				Success: err == nil,
				Message: errorMessage,
				Trace:   nil,
			},
			Id: ID,
		},
		nil
}

var (
	dockerSys iotmakerDocker.DockerSystem
)

func main() {
	var err error
	err = dockerSys.Init()
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDockerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
