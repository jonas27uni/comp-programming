package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func exec(rw *bufio.ReadWriter) {
	unity := []int{1, 0, 0, 1}
	top := convert(readLine(rw))
	bottom := convert(readLine(rw))
	result := ""
	for toMatrix(top, bottom) != unity {
		if sum(bottom) > sum(top) {
			bottom = substract(bottom, top)
		} else {

		}
	}

}

func toMatrix(top []int, bottom []int) []int {
	return append(top, bottom...)
}

func multiple(left []int, right []int) []int {
	result := []int{0, 0, 0, 0}
	result[0] = left[0]*right[0] + left[1]*right[2]
	result[1] = left[0]*right[1] + left[1]*right[3]
	result[2] = left[2]*right[0] + left[3]*right[2]
	result[3] = left[2]*right[1] + left[3]*right[3]
	return result
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

func convert(in string) []int {
	i := []int{}
	sarr := strings.Split(in, " ")
	for _, s := range sarr {
		out, err := strconv.Atoi(s)
		if err != nil {
			log.Println(err)
		}
		i = append(i, out)
	}
	return i
}

func readLine(rw *bufio.ReadWriter) string {
	index, err := rw.ReadString('\n')
	if err != nil {
		log.Println(err)
	}
	return strings.Replace(index, "\n", "", -1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	exec(bufio.NewReadWriter(reader, writer))
}
