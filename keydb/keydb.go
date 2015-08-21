package keydb

import "regexp"

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

func (db *DB) Scan(pattern string) []interface{} {
	matches := make([]interface{}, 0, len(db.data))
	for key, v := range db.data {
		if match, _ := regexp.MatchString(pattern, key); match {
			matches = append(matches, v)
		}
	}
	return matches
}
