package utils

func Concat(delimiter string, args ...string) string {
	concatStr := ""
	for _, v := range args {
		concatStr += (v + delimiter)
	}

	return concatStr
}
