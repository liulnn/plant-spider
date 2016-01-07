package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/scheduler"
	"github.com/hu17889/go_spider/core/spider"
	"regexp"
	"strings"
)

type MyPageProcesser struct {
}

func NewMyPageProcesser() *MyPageProcesser {
	return &MyPageProcesser{}
}

var baikeUrlReg = regexp.MustCompile(`^http://baike\.baidu\.com/view/.*?`)

func (this *MyPageProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		println(p.Errormsg())
		return
	}

	var urls []string

	query := p.GetHtmlParser()

	name := query.Find(".lemmaWgt-lemmaTitle-title").Text()
	name = strings.Trim(name, " \t\n")
	p.AddField("name", name)

	summary := query.Find(".lemma-summary .para").Text()
	summary = strings.Trim(summary, " \t\n")
	p.AddField("summary", summary)

	query.Find("a").Each(func(i int, s *goquery.Selection) {
		url, isExist := s.Attr("href")
		if isExist {
			if baikeUrlReg.MatchString(url) {
				urls = append(urls, url)
			}
		}
	})

	p.AddTargetRequests(urls, "html")
}

func (this *MyPageProcesser) Finish() {
	fmt.Printf("TODO:before end spider \r\n")
}

func main() {
	sp := spider.NewSpider(NewMyPageProcesser(), "TaskName").
		SetScheduler(scheduler.NewQueueScheduler(true)).
		AddPipeline(pipeline.NewPipelineFile("result.txt"))
	sp.AddUrl("http://baike.baidu.com/view/39352.htm", "html")
	sp.Run()
}
