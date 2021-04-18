package pintu

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func MaxProfit() int {
	content := ReadFile("gistfile1.txt")
	s := strings.Split(content, " ")
	var profit, tmp int
	prev, _ := strconv.Atoi(s[0])
	for _, v := range s {
		n, _ := strconv.Atoi(v)
		if n <= prev {
			if tmp > profit {
				profit = tmp
			}
			tmp = 0
		} else {
			tmp = tmp + (n - prev)
		}
		prev = n
	}
	if tmp > profit {
		profit = tmp
	}
	return profit
}

func ReadFile(filename string) string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
