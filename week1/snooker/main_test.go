package main

import (
	"bufio"
	"log"
	"strings"
	"testing"
)

func Test_Case1(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	input := "5\nred\nblack\npink\nred\nred\n"
	expected := "37"
	builder := &strings.Builder{}
	exec(GetRW(input, builder))
	if builder.String() != expected {
		log.Panicf("Expected %s but actual was %s", expected, builder.String())
	}
}

func Test_Case2(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	input := "3\nblue\nblack\npink\n"
	expected := "18"
	builder := &strings.Builder{}
	exec(GetRW(input, builder))
	if builder.String() != expected {
		log.Panicf("Expected %s but actual was %s", expected, builder.String())
	}
}

func Test_Case3(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	input := "8\nyellow\ngreen\nbrown\nred\nred\nred\nred\nred\n"
	expected := "34"
	builder := &strings.Builder{}
	exec(GetRW(input, builder))
	if builder.String() != expected {
		log.Panicf("Expected %s but actual was %s", expected, builder.String())
	}
}
func Test_Case4(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	input := "2\nred\nred\n"
	expected := "1"
	builder := &strings.Builder{}
	exec(GetRW(input, builder))
	if builder.String() != expected {
		log.Panicf("Expected %s but actual was %s", expected, builder.String())
	}
}

func GetRW(input string, builder *strings.Builder) *bufio.ReadWriter {
	reader := bufio.NewReader(strings.NewReader(input))
	writer := bufio.NewWriter(builder)
	return bufio.NewReadWriter(reader, writer)
}
