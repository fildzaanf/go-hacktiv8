package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"time"
)

func IsDataEmpty(fields []string, data ...interface{}) error {
	if len(fields) != len(data) {
		return errors.New("column names and data length mismatch")
	}

	for i, value := range data {
		switch v := value.(type) {
		case string:
			if v == "" {
				return fmt.Errorf("%s is empty", fields[i])
			}
		case int:
			if v == 0 {
				return fmt.Errorf("%s is empty", fields[i])
			}
		case time.Time:
			if v.IsZero() {
				return fmt.Errorf("%s is empty", fields[i])
			}
		case []interface{}:
			if len(v) == 0 {
				return fmt.Errorf("%s is empty", fields[i])
			}
		case []string:
			if len(v) == 0 {
				return fmt.Errorf("%s is empty", fields[i])
			}
		case []int:
			if len(v) == 0 {
				return fmt.Errorf("%s is empty", fields[i])
			}
		default:
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				slice := reflect.ValueOf(v)
				if slice.Len() == 0 {
					return fmt.Errorf("%s is empty", fields[i])
				}
			} else {
				return fmt.Errorf("unsupported data type for %s: %T", fields[i], v)
			}
		}
	}
	return nil
}


func IsEmailValid(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New("invalid email format. example: emailname@gmail.com")
	}
	return nil
}


func IsMinLengthValid(minLength int, fields map[string]string) error {
    for fieldName, fieldValue := range fields {
        if len(fieldValue) < minLength {
            return fmt.Errorf("minimum length for field %s is %d characters", fieldName, minLength)
        }
    }
    return nil
}

func IsMaxLengthValid(maxLength int, fields map[string]string) error {
    for fieldName, fieldValue := range fields {
        if len(fieldValue) > maxLength {
            return fmt.Errorf("maximum length for field %s is %d characters",  fieldName, maxLength)
        }
    }
    return nil
}

