package server

import (
	"fmt"
	"sync"
)

// ErrorOffsetNotFound is a custom error used to express that
// the record of interest is not found.
var ErrorOffsetNotFound = fmt.Errorf("offset not found")

// Log struct.
type Log struct {
	mu      sync.Mutex
	records []Record
}

// Record struct.
type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

// NewLog returns a log instance.
func NewLog() *Log {
	return &Log{}
}

// Append appends a record at the end of the records slice of the log.
func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrorOffsetNotFound
	}

	return c.records[offset], nil
}
