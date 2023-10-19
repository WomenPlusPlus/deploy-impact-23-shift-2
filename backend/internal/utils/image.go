package utils

import (
	"context"
	"github.com/sirupsen/logrus"
	"shift/internal/entity"
)

func ReplaceWithImageUrl(ctx context.Context, bucket entity.BucketDB, image *string) {
	if image == nil || *image == "" {
		return
	}
	imageUrl, err := bucket.SignUrl(ctx, *image)
	if err != nil {
		logrus.Errorf("could not sign url for user image: %v", err)
	} else {
		logrus.Tracef("Signed url for image: path=%s, url=%s", *image, imageUrl)
		*image = imageUrl
	}
}
