package db

import (
	"github.com/nknorg/nkn/v2/config"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type IStore interface {
	Put(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Has(key []byte) (bool, error)
	Delete(key []byte) error
	NewBatch() error
	BatchPut(key []byte, value []byte) error
	BatchDelete(key []byte) error
	BatchCommit() error
	Compact() error
	Close() error
	NewIterator(prefix []byte) IIterator
}

type LevelDBStore struct {
	db    *leveldb.DB // LevelDB instance
	batch *leveldb.Batch
}

// used to compute the size of bloom filter bits array .
// too small will lead to high false positive rate.
const BITSPERKEY = 10

func NewLevelDBStore(file string) (*LevelDBStore, error) {
	// default Options
	o := opt.Options{
		NoSync:                 false,
		OpenFilesCacheCapacity: int(config.Parameters.DBFilesCacheCapacity),
		Filter:                 filter.NewBloomFilter(BITSPERKEY),
	}

	db, err := leveldb.OpenFile(file, &o)

	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(file, nil)
	}

	if err != nil {
		return nil, err
	}

	return &LevelDBStore{
		db:    db,
		batch: nil,
	}, nil
}

func (self *LevelDBStore) Put(key []byte, value []byte) error {
	return self.db.Put(key, value, nil)
}

func (self *LevelDBStore) Get(key []byte) ([]byte, error) {
	dat, err := self.db.Get(key, nil)
	return dat, err
}

func (self *LevelDBStore) Has(key []byte) (bool, error) {
	return self.db.Has(key, nil)
}

func (self *LevelDBStore) Delete(key []byte) error {
	return self.db.Delete(key, nil)
}

func (self *LevelDBStore) NewBatch() error {
	self.batch = new(leveldb.Batch)
	return nil
}

func (self *LevelDBStore) BatchPut(key []byte, value []byte) error {
	self.batch.Put(key, value)
	return nil
}

func (self *LevelDBStore) BatchDelete(key []byte) error {
	self.batch.Delete(key)
	return nil
}

func (self *LevelDBStore) BatchCommit() error {
	err := self.db.Write(self.batch, nil)
	if err != nil {
		return err
	}
	return nil
}

func (self *LevelDBStore) Compact() error {
	r := util.Range{
		Start: nil,
		Limit: nil,
	}
	return self.db.CompactRange(r)
}

func (self *LevelDBStore) Close() error {
	err := self.db.Close()
	return err
}

func (self *LevelDBStore) NewIterator(prefix []byte) IIterator {

	iter := self.db.NewIterator(util.BytesPrefix(prefix), nil)

	return &Iterator{
		iter: iter,
	}
}
