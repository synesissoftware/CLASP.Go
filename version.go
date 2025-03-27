// Copyright 2019-2025 Matthew Wilson and Synesis Information Systems.
// Copyright 2015-2019 Matthew Wilson. All rights reserved. Use of this
// source code is governed by a BSD-style license that can be found in the
// LICENSE file.

/*
 * Created: 15th August 2015
 * Updated: 27th March 2025
 */

package clasp

import "github.com/synesissoftware/ver2go"

const (
	VersionMajor uint16 = 0
	VersionMinor uint16 = 17
	VersionPatch uint16 = 0
	VersionAB    uint16 = 0xFFFF
	Version      uint64 = (uint64(VersionMajor) << 48) + (uint64(VersionMinor) << 32) + (uint64(VersionPatch) << 16) + (uint64(VersionAB) << 0)
)

var (
	versionString string = ver2go.CalcVersionString(VersionMajor, VersionMinor, VersionPatch, VersionAB)
)

func VersionString() string {
	return versionString
}

/* ///////////////////////////// end of file //////////////////////////// */
