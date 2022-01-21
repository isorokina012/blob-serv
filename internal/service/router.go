package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokend/blob-serv/internal/data/pg"
	"gitlab.com/tokend/blob-serv/internal/service/handlers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxDB(pg.NewBlobQ(s.db)),
		),
	)

	r.Route("/blobs", func(r chi.Router) {
		r.Post("/", handlers.CreateBlob)
		r.Get("/", handlers.GetBlobList)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetBlob)
			r.Delete("/", handlers.DeleteBlob)
		})
	})

	return r
}
