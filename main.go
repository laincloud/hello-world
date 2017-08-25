package main

import (
	"fmt"
	"net/http"
	"strconv"
	"os"

	"github.com/go-redis/redis"
)

const visitCountKey = "visitCountKey"

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	containerID, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := redisClient.Incr(visitCountKey).Err(); err != nil {
			fmt.Fprintf(w, "redisClient.Incr() failed, error: %s.", err)
			return
		}

		visitCountStr, err := redisClient.Get(visitCountKey).Result()
		if err != nil {
			fmt.Fprintf(w, "redisClient.Get() failed, error: %s.", err)
			return
		}

		visitCount, err := strconv.Atoi(visitCountStr)
		if err != nil {
			fmt.Fprintf(w, "strconv.Atoi() failed, error: %s.", err)
			return
		}
		fmt.Fprintf(w, "Hello, LAIN. I'm container: %s. " +
			       "You are the %dth visitor.", containerID, visitCount)
	})

	http.ListenAndServe(":8080", nil)
}
