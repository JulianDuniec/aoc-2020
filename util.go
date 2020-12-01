package aoc

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func inputIntSlice(file string) []int {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("unable to open file %v", err)
	}
	defer f.Close()
	res := make([]int, 0)
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("unable to read line %v", err)
		}
		num, err := strconv.ParseInt(strings.TrimSpace(line), 10, 32)
		if err != nil {
			log.Fatalf("line did not contain a number %v", err)
		}
		res = append(res, int(num))
	}
	return res
}
