package utils_usecases

import (
	"encoding/base32"
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomSecret(length int) (string, error) {
	randomBytes := make([]byte, length)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}
	secret := base32.StdEncoding.EncodeToString(randomBytes)
	return secret, nil
}

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
