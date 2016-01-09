package main

import (
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/scheduler"
	"github.com/hu17889/go_spider/core/spider"
	"plant-spider/plant"
)

func main() {
	var startUrl = "http://baike.baidu.com/subview/412610/19548276.htm"
	sp := spider.NewSpider(plant.NewPlantProcesser(), "TaskName").
		SetScheduler(scheduler.NewQueueScheduler(true)).
		AddPipeline(pipeline.NewPipelineFile("plants.txt"))
	sp.AddUrl(startUrl, "html")
	sp.Run()
}
