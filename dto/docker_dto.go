package dto

const (
	CONTAINER_TYPE ElementType = "container"
	IMAGE_TYPE     ElementType = "image"
	VOLUME_TYPE    ElementType = "volume"
	NETWORK_TYPE   ElementType = "network"
)

type ElementType string

type Resume struct {
	Type      ElementType
	NbElement int
}

type ContainerResume struct {
	Resume
	NbActive int
}

type ImageResume struct {
	Resume
	TotalSize float64
}

type VolumeResume struct {
	Resume
}

type NetworkResume struct {
	Resume
	NbDriver int
	NbHost   int
	NbBridge int
}
