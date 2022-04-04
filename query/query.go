package query

import "reflect"

// the query interface
type Query interface {
	// return true is query is ready.
	IsReady() bool
	// describe the query
	Desc() string
	// get query model
	GetModelType() reflect.Type
}
