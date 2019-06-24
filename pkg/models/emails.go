package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

const (
	emailListKey  = "EMAILS:ZSET"
	emailStoreKey = "EMAILS:"
)

// Email struct for data representation
type Email struct {
	ID          string `json:"id,string"`
	FromAddress string `query:"from_address" json:"from_address" form:"from_address"`
	ToAddress   string `query:"to_address" json:"to_address" form:"to_address"`
	Subject     string `query:"subject" json:"subject" form:"subject"`
	Body        string `query:"body" json:"body" form:"body"`
	WhenSent    int64  `json:"when_sent,string"` // stored in unix timestamp
	When        int64  `json:"when,string"`      // stored in unix timestamp
	Status      string `json:"status"`
	Remarks     string `json:"remarks,string"`
}

// Save : save method for email struct
func (email *Email) Save(client *redis.Client) error {
	email.When = time.Now().Unix()
	uid, _ := uuid.NewV4()
	email.ID = uid.String()
	email.Status = "Pending"

	var emailMap map[string]interface{}
	emailJSON, _ := json.Marshal(email)
	json.Unmarshal(emailJSON, &emailMap)

	emailKey := emailStoreKey + uid.String()
	err := client.HMSet(emailKey, emailMap).Err()
	if err != nil {
		return err
	}
	err = client.ZAdd(emailListKey, &redis.Z{
		Score:  float64(email.When),
		Member: emailKey,
	}).Err()
	if err != nil {
		client.Del(emailKey)
		return err
	}
	return nil
}

// GetAll : Get all data from redis store
func (email *Email) GetAll(client *redis.Client) []string {
	keys, _ := client.ZRangeByScore(emailListKey, &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()

	data := []string{}
	for _, emailKey := range keys {
		val, _ := client.HGetAll(emailKey).Result()
		m, _ := json.Marshal(val)
		data = append(data, string(m))
	}

	return data
}

// MarkSent : set email status to success
func (email *Email) MarkSent(client *redis.Client) {
	fmt.Println("Email Sent")
}

// MarkFail : set email status to failed
func (email *Email) MarkFail(client *redis.Client) {
	fmt.Println("Email Sent")
}
