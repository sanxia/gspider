# gspider
Html Page Download Analyzer Spider
==========================

import (
    "log"
    "github.com/sanxia/gspider"
)

/** ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 ** 初始化爬虫
 ** ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ **/
func NewSpider() gspider.ISpider {

    request := getHtmlRequest()

    customAnalyzer := analyzer.NewCustomAnalyzer(request)

    spider := gspider.NewSpider(10)

    spider.SetRequest(request)

    spider.RegisterAnalyzer("custom", customAnalyzer)

    page := new(gspider.Page)

    page.Name = "custom"

    page.Url = "http://www.baidu.com"

    spider.AddPage(page)

    go func() {

        select {

        case spiderResult := <-spider.GetResult():

            outputResult(spiderResult)

            log.Printf("spider result name: %v", spiderResult.Page.GetName())

        }

    }()

    return spider

}

/** ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 ** 自定义分析器
 ** ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ **/
type (

    CustomAnalyzer struct {

        *gspider.Analyzer

    }

)

/** ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 ** 初始化分析器
 ** ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ **/
func NewCustomAnalyzer(reqeust gspider.IHtmlRequest) gspider.IAnalyzer {

    return &CustomAnalyzer{

        &gspider.Analyzer{

            Request: reqeust,

        },

    }

}

/** ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 ** 解析页面
 ** ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ **/
func (s *CustomAnalyzer) Parse(page gspider.IPage) interface{} {

    result := page.GetContent(s.Request)

    return result

}

/** ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 ** 输出爬取结果
 ** ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ **/
func outputResult(result *gspider.SpiderResult) {

    results, ok := result.Data.(gspider.HtmlTagList)

    if !ok {

        return

    }


    for _, custom := range results {

        log.Printf("===== %s - %s =====\r\n", custom.Code, custom.Title)

    }

}

/** ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 ** 获取Http请求对象
 ** ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ **/
func getHtmlRequest() gspider.IHtmlRequest {

    headers := map[string]string{

        "Referer":    "http://www.baidu.com",

        "User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36",

    }

    cookies := map[string]string{

        "JSESSIONID": "35A39472B18BC213B6288F3B6BAA3ABC",

    }

    htmlRequest := gspider.NewHtmlRequest()

    htmlRequest.SetHeaders(headers)

    htmlRequest.SetCookies(cookies)

    return htmlRequest

}