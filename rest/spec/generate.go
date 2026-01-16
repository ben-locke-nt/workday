// // Copyright 2025 Nametag Inc.
// //
// // All information contained herein is the property of Nametag Inc.. The
// // intellectual and technical concepts contained herein are proprietary, trade
// // secrets, and/or confidential to Nametag, Inc. and may be covered by U.S.
// // and Foreign Patents, patents in process, and are protected by trade secret or
// // copyright law. Reproduction or distribution, in whole or in part, is
// // forbidden except by express written permission of Nametag, Inc.

// //go:build generate
// // +build generate

package main

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"path/filepath"
// 	"strings"

// 	"github.com/nametaginc/nt/pkg/genx"
// 	"github.com/nametaginc/nt/pkg/must"
// )

// //go:generate go run generate.go

// func main() {
// 	defer genx.Log()()

// 	must.NotFail(genx.Cached(
// 		[]string{"./*"},
// 		[]string{"../model/recruiting.go"},
// 		generateCode))
// }

// func generateCode() error {
// 	paths, err := filepath.Glob("./*.json")
// 	if err != nil {
// 		return err
// 	}

// 	for _, path := range paths {
// 		if !strings.Contains(path, "recruiting") {
// 			continue
// 		}
// 		if err := generateModel(path); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func generateModel(path string) error {
// 	filename := filepath.Base(path)

// 	// The filenames are in the format <component>_whatever.json
// 	parts := strings.SplitN(filename, "_", 2)
// 	if len(parts) < 2 {
// 		return fmt.Errorf("expected 2 parts, not %d, from filename %s", len(parts), filename)
// 	}
// 	modelName := parts[0]

// 	cmd := exec.Command("go", "run",
// 		"github.com/go-swagger/go-swagger/cmd/swagger@latest",
// 		"generate", "model",
// 		"--model-package", "model",
// 		"--target", "../model",
// 		"--name", modelName,
// 		"--spec", path)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	return cmd.Run()
// }
