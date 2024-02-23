package main

import (
	"errors"
	"fmt"
	"reflect"
)

func Pop(array interface{}, front bool) (interface{}, error) {
	var ans []interface{}
	arr := reflect.ValueOf(array)

	if arr.Kind() == reflect.Slice || arr.Kind() == reflect.Array {
		if arr.Len() == 0 {
			return nil, errors.New("empty array")
		} else {
			for i := 0; i < arr.Len(); i++ {
				if i == 0 && !front {
					ans = append(ans, arr.Index(i).Interface())

				} else if i == arr.Len()-1 && front {
					ans = append(ans, arr.Index(i).Interface())

				} else if i != 0 && i != arr.Len()-1 {
					ans = append(ans, arr.Index(i).Interface())
				}
			}
		}
	} else {
		return nil, errors.New("invalid input")
	}
	return ans, nil
}

func main() {
	arr := [][]int{{1, 2}, {2}, {3}, {4}, {5, 6}}
	ans, err := Pop(arr, false)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(ans)
	}
}
