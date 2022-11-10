package text

import (
	"reflect"
)

func ContainsAny(entity interface{}, field, value string) bool {
    var boolean, _ = any(entity, field, value)

    return boolean
}

func any(entity interface{}, field, value string) (bool, int) {
    var valueOf = reflect.ValueOf(entity)

    for i := 0; i < valueOf.Len(); i++ {
        var fieldByName = valueOf.Index(i).FieldByName(field)

        if fieldByName.IsValid() && fieldByName.String() == value {
            return true, i
        }
    }

    return false, 0
}