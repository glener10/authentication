package utils_usecases

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomCode() string {
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)
	var digits []int
	for i := 0; i < 6; i++ {
		digit := randGen.Intn(10)
		digits = append(digits, digit)
	}
	var result string
	for _, digit := range digits {
		result += strconv.Itoa(digit)
	}
	return result
}

func GetExpirationTime() time.Time {
	currentTime := time.Now()
	expirationTime := currentTime.Add(5 * time.Minute)
	return expirationTime
}
