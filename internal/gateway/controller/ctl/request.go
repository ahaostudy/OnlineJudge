package ctl

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

var zero any

type Request map[string]any

func (r *Request) ReadRawData(c *gin.Context) error {
	bytes, err := c.GetRawData()
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, &r)
}

func (r *Request) Get(key string) (any, bool) {
	v, ok := (*r)[key]
	return v, ok
}

func (r *Request) GetString(key string) string {
	v, ok := (*r)[key]
	if !ok {
		return zero.(string)
	}
	res, _ := v.(string)
	return res
}

func (r *Request) GetFloat64(key string) float64 {
	v, ok := (*r)[key]
	if !ok {
		return zero.(float64)
	}
	res, _ := v.(float64)
	return res
}

func (r *Request) GetFloat32(key string) float32 {
	return float32(r.GetFloat64(key))
}

func (r *Request) GetInt(key string) int {
	return int(r.GetFloat64(key))
}

func (r *Request) GetInt8(key string) int8 {
	return int8(r.GetFloat64(key))
}

func (r *Request) GetInt32(key string) int32 {
	return int32(r.GetFloat64(key))
}

func (r *Request) GetInt64(key string) int64 {
	return int64(r.GetFloat64(key))
}

func (r *Request) GetBool(key string) bool {
	return (*r)[key].(bool)
}

func (r *Request) Exists(key string) bool {
	_, ok := (*r)[key]
	return ok
}

func (r *Request) Map() map[string]any {
	return *r
}
