package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersionCmdPrintsAVersion(t *testing.T) {
	cmd := versionCmd()
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetArgs(nil)
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	got := strings.TrimSpace(buf.String())
	if got == "" {
		t.Fatal("version printed nothing")
	}
}

func TestModuleVersionNeverEmpty(t *testing.T) {
	if got := moduleVersion(); got == "" {
		t.Fatal("moduleVersion() is empty")
	}
}
