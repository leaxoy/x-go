package sequtil

import (
	"encoding/hex"
	"math/rand"
	"strings"
)

func RandomStringID() (string, error) {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString(buf)), nil
}

func MustRandomStringID() string {
	id, err := RandomStringID()
	if err != nil {
		panic(err)
	}
	return id
}

var GlobalID string

func init() {
	GlobalID = MustRandomStringID()
}
