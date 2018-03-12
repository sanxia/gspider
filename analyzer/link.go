package analyzer

import (
	"strings"
)

import (
	"github.com/PuerkitoBio/goquery"
)

import (
	"github.com/sanxia/gspider"
)

/* ================================================================================
 * 超链接分析器
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

type (
	LinkAnalyzer struct {
		*gspider.Analyzer
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 初始化链接解析器
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewLinkAnalyzer(reqeust gspider.IHtmlRequest) gspider.IAnalyzer {
	return &LinkAnalyzer{
		&gspider.Analyzer{
			Request: reqeust,
		},
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 解析
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *LinkAnalyzer) Parse(page gspider.IPage) interface{} {
	result := page.GetContent(s.Request)

	htmlLinkList := make(gspider.HtmlLinkList, 0)

	htmlQuery := new(gspider.HtmlQuery)
	htmlQuery.FromReader(strings.NewReader(string(result)))
	htmlQuery.Document.Find("body a").Each(func(index int, sel *goquery.Selection) {
		title := sel.Text()
		url, _ := sel.Attr("href")

		if len(title) > 0 {
			htmlLink := &gspider.HtmlLink{
				Title: title,
				Url:   url,
			}
			htmlLinkList = append(htmlLinkList, htmlLink)
		}
	})

	return htmlLinkList
}
