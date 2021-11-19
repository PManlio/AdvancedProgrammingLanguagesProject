package utils

import "strings"

func GenerateArray(s *string) []string {
	arr := strings.Split(*s, ",")

	// necessario perché da MySQL otteniamo sempre l'ultimo elemento come stringa vuota
	// quindi returno uno slice che va dal primo elemento dell'array arr fino al penultimo
	return arr[:len(arr)-1]
}
