// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// AUTOGENERATED. DO NOT EDIT.

// Package starlark is generated by go.chromium.org/luci/tools/cmd/assets.
//
// It contains all [*.css *.html *.js *.star *.tmpl] files found in the package as byte arrays.
package starlark

// GetAsset returns an asset by its name. Returns nil if no such asset exists.
func GetAsset(name string) []byte {
	return []byte(files[name])
}

// GetAssetString is version of GetAsset that returns string instead of byte
// slice. Returns empty string if no such asset exists.
func GetAssetString(name string) string {
	return files[name]
}

// GetAssetSHA256 returns the asset checksum. Returns nil if no such asset
// exists.
func GetAssetSHA256(name string) []byte {
	data := fileSha256s[name]
	if data == nil {
		return nil
	}
	return append([]byte(nil), data...)
}

// Assets returns a map of all assets.
func Assets() map[string]string {
	cpy := make(map[string]string, len(files))
	for k, v := range files {
		cpy[k] = v
	}
	return cpy
}

var files = map[string]string{
	"stdlib/builtins.star": string([]byte{35, 32,
		67, 111, 112, 121, 114, 105, 103, 104, 116, 32, 50, 48, 49, 56,
		32, 84, 104, 101, 32, 76, 85, 67, 73, 32, 65, 117, 116, 104,
		111, 114, 115, 46, 10, 35, 10, 35, 32, 76, 105, 99, 101, 110,
		115, 101, 100, 32, 117, 110, 100, 101, 114, 32, 116, 104, 101, 32,
		65, 112, 97, 99, 104, 101, 32, 76, 105, 99, 101, 110, 115, 101,
		44, 32, 86, 101, 114, 115, 105, 111, 110, 32, 50, 46, 48, 32,
		40, 116, 104, 101, 32, 34, 76, 105, 99, 101, 110, 115, 101, 34,
		41, 59, 10, 35, 32, 121, 111, 117, 32, 109, 97, 121, 32, 110,
		111, 116, 32, 117, 115, 101, 32, 116, 104, 105, 115, 32, 102, 105,
		108, 101, 32, 101, 120, 99, 101, 112, 116, 32, 105, 110, 32, 99,
		111, 109, 112, 108, 105, 97, 110, 99, 101, 32, 119, 105, 116, 104,
		32, 116, 104, 101, 32, 76, 105, 99, 101, 110, 115, 101, 46, 10,
		35, 32, 89, 111, 117, 32, 109, 97, 121, 32, 111, 98, 116, 97,
		105, 110, 32, 97, 32, 99, 111, 112, 121, 32, 111, 102, 32, 116,
		104, 101, 32, 76, 105, 99, 101, 110, 115, 101, 32, 97, 116, 10,
		35, 10, 35, 32, 32, 32, 32, 32, 32, 104, 116, 116, 112, 58,
		47, 47, 119, 119, 119, 46, 97, 112, 97, 99, 104, 101, 46, 111,
		114, 103, 47, 108, 105, 99, 101, 110, 115, 101, 115, 47, 76, 73,
		67, 69, 78, 83, 69, 45, 50, 46, 48, 10, 35, 10, 35, 32,
		85, 110, 108, 101, 115, 115, 32, 114, 101, 113, 117, 105, 114, 101,
		100, 32, 98, 121, 32, 97, 112, 112, 108, 105, 99, 97, 98, 108,
		101, 32, 108, 97, 119, 32, 111, 114, 32, 97, 103, 114, 101, 101,
		100, 32, 116, 111, 32, 105, 110, 32, 119, 114, 105, 116, 105, 110,
		103, 44, 32, 115, 111, 102, 116, 119, 97, 114, 101, 10, 35, 32,
		100, 105, 115, 116, 114, 105, 98, 117, 116, 101, 100, 32, 117, 110,
		100, 101, 114, 32, 116, 104, 101, 32, 76, 105, 99, 101, 110, 115,
		101, 32, 105, 115, 32, 100, 105, 115, 116, 114, 105, 98, 117, 116,
		101, 100, 32, 111, 110, 32, 97, 110, 32, 34, 65, 83, 32, 73,
		83, 34, 32, 66, 65, 83, 73, 83, 44, 10, 35, 32, 87, 73,
		84, 72, 79, 85, 84, 32, 87, 65, 82, 82, 65, 78, 84, 73,
		69, 83, 32, 79, 82, 32, 67, 79, 78, 68, 73, 84, 73, 79,
		78, 83, 32, 79, 70, 32, 65, 78, 89, 32, 75, 73, 78, 68,
		44, 32, 101, 105, 116, 104, 101, 114, 32, 101, 120, 112, 114, 101,
		115, 115, 32, 111, 114, 32, 105, 109, 112, 108, 105, 101, 100, 46,
		10, 35, 32, 83, 101, 101, 32, 116, 104, 101, 32, 76, 105, 99,
		101, 110, 115, 101, 32, 102, 111, 114, 32, 116, 104, 101, 32, 115,
		112, 101, 99, 105, 102, 105, 99, 32, 108, 97, 110, 103, 117, 97,
		103, 101, 32, 103, 111, 118, 101, 114, 110, 105, 110, 103, 32, 112,
		101, 114, 109, 105, 115, 115, 105, 111, 110, 115, 32, 97, 110, 100,
		10, 35, 32, 108, 105, 109, 105, 116, 97, 116, 105, 111, 110, 115,
		32, 117, 110, 100, 101, 114, 32, 116, 104, 101, 32, 76, 105, 99,
		101, 110, 115, 101, 46, 10, 10, 10, 100, 101, 102, 32, 95, 103,
		101, 110, 101, 114, 97, 116, 111, 114, 40, 105, 109, 112, 108, 41,
		58, 10, 32, 32, 34, 34, 34, 82, 101, 103, 105, 115, 116, 101,
		114, 115, 32, 97, 32, 99, 97, 108, 108, 98, 97, 99, 107, 32,
		116, 104, 97, 116, 32, 105, 115, 32, 99, 97, 108, 108, 101, 100,
		32, 97, 116, 32, 116, 104, 101, 32, 101, 110, 100, 32, 111, 102,
		32, 116, 104, 101, 32, 99, 111, 110, 102, 105, 103, 32, 103, 101,
		110, 101, 114, 97, 116, 105, 111, 110, 10, 32, 32, 115, 116, 97,
		103, 101, 32, 116, 111, 32, 109, 111, 100, 105, 102, 121, 47, 97,
		112, 112, 101, 110, 100, 47, 100, 101, 108, 101, 116, 101, 32, 103,
		101, 110, 101, 114, 97, 116, 101, 100, 32, 99, 111, 110, 102, 105,
		103, 115, 32, 105, 110, 32, 97, 110, 32, 97, 114, 98, 105, 116,
		114, 97, 114, 121, 32, 119, 97, 121, 46, 10, 10, 32, 32, 84,
		104, 101, 32, 99, 97, 108, 108, 98, 97, 99, 107, 32, 97, 99,
		99, 101, 112, 116, 115, 32, 115, 105, 110, 103, 108, 101, 32, 97,
		114, 103, 117, 109, 101, 110, 116, 32, 39, 99, 116, 120, 39, 32,
		119, 104, 105, 99, 104, 32, 105, 115, 32, 97, 32, 115, 116, 114,
		117, 99, 116, 32, 119, 105, 116, 104, 32, 102, 111, 108, 108, 111,
		119, 105, 110, 103, 10, 32, 32, 102, 105, 101, 108, 100, 115, 58,
		10, 32, 32, 32, 32, 39, 99, 111, 110, 102, 105, 103, 95, 115,
		101, 116, 39, 58, 32, 97, 32, 100, 105, 99, 116, 32, 123, 99,
		111, 110, 102, 105, 103, 32, 102, 105, 108, 101, 32, 110, 97, 109,
		101, 32, 45, 62, 32, 40, 115, 116, 114, 32, 124, 32, 112, 114,
		111, 116, 111, 41, 125, 46, 10, 10, 32, 32, 84, 104, 101, 32,
		99, 97, 108, 108, 98, 97, 99, 107, 32, 105, 115, 32, 102, 114,
		101, 101, 32, 116, 111, 32, 109, 111, 100, 105, 102, 121, 32, 99,
		116, 120, 46, 99, 111, 110, 102, 105, 103, 95, 115, 101, 116, 32,
		105, 110, 32, 119, 104, 97, 116, 101, 118, 101, 114, 32, 119, 97,
		121, 32, 105, 116, 32, 119, 97, 110, 116, 115, 44, 32, 101, 46,
		103, 46, 10, 32, 32, 98, 121, 32, 97, 100, 100, 105, 110, 103,
		32, 110, 101, 119, 32, 118, 97, 108, 117, 101, 115, 32, 116, 104,
		101, 114, 101, 32, 111, 114, 32, 109, 117, 116, 97, 116, 105, 110,
		103, 47, 100, 101, 108, 101, 116, 105, 110, 103, 32, 101, 120, 105,
		115, 116, 105, 110, 103, 32, 111, 110, 101, 115, 46, 10, 10, 32,
		32, 65, 114, 103, 115, 58, 10, 32, 32, 32, 32, 105, 109, 112,
		108, 58, 32, 97, 32, 99, 97, 108, 108, 98, 97, 99, 107, 32,
		102, 117, 110, 99, 40, 99, 116, 120, 41, 32, 45, 62, 32, 78,
		111, 110, 101, 46, 10, 32, 32, 34, 34, 34, 10, 32, 32, 95,
		95, 110, 97, 116, 105, 118, 101, 95, 95, 46, 97, 100, 100, 95,
		103, 101, 110, 101, 114, 97, 116, 111, 114, 40, 105, 109, 112, 108,
		41, 10, 10, 10, 35, 32, 80, 117, 98, 108, 105, 99, 32, 65,
		80, 73, 46, 10, 99, 111, 114, 101, 32, 61, 32, 115, 116, 114,
		117, 99, 116, 40, 10, 32, 32, 32, 32, 103, 101, 110, 101, 114,
		97, 116, 111, 114, 32, 61, 32, 95, 103, 101, 110, 101, 114, 97,
		116, 111, 114, 44, 10, 41, 10}),
	"stdlib/internal/error.star": string([]byte{35, 32,
		67, 111, 112, 121, 114, 105, 103, 104, 116, 32, 50, 48, 49, 56,
		32, 84, 104, 101, 32, 76, 85, 67, 73, 32, 65, 117, 116, 104,
		111, 114, 115, 46, 10, 35, 10, 35, 32, 76, 105, 99, 101, 110,
		115, 101, 100, 32, 117, 110, 100, 101, 114, 32, 116, 104, 101, 32,
		65, 112, 97, 99, 104, 101, 32, 76, 105, 99, 101, 110, 115, 101,
		44, 32, 86, 101, 114, 115, 105, 111, 110, 32, 50, 46, 48, 32,
		40, 116, 104, 101, 32, 34, 76, 105, 99, 101, 110, 115, 101, 34,
		41, 59, 10, 35, 32, 121, 111, 117, 32, 109, 97, 121, 32, 110,
		111, 116, 32, 117, 115, 101, 32, 116, 104, 105, 115, 32, 102, 105,
		108, 101, 32, 101, 120, 99, 101, 112, 116, 32, 105, 110, 32, 99,
		111, 109, 112, 108, 105, 97, 110, 99, 101, 32, 119, 105, 116, 104,
		32, 116, 104, 101, 32, 76, 105, 99, 101, 110, 115, 101, 46, 10,
		35, 32, 89, 111, 117, 32, 109, 97, 121, 32, 111, 98, 116, 97,
		105, 110, 32, 97, 32, 99, 111, 112, 121, 32, 111, 102, 32, 116,
		104, 101, 32, 76, 105, 99, 101, 110, 115, 101, 32, 97, 116, 10,
		35, 10, 35, 32, 32, 32, 32, 32, 32, 104, 116, 116, 112, 58,
		47, 47, 119, 119, 119, 46, 97, 112, 97, 99, 104, 101, 46, 111,
		114, 103, 47, 108, 105, 99, 101, 110, 115, 101, 115, 47, 76, 73,
		67, 69, 78, 83, 69, 45, 50, 46, 48, 10, 35, 10, 35, 32,
		85, 110, 108, 101, 115, 115, 32, 114, 101, 113, 117, 105, 114, 101,
		100, 32, 98, 121, 32, 97, 112, 112, 108, 105, 99, 97, 98, 108,
		101, 32, 108, 97, 119, 32, 111, 114, 32, 97, 103, 114, 101, 101,
		100, 32, 116, 111, 32, 105, 110, 32, 119, 114, 105, 116, 105, 110,
		103, 44, 32, 115, 111, 102, 116, 119, 97, 114, 101, 10, 35, 32,
		100, 105, 115, 116, 114, 105, 98, 117, 116, 101, 100, 32, 117, 110,
		100, 101, 114, 32, 116, 104, 101, 32, 76, 105, 99, 101, 110, 115,
		101, 32, 105, 115, 32, 100, 105, 115, 116, 114, 105, 98, 117, 116,
		101, 100, 32, 111, 110, 32, 97, 110, 32, 34, 65, 83, 32, 73,
		83, 34, 32, 66, 65, 83, 73, 83, 44, 10, 35, 32, 87, 73,
		84, 72, 79, 85, 84, 32, 87, 65, 82, 82, 65, 78, 84, 73,
		69, 83, 32, 79, 82, 32, 67, 79, 78, 68, 73, 84, 73, 79,
		78, 83, 32, 79, 70, 32, 65, 78, 89, 32, 75, 73, 78, 68,
		44, 32, 101, 105, 116, 104, 101, 114, 32, 101, 120, 112, 114, 101,
		115, 115, 32, 111, 114, 32, 105, 109, 112, 108, 105, 101, 100, 46,
		10, 35, 32, 83, 101, 101, 32, 116, 104, 101, 32, 76, 105, 99,
		101, 110, 115, 101, 32, 102, 111, 114, 32, 116, 104, 101, 32, 115,
		112, 101, 99, 105, 102, 105, 99, 32, 108, 97, 110, 103, 117, 97,
		103, 101, 32, 103, 111, 118, 101, 114, 110, 105, 110, 103, 32, 112,
		101, 114, 109, 105, 115, 115, 105, 111, 110, 115, 32, 97, 110, 100,
		10, 35, 32, 108, 105, 109, 105, 116, 97, 116, 105, 111, 110, 115,
		32, 117, 110, 100, 101, 114, 32, 116, 104, 101, 32, 76, 105, 99,
		101, 110, 115, 101, 46, 10, 10, 100, 101, 102, 32, 101, 114, 114,
		111, 114, 40, 109, 115, 103, 44, 32, 42, 97, 114, 103, 115, 44,
		32, 42, 42, 107, 119, 97, 114, 103, 115, 41, 58, 10, 32, 32,
		34, 34, 34, 69, 109, 105, 116, 115, 32, 97, 110, 32, 101, 114,
		114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 32, 97, 110,
		100, 32, 99, 111, 110, 116, 105, 110, 117, 101, 115, 32, 116, 104,
		101, 32, 101, 120, 101, 99, 117, 116, 105, 111, 110, 46, 10, 10,
		32, 32, 73, 102, 32, 97, 116, 32, 116, 104, 101, 32, 101, 110,
		100, 32, 111, 102, 32, 116, 104, 101, 32, 101, 120, 101, 99, 117,
		116, 105, 111, 110, 32, 116, 104, 101, 114, 101, 39, 115, 32, 97,
		116, 32, 108, 101, 97, 115, 116, 32, 111, 110, 101, 32, 101, 114,
		114, 111, 114, 32, 114, 101, 99, 111, 114, 100, 101, 100, 44, 32,
		116, 104, 101, 10, 32, 32, 101, 120, 101, 99, 117, 116, 105, 111,
		110, 32, 105, 115, 32, 99, 111, 110, 115, 105, 100, 101, 114, 101,
		100, 32, 102, 97, 105, 108, 101, 100, 46, 10, 10, 32, 32, 69,
		105, 116, 104, 101, 114, 32, 99, 97, 112, 116, 117, 114, 101, 115,
		32, 116, 104, 101, 32, 99, 117, 114, 114, 101, 110, 116, 32, 115,
		116, 97, 99, 107, 32, 116, 114, 97, 99, 101, 32, 102, 111, 114,
		32, 116, 104, 101, 32, 116, 114, 97, 99, 101, 98, 97, 99, 107,
		32, 111, 114, 32, 117, 115, 101, 115, 32, 97, 10, 32, 32, 112,
		114, 101, 118, 105, 111, 117, 115, 108, 121, 32, 99, 97, 112, 116,
		117, 114, 101, 100, 32, 111, 110, 101, 32, 105, 102, 32, 105, 116,
		32, 119, 97, 115, 32, 112, 97, 115, 115, 101, 100, 32, 118, 105,
		97, 32, 39, 115, 116, 97, 99, 107, 39, 32, 107, 101, 121, 119,
		111, 114, 100, 32, 97, 114, 103, 117, 109, 101, 110, 116, 58, 10,
		10, 32, 32, 32, 32, 115, 116, 97, 99, 107, 32, 61, 32, 115,
		116, 97, 99, 107, 116, 114, 97, 99, 101, 40, 41, 10, 32, 32,
		32, 32, 46, 46, 46, 10, 32, 32, 32, 32, 101, 114, 114, 111,
		114, 40, 39, 66, 111, 111, 109, 44, 32, 37, 115, 39, 44, 32,
		39, 97, 114, 103, 39, 44, 32, 115, 116, 97, 99, 107, 61, 115,
		116, 97, 99, 107, 41, 10, 10, 32, 32, 65, 114, 103, 115, 58,
		10, 32, 32, 32, 32, 109, 115, 103, 58, 32, 101, 114, 114, 111,
		114, 32, 109, 101, 115, 115, 97, 103, 101, 32, 102, 111, 114, 109,
		97, 116, 32, 115, 116, 114, 105, 110, 103, 46, 10, 32, 32, 32,
		32, 42, 97, 114, 103, 115, 58, 32, 97, 114, 103, 117, 109, 101,
		110, 116, 115, 32, 102, 111, 114, 32, 116, 104, 101, 32, 102, 111,
		114, 109, 97, 116, 32, 115, 116, 114, 105, 110, 103, 46, 10, 32,
		32, 32, 32, 42, 42, 107, 119, 97, 114, 103, 115, 58, 32, 101,
		105, 116, 104, 101, 114, 32, 101, 109, 112, 116, 121, 32, 111, 102,
		32, 99, 111, 110, 116, 97, 105, 110, 115, 32, 115, 105, 110, 103,
		108, 101, 32, 39, 115, 116, 97, 99, 107, 39, 32, 118, 97, 114,
		32, 119, 105, 116, 104, 32, 97, 32, 115, 116, 97, 99, 107, 32,
		116, 114, 97, 99, 101, 46, 10, 32, 32, 34, 34, 34, 10, 32,
		32, 115, 116, 97, 99, 107, 32, 61, 32, 107, 119, 97, 114, 103,
		115, 46, 112, 111, 112, 40, 39, 115, 116, 97, 99, 107, 39, 44,
		32, 78, 111, 110, 101, 41, 32, 111, 114, 32, 115, 116, 97, 99,
		107, 116, 114, 97, 99, 101, 40, 115, 107, 105, 112, 61, 49, 41,
		10, 32, 32, 105, 102, 32, 108, 101, 110, 40, 107, 119, 97, 114,
		103, 115, 41, 32, 33, 61, 32, 48, 58, 10, 32, 32, 32, 32,
		102, 97, 105, 108, 40, 39, 101, 120, 112, 101, 99, 116, 105, 110,
		103, 32, 115, 116, 97, 99, 107, 61, 46, 46, 46, 32, 107, 119,
		97, 114, 103, 115, 32, 111, 110, 108, 121, 39, 41, 10, 32, 32,
		95, 95, 110, 97, 116, 105, 118, 101, 95, 95, 46, 101, 109, 105,
		116, 95, 101, 114, 114, 111, 114, 40, 109, 115, 103, 32, 37, 32,
		97, 114, 103, 115, 44, 32, 115, 116, 97, 99, 107, 41, 10}),
	"stdlib/internal/graph.star": string([]byte{35, 32,
		67, 111, 112, 121, 114, 105, 103, 104, 116, 32, 50, 48, 49, 56,
		32, 84, 104, 101, 32, 76, 85, 67, 73, 32, 65, 117, 116, 104,
		111, 114, 115, 46, 10, 35, 10, 35, 32, 76, 105, 99, 101, 110,
		115, 101, 100, 32, 117, 110, 100, 101, 114, 32, 116, 104, 101, 32,
		65, 112, 97, 99, 104, 101, 32, 76, 105, 99, 101, 110, 115, 101,
		44, 32, 86, 101, 114, 115, 105, 111, 110, 32, 50, 46, 48, 32,
		40, 116, 104, 101, 32, 34, 76, 105, 99, 101, 110, 115, 101, 34,
		41, 59, 10, 35, 32, 121, 111, 117, 32, 109, 97, 121, 32, 110,
		111, 116, 32, 117, 115, 101, 32, 116, 104, 105, 115, 32, 102, 105,
		108, 101, 32, 101, 120, 99, 101, 112, 116, 32, 105, 110, 32, 99,
		111, 109, 112, 108, 105, 97, 110, 99, 101, 32, 119, 105, 116, 104,
		32, 116, 104, 101, 32, 76, 105, 99, 101, 110, 115, 101, 46, 10,
		35, 32, 89, 111, 117, 32, 109, 97, 121, 32, 111, 98, 116, 97,
		105, 110, 32, 97, 32, 99, 111, 112, 121, 32, 111, 102, 32, 116,
		104, 101, 32, 76, 105, 99, 101, 110, 115, 101, 32, 97, 116, 10,
		35, 10, 35, 32, 32, 32, 32, 32, 32, 104, 116, 116, 112, 58,
		47, 47, 119, 119, 119, 46, 97, 112, 97, 99, 104, 101, 46, 111,
		114, 103, 47, 108, 105, 99, 101, 110, 115, 101, 115, 47, 76, 73,
		67, 69, 78, 83, 69, 45, 50, 46, 48, 10, 35, 10, 35, 32,
		85, 110, 108, 101, 115, 115, 32, 114, 101, 113, 117, 105, 114, 101,
		100, 32, 98, 121, 32, 97, 112, 112, 108, 105, 99, 97, 98, 108,
		101, 32, 108, 97, 119, 32, 111, 114, 32, 97, 103, 114, 101, 101,
		100, 32, 116, 111, 32, 105, 110, 32, 119, 114, 105, 116, 105, 110,
		103, 44, 32, 115, 111, 102, 116, 119, 97, 114, 101, 10, 35, 32,
		100, 105, 115, 116, 114, 105, 98, 117, 116, 101, 100, 32, 117, 110,
		100, 101, 114, 32, 116, 104, 101, 32, 76, 105, 99, 101, 110, 115,
		101, 32, 105, 115, 32, 100, 105, 115, 116, 114, 105, 98, 117, 116,
		101, 100, 32, 111, 110, 32, 97, 110, 32, 34, 65, 83, 32, 73,
		83, 34, 32, 66, 65, 83, 73, 83, 44, 10, 35, 32, 87, 73,
		84, 72, 79, 85, 84, 32, 87, 65, 82, 82, 65, 78, 84, 73,
		69, 83, 32, 79, 82, 32, 67, 79, 78, 68, 73, 84, 73, 79,
		78, 83, 32, 79, 70, 32, 65, 78, 89, 32, 75, 73, 78, 68,
		44, 32, 101, 105, 116, 104, 101, 114, 32, 101, 120, 112, 114, 101,
		115, 115, 32, 111, 114, 32, 105, 109, 112, 108, 105, 101, 100, 46,
		10, 35, 32, 83, 101, 101, 32, 116, 104, 101, 32, 76, 105, 99,
		101, 110, 115, 101, 32, 102, 111, 114, 32, 116, 104, 101, 32, 115,
		112, 101, 99, 105, 102, 105, 99, 32, 108, 97, 110, 103, 117, 97,
		103, 101, 32, 103, 111, 118, 101, 114, 110, 105, 110, 103, 32, 112,
		101, 114, 109, 105, 115, 115, 105, 111, 110, 115, 32, 97, 110, 100,
		10, 35, 32, 108, 105, 109, 105, 116, 97, 116, 105, 111, 110, 115,
		32, 117, 110, 100, 101, 114, 32, 116, 104, 101, 32, 76, 105, 99,
		101, 110, 115, 101, 46, 10, 10, 10, 100, 101, 102, 32, 95, 107,
		101, 121, 40, 42, 97, 114, 103, 115, 41, 58, 10, 32, 32, 34,
		34, 34, 82, 101, 116, 117, 114, 110, 115, 32, 97, 32, 107, 101,
		121, 32, 119, 105, 116, 104, 32, 103, 105, 118, 101, 110, 32, 91,
		40, 107, 105, 110, 100, 44, 32, 110, 97, 109, 101, 41, 93, 32,
		112, 97, 116, 104, 46, 10, 10, 32, 32, 65, 114, 103, 115, 58,
		10, 32, 32, 32, 32, 42, 97, 114, 103, 115, 58, 32, 101, 118,
		101, 110, 32, 110, 117, 109, 98, 101, 114, 32, 111, 102, 32, 115,
		116, 114, 105, 110, 103, 115, 58, 32, 107, 105, 110, 100, 49, 44,
		32, 110, 97, 109, 101, 49, 44, 32, 107, 105, 110, 100, 50, 44,
		32, 110, 97, 109, 101, 50, 44, 32, 46, 46, 46, 10, 10, 32,
		32, 82, 101, 116, 117, 114, 110, 115, 58, 10, 32, 32, 32, 32,
		103, 114, 97, 112, 104, 46, 107, 101, 121, 32, 111, 98, 106, 101,
		99, 116, 32, 114, 101, 112, 114, 101, 115, 101, 110, 116, 105, 110,
		103, 32, 116, 104, 105, 115, 32, 112, 97, 116, 104, 46, 10, 32,
		32, 34, 34, 34, 10, 32, 32, 114, 101, 116, 117, 114, 110, 32,
		95, 95, 110, 97, 116, 105, 118, 101, 95, 95, 46, 103, 114, 97,
		112, 104, 40, 41, 46, 107, 101, 121, 40, 42, 97, 114, 103, 115,
		41, 10, 10, 10, 100, 101, 102, 32, 95, 97, 100, 100, 95, 110,
		111, 100, 101, 40, 107, 101, 121, 44, 32, 112, 114, 111, 112, 115,
		61, 78, 111, 110, 101, 44, 32, 116, 114, 97, 99, 101, 61, 78,
		111, 110, 101, 41, 58, 10, 32, 32, 34, 34, 34, 65, 100, 100,
		115, 32, 97, 32, 110, 111, 100, 101, 32, 116, 111, 32, 116, 104,
		101, 32, 103, 114, 97, 112, 104, 32, 111, 114, 32, 102, 97, 105,
		108, 115, 32, 105, 102, 32, 115, 117, 99, 104, 32, 110, 111, 100,
		101, 32, 97, 108, 114, 101, 97, 100, 121, 32, 101, 120, 105, 115,
		116, 115, 46, 10, 10, 32, 32, 65, 108, 115, 111, 32, 102, 97,
		105, 108, 115, 32, 105, 102, 32, 117, 115, 101, 100, 32, 102, 114,
		111, 109, 32, 97, 32, 103, 101, 110, 101, 114, 97, 116, 111, 114,
		32, 99, 97, 108, 108, 98, 97, 99, 107, 58, 32, 97, 116, 32,
		116, 104, 105, 115, 32, 112, 111, 105, 110, 116, 32, 116, 104, 101,
		32, 103, 114, 97, 112, 104, 32, 105, 115, 10, 32, 32, 102, 114,
		111, 122, 101, 110, 32, 97, 110, 100, 32, 99, 97, 110, 39, 116,
		32, 98, 101, 32, 101, 120, 116, 101, 110, 100, 101, 100, 46, 10,
		10, 32, 32, 65, 114, 103, 115, 58, 10, 32, 32, 32, 32, 107,
		101, 121, 58, 32, 97, 32, 110, 111, 100, 101, 32, 107, 101, 121,
		44, 32, 97, 115, 32, 114, 101, 116, 117, 114, 110, 101, 100, 32,
		98, 121, 32, 103, 114, 97, 112, 104, 46, 107, 101, 121, 40, 46,
		46, 46, 41, 46, 10, 32, 32, 32, 32, 112, 114, 111, 112, 115,
		58, 32, 97, 32, 100, 105, 99, 116, 32, 119, 105, 116, 104, 32,
		110, 111, 100, 101, 32, 112, 114, 111, 112, 101, 114, 116, 105, 101,
		115, 44, 32, 119, 105, 108, 108, 32, 98, 101, 32, 102, 114, 111,
		122, 101, 110, 46, 10, 32, 32, 32, 32, 116, 114, 97, 99, 101,
		58, 32, 97, 32, 115, 116, 97, 99, 107, 32, 116, 114, 97, 99,
		101, 32, 116, 111, 32, 97, 115, 115, 111, 99, 105, 97, 116, 101,
		32, 119, 105, 116, 104, 32, 116, 104, 101, 32, 110, 111, 100, 101,
		46, 10, 10, 32, 32, 82, 101, 116, 117, 114, 110, 115, 58, 10,
		32, 32, 32, 32, 103, 114, 97, 112, 104, 46, 110, 111, 100, 101,
		32, 111, 98, 106, 101, 99, 116, 32, 114, 101, 112, 114, 101, 115,
		101, 110, 116, 105, 110, 103, 32, 116, 104, 101, 32, 110, 111, 100,
		101, 46, 10, 32, 32, 34, 34, 34, 10, 32, 32, 114, 101, 116,
		117, 114, 110, 32, 95, 95, 110, 97, 116, 105, 118, 101, 95, 95,
		46, 103, 114, 97, 112, 104, 40, 41, 46, 97, 100, 100, 95, 110,
		111, 100, 101, 40, 10, 32, 32, 32, 32, 32, 32, 107, 101, 121,
		44, 32, 112, 114, 111, 112, 115, 32, 111, 114, 32, 123, 125, 44,
		32, 116, 114, 97, 99, 101, 32, 111, 114, 32, 115, 116, 97, 99,
		107, 116, 114, 97, 99, 101, 40, 115, 107, 105, 112, 61, 49, 41,
		41, 10, 10, 10, 100, 101, 102, 32, 95, 110, 111, 100, 101, 40,
		107, 101, 121, 41, 58, 10, 32, 32, 34, 34, 34, 82, 101, 116,
		117, 114, 110, 115, 32, 97, 32, 110, 111, 100, 101, 32, 98, 121,
		32, 116, 104, 101, 32, 107, 101, 121, 32, 111, 114, 32, 78, 111,
		110, 101, 32, 105, 102, 32, 105, 116, 32, 119, 97, 115, 110, 39,
		116, 32, 97, 100, 100, 101, 100, 32, 98, 121, 32, 97, 100, 100,
		95, 110, 111, 100, 101, 32, 121, 101, 116, 46, 10, 10, 32, 32,
		65, 114, 103, 115, 58, 10, 32, 32, 32, 32, 107, 101, 121, 58,
		32, 97, 32, 110, 111, 100, 101, 32, 107, 101, 121, 44, 32, 97,
		115, 32, 114, 101, 116, 117, 114, 110, 101, 100, 32, 98, 121, 32,
		103, 114, 97, 112, 104, 46, 107, 101, 121, 40, 46, 46, 46, 41,
		46, 10, 10, 32, 32, 82, 101, 116, 117, 114, 110, 115, 58, 10,
		32, 32, 32, 32, 103, 114, 97, 112, 104, 46, 110, 111, 100, 101,
		32, 111, 98, 106, 101, 99, 116, 32, 114, 101, 112, 114, 101, 115,
		101, 110, 116, 105, 110, 103, 32, 116, 104, 101, 32, 110, 111, 100,
		101, 46, 10, 32, 32, 34, 34, 34, 10, 32, 32, 114, 101, 116,
		117, 114, 110, 32, 95, 95, 110, 97, 116, 105, 118, 101, 95, 95,
		46, 103, 114, 97, 112, 104, 40, 41, 46, 110, 111, 100, 101, 40,
		107, 101, 121, 41, 10, 10, 10, 35, 32, 80, 117, 98, 108, 105,
		99, 32, 65, 80, 73, 32, 111, 102, 32, 116, 104, 105, 115, 32,
		109, 111, 100, 117, 108, 101, 46, 10, 103, 114, 97, 112, 104, 32,
		61, 32, 115, 116, 114, 117, 99, 116, 40, 10, 32, 32, 32, 32,
		107, 101, 121, 32, 61, 32, 95, 107, 101, 121, 44, 10, 32, 32,
		32, 32, 97, 100, 100, 95, 110, 111, 100, 101, 32, 61, 32, 95,
		97, 100, 100, 95, 110, 111, 100, 101, 44, 10, 32, 32, 32, 32,
		110, 111, 100, 101, 32, 61, 32, 95, 110, 111, 100, 101, 44, 10,
		41, 10}),
}

var fileSha256s = map[string][]byte{
	"stdlib/builtins.star": {18, 209,
		53, 134, 61, 133, 100, 71, 17, 158, 27, 42, 206, 238, 237, 217,
		152, 60, 175, 220, 45, 217, 129, 154, 88, 154, 205, 142, 252, 68,
		204, 66},
	"stdlib/internal/error.star": {28, 176,
		196, 198, 201, 60, 241, 232, 166, 147, 243, 145, 171, 234, 210, 55,
		95, 17, 196, 214, 247, 196, 95, 88, 10, 62, 153, 127, 126, 141,
		192, 66},
	"stdlib/internal/graph.star": {255, 81,
		5, 166, 65, 195, 9, 50, 59, 173, 149, 115, 220, 62, 199, 226,
		203, 227, 139, 210, 129, 173, 207, 240, 81, 56, 122, 59, 126, 38,
		133, 15},
}
