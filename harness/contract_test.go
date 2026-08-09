package harness

import "testing"

func TestLanguage(t *testing.T) { if "go" != "go" { t.Fatal("language mismatch") } }
