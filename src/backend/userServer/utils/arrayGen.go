package utils

import "strings"

func GenerateArray(s *string) []string {
	return strings.Split(*s, ",")
}
