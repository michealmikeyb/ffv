package main

import (
	"log"

	tags_pb "github.com/michealmikeyb/ffv/tags"
)

func main() {
	err := tags_pb.UpdateAllBuffers()
	if err != nil {
		log.Fatal("errored out")
	}
}
