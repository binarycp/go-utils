package to

import (
	_ "embed"
	"fmt"
	"github.com/binarycp/gutils/slices"
	"testing"
)

var (
	s = Sql{}
	//go:embed data/sql.test
	sqlData string

	e = Env{}
)

type SqlStruct struct {
	Province string `to:"province"`
	Count    int    `to:"count"`
}

func TestMarshaSql(t *testing.T) {
	s.Encode(make(map[string]string))
	sqls := make([]SqlStruct, 3)
	sqls[0] = SqlStruct{
		Province: "太原",
		Count:    23,
	}
	sqls[1] = SqlStruct{
		Province: "故宫",
		Count:    -10,
	}
	fmt.Println(Marshal(sqls, s))
}

func TestUnMarsha(t *testing.T) {
}

func TestCancel(t *testing.T) {
	fmt.Println(!slices.EqualToByte([]byte{0x06, 0xa5, 0x00, 0x55}, []byte{0x06, 0xa5, 0x00, 0x55}))
}
