package gspider

import (
	"github.com/sanxia/glib"
)

/* ================================================================================
 * 爬虫处理器
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

type (
	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * 爬虫处理器接口
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	ISpiderHandler interface {
		GetSpider() ISpider
		Handler(result *SpiderResult)
	}

	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * 爬虫处理器数据结构
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	SpiderHandler struct {
		Request glib.IHttpRequest
	}
)
