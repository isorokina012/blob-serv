package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"gitlab.com/tokend/blob-serv/resources"
	"net/http"
)

func NewCreateBlobRequest(r *http.Request) (CreateBlobRequest, error) {
	request := CreateBlobRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}
	return request, request.validationCreateBlob()
}

func (r *CreateBlobRequest) validationCreateBlob() error {
	return validation.Errors{
		"/data/type": validation.Validate(&r.Data.Type, validation.Required),
	}.Filter()
}

type CreateBlobRequest struct {
	Data resources.Blob
}
