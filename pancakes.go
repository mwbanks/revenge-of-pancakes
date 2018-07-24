package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getStackFlips(pos *int, dat []byte) int64 {
	var count int64
	prev := dat[*pos]
	*pos++
	for ; *pos < len(dat); *pos++ {
		current := dat[*pos]
		if current == '\n' {
			// Finally stack must be '+'
			if prev == '-' {
				count++
			}
			*pos++
			break
		} else if prev != current {
			count++
			prev = current
		}
	}
	return count
}

// Iterate through file and find stacks
func findFlips(dat []byte) {
	stackCount, pos := getCount(dat)
	countCounter := int64(0)
	for ; countCounter < stackCount; countCounter++ {
		count := getStackFlips(&pos, dat)
		fmt.Printf("Case #%d: %d\n", countCounter+1, count)
	}

}

// Returns the number of stacks, and the position in the file to start reading stacks
func getCount(dat []byte) (int64, int) {
	for i, d := range dat {
		if d == '\n' {
			j, err := strconv.ParseInt(string(dat[:i]), 10, 64)
			check(err)
			i++
			return j, i
		}
	}
	return -1, -1
}

func main() {
	var pathVar string
	flag.StringVar(&pathVar, "stack-path", "stacks.txt", "Path to the pancake stacks")
	flag.Parse()

	dat, err := ioutil.ReadFile(pathVar)
	check(err)

	findFlips(dat)
}
