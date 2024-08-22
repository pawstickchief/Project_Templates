package models

type HuaweiNeighbor struct {
	SelectNumber   int
	SInterface     string
	SwitchName     string
	DInterface     string
	SwitchPlatform string
	SwitchMac      string
	IsUpperDevice  bool
}

// HuaweiInterfaceBrief 定义结构体
type HuaweiInterfaceBrief struct {
	Interface string
	PHY       string
	Protocol  string
	InUti     string
	OutUti    string
	InErrors  int
	OutErrors int
}

// HuaweiPortVlan 定义结构体
type HuaweiPortVlan struct {
	Interface     string
	LinkType      string
	PVID          string
	TrunkVLANList string
}

// InterfaceVlanInfo 定义组合后的结构体
type InterfaceVlanInfo struct {
	Interface string
	PVID      string
	LinkType  string
	PHY       string
}
