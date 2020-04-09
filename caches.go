package styx

import (
	"encoding/binary"
	"fmt"

	badger "github.com/dgraph-io/badger/v2"
)

// A valueCache associates uint64 ids with a value.
// The Value struct definition is generated by protobuf.
type valueCache struct {
	values map[iri]string
	ids    map[string]iri
}

// newValueCache creates a new valueCache
func newValueCache() valueCache {
	return valueCache{
		values: map[iri]string{},
		ids:    map[string]iri{},
	}
}

// Commit writes the contents of the value cache to badger
func (values *valueCache) Commit(id iri, value string, db *badger.DB, t *badger.Txn) (txn *badger.Txn, err error) {
	txn = t
	v := []byte(value)
	idKey := make([]byte, 1+len(id))
	idKey[0] = IDToValuePrefix
	copy(idKey[1:], id)
	txn, err = setSafe(idKey, v, txn, db)
	if err != nil {
		return
	}

	valueKey := make([]byte, 1+len(v))
	valueKey[0] = ValueToIDPrefix
	copy(valueKey[1:], v)
	txn, err = setSafe(valueKey, idKey[1:], t, db)

	return
}

// Set a Value in the valueCache
func (values *valueCache) Set(id iri, term string) {
	values.values[id] = term
	values.ids[term] = id
}

// Get a Value from the valueCache
func (values *valueCache) GetValue(id iri, txn *badger.Txn) (string, error) {
	value, has := values.values[id]
	if has {
		return value, nil
	}

	key := make([]byte, 1+len(id))
	key[0] = IDToValuePrefix
	copy(key[1:], id)
	item, err := txn.Get(key)
	if err != nil {
		return "", err
	}

	var val []byte
	val, err = item.ValueCopy(nil)
	if err != nil {
		return "", err
	}

	value = string(val)
	values.values[id] = value
	values.ids[value] = id
	return value, nil
}

func (values *valueCache) GetID(value string, txn *badger.Txn) (id iri, err error) {
	var has bool
	id, has = values.ids[value]
	if has {
		return
	}

	v := []byte(value)
	key := make([]byte, len(v)+1)
	key[0] = ValueToIDPrefix
	copy(key[1:], v)
	item, err := txn.Get(key)
	if err != nil {
		return "", err
	}

	err = item.Value(func(val []byte) error { id = iri(val); return nil })
	if err != nil {
		return
	}

	values.ids[value] = id
	values.values[id] = value
	return
}

type unaryCache map[Term]*[6]uint32

// newUnaryCache creates a new IndexCache
func newUnaryCache() unaryCache {
	return unaryCache{}
}

// getUnaryIndex returns the 6-tuple of counts from an item
func getUnaryIndex(item *badger.Item) (*[6]uint32, error) {
	result := &[6]uint32{}
	return result, item.Value(func(val []byte) error {
		if len(val) != 24 {
			return fmt.Errorf("Unexpected index value: %v", val)
		}
		for i := 0; i < 6; i++ {
			result[i] = binary.BigEndian.Uint32(val[i*4 : (i+1)*4])
		}
		return nil
	})
}

