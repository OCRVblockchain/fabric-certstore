package certstore

import (
	"crypto/sha256"
	//b64 "encoding/base64"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/syndtr/goleveldb/leveldb"
)

const dbFile = ".cert_cache"

type storage struct {
	sync.RWMutex
	db             *leveldb.DB
	cache          *cache.Cache
	transientCache *cache.Cache
}

var certCache = func() *storage {
	db, err := leveldb.OpenFile(dbFile, nil)
	if err != nil {
		logger.Fatal(err)
	}
	return &storage{db: db, cache: cache.New(5*time.Minute, 10*time.Minute), transientCache: cache.New(1*time.Minute, 2*time.Minute)}
}()

func storeCert(cert []byte) {
	certCache.Lock()
	defer certCache.Unlock()
	id := makeID(cert)
	if err := certCache.db.Put(id, cert, nil); err != nil {
		logger.Warn(err)
	}
	certCache.cache.Add(string(id), cert, cache.DefaultExpiration)
}

func getCert(id []byte) ([]byte, error) {
	certCache.RLock()
	defer certCache.RUnlock()

	if cert, exists := certCache.cache.Get(string(id)); exists {
		return cert.([]byte), nil
	}

	exists, err := certCache.db.Has(id, nil)
	if err != nil {
		return nil, err
	}
	if !exists {
		if cert, exists := certCache.transientCache.Get(string(id)); exists {
			return cert.([]byte), nil
		}
		return nil, nil
	}
	return certCache.db.Get(id, nil)
}

func makeID(cert []byte) []byte {
	id := sha256.Sum256(cert)
	return id[:]
}
