# time
Custom time implementation for unmarshalling of time format with only date - ```"2006-01-02"``` in json to time.Time

Deserialization from json to struct (which consists of a time.Time property) fails when the time format in the json isn't 
```RFC3339 ≈ "2006-01-02T15:04:05Z07:00"```

Hence use this wrapper implementation to support time format with only date in your json payload.
