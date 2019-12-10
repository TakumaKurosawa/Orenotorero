package requestBody

type BoardChangePublish struct {
	Id      int  `json:"id"`
	Publish bool `json:"publish"`
}

type BoardSendInviteMail struct {
	Id      int    `json:"id"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type BoardCreate struct {
	Title string `json:"title"`
	Img   string `json:"img"`
}
