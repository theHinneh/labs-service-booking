package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type myTime time.Time

func (mt *myTime) UnmarshalJSON(bs []byte) error {
	var timestamp int64
	err := json.Unmarshal(bs, &timestamp)
	if err != nil {
		return err
	}

	*mt = myTime(time.Unix(timestamp/1000, timestamp%1000*1e6))
	return nil
}

func (mt myTime) MarshalJSON() ([]byte, error) {
	timestamp := time.Time(mt).UnixNano() / 1e6
	log.Println(time.Time(mt).UnixNano())
	return json.Marshal(timestamp)
}

type Timestamp struct {
	OneDay     myTime    `json:"oneDay" form:"oneDay"`
	AnotherDay time.Time `json:"anotherDay" form:"anotherDay" time_format:"unix"`
}

func ParseTime(c *gin.Context) {
	var example Timestamp
	if err := c.Bind(&example); err != nil {
		log.Printf("bind timestamp error: %s", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": example})
}
