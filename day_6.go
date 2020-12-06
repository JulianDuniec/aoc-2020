package aoc

func sumDistinctPositiveAnswers(reader lineReader) int {
	sum := 0
	distinct := make(map[rune]struct{}, 0)
	for {
		line, eof := reader.readLine()
		if line == "" {
			sum += len(distinct)
			distinct = make(map[rune]struct{}, 0)
			continue
		}
		for _, char := range line {
			distinct[char] = struct{}{}
		}
		if eof {
			sum += len(distinct)
			break
		}
	}
	return sum
}

func sumIntersectingPositiveAnswers(reader lineReader) int {
	sum := 0
	var intersecting []rune
	for {
		line, eof := reader.readLine()
		if line == "" {
			sum += len(intersecting)
			intersecting = nil
		} else if intersecting == nil {
			intersecting = []rune(line)
		} else {
		removeNonIntersecting:
			for i := len(intersecting) - 1; i >= 0; i-- {
				c1 := intersecting[i]
				for _, c2 := range line {
					if c1 == c2 {
						continue removeNonIntersecting
					}
				}
				intersecting = append(intersecting[:i], intersecting[i+1:]...)
			}
		}
		if eof {
			break
		}
	}
	sum += len(intersecting)
	return sum
}
