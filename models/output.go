package models

import "errors"

type (
	Output int
)

const (
	UnknownOutput Output = iota
	Stout
	File
	AwsS3
	OtherDatabase
)

const UnknownStr = "unknown"

func (*Output) GetList() []Output {
	return []Output{Stout, File, AwsS3}
}

func (o Output) String() string {
	switch o {
	case Stout:
		return "stdout"
	case File:
		return "file"
	case AwsS3:
		return "aws-s3"
	case OtherDatabase:
		return "other-database"
	default:
		return UnknownStr
	}
}

func ToOutput(str string) (Output, error) {
	switch str {
	case Stout.String():
		return Stout, nil
	case File.String():
		return File, nil
	case AwsS3.String():
		return AwsS3, nil
	case OtherDatabase.String():
		return OtherDatabase, nil
	default:
		return UnknownOutput, errors.New("Unknown output type string")
	}
}
