package server

const (
	ListenAddress = ":8080"
)

type Data struct {
	GitSource   string `form:"gitsource" json:"gitsource" binding:"required"`
	GitPath     string `form:"gitpath" json:"gitpath" binding:"required"`
	CallBack    string `form:"callback" json:"callback" binding:"required"`
	ClusterName string `form:"clustername" json:"clustername"`
}
