package config

import (
	"bufio"
	"io"
	"os"

	"github.com/ethereum/go-ethereum/log"

	"chainmaker.org/chainmaker/common/v3/random/uuid"
)

// CMMID get cmmid
var CMMID, _ = GetCMMId()

// IdFile id file
const IdFile = "CMMID"

// GetCMMId get cmmid
func GetCMMId() (string, error) {

	f, err := os.OpenFile(IdFile, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := f.Close(); r != nil {
			panic(r)
		}
	}()

	buf := make([]byte, 32)

	r := bufio.NewReader(f)
	n, err := r.Read(buf)

	if err != nil && err != io.EOF {
		panic(err)
	}

	id := string(buf[:n])

	if n < 5 {
		id = uuid.GetUUID()
		_, err = f.WriteAt([]byte(id), 0)
		if err != nil {
			log.Error("err:%v", err)
		}
	}

	return id, nil
}
