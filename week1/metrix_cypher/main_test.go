package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
	"testing"
)

func Test_Case1(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	input := "2 1\n3 2\n"
	expected := "010"
	builder := &strings.Builder{}
	exec(GetRW(input, builder))
	if builder.String() != expected {
		log.Panicf("Expected %s but actual was %s", expected, builder.String())
	}
}
func Test_Case2(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	input := "18 29\n13 21\n"
	expected := "10010101"
	builder := &strings.Builder{}
	exec(GetRW(input, builder))
	if builder.String() != expected {
		log.Panicf("Expected %s but actual was %s", expected, builder.String())
	}
}
func Test_Case3(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	input := "1 0\n1 1\n"
	expected := "0"
	builder := &strings.Builder{}
	exec(GetRW(input, builder))
	if builder.String() != expected {
		log.Panicf("Expected %s but actual was %s", expected, builder.String())
	}
}
func Test_Case4(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	input := "1 1\n0 1\n"
	expected := "1"
	builder := &strings.Builder{}
	exec(GetRW(input, builder))
	if builder.String() != expected {
		log.Panicf("Expected %s but actual was %s", expected, builder.String())
	}
}

func Test_Case5Unity(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	input := "1 0\n0 1\n"
	expected := ""
	builder := &strings.Builder{}
	exec(GetRW(input, builder))
	if builder.String() != expected {
		log.Panicf("Expected %s but actual was %s", expected, builder.String())
	}
}

func Test_sd(t *testing.T) {
	if math.MaxUint64 > math.Pow(2, 128)-1 {
		log.Panicln("test")
	}
}

// func Test_BigNumbers(t *testing.T) {
// 	test("big numbers")
// 	i := math.Pow(2, 16) - 1
// 	s := strconv.FormatFloat(i, 'f', 6, 64)
// 	input := "1 " + s + "\n0 1\n"
// 	log.Println(input)
// 	expected := "1"
// 	builder := &strings.Builder{}
// 	exec(GetRW(input, builder))
// 	if builder.String() != expected {
// 		log.Panicf("Expected %s but actual was %s", expected, builder.String())
// 	}
// 	// log.SetFlags(log.Lshortfile)
// 	// i := math.Pow(2, 126) - 1
// 	// // log.Panicln(i)
// 	// ii := int64(i)
// 	// if i != float64(ii) {

// 	// 	log.Panicln(i)
// 	// }
// }

func GetRW(input string, builder *strings.Builder) *bufio.ReadWriter {
	reader := bufio.NewReader(strings.NewReader(input))
	writer := bufio.NewWriter(builder)
	return bufio.NewReadWriter(reader, writer)
}

func test(name string) {
	log.SetFlags(log.Lshortfile)
	log.SetOutput(os.Stdout)
	log.Println(name)
}
