// Code generated by gtag. DO NOT EDIT.
// See: https://github.com/wolfogre/gtag

//go:generate gtag -files ../../test/internal/regular/user.go,../../test/internal/regular/empty.go -types User,Empty
package regular

import "reflect"

var (
	valueOfUser = User{}
	typeOfUser  = reflect.TypeOf(valueOfUser)

	_                = valueOfUser.Id
	fieldOfUserId, _ = typeOfUser.FieldByName("Id")
	tagOfUserId      = fieldOfUserId.Tag

	_                  = valueOfUser.Name
	fieldOfUserName, _ = typeOfUser.FieldByName("Name")
	tagOfUserName      = fieldOfUserName.Tag

	_                   = valueOfUser.Email
	fieldOfUserEmail, _ = typeOfUser.FieldByName("Email")
	tagOfUserEmail      = fieldOfUserEmail.Tag

	_                 = valueOfUser.age
	fieldOfUserage, _ = typeOfUser.FieldByName("age")
	tagOfUserage      = fieldOfUserage.Tag
)

type UserTags struct {
	Id    string
	Name  string
	Email string
	age   string
}

func (User) Tags(tag string) UserTags {
	return UserTags{
		Id:    tagOfUserId.Get(tag),
		Name:  tagOfUserName.Get(tag),
		Email: tagOfUserEmail.Get(tag),
		age:   tagOfUserage.Get(tag),
	}
}
