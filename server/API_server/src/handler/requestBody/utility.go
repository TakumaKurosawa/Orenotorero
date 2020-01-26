package requestBody

type EmailCheck struct {
	Email string `json:"email"`
}

type FileUpload struct {
	Img string `json:"img"`
}

type List struct {
	Position []UpdatePosition `binding:"required"`
}
type UpdatePosition struct {
	kanbanId  int   `json:"kanbanId binding:"required"`
	cardArray []int `json:"cardArray" binding:"required"`
}
