package manager

import (
	"fmt"
	"strconv"
)

func NewRecordCollection() RecordCollection {
	return RecordCollection{}
}

type RecordCollection []Record

func (r *RecordCollection) Add(id int) *Record {
	record := NewRecord(id)
	*r = append(*r, record)
	return &record
}

func (r *RecordCollection) Len() int {
	return len(*r)
}

func NewRecord(id int) Record {
	return Record{ID: id, keys: []string{}, keyMap: map[string]any{}}
}

type Record struct {
	ID int
	keys []string
	keyMap map[string]any
}

func (r Record) Keys() []string {
	return r.keys
}

func (r Record) HasKeys() bool {
	return len(r.keys) > 0
}

func (r Record) Get(name string) any {
	v, ok := r.keyMap[name]
	if ok {
		return v
	}
	return nil
}

func (r Record) GetInt(name string) int {
	if v := r.Get(name); v != nil {
		vv, _ := strconv.Atoi(fmt.Sprintf("%v", v))
		return vv
	}
	return 0
}

func (r Record) GetString(name string) string {
	if v := r.Get(name); v != nil {
		if i, ok := v.(string); ok {
			return i
		}
		return fmt.Sprintf("%v", v)
	}
	return ""
}

func (r *Record) Set(name string, value any) *Record {
	r.keys = append(r.keys, name)
	r.keyMap[name] = value
	return r
}