// 通用唯一识别码
package uuid

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var sq uint16
var rnd *rand.Rand
var mutex sync.Mutex

func init() {
	rnd = rand.New(rand.NewSource(time.Now().Unix()))
}

// 通用唯一识别码
type UUID [16]byte

// 创建通用唯一识别码
func New() UUID {
	mutex.Lock()
	defer mutex.Unlock()

	t := time.Now()
	x := [16]byte{}
	binary.BigEndian.PutUint64(x[0:8], 1000000*uint64(t.Unix())+uint64(t.Nanosecond())/1000)
	binary.BigEndian.PutUint16(x[8:10], sq)
	rnd.Read(x[10:])
	sq += 1
	return UUID(x)
}

// 二进制序列
func (self UUID) Bytes() []byte {
	return self[:]
}

// 字符串
func (self UUID) String() string {
	return fmt.Sprintf("%X-%X-%X-%X-%X", self[0:4], self[4:6], self[6:8], self[8:10], self[10:16])
}
