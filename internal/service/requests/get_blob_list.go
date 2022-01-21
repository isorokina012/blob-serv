package requests

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
	"net/http"
)

func NewGetBlobListRequest(r *http.Request) (GetBlobListRequest, error) {
	request := GetBlobListRequest{}
	err := urlval.DecodeSilently(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}
	return request, nil
}

type GetBlobListRequest struct {
	pgdb.OffsetPageParams
	FilterUserId []string `filter:"user_id"`
}
