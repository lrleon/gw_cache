package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
	"time"

	gw_cache "github.com/lrleon/gateway_cache/v2"
	"github.com/lrleon/gateway_cache/v2/models"
)

type Processor struct {
}

// receive the key defined as int and return a string
func (p *Processor) ToMapKey(someValue int) (string, error) {
	return fmt.Sprint(someValue), nil
}

// receive the value that will be used as a key and return a string, that will be used as a value
func (p *Processor) CacheMissSolver(someValue int, _ ...interface{}) (string, *models.RequestError) {
	time.Sleep(time.Second * 1)
	return fmt.Sprintf("%d processed", someValue), nil
}

type GOBTransformer[T any] struct{}

// BytesToValue decodes GOB bytes into a value of type T.
func (t GOBTransformer[T]) BytesToValue(data []byte) (T, error) {
	var v T
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&v)
	return v, err
}

// ValueToBytes encodes a value of type T into GOB bytes.
func (t GOBTransformer[T]) ValueToBytes(val T) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(val)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Val size = %.4f MB\n", float64(len(buf.Bytes()))/(1024*1024))

	return buf.Bytes(), nil
}

func main() {
	//create the cache
	capacity := 10
	capFactor := 0.6
	ttl := time.Minute * 5
	p := &Processor{}

	cache := gw_cache.New[int, string](
		capacity,
		capFactor,
		ttl,
		ttl,
		p,
	)

	// compute and set the value
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func(i int, wg *sync.WaitGroup) {
				start := time.Now()
				value, err := cache.RetrieveFromCacheOrCompute(i)
				fmt.Printf("with i:%d v:%s, e:%v\n elapsed time %s\n", i, value, err, time.Since(start).Abs())
				wg.Done()
			}(i, &wg)
		}
	}
	wg.Wait()
	metrics := cache.Metrics()
	fmt.Printf("cache.Misses(): %v\n", metrics.Misses())
	fmt.Printf("cache.Hits(): %v\n", metrics.Hits())

	cacheCompression := gw_cache.NewWithCompression[int, string](
		capacity,
		capFactor,
		ttl,
		ttl,
		p,
		GOBTransformer[string]{},
	)

	// compute and set the value
	wgCompression := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			wgCompression.Add(1)
			go func(i int, wgCompression *sync.WaitGroup) {
				start := time.Now()
				value, err := cacheCompression.RetrieveFromCacheOrCompute(i)
				fmt.Printf("with i:%d v:%s, e:%v\n elapsed time %s\n", i, value, err, time.Since(start).Abs())
				wgCompression.Done()
			}(i, &wgCompression)
		}
	}
	wgCompression.Wait()
	compressedMetrics := cacheCompression.Metrics()
	fmt.Printf("cache.Misses(): %v\n", compressedMetrics.Misses())
	fmt.Printf("cache.Hits(): %v\n", compressedMetrics.Hits())

}
