package repository

import "github.com/gomodule/redigo/redis"

type ChunkRepository struct {
	Repository

	DB redis.Conn
}
