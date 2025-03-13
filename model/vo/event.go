package vo

type EventListRequest struct {
	Name    string `json:"name" form:"name"`
	Size    int    `form:"size" binding:"required"`
	Page    int    `form:"page" binding:"required"`
	Cluster string `json:"cluster" form:"cluster" binding:"required"`
}
