package bookvalidations

import (
	"net/http"

	"github.com/gujral1997/book-list/models"
	"github.com/thedevsaddam/govalidator"
)

type BookValidations struct{}

func (b BookValidations) CreateBookValidation(r *http.Request, book models.Book) (map[string]interface{}, bool) {
	var isValid bool = false
	rules := govalidator.MapData{
		"year":   []string{"required"},
		"title":  []string{"required"},
		"author": []string{"required"},
	}
	opts := govalidator.Options{
		Request: r,
		Rules:   rules,
		Data:    &book,
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()
	err := map[string]interface{}{"validationError": e}
	if len(e) == 0 {
		isValid = true
	}
	return err, isValid
}

func (b BookValidations) UpdateBookValidation(r *http.Request, book models.Book) (map[string]interface{}, bool) {
	var isValid bool = false
	rules := govalidator.MapData{
		"id":     []string{"required"},
		"year":   []string{"required"},
		"title":  []string{"required"},
		"author": []string{"required"},
	}
	opts := govalidator.Options{
		Request: r,
		Rules:   rules,
		Data:    &book,
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()
	err := map[string]interface{}{"validationError": e}
	if len(e) == 0 {
		isValid = true
	}
	return err, isValid
}
