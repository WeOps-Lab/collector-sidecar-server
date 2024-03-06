package entity

type SidecarTokenEntity struct {
	Token string `json:"token"`
}

type SidecarTokenWrapperEntity struct {
	Id uint `json:"id"`
	SidecarTokenEntity
}

type SidecarTokenListEntity struct {
	PagerEntity
	Items []SidecarTokenWrapperEntity `json:"items"`
}
