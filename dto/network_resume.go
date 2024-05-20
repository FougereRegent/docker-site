package dto

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
