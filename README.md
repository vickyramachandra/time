# time
Custom time implementation for unMarshalling different time formats in JSON to time.Time

Deserialization from json to struct (which consists of a time.Time property) fails when the time format in the json isn't 
RFC3339 ≈ "2006-01-02T15:04:05Z07:00"

Hence use this wrapper implementation of time to support other time formats in your json payload
