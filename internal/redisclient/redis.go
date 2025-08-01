/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/

package redisclient

import (
	"context"
	"fmt"
	"log"

	"github.com/Abiji-2020/PesudoCLI/pkg/utils"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	ctx    context.Context
	client *redis.Client
}

type RedisClientInterface interface {
	Client() *redis.Client
	Close() error
	Context() context.Context
}

func NewRedisClient(addr string) *RedisClient {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	return &RedisClient{
		ctx:    ctx,
		client: client,
	}
}

func (r *RedisClient) Client() *redis.Client {
	return r.client
}

func (r *RedisClient) Close() error {
	if err := r.client.Close(); err != nil {
		log.Printf("Error closing Redis client: %v", err)
		return err
	}
	return nil
}

func (r *RedisClient) Context() context.Context {
	return r.ctx
}

func (r *RedisClient) CreateVectorIndex(indexName string, dim int) error {
	_, err := r.client.Do(r.ctx, "FT.INFO", indexName).Result()
	if err == nil {
		log.Printf("Index %s already exists", indexName)
		return nil
	}

	commandToRun := []interface{}{
		"FT.CREATE", indexName,
		"ON", "HASH",
		"PREFIX", "1", "doc:",
		"SCHEMA",
		"command", "TEXT",
		"os", "TEXT",
		"text_chunk", "TEXT",
		"embedding", "VECTOR", "HNSW", "6",
		"DIM", dim,
		"TYPE", "FLOAT32",
		"DISTANCE_METRIC", "COSINE",
	}
	_, err = r.client.Do(r.ctx, commandToRun...).Result()

	if err != nil {
		log.Printf("Error creating index %s: %v", indexName, err)
		return err
	}

	return nil
}

func (r *RedisClient) AddDocument(id, command, os, textChunk string, embedding []float32) error {

	bytesEmbedding, err := utils.Float32SliceToBytes(embedding)
	if err != nil {
		return fmt.Errorf("failed to convert embedding to bytes: %w", err)
	}
	key := "doc:" + id

	fields := map[string]interface{}{
		"command":    command,
		"os":         os,
		"text_chunk": textChunk,
		"embedding":  bytesEmbedding,
	}
	if err := r.client.HSet(r.ctx, key, fields).Err(); err != nil {
		return fmt.Errorf("failed to add document to Redis: %w", err)
	}
	return nil
}
