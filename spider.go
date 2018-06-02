package gspider

import (
	"log"
	"time"
)

import (
	"github.com/sanxia/glib"
)

/* ================================================================================
 * 爬虫
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

type (
	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * 爬虫接口
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	ISpider interface {
		Start()
		SetRequest(request IHtmlRequest)
		RegisterAnalyzer(name string, analyzer IAnalyzer)
		AddPage(page IPage)
		GetResult() chan *SpiderResult
		Stop()
	}

	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * 爬虫数据结构
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	Spider struct {
		request              IHtmlRequest
		analyzers            map[string]IAnalyzer
		downloadChan         chan IPage
		parseChan            chan IPage
		resultChan           chan *SpiderResult
		doneChan             chan bool
		perSecondDownload    int
		currentDownloadCount int
		downloadCount        int
		isDone               bool
	}

	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * 爬虫结果数据结构
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	SpiderResult struct {
		Page IPage
		Data interface{}
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 初始化爬虫
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewSpider(maxCount int) ISpider {
	if maxCount == 0 {
		maxCount = 10
	}

	spider := &Spider{
		analyzers:         make(map[string]IAnalyzer, 0),
		perSecondDownload: 6,
		downloadChan:      make(chan IPage, maxCount),
		parseChan:         make(chan IPage, maxCount),
		resultChan:        make(chan *SpiderResult, maxCount),
		doneChan:          make(chan bool),
	}

	return spider
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 开始爬取数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Spider) Start() {
	go s.download()
	go s.parse()

	for {
		select {
		case isDone := <-s.doneChan:
			s.isDone = isDone
			if s.isDone {
				close(s.downloadChan)
				close(s.parseChan)
				close(s.resultChan)
				close(s.doneChan)

				return
			}
		}
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置请求
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Spider) SetRequest(request IHtmlRequest) {
	s.request = request
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 注册分析器
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Spider) RegisterAnalyzer(name string, analyzer IAnalyzer) {
	if _, isOk := s.analyzers[name]; !isOk {
		s.analyzers[name] = analyzer
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 新增Page对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Spider) AddPage(page IPage) {
	if s.isDone {
		return
	}

	s.downloadChan <- page
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 下载任务
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Spider) download() {
	for {
		select {
		case <-time.After(time.Minute * 5):
			log.Println("download task idle")
		case page := <-s.downloadChan:
			sleepMillisecond := 15
			log.Println("download per sleep\r\n")
			time.Sleep(time.Duration(sleepMillisecond) * time.Millisecond)

			s.downloadData(page)

			s.currentDownloadCount = s.currentDownloadCount + 1
			s.downloadCount = s.downloadCount + 1

			if s.currentDownloadCount == s.perSecondDownload {
				log.Println("download limit speed sleep\r\n")
				time.Sleep(time.Duration(sleepMillisecond*2) * time.Millisecond)
				s.currentDownloadCount = 0
			}

			//每下载50条数据休眠随机时间
			if s.currentDownloadCount%5000 == 0 {
				rndSleep := glib.RandIntRange(1000, 2000)
				log.Printf("download rand sleep: %d\r\n", rndSleep)
				time.Sleep(time.Duration(rndSleep) * time.Millisecond)
			}
		}
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 下载数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Spider) downloadData(page IPage) {
	if !s.isDone {
		//下载数据
		page.GetData(s.request)

		s.parseChan <- page
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 分析数据任务
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Spider) parse() {
	for {
		select {
		case <-time.After(time.Minute * 5):
			log.Println("parse task idle")
		case page := <-s.parseChan:
			go s.parseData(page)
		}
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 分析数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Spider) parseData(page IPage) {
	if analyzer, isOk := s.analyzers[page.GetName()]; isOk {
		//解析结果
		data := analyzer.Parse(page)

		//结果
		spiderResult := &SpiderResult{
			Page: page,
			Data: data,
		}

		s.resultChan <- spiderResult
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取结果通道数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Spider) GetResult() chan *SpiderResult {
	return s.resultChan
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 结束爬取
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Spider) Stop() {
	s.doneChan <- true
}