func (uc unaryCache) getIndex(a Term, txn *badger.Txn) (*[6]uint32, error) {
	index, has := uc[a]
	if has {
		return index, nil
	}

	key := assembleKey(UnaryPrefix, false, a)
	item, err := txn.Get(key)
	if err != nil {
		return nil, err
	}

	uc[a] = &[6]uint32{}
	err = item.Value(func(val []byte) error {
		if len(val) != 24 {
			return fmt.Errorf("Unexpected index value: %v", val)
		}
		for i := 0; i < 6; i++ {
			uc[a][i] = binary.BigEndian.Uint32(val[i*4 : (i+1)*4])
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return uc[a], nil
}

func (uc unaryCache) Get(p Permutation, a Term, txn *badger.Txn) (uint32, error) {
	index, err := uc.getIndex(a, txn)
	if err == badger.ErrKeyNotFound {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return index[p], nil
}

func (uc unaryCache) Increment(p Permutation, a Term, txn *badger.Txn) error {
	index, err := uc.getIndex(a, txn)
	if err == badger.ErrKeyNotFound {
		index = &[6]uint32{}
		uc[a] = index
	} else if err != nil {
		return err
	}
	uc[a][p]++
	return nil
}

func (uc unaryCache) Decrement(p Permutation, a Term, txn *badger.Txn) error {
	index, err := uc.getIndex(a, txn)
	if err == badger.ErrKeyNotFound {
		index = &[6]uint32{}
		uc[a] = index
	} else if err != nil {
		return err
	}
	if uc[a][p] > 0 {
		uc[a][p]--
	}
	return nil
}

// Commit writes the contents of the index map to badger
func (uc unaryCache) Commit(db *badger.DB, t *badger.Txn) (txn *badger.Txn, err error) {
	txn = t
	for term, index := range uc {
		key := assembleKey(UnaryPrefix, false, term)
		zero := true
		for _, c := range index {
			if c > 0 {
				zero = false
				break
			}
		}
		if zero {
			txn, err = deleteSafe(key, txn, db)
			if err == badger.ErrKeyNotFound {
				return txn, nil
			}
		} else {
			val := make([]byte, 24)
			for i, c := range index {
				binary.BigEndian.PutUint32(val[i*4:(i+1)*4], c)
			}
			txn, err = setSafe(key, val, txn, db)
			if err != nil {
				return
			}
		}
	}
	return
}

type binaryCache map[string]uint32

// newBinaryCache returns a new binary cache
func newBinaryCache() binaryCache {
	return binaryCache{}
}

func (bc binaryCache) Get(p Permutation, a, b Term, txn *badger.Txn) (uint32, error) {
	key := assembleKey(BinaryPrefixes[p], false, a, b)
	s := string(key)
	count, has := bc[s]
	if has {
		return count, nil
	}

	item, err := txn.Get(key)
	if err == badger.ErrKeyNotFound {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	err = item.Value(func(val []byte) error {
		bc[s] = binary.BigEndian.Uint32(val)
		return nil
	})
	if err != nil {
		return 0, err
	}
	return bc[s], nil
}

func (bc binaryCache) delta(p Permutation, a, b Term, increment bool, uc unaryCache, txn *badger.Txn) error {
	key := assembleKey(BinaryPrefixes[p], false, a, b)
	s := string(key)
	_, has := bc[s]
	if has {
		if increment {
			bc[s]++
			if bc[s] == 1 {
				return uc.Increment(p, a, txn)
			}
		} else if bc[s] > 0 {
			bc[s]--
			if bc[s] == 0 {
				return uc.Decrement(p, a, txn)
			}
		} else {
			// ??
		}
		return nil
	}

	item, err := txn.Get(key)
	if err == badger.ErrKeyNotFound && increment { // Hmm
		bc[s] = 1
		return uc.Increment(p, a, txn)
	} else if err != nil {
		return err
	}

	err = item.Value(func(val []byte) error {
		if len(val) != 4 {
			return fmt.Errorf("Unexpected binary value: %v", val)
		}
		bc[s] = binary.BigEndian.Uint32(val)
		return nil
	})
	if err != nil {
		return err
	}

	if increment {
		bc[s]++
	} else if bc[s] == 1 {
		bc[s] = 0
		return uc.Decrement(p, a, txn)
	} else if bc[s] > 0 {
		bc[s]--
	} else {
		// ???
	}
	return nil
}

func (bc binaryCache) Increment(p Permutation, a, b Term, uc unaryCache, txn *badger.Txn) error {
	return bc.delta(p, a, b, true, uc, txn)
}

func (bc binaryCache) Decrement(p Permutation, a, b Term, uc unaryCache, txn *badger.Txn) error {
	return bc.delta(p, a, b, false, uc, txn)
}

// Commit writes the contents of the index map to badger
func (bc binaryCache) Commit(db *badger.DB, t *badger.Txn) (txn *badger.Txn, err error) {
	txn = t
	for key, count := range bc {
		if count == 0 {
			txn, err = deleteSafe([]byte(key), txn, db)
			if err == badger.ErrKeyNotFound {
			} else if err != nil {
				return
			}
		} else {
			val := make([]byte, 4)
			binary.BigEndian.PutUint32(val, count)
			txn, err = setSafe([]byte(key), val, txn, db)
			if err != nil {
				return
			}
		}
	}
	return
}