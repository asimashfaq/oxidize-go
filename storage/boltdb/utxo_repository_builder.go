package boltdb

import (
	"log"

	"github.com/tclchiam/oxidize-go/blockchain/utxo"
	"github.com/tclchiam/oxidize-go/storage"
)

type utxoRepoBuilder struct {
	path        string
	name        string
	withCache   bool
	withLogger  bool
	withMetrics bool
}

func UtxoBuilder(name string) storage.UtxoRepositoryBuilder {
	return &utxoRepoBuilder{name: name, path: "./"}
}

func (b *utxoRepoBuilder) WithPath(path string) storage.UtxoRepositoryBuilder {
	return &utxoRepoBuilder{
		path:        path,
		name:        b.name,
		withCache:   b.withCache,
		withLogger:  b.withLogger,
		withMetrics: b.withMetrics,
	}
}

func (b *utxoRepoBuilder) WithCache() storage.UtxoRepositoryBuilder {
	return &utxoRepoBuilder{
		path:        b.path,
		name:        b.name,
		withCache:   true,
		withLogger:  b.withLogger,
		withMetrics: b.withMetrics,
	}
}

func (b *utxoRepoBuilder) WithLogger() storage.UtxoRepositoryBuilder {
	return &utxoRepoBuilder{
		path:        b.path,
		name:        b.name,
		withCache:   b.withCache,
		withLogger:  true,
		withMetrics: b.withMetrics,
	}
}

func (b *utxoRepoBuilder) WithMetrics() storage.UtxoRepositoryBuilder {
	return &utxoRepoBuilder{
		path:        b.path,
		name:        b.name,
		withCache:   b.withCache,
		withLogger:  b.withLogger,
		withMetrics: true,
	}
}

func (b *utxoRepoBuilder) Build() utxo.Repository {
	repository, err := NewUtxoRepository(b.path, b.name)
	if err != nil {
		log.Panic(err)
	}

	/*if b.withCache {
		repository = storage.WrapWithCache(repository)
	}
	if b.withMetrics {
		repository = storage.WrapWithMetrics(repository)
	}*/
	if b.withLogger {
		repository = storage.WrapUtxoWithLogger(repository)
	}

	return repository
}
