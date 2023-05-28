package backup

import (
	"backups/internal/db"
	"backups/internal/minio"
	"backups/pkg/model"
	"fmt"
	"log"
	"os"
)

type BackupManager struct {
	Provider    *db.DataBaseProvider
	MinioClient *minio.MinioClient
}

func NewBackupManager(provider *db.DataBaseProvider, minio *minio.MinioClient) (*BackupManager, error) {
	return &BackupManager{
		Provider:    provider,
		MinioClient: minio,
	}, nil
}

func (m *BackupManager) UploadDump(dumpName string) error {
	log.Println("Try to Upload dump to minio")

	dumpPath := m.Provider.Config.Db.Path
	bucketName := m.Provider.Config.Minio.MinioBucket

	dumpFile := fmt.Sprintf("%s/%s", dumpPath, dumpName)
	file, err := os.Open(dumpFile)
	if err != nil {
		log.Println("Cant open file")
		model.ErrorDuringBackups.Inc()
		return err
	}
	defer file.Close()

	// Upload the file to MinIO
	err = m.MinioClient.AddFileToBucket(bucketName, dumpName, file)
	if err != nil {
		log.Println("Cant add file to bucket")
		model.ErrorDuringBackups.Inc()
		return err
	}

	log.Printf("Backup uploaded to MinIO bucket '%s' with object name '%s'\n", bucketName, dumpName)
	return nil
}

func (m *BackupManager) deleteOldDumps() error {
	return nil
}
