package database

import "sync"

/*
	I am using the mutex to be able to block writing in memory at the same time, this will prevent the wprecon from crashing.
	In short: "self.mutex.Lock()" blocks writing and "self.mutex.Unlock()" releases writing. Doing so when something is already being written he will wait to finish before he can write other information.
*/

var (
	// Memory :: The saved information needs to be changed or searched anywhere in the code, so I am exporting the variable that instantiates NewMemory.
	Memory = NewMemory()
)

type memory struct {
	stringx   map[string]string
	intx      map[string]int
	slice     map[string][]string
	boolx     map[string]bool
	mapstring map[string]map[string]string
	mutex     sync.RWMutex
}

// NewMemory :: To avoid having to use sqlite or json, I chose to temporarily save the target's information in memory.
func NewMemory() *memory {
	database := &memory{
		stringx:   map[string]string{},
		intx:      map[string]int{},
		boolx:     map[string]bool{},
		slice:     map[string][]string{},
		mapstring: map[string]map[string]string{},
	}

	database.mapstring["HTTP Plugins Versions"] = map[string]string{}
	database.mapstring["HTTP Themes Versions"] = map[string]string{}

	return database
}

func (self *memory) SetString(key, value string) {
	self.mutex.Lock()
	self.stringx[key] = value
	defer self.mutex.Unlock()
}

func (self *memory) SetSlice(key string, value []string) {
	self.mutex.Lock()
	self.slice[key] = value
	defer self.mutex.Unlock()
}

func (self *memory) SetInt(key string, value int) {
	self.mutex.Lock()
	self.intx[key] = value
	defer self.mutex.Unlock()
}

func (self *memory) SetBool(key string, value bool) {
	self.mutex.Lock()
	self.boolx[key] = value
	defer self.mutex.Unlock()
}

func (self *memory) SetMapString(key string, value map[string]string) {
	self.mutex.Lock()
	self.mapstring[key] = value
	defer self.mutex.Unlock()
}

func (self *memory) SetMapMapString(key, key2, value string) {
	self.mutex.Lock()
	self.mapstring[key][key2] = value
	defer self.mutex.Unlock()
}

func (self *memory) AddInString(key, value string) {
	self.mutex.Lock()
	self.stringx[key] += value
	defer self.mutex.Unlock()
}

func (self *memory) AddInSlice(key, value string) {
	self.mutex.Lock()
	self.slice[key] = append(self.slice[key], value)
	defer self.mutex.Unlock()
}

func (self *memory) AddCalcInt(key string, value int) {
	self.mutex.Lock()
	self.intx[key] = self.intx[key] + value
	defer self.mutex.Unlock()
}

func (self *memory) AddInt(key string) {
	self.mutex.Lock()
	self.intx[key]++
	defer self.mutex.Unlock()
}

func (self *memory) GetString(key string) string {
	defer self.mutex.Unlock()
	self.mutex.Lock()
	return self.stringx[key]
}
func (self *memory) GetSlice(key string) []string {
	defer self.mutex.Unlock()
	self.mutex.Lock()
	return self.slice[key]
}
func (self *memory) GetInt(key string) int {
	defer self.mutex.Unlock()
	self.mutex.Lock()
	return self.intx[key]
}
func (self *memory) GetBool(key string) bool {
	defer self.mutex.Unlock()
	self.mutex.Lock()
	return self.boolx[key]
}
func (self *memory) GetMapString(key string) map[string]string {
	defer self.mutex.Unlock()
	self.mutex.Lock()
	return self.mapstring[key]
}
func (self *memory) GetMapMapString(key, key2 string) string {
	defer self.mutex.Unlock()
	self.mutex.Lock()
	return self.mapstring[key][key2]
}
