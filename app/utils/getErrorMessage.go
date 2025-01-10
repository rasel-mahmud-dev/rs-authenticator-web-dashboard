package utils

import "fmt"

func Error(str1 string, rest ...any) error {
	err := fmt.Errorf(str1, rest...)
	return err
}
