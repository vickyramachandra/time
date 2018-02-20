package time

import (
	"fmt"
	"strings"
	"time"
)

// JSON unmarshal fails if the time format isn't RFC3339.
// Hence writing custom Time with implementations of MarshalJSON and UnMarshalJSON

// Ctime - type wrapper for time.Time
type Ctime struct {
	time.Time
}

var nilTime = time.Time{}
var justDateFormat = "2006-01-02"

// UnmarshalJSON - impl
func (c *Ctime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	// Ignore null, like in the main JSON package.
	if string(b) == "null" {
		return nil
	}
	t, err := time.Parse(justDateFormat, s)
	if err != nil {
		return err
	}
	*c = Ctime{t}
	return nil
}

// MarshalJSON - impl
func (c *Ctime) MarshalJSON() ([]byte, error) {
	if c.Time == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", c.Time.Format(justDateFormat))), nil
}

// IsSet - Util method to check if a valid time is set
func (c *Ctime) IsSet() bool {
	return c.Time != nilTime
}
