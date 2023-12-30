package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	hFormat = "15:04"
)

type Interval struct {
	From time.Time
	To   time.Time
}

func (i *Interval) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%s-%s", i.From.Format(hFormat), i.To.Format(hFormat)))
}

func (i *Interval) UnmarshalJSON(bytes []byte) error {
	var interval string

	err := json.Unmarshal(bytes, &interval)
	if err != nil {
		return err
	}

	fromTo := strings.Split(interval, "-")
	i.From, err = time.Parse(hFormat, fromTo[0])
	if err != nil {
		return err
	}
	i.To, err = time.Parse(hFormat, fromTo[1])
	if err != nil {
		return err
	}
	return nil
}
