package validator

import v "github.com/go-playground/validator/v10"

type Validator struct{ v *v.Validate }

func New() *Validator                      { return &Validator{v: v.New()} }
func (vd *Validator) Validate(i any) error { return vd.v.Struct(i) }
