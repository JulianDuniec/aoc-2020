package aoc

import (
	"log"
	"strconv"
	"strings"
)

func findPossibleContainersCount(reader lineReader, query string) int {
	parents := map[string][]string{}
	for rule := range parseRules(reader) {
		p, ok := parents[rule.child]
		if !ok {
			p = make([]string, 0)
		}
		p = append(p, rule.container)
		parents[rule.child] = p
	}

	set := make(map[string]struct{}, 0)
	var search func(string)
	search = func(k string) {
		for _, v := range parents[k] {
			set[v] = struct{}{}
			search(v)
		}
	}
	search(query)
	return len(set)
}

func findTotalChildCount(reader lineReader, query string) int {
	children := make(map[string]map[string]int, 0)
	for rule := range parseRules(reader) {
		v, ok := children[rule.container]
		if !ok {
			v = make(map[string]int, 0)
		}
		v[rule.child] = rule.count
		children[rule.container] = v
	}

	var search func(string) int
	search = func(k string) int {
		count := 0
		for k, v := range children[k] {
			count += v + v*search(k)
		}
		return count
	}
	return search(query)
}

type rule struct {
	container string
	count     int
	child     string
}

func parseRules(reader lineReader) <-chan rule {
	chn := make(chan rule)
	go func() {
		for {
			line, eof := reader.readLine()
			s := strings.Split(line, " bags contain ")
			ruleColor := s[0]
			ss := strings.Split(s[1], ", ")
			for _, sr := range ss {
				if sr == "no other bags." {
					continue
				}
				words := strings.Split(sr, " ")
				color := ""
				count, err := strconv.ParseInt(words[0], 10, 32)
				if err != nil {
					log.Fatalf("unable to parse number %v", err)
				}
				for i, w := range words[1:] {
					if strings.Index(w, "bag") == 0 {
						color = strings.Join(words[1:i+1], " ")
						break
					}
				}
				chn <- rule{ruleColor, int(count), color}
			}
			if eof {
				break
			}
		}
		close(chn)
	}()
	return chn
}
