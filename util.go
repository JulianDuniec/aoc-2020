package aoc

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
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
