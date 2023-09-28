package kgorm

type Versioner interface {
	Version() Msec
}
