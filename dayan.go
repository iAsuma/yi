package yi

import (
	"math/bits"
	"strconv"
	"strings"
)

// DaYan ...
type DaYan struct {
	Number  int
	Lucky   string
	NvMing  string
	Max     bool
	Gua     string
	SkyNine string
	YiXiang string
	Basis   string
	Family  string
	Health  string
	Comment string
}

var daYanList map[int]*DaYan

func init() {
	daYanList = make(map[int]*DaYan)

	file81shu, err := DataFiles.Open("data/81shu.csv")
	if err != nil {
		panic(err)
	}

	records, err := readData(file81shu)

	if err != nil {
		panic(err)
	}

	for _, record := range records {
		bihua, err := strconv.ParseInt(record[0], 10, bits.UintSize)
		if err != nil {
			panic(err)
		}

		var max bool = false
		maxstr := strings.TrimSpace(record[3])
		if maxstr == "最吉" {
			max = true
		}

		dayan := DaYan{
			Number:  int(bihua),
			Lucky:   record[1],
			NvMing:  record[2],
			Max:     max,
			Gua:     record[4],
			SkyNine: record[5],
			YiXiang: record[6],
			Basis:   record[7],
			Family:  record[8],
			Health:  record[9],
			Comment: record[10],
		}

		daYanList[dayan.Number-1] = &dayan
	}
}

// IsNotSuitableGirl 女性不宜此数
func (dy DaYan) IsNotSuitableGirl() bool {
	return dy.NvMing == "凶"
}

// IsMax 是否最大好运数
func (dy DaYan) IsMax() bool {
	return dy.Max
}

// GetDaYan 获取大衍之数
func GetDaYan(idx int) DaYan {
	if idx <= 0 {
		panic("wrong idx")
	}
	i := (idx - 1) % 81

	return *daYanList[i]
}
