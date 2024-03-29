package proxy

type AbstractEmployee interface {
	DoWork()
	CheckWork()
	GetLaidOff() bool
	SetLaidOff(LaidOff bool)
}
