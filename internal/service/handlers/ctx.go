package handlers

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/blob-serv/internal/data"
	"net/http"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	dbCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxDB(entry data.BlobQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, dbCtxKey, entry)
	}
}

func DB(r *http.Request) data.BlobQ {
	return r.Context().Value(dbCtxKey).(data.BlobQ).New()
}
