package gw_cache

import (
	"encoding/json"

	"github.com/lrleon/gateway_cache/v2/models"
)

// ProcessorI is the interface used to map the key and get the value in case it is missing.
//
// # ToMapKey: is used to convert the input to a string key
//
// # CacheMissSolver: is used to call the upstream services
//
// K represents the input's type to get a value, this will be used as a key
// T represents the value's type  itself, this will be used as a value
//
//go:generate mockery --name ProcessorI --with-expecter=true --filename=processor_mock.go
type ProcessorI[K, T any] interface {
	ToMapKey(K) (string, error)
	CacheMissSolver(K, ...interface{}) (T, *models.RequestError) //we will leave the pre process logic for this function
}

// CompressorI is the interface that wraps the basic Compress and Decompress methods.
//
// # Compress is used to compress the input
//
// # Decompress is used to decompress the input
//
//go:generate mockery --name CompressorI --with-expecter=true --filename=compressor_mock.go
type CompressorI interface {
	Compress([]byte) ([]byte, error)
	Decompress([]byte) ([]byte, error)
}

// TransformerI is the interface that wraps the basic BytesToValue and ValueToBytes methods.
//
// # BytesToValue is used to convert the input to a value
//
// # ValueToBytes is used to convert the value to a byte array
//
// T represents the value's type
//
//go:generate mockery --name TransformerI --with-expecter=true --filename=transformer_mock.go
type TransformerI[T any] interface {
	BytesToValue([]byte) (T, error)
	ValueToBytes(T) ([]byte, error)
}

// Reporter is the interface that wraps the basic ReportMiss and ReportHit methods
//
// # ReportMiss is used to report a cache miss
//
// # ReportHit is used to report a cache hit
//
//go:generate mockery --name Reporter --with-expecter=true --filename=reporter_mock.go
type Reporter interface {
	ReportMiss()
	ReportHit()
}

type DefaultTransformer[T any] struct{}

func (_ *DefaultTransformer[T]) BytesToValue(in []byte) (T, error) {
	var out T
	err := json.Unmarshal(in, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

func (_ *DefaultTransformer[T]) ValueToBytes(in T) ([]byte, error) {

	return json.Marshal(in)
}
