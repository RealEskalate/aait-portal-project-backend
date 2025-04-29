package entities

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

// StringArray is a custom type for handling string arrays in PostgreSQL
type StringArray []string

func (a StringArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}

	// Convert array to PostgreSQL array format
	result := "{"
	for i, tag := range a {
		if i > 0 {
			result += ","
		}
		escapedTag := strings.ReplaceAll(tag, "\"", "\\\"")
		result += "\"" + escapedTag + "\""
	}
	result += "}"
	return result, nil
}

func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = StringArray{}
		return nil
	}

	switch v := value.(type) {
	case []byte:
		// Remove the curly braces and split by comma
		str := string(v)
		str = strings.Trim(str, "{}")
		if str == "" {
			*a = StringArray{}
			return nil
		}

		tags := strings.Split(str, ",")
		result := make(StringArray, len(tags))
		for i, tag := range tags {
			tag = strings.Trim(tag, "\"")
			tag = strings.ReplaceAll(tag, "\\\"", "\"")
			result[i] = tag
		}
		*a = result
		return nil
	case string:
		return a.Scan([]byte(v))
	default:
		return fmt.Errorf("cannot scan %T into StringArray", value)
	}
}

type Problem struct {
	gorm.Model
	Title      string      `json:"title"`
	Tags       StringArray `json:"tags" gorm:"type:text[]"`
	Difficulty string      `json:"difficulty"`
	Track      string      `json:"track"`
	Link       string      `json:"link"`
	PostedAt   time.Time   `json:"posted_at"`
}
