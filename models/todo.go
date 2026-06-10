package models

type Todo struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

type UpdateTodoInput struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
	Done    *bool   `json:"done"`
}
