package helper

import "strconv"

func StrSliceToInt64Slice(values []string) ([]int64, error) {
	var res = make([]int64, 0)

	for i := 0; i < len(values); i++ {
		v, err := strconv.ParseInt(values[i], 10, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, v)
	}

	return res, nil
}
