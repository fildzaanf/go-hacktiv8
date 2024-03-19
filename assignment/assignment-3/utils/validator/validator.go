package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"time"
)

func IsDataEmpty(data ...interface{}) error {
	for _, value := range data {
		switch v := value.(type) {
		case string:
			if v == "" {
				return errors.New("string data is empty")
			}
		case int:
			if v == 0 {
				return errors.New("integer data is empty")
			}
		case time.Time:
			if v.IsZero() {
				return errors.New("time data is empty")
			}
		case []interface{}:
			if len(v) == 0 {
				return errors.New("slice data is empty")
			}
		case []string:
			if len(v) == 0 {
				return errors.New("slice of strings is empty")
			}
		case []int:
			if len(v) == 0 {
				return errors.New("slice of integers is empty")
			}
		default:
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				slice := reflect.ValueOf(v)
				if slice.Len() == 0 {
					return errors.New("slice data is empty")
				}
			} else {
				return fmt.Errorf("unsupported data type: %T", v)
			}
		}
	}
	return nil
}

func IsEmailValid(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func IsMinLengthValid(data string, minLength int) error {
	if len(data) < minLength {
		return fmt.Errorf("minimum length is %d characters", minLength)
	}
	return nil
}

func IsMaxLengthValid(data string, maxLength int) error {
	if len(data) > maxLength {
		return fmt.Errorf("maximum length is %d characters", maxLength)
	}
	return nil
}
