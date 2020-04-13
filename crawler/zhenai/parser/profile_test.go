package parser

import (
	"learn_go/crawler/model"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("./profile_test_data.json")
	if err != nil {
		panic(err)
	}

	result := Profile(contents, "张韶勇")
	fmt.Println(result)

	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Age:        27,
		Height:     165,
		Weight:     70,
		Income:     "月收入:1.2-2万",
		Gender:     "男士",
		Name:       "张韶勇",
		Xinzuo:     "射手座(11.22-12.21)",
		Occupation: "工程师",
		Marriage:   "未婚",
		House:      "已购房",
		Hokou:      "籍贯:重庆",
		Education:  "大专",
		Car:        "已买车",
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}