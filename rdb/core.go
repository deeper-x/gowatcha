package rdb

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// RDB is a wrapper for redis.Client
type RDB struct {
	cli *redis.Client
	Ctx context.Context
}

// New connect to redis server
func New() RDB {
	c := context.Background()
	return RDB{
		cli: redis.NewClient(
			&redis.Options{
				Addr:     ":6379",
				Password: "", // no password set
				DB:       0,  // use default DB
			},
		),
		Ctx: c,
	}
}

// GetOffset return offset
func (r RDB) GetOffset(inst string) (int64, error) {
	// start reading from offset
	offs, err := r.cli.Get(r.Ctx, inst).Int64()

	if err == redis.Nil {
		log.Println("key not set yet")
		offs = 0
	} else if err != nil {
		return 0, err
	}

	return offs, nil
}

// SetOffSet write offset
func (r RDB) SetOffSet(inst string, size int64) (bool, error) {
	if err := r.cli.Set(r.Ctx, inst, size, 0).Err(); err != nil {
		return false, err
	}

	return true, nil
}

// WriteLastSent write last sent timestamp to given recipient
func (r RDB) WriteLastSent(email string) (bool, error) {
	now := time.Now().Unix()

	if err := r.cli.Set(r.Ctx, email, now, 0).Err(); err != nil {
		return false, err
	}

	return true, nil
}

// GetLastSent return unix timestamp
func (r RDB) GetLastSent(email string) (int64, error) {
	ux, err := r.cli.Get(r.Ctx, email).Int64()

	if err != nil {
		return 0, err
	}

	return ux, nil
}
