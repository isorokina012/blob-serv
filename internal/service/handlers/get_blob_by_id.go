package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/blob-serv/internal/service/models"
	"gitlab.com/tokend/blob-serv/internal/service/requests"
	"net/http"
)

func GetBlob(w http.ResponseWriter, r *http.Request) {
	log := Log(r)
	request, err := requests.NewGetBlobIdRequest(r)
	if err != nil {
		log.WithError(err).Info("Bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	db := DB(r)
	blob, err := db.FilteredById(request.Id).Get()
	if err != nil {
		log.WithError(err).Info("Failed to get blob with this Id from DB")
		ape.Render(w, problems.InternalError())
		return
	}

	if blob == nil {
		log.WithError(err).Error("Blob with this Id doesn't exist")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, models.NewBlobModel(*blob))
}
