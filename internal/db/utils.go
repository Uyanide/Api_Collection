package db

import (
	"strconv"
)

// Return value for given key and create with defaultValue if key is not found
func GetOrCreate(db DBIntf, key string, defaultValue string) (string, error) {
	value, err := db.Get(key)
	// OK
	if err == nil {
		return value, nil
	}
	// Other errors
	if err != ErrKeyNotFound {
		return "", err
	}
	// Not found, create with defaultValue
	if err := db.Set(key, defaultValue); err != nil {
		return "", err
	}
	return defaultValue, nil
}

// Return value for given key, create with defaultValue if key is not found or not an integer
func GetOrCreateInt(db DBIntf, key string, defaultValue int64) (int64, error) {
	value, err := db.Get(key)
	if err == nil {
		numValue, convErr := strconv.ParseInt(value, 10, 64)
		// OK
		if convErr == nil {
			return numValue, nil
		}
		// Conversion error, reset to default
		if err := db.Set(key, strconv.FormatInt(defaultValue, 10)); err != nil {
			return 0, err
		}
		return defaultValue, nil
	}
	if err != ErrKeyNotFound {
		return 0, err
	}
	if err := db.Set(key, strconv.FormatInt(defaultValue, 10)); err != nil {
		return 0, err
	}
	return defaultValue, nil
}

func IncrementInt(db DBIntf, key string, defaultValue int64, increaseValue int64) (int64, error) {
	value, err := GetOrCreateInt(db, key, 0)
	if err != nil {
		return 0, err
	}
	value += increaseValue
	if err := db.Set(key, strconv.FormatInt(value, 10)); err != nil {
		return 0, err
	}
	return value, nil
}
