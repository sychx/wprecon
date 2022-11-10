package memory

/*
   I am using the mutex to be able to block writing in memory at the same time, model will prevent the wprecon from crashing.
   In short: "model.mutex.Lock()" blocks writing and "model.mutex.Unlock()" releases writing. Doing so when something is already being written he will wait to finish before he can write other information.
*/

var (
    // Memory :: The saved information needs to be changed or searched anywhere in the code, so I am exporting the variable that instantiates NewMemory.
    Memory = NewMemory()
)

// NewMemory :: To avoid having to use sqlite or json, I chose to temporarily save the target's information in _Memory.
func NewMemory() *_Memory {
    database := &_Memory{
        stringx:   map[string]string{},
        intx:      map[string]int{},
        boolx:     map[string]bool{},
        slice:     map[string][]string{},
        mapstring: map[string]map[string]string{},
    }

    return database
}

func (model *_Memory) SetString(key, value string) {
    model.mutex.Lock()
    model.stringx[key] = value
    model.mutex.Unlock()
}

func (model *_Memory) SetSlice(key string, value []string) {
    model.mutex.Lock()
    model.slice[key] = value
    model.mutex.Unlock()
}

func (model *_Memory) SetInt(key string, value int) {
    model.mutex.Lock()
    model.intx[key] = value
    model.mutex.Unlock()
}

func (model *_Memory) SetBool(key string, value bool) {
    model.mutex.Lock()
    model.boolx[key] = value
    model.mutex.Unlock()
}

func (model *_Memory) SetMapString(key string, value map[string]string) {
    model.mutex.Lock()
    model.mapstring[key] = value
    model.mutex.Unlock()
}

func (model *_Memory) SetMapMapString(key, key2, value string) {
    model.mutex.Lock()
    model.mapstring[key][key2] = value
    model.mutex.Unlock()
}

func (model *_Memory) AddInString(key, value string) {
    model.mutex.Lock()
    model.stringx[key] += value
    model.mutex.Unlock()
}

func (model *_Memory) AddInSlice(key, value string) {
    model.mutex.Lock()
    model.slice[key] = append(model.slice[key], value)
    model.mutex.Unlock()
}

func (model *_Memory) AddCalcInt(key string, value int) {
    model.mutex.Lock()
    model.intx[key] = model.intx[key] + value
    model.mutex.Unlock()
}

func (model *_Memory) AddInt(key string) {
    model.mutex.Lock()
    model.intx[key]++
    model.mutex.Unlock()
}

func (model *_Memory) GetString(key string) string {
    defer model.mutex.Unlock()
    model.mutex.Lock()
    return model.stringx[key]
}
func (model *_Memory) GetSlice(key string) []string {
    defer model.mutex.Unlock()
    model.mutex.Lock()
    return model.slice[key]
}
func (model *_Memory) GetInt(key string) int {
    defer model.mutex.Unlock()
    model.mutex.Lock()
    return model.intx[key]
}
func (model *_Memory) GetBool(key string) bool {
    defer model.mutex.Unlock()
    model.mutex.Lock()
    return model.boolx[key]
}
func (model *_Memory) GetMapString(key string) map[string]string {
    defer model.mutex.Unlock()
    model.mutex.Lock()
    return model.mapstring[key]
}
func (model *_Memory) GetMapMapString(key, key2 string) string {
    defer model.mutex.Unlock()
    model.mutex.Lock()
    return model.mapstring[key][key2]
}
