package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestPathTransFormFunc(t *testing.T) {
	key := "momsbestpicture"
	pathKey := CASPathTransFormFunc(key)

	expectedHash := "6804429f74181a63c50c3d81d733a12f14a353ff"
	expectedPath := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"

	if pathKey.Pathname != expectedPath {
		t.Errorf("Have %s want %s", pathKey.Pathname, expectedPath)
	}

	if pathKey.Filename != expectedHash {
		t.Errorf("Have %s want %s", pathKey.Filename, expectedHash)
	}
}

func TestStoreDeleteKey(t *testing.T) {
	opts := storeOpts{
		PathtransFormFunc: CASPathTransFormFunc,
	}
	s := NewStore(opts)

	key := "momsbestpicture"
	data := []byte("some jpg bytes")

	// Step 1: Write file
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Fatalf("write failed: %v", err)
	}

	// Step 2: Ensure file exists before delete
	pathKey := CASPathTransFormFunc(key)
	if _, err := os.Stat(pathKey.FullPath()); os.IsNotExist(err) {
		t.Fatalf("file does not exist before delete")
	}

	// Step 3: Delete file
	if err := s.Delete(key); err != nil {
		t.Fatalf("delete failed: %v", err)
	}

	// Step 4: Ensure file is deleted
	if _, err := os.Stat(pathKey.FullPath()); !os.IsNotExist(err) {
		t.Fatalf("file still exists after delete")
	}
}

func TestStore(t *testing.T) {
	opts := storeOpts{
		PathtransFormFunc: CASPathTransFormFunc,
	}

	s := NewStore(opts)

	key := "momsSpecials"
	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := ioutil.ReadAll(r)

	if string(b) != string(data) {
		t.Errorf("Want %s have %s", data, b)
	}
}
