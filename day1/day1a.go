package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	sc := bufio.NewScanner(input)
	var str string
	var sum int64
	var value int64

	for sc.Scan() {
		str = sc.Text()
		var calibrationStr string
		for i := 0; i < len(str); i++ {
			if int(str[i]) >= 48 && int(str[i]) < 58 {
				calibrationStr += string(str[i])
			}
		}
		if len(calibrationStr) == 1 {
			calibrationStr += string(calibrationStr[0])
			value, _ = strconv.ParseInt(calibrationStr, 0, 64)
		} else {
			var tmp string = string(calibrationStr[0]) + string(calibrationStr[len(calibrationStr)-1])
			value, _ = strconv.ParseInt(tmp, 0, 64)
		}
		sum += value
	}
	fmt.Println(sum)
}
