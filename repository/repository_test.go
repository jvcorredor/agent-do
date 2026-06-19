package repository

import (
	"testing"
	"os"
	"path/filepath"
)

func cleanup(t *testing.T, dir string) {
	t.Helper()
	if err := os.Remove(dir); err != nil {
		t.Logf("failed to clean up directory: %s", dir)
	}
}

func TestClone(t *testing.T) {
	r := New("git@github.com:usememos/memos.git")
	dir, err := r.Clone()
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	fileInfo, err := os.Stat(filepath.Join(dir, "README.md"))
	if err != nil {
		t.Fatalf("could not stat: %v", err)
	}

	if fileInfo.Size() <= 0 {
		t.Fatalf("file size <= 0: %d", fileInfo.Size())
	}

	sha := "797f1ff15dcb94543ce15462f7cfc8d292f2ffa7"
	if err := r.Checkout(sha); err != nil {
		t.Fatalf("failed to checkout commit hash: %s", sha)
	}

	cleanup(t, dir)
}

