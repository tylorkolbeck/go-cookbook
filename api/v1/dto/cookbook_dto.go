package dto

import "github.com/go-playground/validator/v10"

// CookbookBase contains fields shared between different cookbook requests.
type CookbookBase struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=250"`
	Public      bool   `json:"public"`
}

// AddCookbookRequest is for creating new cookbooks.
type AddCookbookRequest struct {
	CookbookBase
}

// UpdateCookbookRequest is for updating existing cookbooks.
// This could have additional or fewer fields than the CreateCookbookRequest.
type UpdateCookbookRequest struct {
	CookbookBase
}

// Validate performs shared validation logic on CookbookBase.
func (cb *CookbookBase) Validate() error {
	validate := validator.New()
	return validate.Struct(cb)
}

func (r AddCookbookRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r UpdateCookbookRequest) Validate() error {
	validate := validator.New()
	// Custom validation logic for UpdateCookbookRequest
	return validate.Struct(r)
}
