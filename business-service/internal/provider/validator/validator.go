package validator

import (
	"log"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/lucasd-coder/fast-feet/pkg/val"
)

type Validation struct {
	once     sync.Once
	validate *validator.Validate
}

func NewValidation() *Validation {
	return &Validation{}
}

func (v *Validation) ValidateStruct(s interface{}) error {
	v.lazyInit()

	if err := v.validate.RegisterValidation("pattern", val.Pattern); err != nil {
		log.Fatal(err)
	}

	if err := v.validate.RegisterValidation("isCPF", val.TagIsCPF); err != nil {
		log.Fatal(err)
	}

	if err := v.validate.RegisterValidation("rfc3339", val.DateTime); err != nil {
		log.Fatal(err)
	}

	if err := v.validate.RegisterValidation("objectID", val.ObjectID); err != nil {
		log.Fatal(err)
	}

	return v.validate.Struct(s)
}

func (v *Validation) lazyInit() {
	v.once.Do(func() {
		v.validate = validator.New()
	})
}
