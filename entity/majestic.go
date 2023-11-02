package entity

type Majestic struct {
	GlobalRank     int
	TldRank        int
	Domain         string
	TLD            string
	RefSubNets     int
	RefIPs         int
	IDNDomain      string
	IDNTLD         string
	PrevGlobalRank int
	PrevTldRank    int
	PrevRefSubNets int
	PrevRefIPs     int
}
