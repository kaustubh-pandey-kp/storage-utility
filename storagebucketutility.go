package mobilestoragebucketutility

import (
	"errors"
	"github.com/LambdatestIncPrivate/mobile-automation/mobile-storage-bucket-utility-initial_changes/pkg/entity"
	"github.com/LambdatestIncPrivate/mobile-automation/mobile-storage-bucket-utility-initial_changes/pkg/repository"
)

type StorageBucketUtility struct {
	wasabiRepo repository.WasabiRepository
}

func NewStorageBucketUtility(wasabiRepo *repository.WasabiRepository) *StorageBucketUtility {
	return &StorageBucketUtility{wasabiRepo: *wasabiRepo}
}

func (s *StorageBucketUtility) UploadArtifact(artifact *entity.Artifact) error {
	if artifact == nil {
		return errors.New("nil artifact provided")
	}

	return s.wasabiRepo.UploadArtifact(artifact)
}

func (s *StorageBucketUtility) DownloadArtifact(artifactName string) (*entity.Artifact, error) {
	if artifactName == "" {
		return nil, errors.New("empty artifact name provided")
	}

	return s.wasabiRepo.DownloadArtifact(artifactName)
}