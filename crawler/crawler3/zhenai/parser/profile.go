package parser

import (
	"shishaoGo/crawler/crawler3/engine"
	"shishaoGo/crawler/crawler3/model"
	"regexp"
		)

//const ageRe = `"gender":0,"genderString":"男士",`
var generRe = regexp.MustCompile(`"gender":[01]{1},"genderString":"([^"]+)"`)
var messagesRe = regexp.MustCompile(`<div[^c]*class="m-btn purple"[^>]*>([^<]+)</div>`)

var numberRe = regexp.MustCompile("[0-9]")

func PaseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = extractString(contents, generRe)

	matches := messagesRe.FindAllSubmatch(contents, -1)
	//fmt.Print(matches)
	i := 0
	for _, m := range matches {
		switch i {
		case 0 :
			profile.Marriage = string(m[1])
		case 1 :
			profile.Age = string(m[1])
		case 2 :
			profile.Xinzuo = string(m[1])
		case 3 :
			profile.Height = string(m[1])
		case 4 :
			profile.Weight = string(m[1])
		case 5 :

		case 6 :
			profile.Income = string(m[1])
		case 7 :
			profile.Occupation = string(m[1])
		case 8 :
			profile.Education = string(m[1])
		default:

		}
		i++
	}

	result := engine.ParseResult{Items: []interface{}{profile}}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
