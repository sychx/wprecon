package database

import "sync"

/*
	I am using the mutex to be able to block writing in memory at the same time, this will prevent the wprecon from crashing.
	In short: "mutex.Lock ()" blocks writing and "mutex.Unlock ()" releases writing. Doing so when something is already being written he will wait to finish before he can write other information.
*/

var (
	// Memory :: The saved information needs to be changed or searched anywhere in the code, so I am exporting the variable that instantiates NewMemory.
	Memory = NewMemory()
	mutex  sync.Mutex
)

type memory struct {
	stringx   map[string]string
	intx      map[string]int
	slice     map[string][]string
	boolx     map[string]bool
	mapstring map[string]map[string]string
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
	mutex.Lock()
	db.stringx[key] = value
	mutex.Unlock()
}
func (db *memory) SetSlice(key string, value []string) {
	mutex.Lock()
	db.slice[key] = value
	mutex.Unlock()
}
func (db *memory) SetInt(key string, value int) {
	mutex.Lock()
	db.intx[key] = value
	mutex.Unlock()
}
func (db *memory) SetBool(key string, value bool) {
	mutex.Lock()
	db.boolx[key] = value
	mutex.Unlock()
}
func (db *memory) SetMapString(key string, value map[string]string) {
	mutex.Lock()
	db.mapstring[key] = value
	mutex.Unlock()
}
func (db *memory) SetMapMapString(key, key2, value string) {
	mutex.Lock()
	db.mapstring[key][key2] = value
	mutex.Unlock()
}

func (db *memory) AddInString(key, value string) {
	mutex.Lock()
	db.stringx[key] += value
	mutex.Unlock()
}
func (db *memory) AddInSlice(key, value string) {
	mutex.Lock()
	db.slice[key] = append(db.slice[key], value)
	mutex.Unlock()
}
func (db *memory) AddCalcInt(key string, value int) {
	mutex.Lock()
	db.intx[key] = db.intx[key] + value
	mutex.Unlock()
}
func (db *memory) AddInt(key string) {
	mutex.Lock()
	db.intx[key]++
	mutex.Unlock()
}

func (db *memory) GetString(key string) string  { return db.stringx[key] }
func (db *memory) GetSlice(key string) []string { return db.slice[key] }
func (db *memory) GetInt(key string) int        { return db.intx[key] }
func (db *memory) GetBool(key string) bool      { return db.boolx[key] }
func (db *memory) GetMapString(key string) map[string]string {
	return db.mapstring[key]
}
func (db *memory) GetMapMapString(key, key2 string) string {
	return db.mapstring[key][key2]
}
