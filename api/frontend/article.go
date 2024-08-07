package frontend

import "github.com/gogf/gf/v2/frame/g"

type ArticleAddReq struct {
	g.Meta `path:"/article/add" tags:"前端文章" method:"post" summary:"添加文章"`
	ArticleCommonAddUpdate
}

type ArticleCommonAddUpdate struct {
	Title  string `json:"title"             description:"文章标题" v:"required#名称必传"`
	Desc   string `json:"desc" dc:"文章概要"`
	PicUrl string `json:"pic_url"           description:"图片"`
	Detail string `json:"detail"            description:"文章详情" v:"required#文章详情必填"`
}

type ArticleAddRes struct {
	Id uint `json:"id"`
}
