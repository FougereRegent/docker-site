package dto

type ElementType string

const (
	CONTAINER_TYPE ElementType = "Container"
	IMAGE_TYPE     ElementType = "Image"
	VOLUME_TYPE    ElementType = "Volume"
	NETWORK_TYPE   ElementType = "Network"
)
