package requests

import (
	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
)

type GetBlobIdRequest struct {
	Id string `url:"id"`
}

func NewGetBlobIdRequest(r *http.Request) (GetBlobIdRequest, error) {
	id := chi.URLParam(r, "id")
	result := GetBlobIdRequest{Id: id}
	return result, result.validateBlobList()
}

func (r *GetBlobIdRequest) validateBlobList() error {
	return validation.Errors{
		"id": validation.Validate(&r.Id, validation.Required),
	}.Filter()
}
