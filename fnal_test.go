package fnal

import (
	"reflect"
	"testing"
)

var (
	testSlice       = []int{1, 2, 3}
	testMapSlice    = []int{2, 3, 4}
	testFilterSlice = []int{2, 3}
	testFoldlSlice  = 0 - 1 - 2 - 3
	testFoldrSlice  = 0 - 3 - 2 - 1

	testMap       = map[string]string{"A": "A0", "B": "B0"}
	testMapMap    = map[string]string{"A": "A0A", "B": "B0B"}
	testFilterMap = map[string]string{"A": "A0"}
	testFoldlMap  = 0 + len("A") + len("A0") + len("B") + len("B0")
	testFoldrMap  = 0 + len("A") + len("A0") + len("B") + len("B0")
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
	return acc.(int) - i.(int)
}

func mapFoldlFunc(acc interface{}, i interface{}) interface{} {
	entry := i.(MapEntry)
	return acc.(int) + len(entry.Key.(string)) + len(entry.Value.(string))
}

func sliceFoldrFunc(i interface{}, acc interface{}) interface{} {
	return sliceFoldlFunc(acc, i)
}

func mapFoldrFunc(i interface{}, acc interface{}) interface{} {
	return mapFoldlFunc(acc, i)
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

	rt = Foldl(testSlice, 0, sliceFoldlFunc)
	if !reflect.DeepEqual(rt, testFoldlSlice) {
		t.Log(rt)
		t.Fatal("Foldl error")
	}
	rt = Foldl(testMap, 0, mapFoldlFunc)
	if !reflect.DeepEqual(rt, testFoldlMap) {
		t.Log(rt)
		t.Fatal("Foldl error")
	}

	rt = Foldr(testSlice, 0, sliceFoldrFunc)
	if !reflect.DeepEqual(rt, testFoldrSlice) {
		t.Log(rt)
		t.Fatal("Foldr error")
	}
	rt = Foldr(testMap, 0, mapFoldrFunc)
	if !reflect.DeepEqual(rt, testFoldrMap) {
		t.Log(rt)
		t.Fatal("Foldr error")
	}
}
