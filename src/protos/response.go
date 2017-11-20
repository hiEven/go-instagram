package protos

type defaultResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type TimelineFeedResponse struct {
	defaultResponse
	MoreAvailable bool            `json:"more_available"`
	NextMaxID     string          `json:"next_max_id"`
	Items         []*TimelineItem `json:"feed_items"`
}

type LoginResponse struct {
	defaultResponse
	LoggedInUser *LoggedInUser `json:"logged_in_user"`
}

type InboxFeedResponse struct {
	defaultResponse
	Inbox                *Inbox `json:"inbox"`
	PendingRequestsTotal int    `json:"pending_requests_total"`
	SeqID                int    `json:"seq_id"`
	// PendingRequestsUsers []string `json:"pending_requests_users"`
}

type ThreadBroadcastTextResponse struct {
	defaultResponse
	Threads []*ThreadItem `json:"threads"`
}

type ThreadBroadcastLinkResponse struct {
	defaultResponse
	Threads []*ThreadItem `json:"threads"`
}

type ThreadShowResponse struct {
	defaultResponse
	Thread *ThreadItem `json:"thread"`
}

type ThreadApproveAllResponse struct {
	defaultResponse
}

type MediaLikeResponse struct {
	defaultResponse
}

type MediaUnlikeResponse struct {
	defaultResponse
}
