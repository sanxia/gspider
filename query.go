package gspider

import (
	"io"
	"net/http"
)

import (
	"github.com/PuerkitoBio/goquery"
)

/* ================================================================================
 * Html页面查询
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

type (
	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * Html查询数据结构
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	HtmlQuery struct {
		Document *goquery.Document
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据url初始化
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlQuery) FromUrl(url string) error {
	var err error
	s.Document, err = goquery.NewDocument(url)
	return err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据http.Response初始化
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlQuery) FromResponse(res *http.Response) error {
	var err error
	s.Document, err = goquery.NewDocumentFromResponse(res)
	return err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据io.Reader初始化
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlQuery) FromReader(reader io.Reader) error {
	var err error
	s.Document, err = goquery.NewDocumentFromReader(reader)
	return err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 克隆对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HtmlQuery) Clone() *HtmlQuery {
	return &HtmlQuery{
		Document: goquery.CloneDocument(s.Document),
	}
}

/*

  type Selection struct {
     Nodes    []*html.Node
     document *Document
     prevSel  *Selection
  }

Selection类型提供的方法
-  Eq(index int) *Selection     //根据索引获取某个节点集
-  First() *Selection          //获取第一个子节点集
-  Last() *Selection         //获取最后一个子节点集
-  Next() *Selection         //获取下一个兄弟节点集
-  NextAll() *Selection      //获取后面所有兄弟节点集
-  Prev() *Selection         //前一个兄弟节点集
-  Get(index int) *html.Node  //根据索引获取一个节点
-  Index() int                //返回选择对象中第一个元素的位置
-  Slice(start, end int) *Selection  //根据起始位置获取子节点集
*/

/*
扩大 Selection 集合（增加选择的节点）
-  Add(selector string) *Selection //将匹配到的节点添加当前节点集合中
-  AndSelf() *Selection    //将堆栈上的前一组元素添加到当前的
-  Union() *Selection    //which is an alias for AddSelection()
*/

/*
过滤方法，减少节点集合
- End() *Selection
- Filter…()     //过滤
- Has…()
- Intersection()   //which is an alias of FilterSelection()
- Not…()
*/

/*
- Each(f func(int, *Selection)) *Selection //遍历
- EachWithBreak(f func(int, *Selection) bool) *Selection  //可中断遍历
- Map(f func(int, *Selection) string) (result []string)  //返回字符串数组
*/

/*
检测或获取节点属性值
- Attr(), RemoveAttr(), SetAttr()  //获取，移除，设置属性的值
- AttrOr 获取对应的标签属性。这个可以设置第二个参数。获取的默认值 如果获取不到默认调用对应默认值
- AddClass(), HasClass(), RemoveClass(), ToggleClass()
- Html()  //获取该节点的html
- Length() //返回该Selection的元素个数
- Size(), which is an alias for Length()
- Text()  //获取该节点的文本值
*/

/*
查询或显示一个节点的身份
- Contains() //包含
- Is…()
*/

/*
在文档树之间来回跳转（常用的查找节点方法）
- Children() 返回所有子元素
- Contents()
- Find() 查找获取当前匹配的每个元素的后代
- Next() 获取下一个元素
- Parent[s]()
- Prev() 获取上一个元素
- Siblings()
- Filter 过滤标签元素
*/

//content.Find(".name").Text()
//content.Find("#name").Text()
//content.Find("input[name='gender']:checked")
/*
for i:=0;i<sex.Length();i++{
	if sex.Eq(i).Attr("checked") != "checked"{
        continue
	}

	if sex.Eq(i).Attr("value")=="0"{
        info.sex = "女"
	}

	if sex.Eq(i).Attr("value")=="1"{
        info.sex = "男"
	}
*/

/*

	query.Find("ul[class=\"publications-list  weekly\"]").Eq(0).Find("li").Each(func(i int, s *goquery.Selection) {
		a := s.Find("a").Eq(0)
		a.Next()
		r := Book{}
		//获取对应 跳转链接
		r.Url, _ = a.Attr("href")
		//获取书的名字
		r.Title = a.Find(".publications-item-title").Eq(0).Text()
		//获取书的图片
		r.Img, _ = a.Find(".publications-item-image").Eq(0).Attr("src")
		//获取书的作者
		r.Author = a.Find(".publications-item-author").Eq(0).Text()
		//获取书的售价
		r.Sell = a.Find(".publications-item-promotion").Eq(0).Find("span").Eq(0).Text()
		data_arr[i] = r
	})
	for k, v := range data_arr {
		fmt.Println(k, v)
	}
*/
