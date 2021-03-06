package parser

import (
	"io/ioutil"
	"testing"
)


func TestCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("./citylist_test_data.html")
	// contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	result := CityList(contents)

	const resultSize = 470
	expectedUrls := []string {
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}
}

