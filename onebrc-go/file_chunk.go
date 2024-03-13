package main

import (
	"bytes"
	"errors"
	"io"
	"math"
	"os"
	"sync"
)

type FileChunkReader struct {
	// f is the underlining file that need to be read.
	fileName string
	chunks   []chunk
}

type chunk struct {
	buf *bytes.Buffer
}

func NewFileChunkReader(fileName string) *FileChunkReader {
	return &FileChunkReader{
		fileName: fileName,
		chunks:   make([]chunk, 0, 100),
	}
}

// NumChunks returns number of chunks in FileChunkReader.
func (fr *FileChunkReader) NumChunks() int {
	return len(fr.chunks)
}

func (fr *FileChunkReader) GetChunk(idx int) (chunk, error) {
	if idx >= len(fr.chunks) || idx < 0 {
		return chunk{}, errors.New("Invalid index")
	}
	return fr.chunks[idx], nil
}

const CHUNK_SIZE = 8192

func (fr *FileChunkReader) ReadAll() ([]byte, error) {
	f, err := os.Open(fr.fileName)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	//fmt.Println(info.Size())
	totalSize := info.Size()
	numOfChunks := int(math.Ceil(float64(totalSize) / float64(CHUNK_SIZE)))
	//fmt.Println("numOfChunks", numOfChunks)
	fr.chunks = make([]chunk, numOfChunks)

	errChan := make(chan error, numOfChunks)

	wg := &sync.WaitGroup{}
	var mu sync.Mutex
	for i := 0; i < numOfChunks; i++ {
		wg.Add(1)
		go func(i int) {

			defer wg.Done()

			offset := i * CHUNK_SIZE
			// What might go wrong ?
			dataBytes := make([]byte, CHUNK_SIZE)
			f, err := os.Open(fr.fileName)
			defer f.Close()
			if err != nil {
				errChan <- err
				return
			}
			n, err := f.ReadAt(dataBytes, int64(offset))
			if err != nil {
				switch err {
				case io.EOF:
					break
				default:
					errChan <- err
					return
				}
			}
			if n < CHUNK_SIZE {
				// shrink the data
				dataBytes = dataBytes[:n]
			}
			mu.Lock()
			defer mu.Unlock()
			fr.chunks[i] = chunk{buf: bytes.NewBuffer(dataBytes)}
			errChan <- nil
		}(i)
	}

	wg.Wait()

	for i := 0; i < numOfChunks; i++ {
		select {
		case err := <-errChan:
			if err != nil {
				return nil, err
			}
		default:
		}
	}

	var buf bytes.Buffer
	for _, ck := range fr.chunks {
		buf.WriteString(ck.buf.String())
	}
	return buf.Bytes(), nil
}
