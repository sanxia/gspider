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
		GetTag() string
		GetName() string
		GetUrl() string
		GetData(request IHtmlRequest) []byte
		GetExtend() interface{}
	}

	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * Page数据结构
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	Page struct {
		Name     string
		Url      string
		Filename string
		Data     []byte
		Tag      string
		Extend   interface{}
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取名称
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Page) GetName() string {
	return s.Name
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取Url
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Page) GetUrl() string {
	return s.Url
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取页面数据
 * 优先级由高到低 Content > Filename > Url
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Page) GetData(request IHtmlRequest) []byte {
	var data []byte

	if len(s.Data) > 0 {
		//直接获取数据
		data = s.Data
	} else if len(s.Filename) > 0 {
		//从磁盘文件获取数据
		currentPath := glib.GetCurrentPath()
		fullPath := filepath.Join(currentPath, s.Filename)

		fileContent, err := glib.GetFileContent(fullPath)
		if err == nil {
			data = fileContent
			s.Data = data
		}

	} else {
		//从Url下载数据
		if request != nil {
			if len(s.Url) > 0 {
				if resp, err := request.Get(s.Url); err == nil {
					data = resp.GetData()
					s.Data = data
				} else {
					log.Printf("GetData request error: %v", err)
				}
			}
		}
	}

	return data
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取Tag
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Page) GetTag() string {
	return s.Tag
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取扩展数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Page) GetExtend() interface{} {
	return s.Extend
}
