package gspider

/* ================================================================================
 * Html标记
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

type (
	HtmlTagList []*HtmlTag
	HtmlTag     struct {
		Code        string      `json:"code"`
		Title       string      `json:"title"`
		Description string      `json:"description"`
		ImageUrl    string      `json:"image_url"`
		LinkUrl     string      `json:"link_url"`
		Childs      HtmlTagList `json:"children"`
	}
)
