package parser

import (
	"log"
	"fmt"
	"strconv"
	"learn_go/crawler/model"
	"encoding/json"
	"learn_go/crawler/engine"
)


// UserInfo type json
type UserInfo struct {
	Data struct {
		Age                      int      `json:"age"`
		BasicInfo                []string `json:"basicInfo"`
		DetailInfo               []string `json:"detailInfo"`
		EducationString          string   `json:"educationString"`
		EmotionStatus            int      `json:"emotionStatus"`
		Gender                   int      `json:"gender"`
		GenderString             string   `json:"genderString"`
		Nickname                 string   `json:"nickname"`
	} `json:"data"`

}


// Profile func
func Profile(contents []byte, name string) engine.ParserResult {
	profile := model.Profile{}
	userInfo := UserInfo{}
	err := json.Unmarshal(contents, &userInfo)
	if err != nil {
		log.Printf("error %v", err)
		return engine.ParserResult{}
	}

	profile.Name = name
	profile.Age = userInfo.Data.Age

	height, err := strconv.Atoi(userInfo.Data.BasicInfo[3][:len(userInfo.Data.BasicInfo[3]) - 2])
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(userInfo.Data.BasicInfo[4][:len(userInfo.Data.BasicInfo[4]) - 2])
	if err == nil {
		profile.Weight = weight
	}	

	profile.Income = userInfo.Data.BasicInfo[6]

	profile.Gender = userInfo.Data.GenderString

	profile.Car = userInfo.Data.DetailInfo[6]

	profile.Education = userInfo.Data.EducationString

	profile.Hokou = userInfo.Data.DetailInfo[1]

	profile.House = userInfo.Data.DetailInfo[5]

	profile.Marriage = userInfo.Data.BasicInfo[0]

	profile.Occupation = userInfo.Data.BasicInfo[7]

	profile.Xinzuo = userInfo.Data.BasicInfo[2]

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}
	fmt.Println(result)
	return result
}