package lib

import (
	"os"
	"path"
	"reflect"
	"strconv"
)

func Getenv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func ParseInt(value string) int {
	if i, e := strconv.Atoi(value); e == nil {
		return i
	}
	return -1
}

func ParseString(value interface{}) string {
	return reflect.TypeOf(value).String()
}

func Pwd() string{
	if dir, err := os.Getwd(); err == nil {
		return path.Clean(dir)
	}
	return ""
}

func DatabasePath(value string) string {
	databasePath := path.Join(Pwd(), "database")
	dir := path.Join(path.Clean(databasePath), value)
	return string(dir)
}
