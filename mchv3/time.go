package mchv3

import (
	"encoding/json"
	"time"
)

const Fmt = time.RFC3339

type Time time.Time

func (jt *Time) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t := new(time.Time)
	t1, err := time.Parse(Fmt, string(data))
	t = &t1
	if err != nil {
		return err
	}
	*jt = (Time)(*t)
	return nil
}

func (jt Time) MarshalJSON() ([]byte, error) {
	return json.Marshal((*time.Time)(&jt).Format(Fmt))
}
