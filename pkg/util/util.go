package util

import (
	"errors"
	"fmt"
	"net"
	"os"
)

func Flatten(items []interface{}) ([]string, error) {
	return doFlatten([]string{}, items)
}

func doFlatten(result []string, items interface{}) ([]string, error) {
	var err error

	switch v := items.(type) {
	case string:
		result = append(result, v)
	case []string:
		result = append(result, v...)
	case []interface{}:
		for _, item := range v {
			result, err = doFlatten(result, item)
			if err != nil {
				return nil, err
			}
		}
	default:
		return nil, errors.New(fmt.Sprintf("Flatten does not support %T", v))
	}

	return result, err
}

func StringSliceContains(slice []string, s string) bool {
	for _, element := range slice {
		if s == element {
			return true
		}
	}
	return false
}

func FileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		return false
	}
	return true
}

func GetRandomPort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	defer listener.Close()
	if err != nil {
		return 0, err
	}

	return listener.Addr().(*net.TCPAddr).Port, nil
}
