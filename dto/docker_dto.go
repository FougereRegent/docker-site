package dto

const (
	CONTAINER_TYPE ElementType = "Container"
	IMAGE_TYPE     ElementType = "Image"
	VOLUME_TYPE    ElementType = "Volume"
	NETWORK_TYPE   ElementType = "Network"
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
	NbDriver  int
	NbHost    int
	NbBridge  int
	NbNone    int
	NbOverlay int
	NbIpVlan  int
	NbMacVlan int
}
