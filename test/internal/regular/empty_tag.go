// Code generated by gtag. DO NOT EDIT.
// See: https://github.com/wolfogre/gtag

//go:generate gtag -files ../../test/internal/regular/user.go,../../test/internal/regular/empty.go -types User,Empty
package regular

import "reflect"

var (
	valueOfEmpty = Empty{}
	typeOfEmpty  = reflect.TypeOf(valueOfEmpty)
)

type EmptyTags struct {
}

func (Empty) Tags(tag string) EmptyTags {
	return EmptyTags{}
}
