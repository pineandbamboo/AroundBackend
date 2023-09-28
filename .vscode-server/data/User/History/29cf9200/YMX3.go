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