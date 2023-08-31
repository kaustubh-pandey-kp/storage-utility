package usecase

import (
	"mobilestoragebucketutility/pkg/entity"
	"mobilestoragebucketutility/pkg/repository"
)


type StorageUsecase interface {
	UploadArtifact(artifact *entity.Artifact, bucketName string) error
	DownloadArtifact(artifactName, bucketName string) (*entity.Artifact, error)
}

type storageUsecase struct {
    wasabiRepo repository.WasabiRepository
}

func NewStorageUsecase(wasabiRepo repository.WasabiRepository) StorageUsecase {
    return &storageUsecase{
        wasabiRepo: wasabiRepo,
    }
}

func (u *storageUsecase) UploadArtifact(artifact *entity.Artifact, bucketName string) error {
	// Call the Wasabi repository's UploadArtifact method
	return u.wasabiRepo.UploadArtifact(artifact)
}

func (u *storageUsecase) DownloadArtifact(artifactName, bucketName string) (*entity.Artifact, error) {
	// Call the Wasabi repository's DownloadArtifact method
	return u.wasabiRepo.DownloadArtifact(artifactName)
}