package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/blob-serv/internal/service/models"
	"gitlab.com/tokend/blob-serv/internal/service/requests"
	"gitlab.com/tokend/blob-serv/resources"
	"net/http"
)

func GetBlobList(w http.ResponseWriter, r *http.Request) {
	log := Log(r)

	request, err := requests.NewGetBlobListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	db := DB(r).Page(request.OffsetPageParams)
	blobQ := db.New()

	blobs, err := blobQ.Select()
	if err != nil {
		log.WithError(err).Errorf("Failed to get blob list")
		ape.Render(w, err)
		return
	}
	response := resources.BlobListResponse{
		Data: models.NewBlobListModel(blobs),
	}
	ape.Render(w, response.Data)
	return
}
