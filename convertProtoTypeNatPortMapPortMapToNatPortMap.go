package iotmaker_docker_communication_grpc

import (
	"github.com/docker/go-connections/nat"
)

func ConvertGoNatPortMapToProtoTypeNatPortMapPortMap(
	mapPort nat.PortMap,
) (protoPortMap []*TypeNatPortMap) {

	protoPortMap = make([]*TypeNatPortMap, 0)

	for keyPort, valuePort := range mapPort {
		var listOfPortBinds = make([]*TypePortBinding, 0)

		for _, portBind := range valuePort {
			listOfPortBinds = append(
				listOfPortBinds,
				&TypePortBinding{
					HostIP:   portBind.HostIP,
					HostPort: portBind.HostPort,
				},
			)
		}

		protoPortMap = append(
			protoPortMap,
			&TypeNatPortMap{
				Port: &TypePort{
					Port:  keyPort.Port(),
					Proto: keyPort.Proto(),
				},
				PortBind: listOfPortBinds,
			},
		)

	}

	return
}

func ConvertProtoTypeNatPortMapPortMapToGoNatPortMap(
	portMap []*TypeNatPortMap,
) (err error, port nat.PortMap) {

	port = make(map[nat.Port][]nat.PortBinding)

	for _, v := range portMap {
		var portKey nat.Port
		portKey, err = nat.NewPort(v.Port.Proto, v.Port.Port)
		if err != nil {
			return
		}
		port[portKey] = make([]nat.PortBinding, 0)

		for _, bind := range v.PortBind {
			port[portKey] = append(
				port[portKey],
				nat.PortBinding{
					HostIP:   bind.HostIP,
					HostPort: bind.HostPort,
				},
			)
		}
	}

	return
}
