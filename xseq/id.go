package xseq

import (
	"github.com/google/uuid"

	"github.com/leaxoy/x-go/xstrings"
)

var globalId string

func init() {
	var err error
	globalId, err = xstrings.FastRandom(32)
	if err != nil {
		globalId = UUID()
	}
}

func GlobalID() string {
	return globalId
}

func UUID() string {
	return uuid.New().String()
}
