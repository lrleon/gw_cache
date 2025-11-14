/*
This is an implementation of a cache system for the Go language.

It allows you to cache any type of data, as long as it is serializable.

It is based on the idea of a processor, which is the one that will be in charge of generating the key and solving the cache miss.
also uses a map under the hood to store the data, it is recommended give a cap factor big enough to avoid resize.

Here is an example of the usage of this package to implement a cache that
will use an int as a key and a string as a value:

	package main

	import (

		"fmt"
		"time"

		gw_cache "github.com/lrleon/gateway_cache/v2"
		"github.com/lrleon/gateway_cache/v2/models"

	)

	type Processor struct {}

	// receive the key defined as int and return a string
	func (p *Processor) ToMapKey(someValue int) (string, error) {
		return fmt.Sprint(someValue), nil
	}

	// receive the value that will be used as a key and return a string, that will be used as a value
	func (p *Processor) CacheMissSolver(someValue int) (string, *models.RequestError) {
		return fmt.Sprintf("%d processed", someValue), nil
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
			p,
		)

		//compute and set the value
		value, err := cache.RetrieveFromCacheOrCompute(3)
		fmt.Println(value, err) // 3 processed <nil>
	}

The above example is present in the main folder of this repository.
*/
package gw_cache
