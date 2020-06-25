package iotmaker_docker_communication_grpc

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
)

func volumeProtoTypeToMountType(t TypeVolumeMountType) mount.Type {
	switch t {
	case TypeVolumeMountType_KVolumeMountTypeBind:
		return "bind"

	case TypeVolumeMountType_KVolumeMountTypeVolume:
		return "volume"

	case TypeVolumeMountType_KVolumeMountTypeTmpfs:
		return "tmpfs"

	case TypeVolumeMountType_KVolumeMountTypeNpipe:
		return "npipe"
	}

	return "bind"
}

func volumeMountTypeToProtoType(t mount.Type) TypeVolumeMountType {
	switch t {
	case "bind":
		return TypeVolumeMountType_KVolumeMountTypeBind

	case "volume":
		return TypeVolumeMountType_KVolumeMountTypeVolume

	case "tmpfs":
		return TypeVolumeMountType_KVolumeMountTypeTmpfs

	case "npipe":
		return TypeVolumeMountType_KVolumeMountTypeNpipe
	}

	return TypeVolumeMountType_KVolumeMountTypeBind
}

func ConvertProtoMountVolumeToGoMountMount(
	volumes []*TypeMountMountVolume,
) (mountList []mount.Mount) {

	mountList = make([]mount.Mount, 0)

	for _, volume := range volumes {
		mountList = append(
			mountList,
			mount.Mount{
				Type:   volumeProtoTypeToMountType(volume.Type),
				Source: volume.Source,
				Target: volume.Target,
			},
		)
	}

	return
}

func ConvertGoMountMountToProtoMountVolume(
	mountList []mount.Mount,
) (volumes []*TypeMountMountVolume) {

	volumes = make([]*TypeMountMountVolume, 0)

	for _, volume := range mountList {
		volumes = append(
			volumes,
			&TypeMountMountVolume{
				Type:   volumeMountTypeToProtoType(volume.Type),
				Source: volume.Source,
				Target: volume.Target,
			},
		)
	}

	return
}

func GoArrNatPortToConvertArrProtoPort(
	portList []nat.Port,
) (list []*TypePort) {

	list = make([]*TypePort, 0)

	for _, p := range portList {
		list = append(
			list,
			&TypePort{
				Port:  p.Port(),
				Proto: p.Proto(),
			},
		)
	}

	return
}

func ConvertArrProtoPortToGoArrNatPort(
	portList []*TypePort,
) (err error, list []nat.Port) {

	list = make([]nat.Port, 0)

	for _, p := range portList {
		var err error
		var port nat.Port

		port, err = nat.NewPort(p.Proto, p.Port)
		if err != nil {
			return
		}

		list = append(list, port)
	}

	return
}
