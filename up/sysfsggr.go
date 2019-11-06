package main

// GeigerCounter object
type SysfsCounter struct {
}

func New(debug bool) *SysfsCounter {
	sc := SysfsCounter{}
	return &sc
}

func (sc *SysfsCounter) Init(pinid int) error {

	return nil
}
