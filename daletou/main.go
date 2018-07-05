package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func isExist(arr []string, target string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return true
		}
	}
	return false
}

func generateBalls(arr []string, ranges int, count int) {
	i := 0
	for {
		if i >= count {
			break
		}
		rand.Seed(time.Now().Unix())
		number := rand.Intn(ranges)
		numberStr := ""
		if number < 10 {
			numberStr = "0" + strconv.Itoa(number)
		} else {
			numberStr = strconv.Itoa(number)
		}
		if !isExist(arr, numberStr) {
			fmt.Print(numberStr + " ")
			arr[i] = numberStr
			i++
		}
	}
}
func main() {
	fmt.Println("大乐透即将开奖:")
	redBall := make([]string, 5)
	blueBall := make([]string, 2)
	generateBalls(redBall, 35, 5)
	fmt.Print("| ")
	generateBalls(blueBall, 12, 2)
	fmt.Println("")
	fmt.Println("恭喜中奖!!!")
	fmt.Println("")
	fmt.Println("双色球即将开奖:")
	redBall2 := make([]string, 6)
	blueBall2 := make([]string, 1)
	generateBalls(redBall2, 33, 6)
	fmt.Print("| ")
	generateBalls(blueBall2, 16, 1)
	fmt.Println("")
	fmt.Println("恭喜中奖!!!")

}
