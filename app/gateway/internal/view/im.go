package view

type Message struct {
	Id        int64  `json:"id"`
	SpaceId   int64  `json:"space_id"`
	From      int64  `json:"from"`
	To        int64  `json:"to"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}

type SendMessageRequest struct {
	SpaceId int64  `json:"space_id"`
	To      int64  `json:"to"`
	Type    string `json:"type"`
	Content string `json:"content"`
}
type SendMessageResponse struct {
	MessageId int64 `json:"message_id"`
}

type AckMessageRequest struct {
	SpaceId    int64   `json:"space_id"`
	MessageIds []int64 `json:"message_ids"`
}

type AckMessageResponse struct {
	Success bool `json:"success"`
}

type PullHistoryRequest struct {
	SpaceId   int64 `json:"space_id"`
	ChannelId int64 `json:"channel_id"`
	From      int64 `json:"from"`
	Cursor    int64 `json:"cursor"`
	Limit     int32 `json:"limit"`
}

type PullHistoryResponse struct {
	Messages []Message `json:"messages"`
	Cursor   int64     `json:"cursor"`
	Limit    int32     `json:"limit"`
}

type GetInboxRequest struct {
	SpaceId int64 `json:"space_id"`
	Limit   int32 `json:"limit"`
}

type GetInboxResponse struct {
	Messages []Message `json:"messages"`
}

type AckChannelMessageRequest struct {
	MessageId int64 `json:"message_id"`
}

type AckChannelMessageResponse struct {
	Success bool `json:"success"`
}

type GetChannelInboxResponse struct {
	Current int64 `json:"current"`
	Last    int64 `json:"last"`
}
