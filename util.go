package aoc

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
)

type closer func() error
type lineReader interface {
	readLine() (string, bool)
}

type bufioLineReader struct {
	r   *bufio.Reader
	eof bool
}

func (blr *bufioLineReader) readLine() (string, bool) {
	if blr.eof {
		panic("read after EOF")
	}
	v, err := blr.r.ReadString('\n')
	blr.eof = err == io.EOF
	if err != nil && blr.eof == false {
		log.Fatalf("unable to read line %v", err)
	}
	lastIndex := len(v) - 1
	if v[lastIndex] == '\n' {
		v = v[0:lastIndex]
	}
	return v, blr.eof
}

func lineReaderFromFile(file string) (lineReader, closer) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("unable to open file %v", err)
	}

	return &bufioLineReader{r: bufio.NewReader(f)}, f.Close
}

func lineReaderFromString(data string) lineReader {
	return &bufioLineReader{r: bufio.NewReader(bytes.NewReader([]byte(data)))}
}

func mustParseInt(s string) int {
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		log.Fatalf("invalid number %s %v", s, err)
	}
	return int(v)
}

func readStrings(reader lineReader) []string {
	lines := make([]string, 0)
	for {
		line, eof := reader.readLine()
		lines = append(lines, line)
		if eof {
			break
		}
	}
	return lines
}

func readNumbers(reader lineReader) []int {
	numbers := make([]int, 0)
	for {
		line, eof := reader.readLine()
		v, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatalf("invalid input %s %v", line, err)
		}
		numbers = append(numbers, int(v))
		if eof {
			break
		}
	}
	return numbers
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
