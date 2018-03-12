package gspider

import (
	"log"
	"path/filepath"
)

import (
	"github.com/sanxia/glib"
)

/* ================================================================================
 * Html页面
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

type (
	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * Page接口
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	IPage interface {
		GetName() string
		GetContent(request IHtmlRequest) []byte
	}

	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * Page数据结构
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	Page struct {
		Name     string
		Url      string
		Content  []byte
		Filename string
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取自定义名称
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Page) GetName() string {
	return s.Name
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取页面数据
 * 优先级由高到低 Content > Filename > Url
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Page) GetContent(request IHtmlRequest) []byte {
	var content []byte

	if len(s.Content) > 0 {
		//直接获取数据
		content = s.Content
	} else if len(s.Filename) > 0 {
		//从磁盘文件获取数据
		currentPath := glib.GetCurrentPath()
		fullPath := filepath.Join(currentPath, s.Filename)

		fileContent, err := glib.GetFileContent(fullPath)
		if err == nil {
			content = fileContent
		}
	} else {
		//从Url下载数据
		log.Printf("GetContent Url: %s", s.Url)
		resp, err := request.Get(s.Url)

		if err == nil {
			content = resp.GetContent()
		} else {
			log.Printf("GetContent error: %v", err)
		}
	}

	s.Content = content

	return s.Content
}
