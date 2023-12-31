// internal/repository/wasabi_repository.go
package repository

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kaustubh-pandey-kp/storage-utility/constants"
	"github.com/kaustubh-pandey-kp/storage-utility/pkg/entity"
	"github.com/kaustubh-pandey-kp/storage-utility/pkg/logger"
	cachedRedis "github.com/kaustubh-pandey-kp/storage-utility/pkg/redis"
)

type WasabiRepository struct {
	client       *s3.S3
	configParams constants.WasabiConfigParams
	ctx          *context.Context
}

func NewWasabiRepository(configParams constants.WasabiConfigParams) (*WasabiRepository, error) {
	cfg := &aws.Config{
		Endpoint:    aws.String(configParams.WasabiEndpoint),
		Credentials: credentials.NewStaticCredentials(configParams.WasabiAccessKey, configParams.WasabiSecretKey, ""),
		Region:      aws.String(configParams.WasabiBucketRegion),
	}
	sess, err := session.NewSession(cfg)
	if err != nil {
		return nil, err
	}

	return &WasabiRepository{
		client:       s3.New(sess),
		configParams: configParams,
	}, nil
}

func (w *WasabiRepository) UploadArtifact(artifact *entity.Artifact) (err error) {
	if artifact == nil {
		return errors.New("nil artifact provided")
	}
	maxRetries := constants.FILE_UPLOAD_MAX_ATTEMPTS
	for attempt := 1; attempt < maxRetries; attempt++ {
		err = w.FileUploadRetry(attempt, artifact)
		if err != nil {
			cachedRedis.IncrementWasabiFailureCounter()
			logger.Errorf("Failed uploading file to wasabi - mobile in attempt #%d, bucket %s, %+v", attempt, w.configParams.WasabiBucketName, err)
		} else {
			logger.Errorf("Success uploading file to wasabi in attempt #%d", attempt)
			break
		}
	}

	if err != nil {
		logger.Infof("Failed uploading file to wasabi - mobile: Bucket %s, %+v", w.configParams.WasabiBucketName, err)
		return err
	}
	return nil
}

func (w *WasabiRepository) DownloadArtifact(artifactName string) (artifact *entity.Artifact, err error) {
	output, err := w.client.GetObjectWithContext(*w.ctx, &s3.GetObjectInput{
		Bucket: aws.String(w.configParams.WasabiBucketName),
		Key:    aws.String(artifactName),
	})
	if err != nil {
		return nil, err
	}
	defer output.Body.Close()

	var buffer bytes.Buffer

	_, err = buffer.ReadFrom(output.Body)

	if err != nil {
		logger.Infof("Error reading the object body %+v", err)
		return nil, err
	}

	content := buffer.Bytes()

	artifact = &entity.Artifact{
		Key:     artifactName,
		Content: content,
	}

	return artifact, nil
}

func (w *WasabiRepository) FileUploadRetry(attempt int, artifact *entity.Artifact) error {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(constants.FILE_UPLOAD_MAX_TIME_MILLISEC*(attempt+1))*time.Millisecond)
	defer cancel()

	_, err := w.client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:          aws.String(w.configParams.WasabiBucketName),
		Key:             aws.String(artifact.Key),
		Body:            bytes.NewReader(artifact.Content),
		ContentType:     aws.String(http.DetectContentType(artifact.Content)),
		ContentEncoding: aws.String("base64"),
		ContentLength:   aws.Int64(int64(len(artifact.Content))),
	})
	if err != nil {
		return err
	}
	return nil
}
