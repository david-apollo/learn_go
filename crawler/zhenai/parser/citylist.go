package parser

import (
	"strings"
	"regexp"
	"learn_go/crawler/engine"
)

// CityList func
func CityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(`"linkContent":"(.*?)","linkURL":"(.*?)"`)
	cityURL := re.FindAllStringSubmatch(string(contents), -1)

	result := engine.ParserResult{}
	limit := 1
	for _, m := range cityURL {
		if strings.Contains(m[1], "征婚"){
			continue
		}
		result.Items = append(result.Items, string(m[1]))
		result.Requests = append(result.Requests, engine.Request{
			Url: strings.Replace(m[2], `\u002F`, "/", -1),
			ParserFunc: City,
		})
		limit --
		if limit == 0 {
			break
		}
	}
	return result
}