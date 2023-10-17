package entity

import (
	"context"
	"io"
)

type BucketDB interface {
	UploadObject(ctx context.Context, objectName string, file io.Reader) error
	DeleteObject(ctx context.Context, objectName string) error
	DeleteObjects(ctx context.Context, objectName string) error
	SignUrl(ctx context.Context, objectName string) (string, error)
}
