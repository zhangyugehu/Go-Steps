package parser

import (
	"study/crawler/engine"
	"regexp"
	"strconv"
	"study/crawler/model"
)

var ageReg = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightReg = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightReg = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([^<]+)</span></td>`)
var nameReg = regexp.MustCompile(`<h1 class="ceiling-name ib fl fs24 lh32 blue">([^<]+)</h1>`)
var genderReg = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var incomeReg = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var educationReg = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var marriageReg = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var occupationReg = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hokouReg = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xinzuoReg = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var houseReg = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carReg = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

var idUrlReg = regexp.MustCompile(`http://album.zhenai.com/u/[\d]+`)


func ParseProfile(contents []byte,
	url string,
	name string) engine.ParseResult {

	profile :=model.Profile{
		Name:name,
	}

	age, err := strconv.Atoi(extracString(contents, ageReg))
	if err==nil{
		profile.Age = age
	}

	height, err := strconv.Atoi(extracString(contents, heightReg))
	if err==nil{
		profile.Height = height
	}

	weight, err := strconv.Atoi(extracString(contents, weightReg))
	if err==nil{
		profile.Weight = weight
	}

	profile.Marriage = extracString(contents, marriageReg)
	//profile.Name = extracString(contents, nameReg)
	profile.Gender = extracString(contents, genderReg)
	profile.Income = extracString(contents, incomeReg)
	profile.Education = extracString(contents, educationReg)
	profile.Occupation = extracString(contents, occupationReg)
	profile.Hokou = extracString(contents, hokouReg)
	profile.Xinzuo = extracString(contents, xinzuoReg)
	profile.House = extracString(contents, houseReg)
	profile.Car = extracString(contents, carReg)

	result := engine.ParseResult{
		Items:[]engine.Item{
			{
				Url:	url,
				Id:		extracString([]byte(url), idUrlReg),
				Payload:profile,
			},
		},
	}

	return result
}

func extracString(contents []byte, re *regexp.Regexp) string{
	match := re.FindSubmatch(contents)
	if len(match) >= 2{
		return string(match[1])
	}
	return ""
}
