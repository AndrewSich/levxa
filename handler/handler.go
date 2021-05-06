package handler

import (
	"levxa/handler/api"
	"levxa/handler/blog"
	"levxa/handler/root"
)

var (
	HandlerMain = root.HandlerMain
	HandlerBlog = blog.HandlerBlog
	HandlerApi  = api.HandlerApi
)
