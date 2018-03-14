package gspider

import (
	"errors"
	"io"
	"net/http"
)

import (
	"github.com/mozillazg/request"
	"github.com/sanxia/glib"
)

/* ================================================================================
 * 网络请求
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

type (
	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * Html请求接口
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	IHtmlRequest interface {
		Get(url string) (IHtmlResponse, error)
		Post(url string) (IHtmlResponse, error)

		SetUserAgent(userAgent string)
		SetHeaders(headers map[string]string)
		SetParams(params map[string]string)
		SetCookies(cookies map[string]string)
		SetJson(json map[string]string)
		SetData(data map[string]string)
		SetFiles(files FormFiles)
	}

	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * Html响应接口
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	IHtmlResponse interface {
		GetData() []byte
		GetHeader() http.Header
		GetStatusCode() int
		GetStatus() string
	}
)

type (
	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * Html表单文件数据结构
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	FormFiles []FormFile
	FormFile  struct {
		FieldName string
		FileName  string
		Datas     io.Reader
	}

	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * Html请求数据结构
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	HtmlRequest struct {
		userAgents []string
		userAgent  string
		headers    map[string]string
		params     map[string]string
		cookies    map[string]string
		json       map[string]string
		data       map[string]string
		files      FormFiles
		request    *request.Request
	}

	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * Html响应数据结构
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	HtmlResponse struct {
		Header     http.Header
		Data       []byte
		StatusCode int
		Status     string
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 初始化Html请求
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewHtmlRequest() IHtmlRequest {
	htmlRequest := &HtmlRequest{}
	htmlRequest.userAgents = []string{
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36",
		"Mozilla/5.0 (iPod; U; CPU iPhone OS 4_3_2 like Mac OS X; zh-cn) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5",
		"Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_3_2 like Mac OS X; zh-cn) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5",
		"Mozilla/5.0 (Linux; U; Android 2.3.3; zh-cn; HTC_DesireS_S510e Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		"Mozilla/5.0 (SymbianOS/9.3; U; Series60/3.2 NokiaE75-1 /110.48.125 Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413",
		"Mozilla/5.0 (iPad; U; CPU OS 4_3_3 like Mac OS X; zh-cn) AppleWebKit/533.17.9 (KHTML, like Gecko) Mobile/8J2",
		"Mozilla/5.0 (Windows NT 5.2) AppleWebKit/534.30 (KHTML, like Gecko) Chrome/12.0.742.122 Safari/534.30",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.202 Safari/535.1",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/534.51.22 (KHTML, like Gecko) Version/5.1.1 Safari/534.51.22",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A5313e Safari/7534.48.3",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A5313e Safari/7534.48.3",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A5313e Safari/7534.48.3",
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.202 Safari/535.1",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; OMNIA7)",
		"Mozilla/5.0 (Windows NT 5.2) AppleWebKit/534.30 (KHTML, like Gecko) Chrome/12.0.742.122 Safari/534.30",
		"Mozilla/5.0 (Windows NT 5.1; rv:5.0) Gecko/20100101 Firefox/5.0",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.2; Trident/4.0; .NET CLR 1.1.4322; .NET CLR 2.0.50727; .NET4.0E; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; .NET4.0E; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C)",
		"Mozilla/4.0 (compatible; MSIE 60; Windows NT 5.1; SV1; .NET CLR 2.0.50727)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E)",
		"Opera/9.80 (Windows NT 5.1; U; zh-cn) Presto/2.9.168 Version/11.50",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1)",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; ) AppleWebKit/534.12 (KHTML, like Gecko) Maxthon/3.0 Safari/534.12",
	}
	htmlRequest.request = request.NewRequest(new(http.Client))

	return htmlRequest
}

func (s *HtmlRequest) getUserAgent() string {
	userAgent := s.userAgent
	if len(userAgent) > 0 {
		return userAgent
	} else {
		//随机选择用户代理
		maxIndex := len(s.userAgents)
		index := glib.RandIntRange(0, maxIndex)
		userAgent = s.userAgents[index]
	}

	return userAgent
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Get请求
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlRequest) Get(url string) (IHtmlResponse, error) {
	if len(url) == 0 {
		return nil, errors.New("argument url error")
	}

	//http头
	if len(s.headers) == 0 {
		s.headers = make(map[string]string, 0)
	}

	//用户代理
	s.headers["User-Agent"] = s.getUserAgent()
	s.request.Headers = s.headers

	if len(s.cookies) > 0 {
		s.request.Cookies = s.cookies
	}

	if len(s.params) > 0 {
		s.request.Params = s.params
	}

	resp, err := s.request.Get(url)
	defer resp.Body.Close()

	httpResponse := new(HtmlResponse)
	httpResponse.Header = resp.Header

	if data, err := resp.Content(); err == nil {
		httpResponse.Data = data
	}

	httpResponse.StatusCode = resp.StatusCode
	httpResponse.Status = resp.Status

	return httpResponse, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Post请求
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlRequest) Post(url string) (IHtmlResponse, error) {
	//http头
	if len(s.headers) == 0 {
		s.headers = make(map[string]string, 0)
	}

	//用户代理
	s.headers["User-Agent"] = s.getUserAgent()
	s.request.Headers = s.headers

	if len(s.cookies) > 0 {
		s.request.Cookies = s.cookies
	}

	if len(s.json) > 0 {
		s.request.Json = s.json
	}

	if len(s.data) > 0 {
		s.request.Data = s.data
	}

	if len(s.files) > 0 {
		fileFields := make([]request.FileField, 0)
		for _, file := range s.files {
			fileField := request.FileField{
				FieldName: file.FieldName, FileName: file.FileName, File: file.Datas,
			}
			fileFields = append(fileFields, fileField)
		}
		s.request.Files = fileFields
	}

	resp, err := s.request.Post(url)
	defer resp.Body.Close()

	httpResponse := new(HtmlResponse)
	httpResponse.Header = resp.Header

	if data, err := resp.Content(); err == nil {
		httpResponse.Data = data
	}

	httpResponse.StatusCode = resp.StatusCode
	httpResponse.Status = resp.Status

	return httpResponse, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置用户代理http头
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlRequest) SetUserAgent(userAgent string) {
	s.userAgent = userAgent
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置头
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlRequest) SetHeaders(headers map[string]string) {
	s.headers = headers
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置参数
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlRequest) SetParams(params map[string]string) {
	s.params = params
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置Cookie
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlRequest) SetCookies(cookies map[string]string) {
	s.cookies = cookies
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置Json数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlRequest) SetJson(json map[string]string) {
	s.json = json
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置字典数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlRequest) SetData(data map[string]string) {
	s.data = data
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置文件数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlRequest) SetFiles(files FormFiles) {
	s.files = files
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取请求内容
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlResponse) GetData() []byte {
	return s.Data
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取请求头
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlResponse) GetHeader() http.Header {
	return s.Header
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取状态码
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlResponse) GetStatusCode() int {
	return s.StatusCode
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取状态描述
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlResponse) GetStatus() string {
	return s.Status
}
