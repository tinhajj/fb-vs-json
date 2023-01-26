package main

import (
	"strings"
	"testing"
)

var small = strings.Repeat("BIG", 10)
var huge = strings.Repeat("BIG", 10000000)

var smallBlockFlat = serializeFlat(&small)
var smallBlockJSON = serializeJSON(&small)

var hugeBlockFlat = serializeFlat(&huge)
var hugeBlockJSON = serializeJSON(&huge)

func BenchmarkSerializeSmallFlat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		serializeFlat(&small)
	}
}

func BenchmarkSerializeHugeFlat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		serializeFlat(&huge)
	}
}

func BenchmarkSerializeSmallJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		serializeJSON(&small)
	}
}

func BenchmarkSerializeHugeJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		serializeJSON(&huge)
	}
}

func BenchmarkDeserializeSmallFlat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		deserializeFlat(smallBlockFlat)
	}
}

func BenchmarkDeserializeHugeFlat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		deserializeFlat(hugeBlockFlat)
	}
}

func BenchmarkDeserializeSmallJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		deserializeJSON(smallBlockJSON)
	}
}

func BenchmarkDeserializeHugeJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		deserializeJSON(hugeBlockJSON)
	}
}
