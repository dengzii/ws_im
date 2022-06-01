package gate

import (
	"encoding/hex"
	"io"
	"math/rand"
	"time"
)

func GenTempID(gateID string) (ID, error) {
	uuid, err := newUUID()
	if err != nil {
		return "", err
	}
	return NewID(gateID, uuid, 0), nil
}

func newUUID() (string, error) {
	var uuid [16]byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	_, err := io.ReadFull(r, uuid[:])
	if err != nil {
		return "", err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	var buf1 [36]byte
	buf := buf1[:]
	hex.Encode(buf, uuid[:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], uuid[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], uuid[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], uuid[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], uuid[10:])
	return string(buf[:]), nil
}
