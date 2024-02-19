package util

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"math/big"
	"reflect"
	"strconv"
)

func GenerateRandomUsername(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	base64Encoded := base64.StdEncoding.EncodeToString(randomBytes)

	return base64Encoded, nil
}

func GenerateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var result string
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		result += string(charset[randomIndex.Int64()])
	}

	return result, nil
}

func ConvertStringsToUint64Array(strArray []string) ([]uint64, error) {
	var uintArray []uint64

	for _, str := range strArray {
		uintVal, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return nil, err
		}
		uintArray = append(uintArray, uintVal)
	}

	return uintArray, nil
}

func ConvertUint64ToStringsArray(arr []uint64) []string {
	strArr := make([]string, len(arr))
	for i, v := range arr {
		strArr[i] = strconv.FormatUint(v, 10)
	}

	return strArr
}

func ArrayIncludes(array []interface{}, value interface{}) (bool, error) {
	arrType := reflect.TypeOf(array)

	if arrType.Elem() != reflect.TypeOf(value) {
		return false, errors.New("wrong element type provided")
	}

	for _, unit := range array {
		if unit == value {
			return true, nil
		}
	}

	return false, nil
}

func Uint64Includes(unint64Array []uint64, value uint64) bool {
	for _, unit := range unint64Array {
		if unit == value {
			return true
		}
	}
	return false
}

func Remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
