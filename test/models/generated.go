// This file was generated by github.com/vektah/gqlgen, DO NOT EDIT

package models

type InnerInput struct {
	ID int
}
type InnerObject struct {
	ID int
}
type OuterInput struct {
	Inner InnerInput
}
type OuterObject struct {
	InnerID int
}