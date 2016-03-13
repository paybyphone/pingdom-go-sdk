package util

import "reflect"

// SimpleCopyStruct does a top-level copy of struct fields from one struct
// to another, if the field in the source is set.
func SimpleCopyStruct(src, dst interface{}) {
	s := reflect.ValueOf(src).Elem()
	d := reflect.ValueOf(dst).Elem()

	for i := 0; i < s.NumField(); i++ {
		if s.Field(i).CanSet() == true {
			if s.Field(i).Interface() != nil {
				for j := 0; j < d.NumField(); j++ {
					if d.Type().Field(j).Name == s.Type().Field(i).Name && s.Field(i).Elem() != reflect.Zero(s.Type().Field(i).Type) {
						d.Field(j).Set(s.Field(i))
					}
				}
			}
		}
	}
}
