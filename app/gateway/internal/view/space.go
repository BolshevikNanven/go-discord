package view

type Space struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Owner  int64  `json:"owner"`
}

type SpaceRequest struct {
	Name   string `json:"name" form:"name"`
	Avatar string `json:"avatar" form:"avatar"`
}
