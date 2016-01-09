package plant

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
)

type PlantProcesser struct {
}

func NewPlantProcesser() *PlantProcesser {
	return &PlantProcesser{}
}

var urlCache = NewUrlCache()
var BaikeUrlReg = regexp.MustCompile(`^http://baike\.baidu\.com/view/.*?`)

func (this *PlantProcesser) getUrls(query *goquery.Document) (urls []string) {
	query.Find("a").Each(func(i int, s *goquery.Selection) {
		url, isExist := s.Attr("href")
		if isExist {
			ex := strings.Index(url, ".htm")
			if ex > 0 {
				url = url[:ex]
			}
			if BaikeUrlReg.MatchString(url) {
				if urlCache.Set(url) {
					urls = append(urls, url)
				}
			}
		}
	})
	return urls
}

func (this *PlantProcesser) isPlant(query *goquery.Document, p *page.Page) bool {

	var isPlant bool = false
	query.Find(".basicInfo-item").Each(func(i int, s *goquery.Selection) {
		if strings.Trim(s.Text(), " \t\n") == "植物界" {
			isPlant = true
		}
	})
	return isPlant
}
func (this *PlantProcesser) getCatalog(query *goquery.Document, p *page.Page) {

	catalog := query.Find(".lemma-catalog").Find("span.text").Text()
	catalog = strings.Trim(catalog, " \t\n")
	p.AddField("catalog", catalog)
}

func (this *PlantProcesser) getName(query *goquery.Document, p *page.Page) {

	name := query.Find(".lemmaWgt-lemmaTitle-title").Find("h1").Text()
	name = strings.Trim(name, " \t\n")
	p.AddField("name", name)

}
func (this *PlantProcesser) getSummary(query *goquery.Document, p *page.Page) {

	summary := query.Find(".lemma-summary .para").Text()
	summary = strings.Trim(summary, " \t\n")
	p.AddField("summary", summary)
}

func (this *PlantProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		println(p.Errormsg())
		return
	}

	query := p.GetHtmlParser()

	if !this.isPlant(query, p) {
		p.SetSkip(true)
	}

	this.getName(query, p)
	this.getSummary(query, p)
	this.getCatalog(query, p)
	p.AddTargetRequests(this.getUrls(query), "html")
}

func (this *PlantProcesser) Finish() {
	fmt.Printf("TODO:before end spider \r\n")
}
