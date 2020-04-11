package main

import (
	"learn_go/crawler/zhenai/parser"
	"learn_go/crawler/engine"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.CityList,
	})
}
