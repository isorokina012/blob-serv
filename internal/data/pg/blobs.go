package pg

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/blob-serv/internal/data"
)

const blobTableName = "blobs"

func NewBlobQ(db *pgdb.DB) data.BlobQ {
	return &blobQ{
		db:  db,
		sql: squirrel.Select("*").From(blobTableName),
	}
}

type blobQ struct {
	db  *pgdb.DB
	sql squirrel.SelectBuilder
}

func (q *blobQ) New() data.BlobQ {
	return NewBlobQ(q.db)
}

func (q *blobQ) Get() (*data.Blob, error) {
	var result data.Blob
	err := q.db.Get(&result, q.sql)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func (q *blobQ) Select() ([]data.Blob, error) {
	var result []data.Blob
	err := q.db.Select(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (q *blobQ) Insert(value data.Blob) (data.Blob, error) {
	clauses := structs.Map(value)
	var result data.Blob
	stmt := squirrel.Insert(blobTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		return data.Blob{}, err
	}
	return result, nil
}

func (q *blobQ) DeleteByID(id string) error {
	stmt := squirrel.Delete(blobTableName).Where(squirrel.Eq{"id": id})
	err := q.db.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}

//func (q *blobQ) DeleteByID() error {
//	squirrel.DeleteByID()
//	err := q.db.Exec(q.sql.DeleteByID(blobTableName))
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (q *blobQ) FilteredById(ids ...string) data.BlobQ {
	q.sql = q.sql.Where(squirrel.Eq{"id": ids})
	return q
}

func (q *blobQ) FilteredByUserId(ids ...string) data.BlobQ {
	q.sql = q.sql.Where(squirrel.Eq{"user_id": ids})
	return q
}

func (q *blobQ) Page(pageParams pgdb.OffsetPageParams) data.BlobQ {
	q.sql = pageParams.ApplyTo(q.sql, "id")
	return q
}
