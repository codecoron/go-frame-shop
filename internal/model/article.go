package model

// ArticleCreateUpdateBase 创建/修改内容基类
type ArticleCreateUpdateBase struct {
	UserId  int
	Title   string
	Desc    string
	PicUrl  string
	IsAdmin int
	Praise  int
	Detail  string
}

// ArticleCreateInput 创建内容
type ArticleCreateInput struct {
	ArticleCreateUpdateBase
}

// ArticleCreateOutput 创建内容返回结果
type ArticleCreateOutput struct {
	Id uint `json:"id"`
}
