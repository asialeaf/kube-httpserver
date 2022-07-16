package server

type Data struct {
	GitSource   string `form:"gitsource" json:"gitsource" binding:"required"`
	CallBack    string `form:"callback" json:"callback" binding:"required"`
	ClusterName string `form:"clustername" json:"clustername"`
}
