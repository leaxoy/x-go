package xhash

import (
	"github.com/speps/go-hashids"
)

var g *hashids.HashID

func init() {
	d := hashids.NewData()
	g, _ = hashids.NewWithData(d)
}

func IntID(ints []int64) string {
	id, err := g.EncodeInt64(ints)
	if err != nil {
		return ""
	}
	return id
}

func StringID(hexStr string) string {
	id, err := g.EncodeHex(hexStr)
	if err != nil {
		return ""
	}
	return id
}

func IDString(id string) string {
	str, _ := g.DecodeHex(id)
	return str
}

func IDInt(id string) []int64 {
	ints, _ := g.DecodeInt64WithError(id)
	return ints
}
