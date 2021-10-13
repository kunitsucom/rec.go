package rec

import "sync"

const (
	bufferPoolCap = 1024
	pcPoolCap     = 64
)

type buffer struct {
	Buffer []byte
}

var bufferPool = &sync.Pool{ // nolint: gochecknoglobals
	New: func() interface{} {
		return &buffer{make([]byte, 0, bufferPoolCap)}
	},
}

type programcounter struct {
	PC []uintptr
}

var pcPool = &sync.Pool{ // nolint: gochecknoglobals
	New: func() interface{} {
		return &programcounter{make([]uintptr, pcPoolCap)}
	},
}
