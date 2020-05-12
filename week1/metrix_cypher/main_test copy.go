package main

import (
	"bufio"
	"log"
	"strings"
	"testing"
)

func Test_Case1(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	input := "2 1\n3 2"
	output := "010"
	builder := &strings.Builder{}

	exec(GetRW(input, builder))
	if builder.String() != output {
		log.Panicf("Expected %s but actual was %s", output, builder.String())
	}

}

func GetRW(input string, builder *strings.Builder) *bufio.ReadWriter {
	reader := bufio.NewReader(strings.NewReader(input))
	writer := bufio.NewWriter(builder)
	return bufio.NewReadWriter(reader, writer)
}
