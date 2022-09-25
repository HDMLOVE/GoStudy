package main

import (
	"encoding/binary"
	"fmt"
)

const entryHeaderSize = 10

// 写入文件记录
type Entry struct {
	Key       []byte
	Value     []byte
	KeySize   uint32
	ValueSize uint32
	Mark      uint16
}

// 常见一个节点
func NewEntry(key, value []byte, mark uint16) *Entry {
	return &Entry{
		Key:       key,
		Value:     value,
		KeySize:   uint32(len(key)),
		ValueSize: uint32(len(value)),
		Mark:      mark,
	}
}

func (e *Entry) GetSize() int64 {
	return int64(entryHeaderSize + e.KeySize + e.ValueSize)
}

func (e *Entry) Encode() ([]byte, error) {
	buf := make([]byte, e.GetSize())
	binary.BigEndian.PutUint32(buf[0:4], e.KeySize)
	binary.BigEndian.PutUint32(buf[4:8], e.ValueSize)
	binary.BigEndian.PutUint16(buf[8:10], e.Mark)
	copy(buf[entryHeaderSize:entryHeaderSize+e.KeySize], e.Key)
	copy(buf[entryHeaderSize+e.KeySize:], e.Value)
	return buf, nil
}

func main() {
	fmt.Printf("dasdadadad\n")
}
