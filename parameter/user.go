package parameter

type User struct {
	Name string `json:"name" validate:"required,min=3"`
}
