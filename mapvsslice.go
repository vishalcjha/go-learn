package main

type campData struct {
	goal uint8
	data []int
}

var data = make([]int, 1000)
var testData = campData{data: data}
var mapStore = make(map[string]map[uint8]*campData)
var sliceStore = make(map[string][]*campData)

func getCampDataFromMap(campID string, goalType uint8) *campData {
	if data, found := mapStore[campID]; found {
		if d, found := data[goalType]; found {
			return d
		}
	}
	return nil
}

func getCampDataFromSlice(campID string, goalType uint8) *campData {
	if data, found := sliceStore[campID]; found {
		for _, d := range data {
			if d.goal == goalType {
				return d
			}
		}
	}
	return nil
}
