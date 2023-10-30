package utils

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Signer interface {
	SignUrl(ctx context.Context, objectName string) (string, error)
}

func ReplaceWithSignedUrl(ctx context.Context, signer Signer, image *string) {
	if image == nil {
		return
	}
	SetSignedUrl(ctx, signer, image, *image)
}

func SetSignedUrl(ctx context.Context, signer Signer, target *string, path string) {
	if path == "" {
		return
	}
	imageUrl, err := signer.SignUrl(ctx, path)
	if err != nil {
		logrus.Errorf("could not sign url for user target: %v", err)
	} else {
		logrus.Tracef("Signed url for target: path=%s, url=%s", path, imageUrl)
		*target = imageUrl
	}
}
