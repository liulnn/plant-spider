package plant

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hu17889/go_spider/core/common/page"
)

type PlantProcesser struct {
}

func NewPlantProcesser() *PlantProcesser {
	return &PlantProcesser{}
}

var BaikeUrlReg = regexp.MustCompile(`^http://baike\.baidu\.com/view/.*?`)

func (this *PlantProcesser) Process(p *page.Page) {
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
			if BaikeUrlReg.MatchString(url) {
				urls = append(urls, url)
			}
		}
	})

	p.AddTargetRequests(urls, "html")
}

func (this *PlantProcesser) Finish() {
	fmt.Printf("TODO:before end spider \r\n")
}
