package keydb

type DB struct {
	data map[string]interface{}
}

func NewDB() *DB {
	db := &DB{}
	db.data = make(map[string]interface{})
	return db
}

func (db *DB) Set(key string, value interface{}) bool {
	_, exists := db.data[key]
	db.data[key] = value
	return exists
}

func (db *DB) Get(key string) (interface{}, bool) {
	v, ok := db.data[key]
	return v, ok
}
