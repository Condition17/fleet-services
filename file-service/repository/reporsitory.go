package repository

import "fmt"

type Repository interface{}

type CacheEntityType string

const (
	fileEntity  CacheEntityType = "file"
	storeEntity CacheEntityType = "store"
)

func composeFileKey(fileId string) string {
	return fmt.Sprintf("%s:%s", fileEntity, fileId)
}
