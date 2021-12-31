package pipedrive

import "time"

const (
	DateLayout     = "2006-01-02"
	DateTimeLayout = "2006-01-02 15:04:05"
)

type Timestamp struct {
	time.Time
}

func (t Timestamp) String() string {
	return t.Time.String()
}

func (t Timestamp) Format() string {
	return t.Time.Format(DateLayout)
}

func (t Timestamp) FormatFull() string {
	return t.Time.Format(DateTimeLayout)
}
