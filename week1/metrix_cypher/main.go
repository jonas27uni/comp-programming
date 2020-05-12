package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func exec(rw *bufio.ReadWriter) {
	top := convert(readLine(rw))
	bottom := convert(readLine(rw))
	result := ""
	for compare(top, bottom) {
		if sum(bottom) > sum(top) {
			bottom = substract(bottom, top)
			result = result + "0"
		} else if sum(bottom) < sum(top) {
			top = substract(top, bottom)
			result = result + "1"
		}
	}
	rw.Writer.WriteString(result)
	rw.Writer.Flush()
}

func substract(m1 []int64, m2 []int64) []int64 {
	return []int64{m1[0] - m2[0], m1[1] - m2[1]}
}

func compare(m1 []int64, m2 []int64) bool {
	unity := []int64{1, 0, 0, 1}
	if m1[0] == unity[0] && m1[1] == unity[1] && m2[0] == unity[2] && m2[1] == unity[3] {
		return false
	}
	return true
}

// func toMatrix(top []int64, bottom []int64) []int64 {
// 	return append(top, bottom...)
// }

// func multiple(left []int64, right []int64) []int64 {
// 	result := []int64{0, 0, 0, 0}
// 	result[0] = left[0]*right[0] + left[1]*right[2]
// 	result[1] = left[0]*right[1] + left[1]*right[3]
// 	result[2] = left[2]*right[0] + left[3]*right[2]
// 	result[3] = left[2]*right[1] + left[3]*right[3]
// 	return result
// }

func sum(arr []int64) int64 {
	var sum big.I
	for _, v := range arr {
		sum += v
	}
	if sum < 1 {
		// log.Panicln("under 1")
	}
	return sum
}

func convert(in string) []int64 {
	i := []int64{}
	sarr := strings.Split(in, " ")
	for _, s := range sarr {
		out, err := strconv.ParseFloat(s, 64)
		if err != nil {
			log.Println(err)
		}
		i = append(i, int64(out))
	}
	return i
}

func readLine(rw *bufio.ReadWriter) string {
	input, err := rw.ReadString('\n')
	if err != nil {
		log.Panicln(err)
	}
	return strings.Replace(input, "\n", "", -1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	exec(bufio.NewReadWriter(reader, writer))
}
