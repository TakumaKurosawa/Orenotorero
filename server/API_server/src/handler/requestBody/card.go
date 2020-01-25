package requestBody

type CardChangeTitle struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type CardChangeDeadline struct {
	Id       int    `json:"id"`
	Deadline string `json:"deadline"`
}

type CardCreate struct {
	Title    string `json:"title"`
	KanbanId int    `json:"kanban_id"`
	Position int    `json:"position"`
}

type FileAdd struct {
	Id       int    `json:"id"`
	FileName string `json:"name"`
	FileData string `json:"data"`
}

type FileDelete struct {
	Id int `json:"id"`
}

type CardDelete struct {
	Id int `json:"id"`
}
