package raw

import (
	"encoding/json"
)

type Raw map[string]any

func (r *Raw) ReadRawData(bytes []byte) error {
	return json.Unmarshal(bytes, &r)
}

func (r *Raw) Set(key string, value any) {
	if r == nil {
		r = new(Raw)
	}
	(*r)[key] = value
}

func (r *Raw) Del(key string) {
	if r == nil {
		return
	}
	delete(*r, key)
}

func (r *Raw) Get(key string) (any, bool) {
	if r == nil {
		return nil, false
	}
	v, ok := (*r)[key]
	return v, ok
}

func (r *Raw) GetString(key string) string {
	v, ok := (*r)[key]
	if !ok {
		return ""
	}
	res, _ := v.(string)
	return res
}

func (r *Raw) GetFloat64(key string) float64 {
	v, ok := (*r)[key]
	if !ok {
		return 0
	}
	res, _ := v.(float64)
	return res
}

func (r *Raw) GetFloat32(key string) float32 {
	return float32(r.GetFloat64(key))
}

func (r *Raw) GetInt(key string) int {
	return int(r.GetFloat64(key))
}

func (r *Raw) GetInt8(key string) int8 {
	return int8(r.GetFloat64(key))
}

func (r *Raw) GetInt32(key string) int32 {
	return int32(r.GetFloat64(key))
}

func (r *Raw) GetInt64(key string) int64 {
	return int64(r.GetFloat64(key))
}

func (r *Raw) GetBool(key string) bool {
	return (*r)[key].(bool)
}

func (r *Raw) Exists(key string) bool {
	_, ok := (*r)[key]
	return ok
}

func (r *Raw) Map() map[string]any {
	return *r
}
