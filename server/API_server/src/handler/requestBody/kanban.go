package requestBody

type KanbanCreate struct {
	BoardId  int    `json:"board_id"`
	Title    string `json:"title"`
	Position int    `json:"position"`
}

type KanbanDelete struct {
	KanbanId int `json:"id"`
}

type KanbanChangeTitle struct {
	KanbanDelete
	Title string `json:"title"`
}

type KanbanChangePosition struct {
	KanbanDelete
	CurrentPosition     int `json:"current_position"`
	DestinationPosition int `json:"destination_position"`
}
