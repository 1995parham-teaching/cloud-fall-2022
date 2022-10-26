package person

import (
	"errors"

	"github.com/1995parham-teaching/cloud-fall-2022/internal/model"
)

var ErrPersonNotFound = errors.New("person not found")

type Person interface {
	Get() ([]model.Person, error)
	GetByName(string) (model.Person, error)
	Save(model.Person) error
}
