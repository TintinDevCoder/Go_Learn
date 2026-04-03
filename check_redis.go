package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	// Check key "user:1"
	key := "user:1"
	t, err := client.Type(ctx, key).Result()
	if err != nil {
		fmt.Printf("Error checking type: %v\n", err)
	} else {
		fmt.Printf("Key %s type: %s\n", key, t)
	}

	// Try to get as string
	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("Key does not exist")
	} else if err != nil {
		fmt.Printf("Get error: %v\n", err)
	} else {
		fmt.Printf("Value: %s\n", val)
	}

	// Get hash fields
	hash, err := client.HGetAll(ctx, key).Result()
	if err != nil {
		fmt.Printf("HGetAll error: %v\n", err)
	} else {
		fmt.Printf("Hash fields: %v\n", hash)
	}

	// List all keys matching "user:*"
	keys, err := client.Keys(ctx, "user:*").Result()
	if err != nil {
		fmt.Printf("Error listing keys: %v\n", err)
	} else {
		fmt.Printf("Keys matching 'user:*': %v\n", keys)
		for _, k := range keys {
			t, _ := client.Type(ctx, k).Result()
			fmt.Printf("  %s -> type %s\n", k, t)
		}
	}
}
