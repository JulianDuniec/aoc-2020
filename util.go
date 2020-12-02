package aoc

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type iterateLineCb func(line string)

func iterateLines(file string, cb iterateLineCb) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("unable to open file %v", err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	eof := false
	for !eof {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			eof = true
		} else if err != nil {
			log.Fatalf("unable to read line %v", err)
		}
		cb(line)
	}
}

func parseInputAsIntSlice(file string) []int {
	res := make([]int, 0)
	iterateLines(file, func(line string) {
		num, err := strconv.ParseInt(strings.TrimSpace(line), 10, 32)
		if err != nil {
			log.Fatalf("line did not contain a number %v", err)
		}
		res = append(res, int(num))
	})
	return res
}
