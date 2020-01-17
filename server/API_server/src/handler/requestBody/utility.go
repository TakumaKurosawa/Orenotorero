package requestBody

type EmailCheck struct {
	Email string `json:"email"`
}

type FileUpload struct {
	Img string `json:"img"`
}

type UpdatePosition struct {
	Position []int `json:"position"`
}
