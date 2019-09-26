package indiction

func NewIndicCalculator(name string) interface{} {

	if name == "ema" {
		ema := NewEMA()
		return &ema
	}

	return nil
}

func myfunc() {
	// ema := NewIndicCalculator("ema")
	// if e, ok := ema.(EMA); ok {

	// }
}
