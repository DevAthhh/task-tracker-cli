package task

type Task struct {
	ID        int64  `json:"id"`
	Desc      string `json:"desc"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
