package xmap

import (
	"github.com/leaxoy/x-go/xcontainer/set"
)

func KeySet(m interface{}) set.AnyHashSet {
	return set.NewAnyHashSet(Keys(m))
}

func ValueSet(m interface{}) set.AnyHashSet {
	return set.NewAnyHashSet(Values(m))
}
