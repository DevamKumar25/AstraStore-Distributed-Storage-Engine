package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func CASPathTransFormFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])

	blocksize := 5
	sliceLen := len(hashStr) / blocksize
	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blocksize, (i*blocksize)+blocksize
		paths[i] = hashStr[from:to]
	}

	return PathKey{
		Pathname: strings.Join(paths, "/"),
		Filename: hashStr,
	}

}

type PathtransFormFunc func(string) PathKey

type PathKey struct {
	Pathname string
	Filename string
}

func (p PathKey) FullPath() string {
	return fmt.Sprintf("%s%s", p.Pathname, p.Filename)
}

type storeOpts struct {
	PathtransFormFunc PathtransFormFunc
}

type Store struct {
	storeOpts
}

var DefaultPathTransFormFunc = func(key string) string {
	return key
}

func NewStore(opts storeOpts) *Store {
	return &Store{
		storeOpts: opts,
	}
}


func (s *Store) Delete(key string) error {
	pathKey := s.PathtransFormFunc(key)

	defer func() {
		log.Printf("Deleted [%s] from the disk",pathKey.Filename)
	}()

	return os.RemoveAll(pathKey.FullPath())
		
}

func (s *Store) Read(key string) (io.Reader, error) {
	f, err := s.readStream(key)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)
	return buf, err
}

func (s *Store) readStream(key string) (io.ReadCloser, error) {
	pathkey := s.PathtransFormFunc(key)

	return os.Open(pathkey.FullPath())
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathKey := s.PathtransFormFunc(key)
	if err := os.MkdirAll(pathKey.Pathname, os.ModePerm); err != nil {
		return err
	}

	fullPath := pathKey.FullPath()
	log.Printf("Writing to disk: %s", fullPath)

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}

	log.Printf("Written (%d) bytes to disk: %s", n, fullPath)

	return nil
}
