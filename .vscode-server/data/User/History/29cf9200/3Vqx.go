package backend

import (
	"context"
	"fmt"
	"io"

	"around/constants"

	"cloud.google.com/go/storage"
)

var (
	GCSBackend *GoogleCloudStorageBackend
)

type GoogleCloudStorageBackend struct {
	client *storage.Client
	bucket string
}

func (backend *ElasticsearchBackend) ReadFromES()
