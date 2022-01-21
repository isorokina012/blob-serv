package data

import (
	"github.com/jmoiron/sqlx/types"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type BlobQ interface {
	New() BlobQ

	Get() (*Blob, error)
	Select() ([]Blob, error)
	Insert(value Blob) (Blob, error)
	DeleteByID(id string) error

	FilteredById(ids ...string) BlobQ
	FilteredByUserId(ids ...string) BlobQ
	Page(pageParams pgdb.OffsetPageParams) BlobQ
}

type Blob struct {
	Id     string         `db:"id" structs:"id"`
	UserId string         `db:"user_id" structs:"user_id"`
	Blob   types.JSONText `db:"blob" structs:"blob"`
}
