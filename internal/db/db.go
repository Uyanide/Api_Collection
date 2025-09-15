package db

import (
	"errors"

	"github.com/dgraph-io/badger/v4"
)

var dbInstance DBIntf

var (
	ErrKeyNotFound = errors.New("key not found")
)

type DBIntf interface {
	Open(dbPath string) error
	Close() error

	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) error
	Exists(key string) (bool, error)
	Keys() ([]string, error)
}

func GetDB() DBIntf {
	if dbInstance == nil {
		dbInstance = &badgerDB{}
	}
	return dbInstance
}

type badgerDB struct {
	dbPath string
	db     *badger.DB
}

func (b *badgerDB) Open(dbPath string) error {
	if b.db != nil {
		return nil
	}

	b.dbPath = dbPath
	opts := badger.DefaultOptions(b.dbPath).WithLogger(nil)
	db, err := badger.Open(opts)
	if err != nil {
		return err
	}
	b.db = db
	return nil
}

func (b *badgerDB) Close() error {
	if b.db == nil {
		return nil
	}
	err := b.db.Close()
	b.db = nil
	return err
}

func (b *badgerDB) Get(key string) (string, error) {
	if b.db == nil {
		return "", errors.New("database not opened")
	}
	var value string
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return ErrKeyNotFound
			}
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		value = string(val)
		return nil
	})
	if err != nil {
		return "", err
	}
	return value, nil
}

func (b *badgerDB) Set(key, value string) error {
	if b.db == nil {
		return errors.New("database not opened")
	}
	err := b.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
	return err
}

func (b *badgerDB) Delete(key string) error {
	if b.db == nil {
		return errors.New("database not opened")
	}
	err := b.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
	return err
}

func (b *badgerDB) Exists(key string) (bool, error) {
	if b.db == nil {
		return false, errors.New("database not opened")
	}
	err := b.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(key))
		return err
	})
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (b *badgerDB) Keys() ([]string, error) {
	if b.db == nil {
		return nil, errors.New("database not opened")
	}
	var keys []string
	err := b.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			keys = append(keys, string(k))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return keys, nil
}
