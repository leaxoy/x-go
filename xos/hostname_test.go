package xos_test

import (
	"testing"

	"github.com/leaxoy/x-go/xos"
)

func TestHostname(t *testing.T) {
	host := xos.Hostname()
	t.Log(host)
}
