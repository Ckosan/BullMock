package utils

import (
	"math/rand"
	"time"
	"unsafe"
)

type Func struct {
	F string
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	standTime     = "2006-01-02 15:04:05"
)

var src = rand.NewSource(time.Now().UnixNano())

func (fun *Func) Str(length int) string {

	b := make([]byte, length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

// FormatDate format传需要格式化的的格式，t传时间 不传的话默认格式化当前
func (fun *Func) FormatDate(format string, t string) (string, error) {
	if t == "" {
		return time.Now().Format(format), nil
	} else {
		tt, error := time.ParseInLocation(format, t, time.Local)
		if error != nil {
			return "", error
		}
		return tt.String(), nil
	}
}

func (fun *Func) FormatNow2YYYYMMDD() string {
	return time.Now().Format(standTime)
}

func (fun *Func) FormatYYYYMMDDNowDateAdd(year, month, day int) string {
	return time.Now().AddDate(year, month, day).Format(standTime)
}

func (fun *Func) FormatNowDateAdd2(format string, year, month, day int) string {
	return time.Now().AddDate(year, month, day).Format(format)
}

func (fun *Func) FormatDateAdd(t, f string, n int) (string, error) {
	_, err := time.ParseInLocation(f, t, time.Local)
	if err != nil {
		return "", err
	}
	return "", nil
}
