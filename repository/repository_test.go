package repository

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func cleanup(t *testing.T, dir string) {
	t.Helper()
	if err := os.RemoveAll(dir); err != nil {
		t.Logf("failed to clean up directory %s: %v", dir, err)
	}
}

// TestCloneCheckout walks the happy path of clone -> checkout
func TestCloneCheckout(t *testing.T) {
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
		t.Fatalf("failed to checkout commit hash %s: %v", sha, err)
	}

	command := exec.Command("git", "log", "--oneline")
	command.Dir = dir
	var buf bytes.Buffer

	command.Stdout = &buf

	if err := command.Run(); err != nil {
		t.Fatalf("could not get git log: %v", err)
	}

	if !strings.Contains(buf.String(), sha[:7]) {
		t.Fatalf("checkout and intended sha do not match: got %s, want %s", buf.String(), sha)
	}

	cleanup(t, dir)
}
