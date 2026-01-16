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
// 		[]string{
// 			"../model/human_resources/human_resources.go",
// 			"../model/recruiting/recruiting.go",
// 			"../model/staffing/staffing.go"},
// 		generateCode))
// }

// func generateCode() error {
// 	paths, err := filepath.Glob("./*.wsdl")
// 	if err != nil {
// 		return err
// 	}

// 	for _, path := range paths {
// 		if err := generateModel(path); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// // Unfortunately, gowsdl enstructifies one of the fundamental types, the Reference ID type, as
// // a raw string, but it actually has a required attribute "type".
// // https://github.com/fiorix/wsdl2go serializes this correctly, but the whole file it outputs
// // is missing definitions for several types, so it doesn't compile!
// func generateModel(path string) error {
// 	modelName := strings.ToLower(strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)))
// 	cmd := exec.Command("go", "run",
// 		"github.com/hooklift/gowsdl/cmd/gowsdl@latest",
// 		"-d", "../model",
// 		"-p", modelName,
// 		"-o", modelName+".go",
// 		path)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	return cmd.Run()
// }
