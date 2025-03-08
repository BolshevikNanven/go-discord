package view

type Channel struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Owner int64  `json:"owner"`
}

type CreateChannelRequest struct {
	SpaceID int64  `json:"space_id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

type UpdateChannelRequest struct {
	Name string `json:"name"`
}
