package gotorrentparser

import (
	"crypto/sha1"
	"fmt"
)

func toSHA1(data []byte) string {
	hash := sha1.New()
	hash.Write(data)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
