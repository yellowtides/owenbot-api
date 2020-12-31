package utils

import (
	"regexp"
)

func inserty(text string) string {
	var re = regexp.MustCompile("([nm])([oO])")
	return re.ReplaceAllString(text, "${1}y$2")
}

func insertY(text string) string {
	var re = regexp.MustCompile("([NM])([oO])")
	return re.ReplaceAllString(text, "${1}Y$2")
}

func owoifyChars(text string) string {
    var reLower = regexp.MustCompile("[lr]")
	var text0 = reLower.ReplaceAllString(text, "w")
	
    var reUpper = regexp.MustCompile("[LR]")
    return reUpper.ReplaceAllString(text0, "W")
}

func Owoify(text string) string {
	return owoifyChars(insertY(inserty(text))) + " owo"
}