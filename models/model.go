package models

type RequestError struct {
	Error    error
	Code     CodeStatus
	UserCode int
	// TODO: add UserInfo as a kind of opaque data pointer. In that way, the user can return to the invoker any kind of additional data
}

type EntryState int8
type CodeStatus int8
