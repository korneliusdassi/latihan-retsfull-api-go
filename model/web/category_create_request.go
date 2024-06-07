package web

type CategoryCreateRequest struct {
	Id   int    `json:"id"`
	Name string `validate:"required,min=1,max=200" json:"name"`
}
