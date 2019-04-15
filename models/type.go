package models

import "errors"

type (
	ExportType int
)

const (
	UnknownType ExportType = iota
	Dump
	Sync
)

func (*ExportType) GetList() []ExportType {
	return []ExportType{Dump, Sync}
}

func (t ExportType) String() string {
	switch t {
	case Dump:
		return "dump"
	case Sync:
		return "sync"
	default:
		return UnknownStr
	}
}

func ToType(str string) (ExportType, error) {
	switch str {
	case Dump.String():
		return Dump, nil
	case Sync.String():
		return Sync, nil
	default:
		return UnknownType, errors.New("Unknown type string")
	}
}
