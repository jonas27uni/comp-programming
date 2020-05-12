package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	RED    = 1
	YELLOW = 2
	GREEN  = 3
	BROWN  = 4
	BLUE   = 5
	PINK   = 6
	BLACK  = 7
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	exec(bufio.NewReadWriter(reader, writer))
}
func exec(rw *bufio.ReadWriter) {
	i, err := strconv.Atoi(readLine(rw))
	if err != nil {
		log.Panicln(err)
	}
	balls := []int{}
	for c := 0; c < i; c++ {
		balls = append(balls, convert(readLine(rw)))
	}

	occ := occurences(balls, 1)
	max := max(balls)
	if max == 1 {
		rw.Writer.WriteString("1")
		rw.Writer.Flush()
		// fmt.Println(len(balls))
	} else {
		rw.Writer.WriteString(strconv.Itoa(occ*(max+1) + sum(balls)))
		rw.Writer.Flush()
		// fmt.Println(occ*(max+1) + sum(balls))
	}

}

func readLine(rw *bufio.ReadWriter) string {
	index, err := rw.ReadString('\n')
	if err != nil {
		log.Panicln(err)
	}
	return strings.Replace(index, "\n", "", -1)
}

func sum(arr []int) int {
	var sum int
	for _, v := range arr {
		if v != 1 {
			sum += v
		}
	}
	return sum
}

func max(arr []int) int {
	var max int
	for i, v := range arr {
		if i == 0 || v > max {
			max = v
		}
	}
	return max
}

func occurences(arr []int, i int) int {
	var occ int
	for _, v := range arr {
		if i == v {
			occ++
		}
	}
	return occ
}

func convert(i string) int {
	if i == "red" {
		return RED
	} else if i == "yellow" {
		return YELLOW
	} else if i == "green" {
		return GREEN
	} else if i == "brown" {
		return BROWN
	} else if i == "blue" {
		return BLUE
	} else if i == "pink" {
		return PINK
	} else {
		return BLACK
	}
}
