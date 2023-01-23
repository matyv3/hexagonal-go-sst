package domain

type TODO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// TODO status
const (
	StatusPending   = "pending"
	StatusDone      = "done"
	StatusCancelled = "cancelled"
)
