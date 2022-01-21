package models

import (
	"encoding/json"
	"gitlab.com/tokend/blob-serv/internal/data"
	"gitlab.com/tokend/blob-serv/resources"
)

func NewBlobResponseModel(blob data.Blob) resources.BlobResponse {
	result := resources.BlobResponse{
		Data: NewBlobModel(blob),
	}
	return result
}

func NewBlobListModel(blobs []data.Blob) []resources.Blob {
	result := make([]resources.Blob, len(blobs))
	for i, blob := range blobs {
		result[i] = NewBlobModel(blob)
	}
	return result
}

func NewBlobModel(blob data.Blob) resources.Blob {
	result := resources.Blob{
		Key: resources.Key{ID: blob.Id, Type: resources.KYC_DATA},
		Attributes: resources.BlobAttributes{
			Value: json.RawMessage(blob.Blob),
		},
	}
	return result
}
