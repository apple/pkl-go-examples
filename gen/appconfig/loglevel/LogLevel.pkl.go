// Code generated from Pkl module `org.pkl.golang.example.AppConfig`. DO NOT EDIT.
package loglevel

import (
	"encoding"
	"fmt"
)

type LogLevel string

const (
	Error LogLevel = "error"
	Warn  LogLevel = "warn"
	Info  LogLevel = "info"
)

// String returns the string representation of LogLevel
func (rcv LogLevel) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(LogLevel)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for LogLevel.
func (rcv *LogLevel) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "error":
		*rcv = Error
	case "warn":
		*rcv = Warn
	case "info":
		*rcv = Info
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid LogLevel`, str)
	}
	return nil
}
