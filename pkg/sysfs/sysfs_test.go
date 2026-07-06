/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sysfs

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestHostUsesHostRoot(t *testing.T) {
	hostRoot := t.TempDir()
	t.Setenv("HOST_ROOT", hostRoot)

	filename := filepath.Join(hostRoot, "sys", "host-root-test")
	if err := os.MkdirAll(filepath.Dir(filename), 0o755); err != nil {
		t.Fatalf("create fake sysfs: %v", err)
	}
	if err := os.WriteFile(filename, []byte("host sysfs\n"), 0o600); err != nil {
		t.Fatalf("write fake sysfs file: %v", err)
	}

	data, err := fs.ReadFile(Host(), "host-root-test")
	if err != nil {
		t.Fatalf("read through Host(): %v", err)
	}
	if got, want := string(data), "host sysfs\n"; got != want {
		t.Errorf("Host() contents = %q, want %q", got, want)
	}
}
