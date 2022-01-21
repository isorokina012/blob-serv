package requests

import (
	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
)

func NewDeleteBlobRequest(r *http.Request) (DeleteBlobRequest, error) {
	id := chi.URLParam(r, "id")
	result := DeleteBlobRequest{Id: id}
	return result, result.validateDeleteBlob()
}

func (r *DeleteBlobRequest) validateDeleteBlob() error {
	return validation.Errors{
		"id": validation.Validate(&r.Id, validation.Required),
	}.Filter()
}

type DeleteBlobRequest struct {
	Id string `url:"id"'`
}
