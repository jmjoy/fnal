package fnal

import (
	"fmt"
	"reflect"
)

const ERR_UNTRAVERSABLE_TYPE = "The type of sequence can't be traversed: %s."

type Traversable interface {
	current() interface{}
	next()
	rewine()
	valid() bool
}

type MapEntry struct {
	Key   interface{}
	Value interface{}
}

func Map(seq interface{}, fn func(interface{}) interface{}) interface{} {
	value := reflect.ValueOf(seq)
	kind := value.Kind()
	typ := value.Type()

	switch kind {
	case reflect.Slice:
		length := value.Len()
		slice := reflect.MakeSlice(typ, 0, length)
		for i := 0; i < length; i++ {
			item := fn(value.Index(i).Interface())
			slice = reflect.Append(slice, reflect.ValueOf(item))
		}
		return slice.Interface()

	case reflect.Map:
		keys := value.MapKeys()
		m := reflect.MakeMap(typ)
		for i := range keys {
			entry := fn(MapEntry{
				Key:   keys[i].Interface(),
				Value: value.MapIndex(keys[i]).Interface(),
			}).(MapEntry)

			m.SetMapIndex(
				reflect.ValueOf(entry.Key),
				reflect.ValueOf(entry.Value),
			)
		}
		return m.Interface()

	default:
		panic(fmt.Sprintf(ERR_UNTRAVERSABLE_TYPE, typ))
	}

	return nil
}

func Filter(seq interface{}, fn func(interface{}) bool) interface{} {
	value := reflect.ValueOf(seq)
	kind := value.Kind()
	typ := value.Type()

	switch kind {
	case reflect.Slice:
		length := value.Len()
		slice := reflect.MakeSlice(typ, 0, length)
		for i := 0; i < length; i++ {
			ok := fn(value.Index(i).Interface())
			if ok {
				slice = reflect.Append(slice, value.Index(i))
			}
		}
		return slice.Interface()

	case reflect.Map:
		keys := value.MapKeys()
		m := reflect.MakeMap(typ)
		for i := range keys {
			ok := fn(MapEntry{
				Key:   keys[i].Interface(),
				Value: value.MapIndex(keys[i]).Interface(),
			})

			if ok {
				m.SetMapIndex(keys[i], value.MapIndex(keys[i]))
			}
		}
		return m.Interface()

	default:
		panic(fmt.Sprintf(ERR_UNTRAVERSABLE_TYPE, typ))
	}

	return nil
}

func Foldl(seq interface{}, acc interface{}, fn func(interface{}, interface{}) interface{}) interface{} {
	value := reflect.ValueOf(seq)
	kind := value.Kind()
	typ := value.Type()

	switch kind {
	case reflect.Slice:
		length := value.Len()
		for i := 0; i < length; i++ {
			acc = fn(acc, value.Index(i).Interface())
		}
		return acc

	case reflect.Map:
		keys := value.MapKeys()
		for i := range keys {
			acc = fn(acc, MapEntry{
				Key:   keys[i].Interface(),
				Value: value.MapIndex(keys[i]).Interface(),
			})
		}
		return acc

	default:
		panic(fmt.Sprintf(ERR_UNTRAVERSABLE_TYPE, typ))
	}

	return nil
}

func Foldr(seq interface{}, acc interface{}, fn func(interface{}, interface{}) interface{}) interface{} {
	value := reflect.ValueOf(seq)
	kind := value.Kind()
	typ := value.Type()

	switch kind {
	case reflect.Slice:
		length := value.Len()
		for i := length - 1; i >= 0; i-- {
			acc = fn(value.Index(i).Interface(), acc)
		}
		return acc

	case reflect.Map:
		keys := value.MapKeys()
		for i := range keys {
			acc = fn(MapEntry{
				Key:   keys[i].Interface(),
				Value: value.MapIndex(keys[i]).Interface(),
			}, acc)
		}
		return acc

	default:
		panic(fmt.Sprintf(ERR_UNTRAVERSABLE_TYPE, typ))
	}

	return nil
}
