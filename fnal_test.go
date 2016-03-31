package fnal

import (
	"reflect"
	"testing"
)

var (
	testSlice       = []int{1, 2, 3}
	testMapSlice    = []int{2, 3, 4}
	testFilterSlice = []int{2, 3}

	testMap       = map[string]string{"A": "A0", "B": "B0"}
	testMapMap    = map[string]string{"A": "A0A", "B": "B0B"}
	testFilterMap = map[string]string{"A": "A0"}
)

func sliceMapFunc(i interface{}) interface{} {
	return 1 + i.(int)
}

func mapMapFunc(i interface{}) interface{} {
	entry := i.(MapEntry)
	return MapEntry{
		Key:   entry.Key,
		Value: entry.Value.(string) + entry.Key.(string),
	}
}

func sliceFilterFunc(i interface{}) bool {
	return i.(int) >= 2
}

func mapFilterFunc(i interface{}) bool {
	entry := i.(MapEntry)
	return entry.Key.(string)+entry.Value.(string) == "AA0"
}

func sliceFoldlFunc(acc interface{}, i interface{}) interface{} {
	return acc - i
}

func mapFoldlFunc(acc interface{}, i interface{}) interface{} {
	return nil
}

func TestMap(t *testing.T) {
	var rt interface{}
	rt = Map(testSlice, sliceMapFunc)
	if !reflect.DeepEqual(rt, testMapSlice) {
		t.Log(rt)
		t.Fatal("Map errro")
	}
	rt = Map(testMap, mapMapFunc)
	if !reflect.DeepEqual(rt, testMapMap) {
		t.Log(rt)
		t.Fatal("Map error")
	}

	rt = Filter(testSlice, sliceFilterFunc)
	if !reflect.DeepEqual(rt, testFilterSlice) {
		t.Log(rt)
		t.Fatal("Filter error")
	}
	rt = Filter(testMap, mapFilterFunc)
	if !reflect.DeepEqual(rt, testFilterMap) {
		t.Log(rt)
		t.Fatal("Filter error")
	}
}
