package database

import "sync"

/*
	I am using the mutex to be able to block writing in memory at the same time, this will prevent the wprecon from crashing.
	In short: "db.mutex.Lock()" blocks writing and "db.mutex.Unlock()" releases writing. Doing so when something is already being written he will wait to finish before he can write other information.
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

func (db *memory) SetString(key, value string) {
	db.mutex.Lock()
	db.stringx[key] = value
	defer db.mutex.Unlock()
}

func (db *memory) SetSlice(key string, value []string) {
	db.mutex.Lock()
	db.slice[key] = value
	defer db.mutex.Unlock()
}

func (db *memory) SetInt(key string, value int) {
	db.mutex.Lock()
	db.intx[key] = value
	defer db.mutex.Unlock()
}

func (db *memory) SetBool(key string, value bool) {
	db.mutex.Lock()
	db.boolx[key] = value
	defer db.mutex.Unlock()
}

func (db *memory) SetMapString(key string, value map[string]string) {
	db.mutex.Lock()
	db.mapstring[key] = value
	defer db.mutex.Unlock()
}

func (db *memory) SetMapMapString(key, key2, value string) {
	db.mutex.Lock()
	db.mapstring[key][key2] = value
	defer db.mutex.Unlock()
}

func (db *memory) AddInString(key, value string) {
	db.mutex.Lock()
	db.stringx[key] += value
	defer db.mutex.Unlock()
}

func (db *memory) AddInSlice(key, value string) {
	db.mutex.Lock()
	db.slice[key] = append(db.slice[key], value)
	defer db.mutex.Unlock()
}

func (db *memory) AddCalcInt(key string, value int) {
	db.mutex.Lock()
	db.intx[key] = db.intx[key] + value
	defer db.mutex.Unlock()
}

func (db *memory) AddInt(key string) {
	db.mutex.Lock()
	db.intx[key]++
	defer db.mutex.Unlock()
}

func (db *memory) GetString(key string) string {
	defer db.mutex.Unlock()
	db.mutex.Lock()
	return db.stringx[key]
}
func (db *memory) GetSlice(key string) []string {
	defer db.mutex.Unlock()
	db.mutex.Lock()
	return db.slice[key]
}
func (db *memory) GetInt(key string) int {
	defer db.mutex.Unlock()
	db.mutex.Lock()
	return db.intx[key]
}
func (db *memory) GetBool(key string) bool {
	defer db.mutex.Unlock()
	db.mutex.Lock()
	return db.boolx[key]
}
func (db *memory) GetMapString(key string) map[string]string {
	defer db.mutex.Unlock()
	db.mutex.Lock()
	return db.mapstring[key]
}
func (db *memory) GetMapMapString(key, key2 string) string {
	defer db.mutex.Unlock()
	db.mutex.Lock()
	return db.mapstring[key][key2]
}
