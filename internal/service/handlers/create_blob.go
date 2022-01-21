package handlers

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"

	"github.com/jmoiron/sqlx/types"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/blob-serv/internal/data"
	"gitlab.com/tokend/blob-serv/internal/service/models"
	"gitlab.com/tokend/blob-serv/internal/service/requests"
)

func CreateBlob(w http.ResponseWriter, r *http.Request) {
	log := Log(r)
	request, err := requests.NewCreateBlobRequest(r)
	if err != nil {
		log.WithError(err).Error("Bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	h := sha1.New()
	h.Write(request.Data.Attributes.Value)
	db := DB(r)
	existingBlob, err := db.FilteredById(hex.EncodeToString(h.Sum(nil))).Get()
	if existingBlob != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	blob := data.Blob{
		Id:     hex.EncodeToString(h.Sum(nil)),
		UserId: string(request.Data.Attributes.Value),
		Blob:   types.JSONText(request.Data.Attributes.Value),
	}
	blob, err = db.Insert(blob)
	if err != nil {
		log.WithError(err).Info("failed to create blob")
		ape.Render(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	ape.Render(w, models.NewBlobModel(blob))
	return
}
