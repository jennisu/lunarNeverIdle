package waste

import "math/rand"

var PartialBuffers []*PartialObject

const (
	KiB = 1024
	MiB = 1024 * KiB
	GiB = 1024 * MiB
)

type PartialObject struct {
	Data []byte
}

func Memory(mib int) {
	if mib <= 0 {
		PartialBuffers = nil
		return
	}
	bytes := mib * MiB
	createPartialBuffer(bytes)
}

//Helper function: creates a single block of the exact size
func createPartialBuffer(bytes int) {
	PartialBuffers = PartialBuffers[:0] // clear previous blocks
	if cap(PartialBuffers) == 0 {
		PartialBuffers = make([]*PartialObject, 0, 1)
	}

	o := &PartialObject{
		Data: make([]byte, bytes),
	}
	rand.Read(o.Data)
	PartialBuffers = append(PartialBuffers, o)
}