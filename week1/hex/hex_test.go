package main

import (
	"bufio"
	"log"
	"strings"
	"testing"
)

func Test_Case1(t *testing.T) {
	test("Case1")
	input := "1\n2 3\nUCM 2 3\nUAM 1 3\nUPM 1 2\n"
	expected := "Case #1: NO PROBLEMS FOUND"
	builder := &strings.Builder{}
	exec(GetRW(input, builder))
	if builder.String() != expected {
		log.Panicf("Expected %s but actual was %s", expected, builder.String())
	}
}
func Test_Case2(t *testing.T) {
	test("Case1")
	input := "1\n2 3\nUCM 2 3\nUAM 1 2\nUPM 2 2\n"
	expected := "Case #1: 3 PROBLEMS FOUND"
	builder := &strings.Builder{}
	exec(GetRW(input, builder))
	if builder.String() != expected {
		log.Panicf("Expected %s but actual was %s", expected, builder.String())
	}
}
func Test_Case3(t *testing.T) {
	test("Case3")
	input := "2\n2 3\nUCM 2 3\nUAM 1 3\nUPM 1 2\n\n2 3\nUCM 2 3\nUAM 1 2\nUPM 2 2\n"
	expected := "Case #1: NO PROBLEMS FOUND\nCase #2: 3 PROBLEMS FOUND"
	builder := &strings.Builder{}
	exec(GetRW(input, builder))
	if builder.String() != expected {
		log.Panicf("Expected %s but actual was %s", expected, builder.String())
	}
}

func Test_Case4(t *testing.T) {
	test("Case4")
	input := "3\n2 3\nUCM 2 3\nUAM 1 3\nUPM 1 2\n\n2 3\nUCM 2 3\nUAM 1 2\nUPM 2 2\n\n2 4\nUAM 2 3\nUAM 3 4\nUPM 1 4\nUPM 1 2\n"
	expected := "Case #1: NO PROBLEMS FOUND\nCase #2: 3 PROBLEMS FOUND\nCase #3: 1 PROBLEM FOUND"
	builder := &strings.Builder{}
	exec(GetRW(input, builder))
	if builder.String() != expected {
		log.Panicf("Expected %s but actual was %s", expected, builder.String())
	}
}

func Test_Case5(t *testing.T) {
	test("Case5")
	input := "1\n2 6\nUCM 4 6\nUAM 4 5\nUPM 1 2\nUCM 2 3\nUAM 1 3\nUPM 1 2\n"
	expected := "Case #1: NO PROBLEMS FOUND"
	builder := &strings.Builder{}
	exec(GetRW(input, builder))
	if builder.String() != expected {
		log.Panicf("Expected %s but actual was %s", expected, builder.String())
	}
}

func Test_Author(t *testing.T) {
	test("author")
	a := parseAuthor("s 1 2")
	if a.institution != "s" || a.reviews[0] != 1 || a.reviews[1] != 2 {
		log.Panicln("author func incorrect")
	}
}

func test(name string) {
	log.SetFlags(log.Lshortfile)
	log.Println(name)
}

func GetRW(input string, builder *strings.Builder) *bufio.ReadWriter {
	reader := bufio.NewReader(strings.NewReader(input))
	writer := bufio.NewWriter(builder)
	return bufio.NewReadWriter(reader, writer)
}
