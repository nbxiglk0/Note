package nfuncions

import "reflect"

func IsContain(item interface{}, items interface{}) bool{ //判断slice是否包含某个item
	switch reflect.TypeOf(items).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(items)
		for i :=0; i < s.Len(); i++{
			if reflect.DeepEqual(item, s.Index(i).Interface()){
				return true
			}
		}
	}
	return false
}

func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}