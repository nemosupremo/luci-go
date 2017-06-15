// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package cipd

import (
	"fmt"
	"testing"

	"github.com/luci/luci-go/vpython/api/vpython"

	. "github.com/smartystreets/goconvey/convey"
)

func mkTag(plat string) *vpython.PEP425Tag {
	return &vpython.PEP425Tag{
		Python:   "cp27",
		Abi:      "none",
		Platform: plat,
	}
}

func TestPlatformForPEP425Tag(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		tag      *vpython.PEP425Tag
		platform string
	}{
		{mkTag("junk_i686"), ""},

		{mkTag("linux_sparc64"), ""},
		{mkTag("linux_i686"), "linux-386"},
		{mkTag("manylinux1_i686"), "linux-386"},
		{mkTag("linux_x86_64"), "linux-amd64"},
		{mkTag("manylinux1_x86_64"), "linux-amd64"},
		{mkTag("linux_arm64"), "linux-arm64"},
		{mkTag("linux_armv6"), "linux-armv6l"},
		{mkTag("linux_armv7"), "linux-armv6l"},
		{mkTag("linux_mips"), "linux-mips32"},
		{mkTag("linux_mips64"), "linux-mips64"},

		{mkTag("macosx_12_12_pants"), ""},
		{mkTag("macosx_12_12_fat32"), "mac-386"},
		{mkTag("macosx_10_10_intel"), "mac-amd64"},
		{mkTag("macosx_10_9_universal"), "mac-amd64"},
		{mkTag("macosx_12_12_fat64"), "mac-amd64"},

		{mkTag("win_pants"), ""},
		{mkTag("win32"), "windows-386"},
		{mkTag("win_amd64"), "windows-amd64"},
	}

	Convey(`Testing PEP425 tag selection`, t, func() {
		for _, tc := range testCases {
			Convey(fmt.Sprintf("Tag %q => %q", tc.tag.TagString(), tc.platform), func() {
				So(PlatformForPEP425Tag(tc.tag), ShouldResemble, tc.platform)
			})
		}
	})
}