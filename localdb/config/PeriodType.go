package config

type PeriodType string

const (
	Period_D PeriodType = "D"
	Period_W PeriodType = "W"
	Period_M PeriodType = "M"
	Period_Q PeriodType = "Q"
	Period_Y PeriodType = "Y"
)

func (this PeriodType) ToString() string {
	return string(this)
}
