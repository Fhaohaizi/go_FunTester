package ftool

import (
	"funtester/base"
	"regexp"
)

func Match(str, regex string) bool {
	compile, err := regexp.Compile(regex)
	if err != nil {
		return false
	}
	return compile.MatchString(str)
}

func Find(str, regex string) string {
	compile, err := regexp.Compile(regex)
	if err != nil {
		return base.Empty
	}
	return compile.FindString(str)
}

func FindAll(str, regex string) []string {
	compile, err := regexp.Compile(regex)
	if err != nil {
		return []string{}
	}
	submatch := compile.FindAllStringSubmatch(str, -1)
	res := make([]string, len(submatch))
	for i, strings := range submatch {
		res[i] = strings[0]
	}
	return res
}
