package parser

import (
	"math"
	"time"
	"strconv"
	"regexp"
	"learn_go/crawler/engine"
)


// City func
func City(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(`"memberId":(.*?),"nickname":"(.*?)"`)
	idName := re.FindAllStringSubmatch(string(contents), -1)
	result := engine.ParserResult{}
	for _, m := range idName {
		name := string(m[2])
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			Url: `https://album.zhenai.com/api/profile/getObjectProfile.do?objectID=` + m[1] + "&_=" + strconv.FormatInt(time.Now().UnixNano() / int64(math.Pow(10, 6)), 10) + `&ua=h5%2F1.0.0%2F1%2F0%2F0%2F0%2F0%2F0%2F%2F0%2F0%2F7da5ff8c-0eab-416e-99f9-2fd5228ee985%2F0%2F0%2F1854436954`,
			ParserFunc: func(c []byte) engine.ParserResult {
				return Profile(contents, name)
			},
		})
	}
	return result
}