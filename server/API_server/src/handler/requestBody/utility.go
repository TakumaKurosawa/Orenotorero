package requestBody

type EmailCheck struct {
	Email string `json:"email"`
}

type FileUpload struct {
	Img string `json:"img"`
}

type UpdatePosition struct {
	KanbanId  int   `json:"kanbanId"`
	CardArray []int `json:"cardArray"`
}
