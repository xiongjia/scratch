package chunk

import (
	"context"
	"sync"
)

const maxParallel = 1000

var decodeContextPool = sync.Pool{
	New: func() interface{} {
		return NewDecodeContext()
	},
}

// GetParallelChunks fetches chunks in parallel (up to maxParallel).
func GetParallelChunks(ctx context.Context, chunks []Chunk, f func(context.Context, *DecodeContext, Chunk) (Chunk, error)) ([]Chunk, error) {
	queuedChunks := make(chan Chunk)

	go func() {
		for _, c := range chunks {
			queuedChunks <- c
		}
		close(queuedChunks)
	}()

	processedChunks := make(chan Chunk)
	errors := make(chan error)

	for i := 0; i < min(maxParallel, len(chunks)); i++ {
		go func() {
			decodeContext := decodeContextPool.Get().(*DecodeContext)
			for c := range queuedChunks {
				c, err := f(ctx, decodeContext, c)
				if err != nil {
					errors <- err
				} else {
					processedChunks <- c
				}
			}
			decodeContextPool.Put(decodeContext)
		}()
	}

	var result = make([]Chunk, 0, len(chunks))
	var lastErr error
	for i := 0; i < len(chunks); i++ {
		select {
		case chunk := <-processedChunks:
			result = append(result, chunk)
		case err := <-errors:
			lastErr = err
		}
	}

	// Return any chunks we did receive: a partial result may be useful
	return result, lastErr
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
