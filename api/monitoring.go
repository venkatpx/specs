package api

import (
	"time"
)

type ServiceStat struct {
	BytesRead    uint64
	NumReads     uint64
	BytesWritten uint64
	NumWrites    uint64
	IoDepth      uint64
	QueueTime    uint64
}

type NetStat struct {
	BytesSent     uint64
	BytesReceived uint64
}

type Alert struct {
	Type      AlertType
	AlertData []byte
	Stamp     Time
	Id        uint64
}
