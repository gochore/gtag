// Code generated by gtag. DO NOT EDIT.
// See: https://github.com/wolfogre/gtag

//go:generate gtag -types Empty,User,UserName -tags bson,json .
package regular

import (
	"reflect"
	"strings"
)

var (
	valueOfEmpty = Empty{}
	typeOfEmpty  = reflect.TypeOf(valueOfEmpty)
)

// EmptyTags indicate tags of type Empty
type EmptyTags struct {
}

// Values return all tags of Empty as slice
func (t *EmptyTags) Values() []string {
	return []string{}
}

// Tags return specified tags of Empty
func (*Empty) Tags(tag string, convert ...func(string) string) EmptyTags {
	conv := func(in string) string { return strings.TrimSpace(strings.Split(in, ",")[0]) }
	if len(convert) > 0 {
		conv = convert[0]
	}
	if conv == nil {
		conv = func(in string) string { return in }
	}
	return EmptyTags{}
}

// TagsBson is alias of Tags("bson")
func (*Empty) TagsBson() EmptyTags {
	var v *Empty
	return v.Tags("bson")
}

// TagsJson is alias of Tags("json")
func (*Empty) TagsJson() EmptyTags {
	var v *Empty
	return v.Tags("json")
}
