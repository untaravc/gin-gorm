package test_controller

import (
	"encoding/json"
	"fmt"
	"gin-gorm/app/models"
	"gin-gorm/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func ConnectToRedis(ctx *gin.Context) {
	auth_token := models.AuthToken{
		KaryawanId:    1,
		KaryawanNama:  "Untara Vivi Chahya",
		KaryawanEmail: "",
		CabangId:      22299,
		Role:          "",
		JabatanId:     "",
	}

	setAuthToken(ctx, "myToken", auth_token)
	data := getToken(ctx, "myToken")

	ctx.JSON(http.StatusOK, gin.H{"msg": data})
}

func setAuthToken(ctx *gin.Context, key string, auth_token models.AuthToken) {

	jsonStr, jsonErr := json.Marshal(auth_token)
	if jsonErr != nil {
		fmt.Println("Error:", jsonErr)
		return
	}

	rdb := database.InitRedisClient()
	expiration := time.Hour * 24 // Token expiration of 24 hours
	err := rdb.Set(ctx, key, jsonStr, expiration).Err()
	if err != nil {
		return
	}
}

func getToken(ctx *gin.Context, key string) map[string]interface{} {
	rdb := database.InitRedisClient()

	// Get the token from Redis
	jsonStr, err := rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Token not found"})
		ctx.Abort()
	} else if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Redis Failed"})
		ctx.Abort()
	}

	var data map[string]interface{}
	errUnmarshal := json.Unmarshal([]byte(jsonStr), &data)

	if errUnmarshal != nil {
		fmt.Println("Error:", errUnmarshal)
	}

	return data
}
