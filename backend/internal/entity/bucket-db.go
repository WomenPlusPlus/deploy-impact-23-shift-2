package entity

import (
	"context"
	"os"
)

type BucketDB interface {
	UploadObject(ctx context.Context, objectName string, file *os.File) error
	DeleteObject(ctx context.Context, objectName string) error
	DeleteObjects(ctx context.Context, objectName string) error
	SignUrl(ctx context.Context, objectName string) (string, error)
}
