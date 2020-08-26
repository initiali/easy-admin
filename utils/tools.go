package utils

import "strconv"

func StrToUint(str string) uint {
	i, _ := strconv.Atoi(str)
	return uint(i)
}
