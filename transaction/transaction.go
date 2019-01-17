package transaction

import (
	"log"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

// These are the TxKind list
const (
	DrawDai TxKind = iota
	ChangeDaiToEth
	LockEth
	deleted1
	deleted2
	ChangeEthToDai
	WipeDai
	FreeEth
	FreeSellWipe
	DrawSellLock
)

var (
	db *bolt.DB
)

// TxKind is the kind of transaction
type TxKind int

// Tx represents a Transaction of a certain kind
type Tx struct {
	Hash string
	Kind TxKind
}

// GetTx retrieves the current transaction
func GetTx() (currentTx *Tx, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Transaction"))
		bkind := b.Get([]byte("TxKind"))
		bhash := b.Get([]byte("TxHash"))
		if len(bhash) == 0 && len(bkind) == 0 {
			return nil
		}
		ikind, err := strconv.Atoi(string(bkind))
		if err != nil {
			return errors.Wrap(err, "cannot parse TxKind from Bolt DB")
		}
		currentTx = &Tx{string(bhash), TxKind(ikind)}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "cannot retrieve Tx from Bolt DB")
	}
	return currentTx, nil
}

// SaveTx saves the current transaction
func SaveTx(currentTx *Tx) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Transaction"))
		if currentTx == nil {
			err := b.Delete([]byte("TxKind"))
			if err != nil {
				return errors.Wrap(err, "cannot save TxKind to Bolt DB")
			}
			err = b.Delete([]byte("TxHash"))
			if err != nil {
				return errors.Wrap(err, "cannot save TxKind to Bolt DB")
			}
			return nil
		}

		err := b.Put([]byte("TxKind"), []byte(strconv.Itoa(int(currentTx.Kind))))
		if err != nil {
			return errors.Wrap(err, "cannot save TxKind to Bolt DB")
		}
		err = b.Put([]byte("TxHash"), []byte(currentTx.Hash))
		if err != nil {
			return errors.Wrap(err, "cannot save TxHash to Bolt DB")
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "cannot save Tx to Bolt DB")
	}
	return nil
}

func init() {
	var err error
	db, err = bolt.Open("transaction.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatalf("cannot open Bolt DB, error: %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Transaction"))
		if err != nil {
			return errors.Wrap(err, "cannot create / get bucket")
		}
		return nil
	})
	if err != nil {
		log.Fatalf("cannot get bucket, error: %v", err)
	}
}
