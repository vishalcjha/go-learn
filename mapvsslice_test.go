package main

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

var campIDs []string

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 9000; i++ {
		newCampID := RandStringBytes(7)
		campIDs = append(campIDs, newCampID)
		var newSliceData []*campData
		newMapData := make(map[uint8]*campData)
		for i := 1; i < 4; i++ {
			newMapData[uint8(i)] = &campData{goal: uint8(i), data: data}
			newSliceData = append(newSliceData, &campData{goal: uint8(i), data: data})
		}
		sliceStore[newCampID] = newSliceData
		mapStore[newCampID] = newMapData
	}
	ret := m.Run()
	os.Exit(ret)
}
func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, campID := range campIDs {
			for j := 0; j < 4; j++ {
				getCampDataFromMap(campID, uint8(i))
			}
		}
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, campID := range campIDs {
			for j := 0; j < 4; j++ {
				getCampDataFromSlice(campID, uint8(i))
			}
		}
	}
}
