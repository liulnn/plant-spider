package main

import (
	"fmt"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/scheduler"
	"github.com/hu17889/go_spider/core/spider"
	"plant"
	"regexp"
	"strings"
)

func main() {
	var startUrl = "http://baike.baidu.com/view/39352.htm"
	sp := spider.NewSpider(plant.NewPlantProcesser(), "TaskName").
		SetScheduler(scheduler.NewQueueScheduler(true)).
		AddPipeline(pipeline.NewPipelineFile("result.txt"))
	sp.AddUrl(startUrl, "html")
	sp.Run()
}
