package kgorm

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Msec time.Time

func (p Msec) Milli() int64 {
	return time.Time(p).UnixMilli()
}

func (p Msec) MarshalJSON() ([]byte, error) {
	msec := time.Time(p)
	if msec.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(msec.UnixMilli())
}

func (p *Msec) UnmarshalJSON(b []byte) error {
	var msec int64

	err := json.Unmarshal(b, &msec)
	if err == nil {
		*p = Msec(time.UnixMilli(msec))
	}

	return err
}

func (p Msec) Value() (driver.Value, error) {
	msec := time.Time(p)
	if msec.IsZero() {
		return nil, nil
	}

	return time.Time(p), nil
}
