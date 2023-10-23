package db

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"shift/internal/entity"
	"time"

	"cloud.google.com/go/storage"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/iterator"
)

type GoogleBucketDB struct {
	bucketName string
	config     *jwt.Config
	client     *storage.Client
}

func NewGoogleBucketDB(ctx context.Context) entity.BucketDB {
	client, err := createStorageClient(ctx)
	if err != nil {
		logrus.Error(err)
		return new(DummyBucketDB)
	}

	sakeyFile := os.Getenv("GCP_SAKEY_FILE")
	saKey, err := os.ReadFile(sakeyFile)
	if err != nil {
		logrus.Errorf("reading service account file: %v", err)
		return new(DummyBucketDB)
	}

	config, err := google.JWTConfigFromJSON(saKey)
	if err != nil {
		logrus.Error(err)
		return new(DummyBucketDB)
	}

	return &GoogleBucketDB{
		bucketName: os.Getenv("GCP_BUCKET_NAME"),
		client:     client,
		config:     config,
	}
}

func createStorageClient(ctx context.Context) (*storage.Client, error) {
	if err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", os.Getenv("GCP_SAKEY_FILE")); err != nil {
		return nil, err
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (b *GoogleBucketDB) UploadObject(ctx context.Context, objectName string, file io.Reader) error {
	bucket := b.client.Bucket(b.bucketName)
	obj := bucket.Object(objectName)

	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, file); err != nil {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}
	return nil
}

func (b *GoogleBucketDB) DeleteObject(ctx context.Context, objectName string) error {
	bucket := b.client.Bucket(b.bucketName)
	obj := bucket.Object(objectName)
	return obj.Delete(ctx)
}

func (b *GoogleBucketDB) DeleteObjects(ctx context.Context, folderName string) error {
	bucket := b.client.Bucket(b.bucketName)
	it := bucket.Objects(ctx, &storage.Query{Prefix: folderName})
	for {
		attrs, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return err
		}

		if err := b.DeleteObject(ctx, attrs.Name); err != nil {
			return err
		}
	}
	return nil
}

func (b *GoogleBucketDB) SignUrl(_ context.Context, objectName string) (string, error) {
	method := "GET"
	expires := time.Now().Add(time.Second * 60)

	url, err := storage.SignedURL(b.bucketName, objectName, &storage.SignedURLOptions{
		GoogleAccessID: b.config.Email,
		PrivateKey:     b.config.PrivateKey,
		Method:         method,
		Expires:        expires,
	})
	if err != nil {
		return "", fmt.Errorf("generating signed URL: %w", err)
	}
	return url, nil
}

type DummyBucketDB struct{}

func (b *DummyBucketDB) UploadObject(_ context.Context, _ string, _ io.Reader) error {
	logrus.Trace("called UploadObject on dummy bucket")
	return nil
}
func (b *DummyBucketDB) DeleteObject(_ context.Context, _ string) error {
	logrus.Trace("called DeleteObject on dummy bucket")
	return nil
}
func (b *DummyBucketDB) DeleteObjects(_ context.Context, _ string) error {
	logrus.Trace("called DeleteObject on dummy bucket")
	return nil
}
func (b *DummyBucketDB) SignUrl(_ context.Context, _ string) (string, error) {
	logrus.Trace("called SignUrl on dummy bucket")
	return "", nil
}
