package handlers

import (
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Post struct {
	Title   string `json:"title" valid:"required"`
	Content string `json:"content" valid:"required"`
}

// Error is a general purpose error view model.
type Error struct {
	Errors []interface{} `json:"errors"`
}

// NewError helps create an Error with multiple error values without the need to manually declare a []interface{} for the Errors field.
func NewError(errs ...interface{}) *Error {
	return &Error{
		Errors: errs,
	}
}

// ValidationErrors receives a govalidator error and properly formats it into a viewmodel.Error
func ValidationErrors(err error) *Error {
	errs := make([]interface{}, 0)
	for k, v := range govalidator.ErrorsByField(err) {
		errs = append(errs, fmt.Sprintf("%s: %s", k, v))
	}
	return NewError(errs...)
}

// BindValid binds the request's body to a vm and then validates the vm through govalidator.
func BindValid(c *gin.Context, vm interface{}) bool {
	if err := c.Bind(&vm); err != nil {
		c.JSON(http.StatusBadRequest, NewError(errors.Wrap(err, "failed to parse the request's body").Error()))
		return false
	}

	if _, err := govalidator.ValidateStruct(vm); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, ValidationErrors(err))
		return false
	}

	return true
}
