package bo

type Common struct {
	Namespace string `json:"namespace" form:"namespace" binding:"required"`
	Cluster   string `json:"cluster" form:"cluster" binding:"required"`
}
