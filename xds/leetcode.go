package xds

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
)

// 用于处理岛屿类题目中的测试case转换为可用结构
func MarshalBytesMatrix(str string) ([][]byte, error) {
	data := [][]string{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		return nil, err
	}
	result := [][]byte{}
	for _, row := range data {
		tmp := []byte{}
		for _, val := range row {
			if len([]byte(val)) == 0 {
				return nil, errors.New("matrix has a empty string")
			}
			tmp = append(tmp, []byte(val)[0])
		}
		result = append(result, tmp)
	}
	return result, nil
}

// 用于牛客的数列类题目的输入输出
func MarshalIntSlice() ([]int, error) {
	result := []int{}
	for {
		tmp := 0
		if _, err := fmt.Scan(&tmp); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result, nil
}
