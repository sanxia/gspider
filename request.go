package gspider

import (
	"io"
	"net/http"
)

import (
	"github.com/mozillazg/request"
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
		GetContent() []byte
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
		headers map[string]string
		params  map[string]string
		cookies map[string]string
		json    map[string]string
		data    map[string]string
		files   FormFiles
		request *request.Request
	}

	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * Html响应数据结构
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	HtmlResponse struct {
		Header     http.Header
		Content    []byte
		StatusCode int
		Status     string
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 初始化Html请求
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewHtmlRequest() IHtmlRequest {
	htmlRequest := &HtmlRequest{}
	htmlRequest.request = request.NewRequest(new(http.Client))

	return htmlRequest
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Get请求
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlRequest) Get(url string) (IHtmlResponse, error) {
	if len(s.headers) > 0 {
		s.request.Headers = s.headers
	}

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

	if content, err := resp.Content(); err == nil {
		httpResponse.Content = content
	}

	httpResponse.StatusCode = resp.StatusCode
	httpResponse.Status = resp.Status

	return httpResponse, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Post请求
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlRequest) Post(url string) (IHtmlResponse, error) {
	if len(s.headers) > 0 {
		s.request.Headers = s.headers
	}

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

	if content, err := resp.Content(); err == nil {
		httpResponse.Content = content
	}

	httpResponse.StatusCode = resp.StatusCode
	httpResponse.Status = resp.Status

	return httpResponse, err
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
func (s *HtmlResponse) GetContent() []byte {
	return s.Content
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
