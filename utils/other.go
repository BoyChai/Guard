package utils

func Trim(src string) (dist string) {
	if len(src) == 0 {
		return
	}

	r, distR := []rune(src), []rune{}
	for i := 0; i < len(r); i++ {

		if r[i] == 10 || r[i] == 32 {
			continue
		}

		distR = append(distR, r[i])
	}

	dist = string(distR)
	return
}
