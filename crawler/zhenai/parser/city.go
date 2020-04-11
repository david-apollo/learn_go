package parser

import (
	"regexp"
	"learn_go/crawler/engine"
)


// City func
func City(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(`"memberId":(.*?),"nickname":"(.*?)"`)
	idName := re.FindAllStringSubmatch(string(contents), -1)
	result := engine.ParserResult{}
	for _, m := range idName {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: "https://album.zhenai.com/u/"+ m[1],
			ParserFunc: engine.NilParser,
		})
	}
	return result
}