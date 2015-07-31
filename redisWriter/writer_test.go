// Package redisHandler provides ...
package redisWriter

import (
	// "github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestWrite(t *testing.T) {
	writer, err := NewRedisWriter(ADDRESS, "logtest", 2)
	if err != nil {
		log.Fatal(err)
	}

	writer.Write([]byte{'g', 'o', 'l'})
	writer.Write([]byte{'o', 'm', 'g'})
	writer.Write([]byte{'g', 'o', 'l'})
	writer.Write([]byte{'o', 'm', 'g'})
	writer.Write([]byte{'g', 'o', 'l'})
	// verify the byte is wriiten

	logs := ReadLog(writer.Conn, writer.Logname, 5)
	assert.True(t, len(logs) == 3, "log limit doesn't work!")
	//cleanup
	writer.Conn.Do("DEL", writer.Logname)

}
