package memory

import "sync"

type _Memory struct {
	stringx   map[string]string
	intx      map[string]int
	slice     map[string][]string
	boolx     map[string]bool
	mapstring map[string]map[string]string
	mutex     sync.RWMutex
}