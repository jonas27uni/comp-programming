package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type author struct {
	institution  string
	reviews      []int
	numReviewers int
}

func check(as []*author, reviews int, papers int) int {
	problems := 0
	// check rule vialotion for each paper
	calcNumReviews(as, reviews)
	for i, a := range as {
		if a.hasSelfReview(i+1) || hasSameInstReview(as, a) || a.numReviewers != reviews {
			problems++
		}
	}
	return problems
}

func calcNumReviews(as []*author, reviews int) {
	for _, v := range as {
		// the reviews they do
		for _, review := range v.reviews {
			as[review-1].numReviewers++
		}
	}
}

func hasSameInstReview(as []*author, a *author) bool {
	for i, v := range as {
		if a != v && a.institution == v.institution && arrContains(a.reviews, i) {
			return true
		}
	}
	return false
}

func (a *author) hasSelfReview(i int) bool {
	for _, r := range a.reviews {
		if i == r {
			return true
		}
	}
	return false
}

func arrContains(arr []int, i int) bool {
	for _, n := range arr {
		if n == i {
			return true
		}
	}
	return false
}

func loadCase(reviews int, papers int, rw *bufio.ReadWriter) []*author {
	as := []*author{}
	for c := 0; c < papers; c++ {
		as = append(as, parseAuthor(readLine(rw)))
	}
	return as
}

func exec(rw *bufio.ReadWriter) {
	i, err := strconv.Atoi(readLine(rw))
	if err != nil {
		log.Panicln(err)
	}

	for c := 1; c <= i; c++ {
		reviews, papers := revsAndPapers(readLine(rw))
		as := loadCase(reviews, papers, rw)
		p := check(as, reviews, papers)
		var ans string
		if p == 0 {
			ans = "Case #" + fmt.Sprint(c) + ": NO PROBLEMS FOUND"
		} else if p == 1 {
			ans = "Case #" + fmt.Sprint(c) + ": " + fmt.Sprint(p) + " PROBLEM FOUND"
		} else {
			ans = "Case #" + fmt.Sprint(c) + ": " + fmt.Sprint(p) + " PROBLEMS FOUND"
		}
		rw.WriteString(ans)
		rw.Flush()
		if c != i {
			rw.WriteString("\n")
			rw.Flush()
			readLine(rw)
		}
	}
}

func revsAndPapers(s string) (int, int) {
	sarr := strings.Split(s, " ")
	return sToInt(sarr[0]), sToInt(sarr[1])
}

func parseAuthor(in string) *author {
	sarr := strings.Split(in, " ")
	a := author{institution: sarr[0], numReviewers: 0}
	sarr = sarr[1:]
	reviews := []int{}
	for _, s := range sarr {
		review, err := strconv.Atoi(s)
		if err != nil {
			log.Println(err)
		}
		reviews = append(reviews, review)
	}
	a.reviews = reviews
	return &a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	exec(bufio.NewReadWriter(reader, writer))
}
func readLine(rw *bufio.ReadWriter) string {
	index, err := rw.ReadString('\n')
	if err != nil {
		log.Panicln(err)
	}
	return strings.Replace(index, "\n", "", -1)
}
func sToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panicln(err)
	}
	return i
}
