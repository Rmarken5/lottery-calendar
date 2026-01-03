package model

import (
	"errors"
	"strconv"
	"strings"
)

type (
	LotteryNumbers []uint8
)

func (l LotteryNumbers) ToDBRepresentation() string {
	sb := strings.Builder{}
	for _, n := range l {
		sb.WriteString(strconv.Itoa(int(n)))
		sb.WriteString(",")
	}
	return sb.String()[:sb.Len()-1]
}

func FromDBRepresentation(numbers string) (LotteryNumbers, error) {
	x := strings.Split(numbers, ",")
	l := make(LotteryNumbers, len(x))
	for i, n := range x {
		num, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		if num < 0 || num > 256 {
			return nil, errors.New("number out of bounds")
		}
		l[i] = uint8(num)
	}
	return l, nil
}
