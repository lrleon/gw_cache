package gw_cache

import "errors"

var (
	ErrCantCastOutput                 = errors.New("can't cast output to desired type")
	ErrEntryExpired                   = errors.New("entry expired")
	ErrEntryAvailableState            = errors.New("entry is in available state")
	ErrEntryComputingState            = errors.New("entry is in computing state")
	ErrLRUComputing                   = errors.New("LRU entry is in COMPUTING state. This could be a bug or a cache misconfiguration")
	ErrNumOfEntriesBiggerThanCapacity = errors.New("number of entries in the cache is greater than given capacity")
)
