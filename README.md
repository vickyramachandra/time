# time
Custom time implementation for unmarshalling of time format with only date ```"2006-01-02"``` in json to ```time.Time```

Deserialization from json to struct (which consists of a ```time.Time``` property) fails when the time format in the json isn't ```RFC3339 ≈ "2006-01-02T15:04:05Z07:00"```

Hence use this simple implementation to support time format with only date in your json payload.

```go
package main

import (
	"encoding/json"
	"fmt"

	mytime "github.com/vickyramachandra/time"
)

type MyStruct struct {
	Time *mytime.Ctime `json:"time"`
}

func main() {
	// json payload with RFC3339 format
	dateFormat := `{"time":"2018-02-18"}`

	var myStruct = new(MyStruct)
	// this will invoke the unmarshal impl defined in this package
	err := json.Unmarshal([]byte(dateFormat), &myStruct)
	if err != nil {
		fmt.Println(err)
		return
	}
	// prints 2018 February 18
	fmt.Println(myStruct.Time.Date())
}
```
