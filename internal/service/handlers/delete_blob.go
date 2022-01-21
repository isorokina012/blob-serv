package handlers

import (
	"fmt"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokend/blob-serv/internal/service/requests"
	"net/http"
)

func DeleteBlob(w http.ResponseWriter, r *http.Request) {
	log := Log(r)
	request, err := requests.NewDeleteBlobRequest(r)
	if err != nil {
		log.WithError(err).Info("Bad request")
		return
	}
	fmt.Println(request.Id)
	db := DB(r)
	err = db.DeleteByID(request.Id)
	if err != nil {
		log.WithError(err).Errorf("Failed to delete blob")
		ape.Render(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return
}
