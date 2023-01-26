package main

import (
	"encoding/json"
	"versus/block"

	flatbuffers "github.com/google/flatbuffers/go"
)

func main() { }

func serializeFlat(name *string) []byte {
	builder := flatbuffers.NewBuilder(0)
	b := block.BlockT{
		Id: 30,
		Name: *name,
	}
	builder.Finish(b.Pack(builder))
	return builder.FinishedBytes()
}

func serializeJSON(name *string) []byte {
	b := block.BlockT{
		Id: 30,
		Name: *name,
	}

	bytes, _ := json.Marshal(b)
	return bytes
}

func deserializeFlat(blockBytes []byte) string {
	b := block.GetRootAsBlock(blockBytes, 0)
	return string(b.Name())
}

func deserializeJSON(blockBytes []byte) string {
	var b = &block.BlockT{}
	json.Unmarshal(blockBytes, b)
	return b.Name
}
