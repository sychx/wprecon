package database

import "sync"

/*
	I am using the mutex to be able to block writing in memory at the same time, this will prevent the wprecon from crashing.
	In short: "this.mutex.Lock()" blocks writing and "this.mutex.Unlock()" releases writing. Doing so when something is already being written he will wait to finish before he can write other information.
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

func (this *memory) SetString(key, value string) {
	this.mutex.Lock()
	this.stringx[key] = value
	defer this.mutex.Unlock()
}

func (this *memory) SetSlice(key string, value []string) {
	this.mutex.Lock()
	this.slice[key] = value
	defer this.mutex.Unlock()
}

func (this *memory) SetInt(key string, value int) {
	this.mutex.Lock()
	this.intx[key] = value
	defer this.mutex.Unlock()
}

func (this *memory) SetBool(key string, value bool) {
	this.mutex.Lock()
	this.boolx[key] = value
	defer this.mutex.Unlock()
}

func (this *memory) SetMapString(key string, value map[string]string) {
	this.mutex.Lock()
	this.mapstring[key] = value
	defer this.mutex.Unlock()
}

func (this *memory) SetMapMapString(key, key2, value string) {
	this.mutex.Lock()
	this.mapstring[key][key2] = value
	defer this.mutex.Unlock()
}

func (this *memory) AddInString(key, value string) {
	this.mutex.Lock()
	this.stringx[key] += value
	defer this.mutex.Unlock()
}

func (this *memory) AddInSlice(key, value string) {
	this.mutex.Lock()
	this.slice[key] = append(this.slice[key], value)
	defer this.mutex.Unlock()
}

func (this *memory) AddCalcInt(key string, value int) {
	this.mutex.Lock()
	this.intx[key] = this.intx[key] + value
	defer this.mutex.Unlock()
}

func (this *memory) AddInt(key string) {
	this.mutex.Lock()
	this.intx[key]++
	defer this.mutex.Unlock()
}

func (this *memory) GetString(key string) string {
	defer this.mutex.Unlock()
	this.mutex.Lock()
	return this.stringx[key]
}
func (this *memory) GetSlice(key string) []string {
	defer this.mutex.Unlock()
	this.mutex.Lock()
	return this.slice[key]
}
func (this *memory) GetInt(key string) int {
	defer this.mutex.Unlock()
	this.mutex.Lock()
	return this.intx[key]
}
func (this *memory) GetBool(key string) bool {
	defer this.mutex.Unlock()
	this.mutex.Lock()
	return this.boolx[key]
}
func (this *memory) GetMapString(key string) map[string]string {
	defer this.mutex.Unlock()
	this.mutex.Lock()
	return this.mapstring[key]
}
func (this *memory) GetMapMapString(key, key2 string) string {
	defer this.mutex.Unlock()
	this.mutex.Lock()
	return this.mapstring[key][key2]
}
