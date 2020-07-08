package GolangForZip

import (
	"reflect"
)
func CreateAnyTypeSlice(slice interface{}) ([]interface{}, bool){
	val, ok := isSlice(slice)

	if !ok {
		return nil, false
	}

	sliceLen := val.Len()

	out := make([]interface{}, sliceLen)

	for i := 0;i < sliceLen;i++ {
		out[i] = val.Index(i).Interface()
	}

	return out, true
}

func isSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)

	if val.Kind() == reflect.Slice {
		ok = true
	}

	return
}

func Zip(lists... interface{}) (chan []interface{}) {
	var (
		size int
		slice [][]interface{}
	)

	for _, i := range lists {
		s, ok := CreateAnyTypeSlice(i)
		if ok{
			if len(slice)<size{
				size=len(slice)
			}else if size==0{
				size=len(slice)
			}
			slice = append(slice, s)
		}
	}

	ch :=make(chan []interface{},size)
	defer close(ch)

	for i:=0;i<size;i++ {
		var s1 []interface{}
		for x:= 0;x<len(lists);x++ {
			s1 = append(s1,slice[x][i])
		}
		ch <-s1
	}
	return ch
}
