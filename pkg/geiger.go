package geiger

type GeigerCounter interface {
	Init(pinid int) error
	PerMinute() int64
	PerHour() int64
	Radiation() float64
	Close() error
}
