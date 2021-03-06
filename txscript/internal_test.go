// Copyright (c) 2013-2015 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package txscript

import "testing"

// TstMaxScriptSize makes the internal maxScriptSize constant available to the
// test package.
const TstMaxScriptSize = maxScriptSize

// TstHasCanoncialPushes makes the internal isCanonicalPush function available
// to the test package.
var TstHasCanonicalPushes = canonicalPush

// TstParseScript makes the internal parseScript function available to the
// test package.
var TstParseScript = parseScript

// this file is present to export some internal interfaces so that we can
// test them reliably.

func TestCheckPubKeyEncoding(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		key     []byte
		isValid bool
	}{
		{
			name: "uncompressed ok",
			key: []byte{0x04, 0x11, 0xdb, 0x93, 0xe1, 0xdc, 0xdb, 0x8a,
				0x01, 0x6b, 0x49, 0x84, 0x0f, 0x8c, 0x53, 0xbc, 0x1e,
				0xb6, 0x8a, 0x38, 0x2e, 0x97, 0xb1, 0x48, 0x2e, 0xca,
				0xd7, 0xb1, 0x48, 0xa6, 0x90, 0x9a, 0x5c, 0xb2, 0xe0,
				0xea, 0xdd, 0xfb, 0x84, 0xcc, 0xf9, 0x74, 0x44, 0x64,
				0xf8, 0x2e, 0x16, 0x0b, 0xfa, 0x9b, 0x8b, 0x64, 0xf9,
				0xd4, 0xc0, 0x3f, 0x99, 0x9b, 0x86, 0x43, 0xf6, 0x56,
				0xb4, 0x12, 0xa3,
			},
			isValid: true,
		},
		{
			name: "compressed ok",
			key: []byte{0x02, 0xce, 0x0b, 0x14, 0xfb, 0x84, 0x2b, 0x1b,
				0xa5, 0x49, 0xfd, 0xd6, 0x75, 0xc9, 0x80, 0x75, 0xf1,
				0x2e, 0x9c, 0x51, 0x0f, 0x8e, 0xf5, 0x2b, 0xd0, 0x21,
				0xa9, 0xa1, 0xf4, 0x80, 0x9d, 0x3b, 0x4d,
			},
			isValid: true,
		},
		{
			name: "compressed ok",
			key: []byte{0x03, 0x26, 0x89, 0xc7, 0xc2, 0xda, 0xb1, 0x33,
				0x09, 0xfb, 0x14, 0x3e, 0x0e, 0x8f, 0xe3, 0x96, 0x34,
				0x25, 0x21, 0x88, 0x7e, 0x97, 0x66, 0x90, 0xb6, 0xb4,
				0x7f, 0x5b, 0x2a, 0x4b, 0x7d, 0x44, 0x8e,
			},
			isValid: true,
		},
		{
			name: "hybrid",
			key: []byte{0x06, 0x79, 0xbe, 0x66, 0x7e, 0xf9, 0xdc, 0xbb,
				0xac, 0x55, 0xa0, 0x62, 0x95, 0xce, 0x87, 0x0b, 0x07,
				0x02, 0x9b, 0xfc, 0xdb, 0x2d, 0xce, 0x28, 0xd9, 0x59,
				0xf2, 0x81, 0x5b, 0x16, 0xf8, 0x17, 0x98, 0x48, 0x3a,
				0xda, 0x77, 0x26, 0xa3, 0xc4, 0x65, 0x5d, 0xa4, 0xfb,
				0xfc, 0x0e, 0x11, 0x08, 0xa8, 0xfd, 0x17, 0xb4, 0x48,
				0xa6, 0x85, 0x54, 0x19, 0x9c, 0x47, 0xd0, 0x8f, 0xfb,
				0x10, 0xd4, 0xb8,
			},
			isValid: false,
		},
		{
			name:    "empty.",
			key:     []byte{},
			isValid: false,
		},
	}
	vm := Engine{flags: ScriptVerifyStrictEncoding}
	for _, test := range tests {
		err := vm.checkPubKeyEncoding(test.key)
		if err != nil && test.isValid {
			t.Errorf("checkSignatureEncoding test '%s' failed "+
				"when it should have succeeded: %v", test.name,
				err)
		} else if err == nil && !test.isValid {
			t.Errorf("checkSignatureEncooding test '%s' succeeded "+
				"when it should have failed", test.name)
		}
	}

}

func TestCheckSignatureEncoding(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		sig     []byte
		isValid bool
	}{
		{
			name: "valid signature.",
			sig: []byte{0x30, 0x44, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
				0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
				0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
				0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
				0x41, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
				0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
				0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
				0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
			},
			isValid: true,
		},
		{
			name:    "empty.",
			sig:     []byte{},
			isValid: false,
		},
		{
			name: "bad magic.",
			sig: []byte{0x31, 0x44, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
				0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
				0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
				0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
				0x41, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
				0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
				0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
				0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
			},
			isValid: false,
		},
		{
			name: "bad 1st int marker magic.",
			sig: []byte{0x30, 0x44, 0x03, 0x20, 0x4e, 0x45, 0xe1, 0x69,
				0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
				0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
				0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
				0x41, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
				0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
				0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
				0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
			},
			isValid: false,
		},
		{
			name: "bad 2nd int marker.",
			sig: []byte{0x30, 0x44, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
				0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
				0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
				0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
				0x41, 0x03, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
				0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
				0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
				0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
			},
			isValid: false,
		},
		{
			name: "short len",
			sig: []byte{0x30, 0x43, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
				0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
				0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
				0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
				0x41, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
				0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
				0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
				0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
			},
			isValid: false,
		},
		{
			name: "long len",
			sig: []byte{0x30, 0x45, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
				0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
				0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
				0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
				0x41, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
				0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
				0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
				0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
			},
			isValid: false,
		},
		{
			name: "long X",
			sig: []byte{0x30, 0x44, 0x02, 0x42, 0x4e, 0x45, 0xe1, 0x69,
				0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
				0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
				0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
				0x41, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
				0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
				0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
				0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
			},
			isValid: false,
		},
		{
			name: "long Y",
			sig: []byte{0x30, 0x44, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
				0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
				0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
				0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
				0x41, 0x02, 0x21, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
				0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
				0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
				0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
			},
			isValid: false,
		},
		{
			name: "short Y",
			sig: []byte{0x30, 0x44, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
				0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
				0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
				0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
				0x41, 0x02, 0x19, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
				0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
				0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
				0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
			},
			isValid: false,
		},
		{
			name: "trailing crap.",
			sig: []byte{0x30, 0x44, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
				0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
				0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
				0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
				0x41, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
				0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
				0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
				0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09, 0x01,
			},
			isValid: false,
		},
		{
			name: "X == N ",
			sig: []byte{0x30, 0x44, 0x02, 0x20, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF, 0xFF, 0xFE, 0xBA, 0xAE, 0xDC, 0xE6, 0xAF, 0x48,
				0xA0, 0x3B, 0xBF, 0xD2, 0x5E, 0x8C, 0xD0, 0x36, 0x41,
				0x41, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
				0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
				0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
				0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
			},
			isValid: false,
		},
		{
			name: "X == N ",
			sig: []byte{0x30, 0x44, 0x02, 0x20, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF, 0xFF, 0xFE, 0xBA, 0xAE, 0xDC, 0xE6, 0xAF, 0x48,
				0xA0, 0x3B, 0xBF, 0xD2, 0x5E, 0x8C, 0xD0, 0x36, 0x41,
				0x42, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
				0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
				0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
				0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
			},
			isValid: false,
		},
		{
			name: "Y == N",
			sig: []byte{0x30, 0x44, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
				0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
				0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
				0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
				0x41, 0x02, 0x20, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFE, 0xBA, 0xAE, 0xDC, 0xE6, 0xAF, 0x48, 0xA0, 0x3B,
				0xBF, 0xD2, 0x5E, 0x8C, 0xD0, 0x36, 0x41, 0x41,
			},
			isValid: false,
		},
		{
			name: "Y > N",
			sig: []byte{0x30, 0x44, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
				0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
				0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
				0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
				0x41, 0x02, 0x20, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFE, 0xBA, 0xAE, 0xDC, 0xE6, 0xAF, 0x48, 0xA0, 0x3B,
				0xBF, 0xD2, 0x5E, 0x8C, 0xD0, 0x36, 0x41, 0x42,
			},
			isValid: false,
		},
		{
			name: "0 len X.",
			sig: []byte{0x30, 0x24, 0x02, 0x00, 0x02, 0x20, 0x18, 0x15,
				0x22, 0xec, 0x8e, 0xca, 0x07, 0xde, 0x48, 0x60, 0xa4,
				0xac, 0xdd, 0x12, 0x90, 0x9d, 0x83, 0x1c, 0xc5, 0x6c,
				0xbb, 0xac, 0x46, 0x22, 0x08, 0x22, 0x21, 0xa8, 0x76,
				0x8d, 0x1d, 0x09,
			},
			isValid: false,
		},
		{
			name: "0 len Y.",
			sig: []byte{0x30, 0x24, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
				0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
				0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
				0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
				0x41, 0x02, 0x00,
			},
			isValid: false,
		},
		{
			name: "extra R padding.",
			sig: []byte{0x30, 0x45, 0x02, 0x21, 0x00, 0x4e, 0x45, 0xe1, 0x69,
				0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
				0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
				0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
				0x41, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
				0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
				0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
				0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
			},
			isValid: false,
		},
		{
			name: "extra S padding.",
			sig: []byte{0x30, 0x45, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
				0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
				0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
				0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
				0x41, 0x02, 0x21, 0x00, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
				0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
				0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
				0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
			},
			isValid: false,
		},
	}

	vm := Engine{flags: ScriptVerifyStrictEncoding}
	for _, test := range tests {
		err := vm.checkSignatureEncoding(test.sig)
		if err != nil && test.isValid {
			t.Errorf("checkSignatureEncoding test '%s' failed "+
				"when it should have succeeded: %v", test.name,
				err)
		} else if err == nil && !test.isValid {
			t.Errorf("checkSignatureEncooding test '%s' succeeded "+
				"when it should have failed", test.name)
		}
	}
}

func TstRemoveOpcode(pkscript []byte, opcode byte) ([]byte, error) {
	pops, err := parseScript(pkscript)
	if err != nil {
		return nil, err
	}
	pops = removeOpcode(pops, opcode)
	return unparseScript(pops)
}

func TstRemoveOpcodeByData(pkscript []byte, data []byte) ([]byte, error) {
	pops, err := parseScript(pkscript)
	if err != nil {
		return nil, err
	}
	pops = removeOpcodeByData(pops, data)
	return unparseScript(pops)
}

// TestSetPC allows the test modules to set the program counter to whatever they
// want.
func (vm *Engine) TstSetPC(script, off int) {
	vm.scriptIdx = script
	vm.scriptOff = off
}

// Internal tests for opcodde parsing with bad data templates.
func TestParseOpcode(t *testing.T) {
	// Deep copy the array.
	fakeArray := opcodeArray
	// wrong length -8.
	fakeArray[OP_PUSHDATA4] = opcode{value: OP_PUSHDATA4,
		name: "OP_PUSHDATA4", length: -8, opfunc: opcodePushData}

	// this script would be fine if -8 was a valid length.
	_, err := parseScriptTemplate([]byte{OP_PUSHDATA4, 0x1, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00}, &fakeArray)
	if err == nil {
		t.Errorf("no error with dodgy opcode array!")
	}
}

type popTest struct {
	name        string
	pop         *parsedOpcode
	expectedErr error
}

var popTests = []popTest{
	{
		name: "OP_FALSE",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_FALSE],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_FALSE long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_FALSE],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_1 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_1],
			data:   nil,
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_1",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_1],
			data:   make([]byte, 1),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_1 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_1],
			data:   make([]byte, 2),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_2 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_2],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_2",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_2],
			data:   make([]byte, 2),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_2 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_2],
			data:   make([]byte, 3),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_3 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_3],
			data:   make([]byte, 2),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_3",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_3],
			data:   make([]byte, 3),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_3 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_3],
			data:   make([]byte, 4),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_4 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_4],
			data:   make([]byte, 3),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_4",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_4],
			data:   make([]byte, 4),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_4 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_4],
			data:   make([]byte, 5),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_5 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_5],
			data:   make([]byte, 4),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_5",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_5],
			data:   make([]byte, 5),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_5 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_5],
			data:   make([]byte, 6),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_6 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_6],
			data:   make([]byte, 5),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_6",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_6],
			data:   make([]byte, 6),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_6 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_6],
			data:   make([]byte, 7),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_7 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_7],
			data:   make([]byte, 6),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_7",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_7],
			data:   make([]byte, 7),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_7 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_7],
			data:   make([]byte, 8),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_8 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_8],
			data:   make([]byte, 7),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_8",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_8],
			data:   make([]byte, 8),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_8 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_8],
			data:   make([]byte, 9),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_9 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_9],
			data:   make([]byte, 8),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_9",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_9],
			data:   make([]byte, 9),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_9 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_9],
			data:   make([]byte, 10),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_10 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_10],
			data:   make([]byte, 9),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_10",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_10],
			data:   make([]byte, 10),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_10 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_10],
			data:   make([]byte, 11),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_11 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_11],
			data:   make([]byte, 10),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_11",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_11],
			data:   make([]byte, 11),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_11 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_11],
			data:   make([]byte, 12),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_12 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_12],
			data:   make([]byte, 11),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_12",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_12],
			data:   make([]byte, 12),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_12 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_12],
			data:   make([]byte, 13),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_13 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_13],
			data:   make([]byte, 12),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_13",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_13],
			data:   make([]byte, 13),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_13 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_13],
			data:   make([]byte, 14),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_14 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_14],
			data:   make([]byte, 13),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_14",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_14],
			data:   make([]byte, 14),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_14 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_14],
			data:   make([]byte, 15),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_15 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_15],
			data:   make([]byte, 14),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_15",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_15],
			data:   make([]byte, 15),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_15 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_15],
			data:   make([]byte, 16),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_16 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_16],
			data:   make([]byte, 15),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_16",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_16],
			data:   make([]byte, 16),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_16 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_16],
			data:   make([]byte, 17),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_17 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_17],
			data:   make([]byte, 16),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_17",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_17],
			data:   make([]byte, 17),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_17 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_17],
			data:   make([]byte, 18),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_18 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_18],
			data:   make([]byte, 17),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_18",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_18],
			data:   make([]byte, 18),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_18 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_18],
			data:   make([]byte, 19),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_19 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_19],
			data:   make([]byte, 18),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_19",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_19],
			data:   make([]byte, 19),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_19 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_19],
			data:   make([]byte, 20),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_20 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_20],
			data:   make([]byte, 19),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_20",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_20],
			data:   make([]byte, 20),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_20 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_20],
			data:   make([]byte, 21),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_21 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_21],
			data:   make([]byte, 20),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_21",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_21],
			data:   make([]byte, 21),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_21 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_21],
			data:   make([]byte, 22),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_22 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_22],
			data:   make([]byte, 21),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_22",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_22],
			data:   make([]byte, 22),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_22 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_22],
			data:   make([]byte, 23),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_23 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_23],
			data:   make([]byte, 22),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_23",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_23],
			data:   make([]byte, 23),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_23 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_23],
			data:   make([]byte, 24),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_24 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_24],
			data:   make([]byte, 23),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_24",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_24],
			data:   make([]byte, 24),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_24 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_24],
			data:   make([]byte, 25),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_25 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_25],
			data:   make([]byte, 24),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_25",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_25],
			data:   make([]byte, 25),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_25 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_25],
			data:   make([]byte, 26),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_26 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_26],
			data:   make([]byte, 25),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_26",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_26],
			data:   make([]byte, 26),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_26 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_26],
			data:   make([]byte, 27),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_27 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_27],
			data:   make([]byte, 26),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_27",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_27],
			data:   make([]byte, 27),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_27 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_27],
			data:   make([]byte, 28),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_28 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_28],
			data:   make([]byte, 27),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_28",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_28],
			data:   make([]byte, 28),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_28 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_28],
			data:   make([]byte, 29),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_29 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_29],
			data:   make([]byte, 28),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_29",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_29],
			data:   make([]byte, 29),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_29 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_29],
			data:   make([]byte, 30),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_30 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_30],
			data:   make([]byte, 29),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_30",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_30],
			data:   make([]byte, 30),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_30 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_30],
			data:   make([]byte, 31),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_31 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_31],
			data:   make([]byte, 30),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_31",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_31],
			data:   make([]byte, 31),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_31 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_31],
			data:   make([]byte, 32),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_32 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_32],
			data:   make([]byte, 31),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_32",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_32],
			data:   make([]byte, 32),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_32 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_32],
			data:   make([]byte, 33),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_33 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_33],
			data:   make([]byte, 32),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_33",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_33],
			data:   make([]byte, 33),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_33 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_33],
			data:   make([]byte, 34),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_34 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_34],
			data:   make([]byte, 33),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_34",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_34],
			data:   make([]byte, 34),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_34 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_34],
			data:   make([]byte, 35),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_35 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_35],
			data:   make([]byte, 34),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_35",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_35],
			data:   make([]byte, 35),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_35 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_35],
			data:   make([]byte, 36),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_36 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_36],
			data:   make([]byte, 35),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_36",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_36],
			data:   make([]byte, 36),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_36 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_36],
			data:   make([]byte, 37),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_37 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_37],
			data:   make([]byte, 36),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_37",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_37],
			data:   make([]byte, 37),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_37 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_37],
			data:   make([]byte, 38),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_38 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_38],
			data:   make([]byte, 37),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_38",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_38],
			data:   make([]byte, 38),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_38 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_38],
			data:   make([]byte, 39),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_39 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_39],
			data:   make([]byte, 38),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_39",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_39],
			data:   make([]byte, 39),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_39 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_39],
			data:   make([]byte, 40),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_40 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_40],
			data:   make([]byte, 39),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_40",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_40],
			data:   make([]byte, 40),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_40 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_40],
			data:   make([]byte, 41),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_41 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_41],
			data:   make([]byte, 40),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_41",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_41],
			data:   make([]byte, 41),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_41 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_41],
			data:   make([]byte, 42),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_42 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_42],
			data:   make([]byte, 41),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_42",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_42],
			data:   make([]byte, 42),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_42 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_42],
			data:   make([]byte, 43),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_43 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_43],
			data:   make([]byte, 42),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_43",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_43],
			data:   make([]byte, 43),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_43 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_43],
			data:   make([]byte, 44),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_44 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_44],
			data:   make([]byte, 43),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_44",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_44],
			data:   make([]byte, 44),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_44 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_44],
			data:   make([]byte, 45),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_45 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_45],
			data:   make([]byte, 44),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_45",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_45],
			data:   make([]byte, 45),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_45 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_45],
			data:   make([]byte, 46),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_46 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_46],
			data:   make([]byte, 45),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_46",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_46],
			data:   make([]byte, 46),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_46 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_46],
			data:   make([]byte, 47),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_47 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_47],
			data:   make([]byte, 46),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_47",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_47],
			data:   make([]byte, 47),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_47 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_47],
			data:   make([]byte, 48),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_48 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_48],
			data:   make([]byte, 47),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_48",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_48],
			data:   make([]byte, 48),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_48 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_48],
			data:   make([]byte, 49),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_49 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_49],
			data:   make([]byte, 48),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_49",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_49],
			data:   make([]byte, 49),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_49 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_49],
			data:   make([]byte, 50),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_50 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_50],
			data:   make([]byte, 49),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_50",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_50],
			data:   make([]byte, 50),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_50 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_50],
			data:   make([]byte, 51),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_51 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_51],
			data:   make([]byte, 50),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_51",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_51],
			data:   make([]byte, 51),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_51 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_51],
			data:   make([]byte, 52),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_52 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_52],
			data:   make([]byte, 51),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_52",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_52],
			data:   make([]byte, 52),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_52 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_52],
			data:   make([]byte, 53),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_53 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_53],
			data:   make([]byte, 52),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_53",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_53],
			data:   make([]byte, 53),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_53 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_53],
			data:   make([]byte, 54),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_54 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_54],
			data:   make([]byte, 53),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_54",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_54],
			data:   make([]byte, 54),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_54 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_54],
			data:   make([]byte, 55),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_55 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_55],
			data:   make([]byte, 54),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_55",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_55],
			data:   make([]byte, 55),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_55 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_55],
			data:   make([]byte, 56),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_56 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_56],
			data:   make([]byte, 55),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_56",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_56],
			data:   make([]byte, 56),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_56 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_56],
			data:   make([]byte, 57),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_57 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_57],
			data:   make([]byte, 56),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_57",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_57],
			data:   make([]byte, 57),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_57 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_57],
			data:   make([]byte, 58),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_58 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_58],
			data:   make([]byte, 57),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_58",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_58],
			data:   make([]byte, 58),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_58 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_58],
			data:   make([]byte, 59),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_59 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_59],
			data:   make([]byte, 58),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_59",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_59],
			data:   make([]byte, 59),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_59 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_59],
			data:   make([]byte, 60),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_60 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_60],
			data:   make([]byte, 59),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_60",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_60],
			data:   make([]byte, 60),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_60 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_60],
			data:   make([]byte, 61),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_61 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_61],
			data:   make([]byte, 60),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_61",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_61],
			data:   make([]byte, 61),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_61 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_61],
			data:   make([]byte, 62),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_62 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_62],
			data:   make([]byte, 61),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_62",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_62],
			data:   make([]byte, 62),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_62 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_62],
			data:   make([]byte, 63),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_63 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_63],
			data:   make([]byte, 62),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_63",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_63],
			data:   make([]byte, 63),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_63 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_63],
			data:   make([]byte, 64),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_64 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_64],
			data:   make([]byte, 63),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_64",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_64],
			data:   make([]byte, 64),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_64 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_64],
			data:   make([]byte, 65),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_65 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_65],
			data:   make([]byte, 64),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_65",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_65],
			data:   make([]byte, 65),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_65 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_65],
			data:   make([]byte, 66),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_66 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_66],
			data:   make([]byte, 65),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_66",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_66],
			data:   make([]byte, 66),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_66 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_66],
			data:   make([]byte, 67),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_67 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_67],
			data:   make([]byte, 66),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_67",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_67],
			data:   make([]byte, 67),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_67 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_67],
			data:   make([]byte, 68),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_68 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_68],
			data:   make([]byte, 67),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_68",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_68],
			data:   make([]byte, 68),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_68 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_68],
			data:   make([]byte, 69),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_69 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_69],
			data:   make([]byte, 68),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_69",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_69],
			data:   make([]byte, 69),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_69 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_69],
			data:   make([]byte, 70),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_70 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_70],
			data:   make([]byte, 69),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_70",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_70],
			data:   make([]byte, 70),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_70 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_70],
			data:   make([]byte, 71),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_71 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_71],
			data:   make([]byte, 70),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_71",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_71],
			data:   make([]byte, 71),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_71 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_71],
			data:   make([]byte, 72),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_72 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_72],
			data:   make([]byte, 71),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_72",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_72],
			data:   make([]byte, 72),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_72 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_72],
			data:   make([]byte, 73),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_73 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_73],
			data:   make([]byte, 72),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_73",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_73],
			data:   make([]byte, 73),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_73 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_73],
			data:   make([]byte, 74),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_74 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_74],
			data:   make([]byte, 73),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_74",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_74],
			data:   make([]byte, 74),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_74 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_74],
			data:   make([]byte, 75),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_75 short",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_75],
			data:   make([]byte, 74),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DATA_75",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_75],
			data:   make([]byte, 75),
		},
		expectedErr: nil,
	},
	{
		name: "OP_DATA_75 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DATA_75],
			data:   make([]byte, 76),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_PUSHDATA1",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_PUSHDATA1],
			data:   []byte{0, 1, 2, 3, 4},
		},
		expectedErr: nil,
	},
	{
		name: "OP_PUSHDATA2",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_PUSHDATA2],
			data:   []byte{0, 1, 2, 3, 4},
		},
		expectedErr: nil,
	},
	{
		name: "OP_PUSHDATA4",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_PUSHDATA1],
			data:   []byte{0, 1, 2, 3, 4},
		},
		expectedErr: nil,
	},
	{
		name: "OP_1NEGATE",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_1NEGATE],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_1NEGATE long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_1NEGATE],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_RESERVED",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_RESERVED],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_RESERVED long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_RESERVED],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_TRUE",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_TRUE],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_TRUE long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_TRUE],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_2",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_2 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_2",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_2 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_3",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_3],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_3 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_3],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_4",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_4],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_4 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_4],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_5",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_5],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_5 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_5],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_6",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_6],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_6 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_6],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_7",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_7],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_7 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_7],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_8",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_8],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_8 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_8],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_9",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_9],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_9 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_9],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_10",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_10],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_10 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_10],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_11",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_11],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_11 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_11],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_12",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_12],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_12 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_12],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_13",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_13],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_13 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_13],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_14",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_14],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_14 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_14],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_15",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_15],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_15 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_15],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_16",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_16],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_16 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_16],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NOP",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NOP long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_VER",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_VER],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_VER long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_VER],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_IF",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_IF],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_IF long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_IF],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NOTIF",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOTIF],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NOTIF long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOTIF],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_VERIF",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_VERIF],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_VERIF long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_VERIF],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_VERNOTIF",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_VERNOTIF],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_VERNOTIF long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_VERNOTIF],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_ELSE",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_ELSE],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_ELSE long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_ELSE],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_ENDIF",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_ENDIF],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_ENDIF long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_ENDIF],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_VERIFY",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_VERIFY],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_VERIFY long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_VERIFY],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_RETURN",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_RETURN],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_RETURN long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_RETURN],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_TOALTSTACK",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_TOALTSTACK],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_TOALTSTACK long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_TOALTSTACK],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_FROMALTSTACK",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_FROMALTSTACK],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_FROMALTSTACK long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_FROMALTSTACK],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_2DROP",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2DROP],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_2DROP long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2DROP],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_2DUP",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2DUP],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_2DUP long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2DUP],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_3DUP",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_3DUP],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_3DUP long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_3DUP],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_2OVER",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2OVER],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_2OVER long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2OVER],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_2ROT",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2ROT],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_2ROT long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2ROT],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_2SWAP",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2SWAP],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_2SWAP long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2SWAP],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_IFDUP",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_IFDUP],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_IFDUP long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_IFDUP],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DEPTH",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DEPTH],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_DEPTH long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DEPTH],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DROP",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DROP],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_DROP long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DROP],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DUP",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DUP],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_DUP long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DUP],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NIP",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NIP],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NIP long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NIP],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_OVER",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_OVER],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_OVER long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_OVER],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_PICK",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_PICK],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_PICK long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_PICK],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_ROLL",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_ROLL],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_ROLL long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_ROLL],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_ROT",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_ROT],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_ROT long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_ROT],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_SWAP",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_SWAP],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_SWAP long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_SWAP],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_TUCK",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_TUCK],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_TUCK long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_TUCK],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_CAT",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_CAT],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_CAT long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_CAT],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_SUBSTR",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_SUBSTR],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_SUBSTR long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_SUBSTR],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_LEFT",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_LEFT],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_LEFT long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_LEFT],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_LEFT",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_LEFT],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_LEFT long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_LEFT],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_RIGHT",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_RIGHT],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_RIGHT long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_RIGHT],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_SIZE",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_SIZE],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_SIZE long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_SIZE],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_INVERT",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_INVERT],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_INVERT long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_INVERT],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_AND",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_AND],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_AND long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_AND],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_OR",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_OR],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_OR long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_OR],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_XOR",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_XOR],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_XOR long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_XOR],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_EQUAL",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_EQUAL],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_EQUAL long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_EQUAL],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_EQUALVERIFY",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_EQUALVERIFY],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_EQUALVERIFY long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_EQUALVERIFY],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_RESERVED1",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_RESERVED1],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_RESERVED1 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_RESERVED1],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_RESERVED2",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_RESERVED2],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_RESERVED2 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_RESERVED2],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_1ADD",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_1ADD],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_1ADD long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_1ADD],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_1SUB",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_1SUB],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_1SUB long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_1SUB],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_2MUL",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2MUL],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_2MUL long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2MUL],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_2DIV",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2DIV],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_2DIV long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_2DIV],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NEGATE",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NEGATE],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NEGATE long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NEGATE],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_ABS",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_ABS],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_ABS long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_ABS],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NOT",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOT],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NOT long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOT],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_0NOTEQUAL",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_0NOTEQUAL],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_0NOTEQUAL long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_0NOTEQUAL],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_ADD",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_ADD],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_ADD long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_ADD],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_SUB",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_SUB],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_SUB long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_SUB],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_MUL",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_MUL],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_MUL long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_MUL],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_DIV",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DIV],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_DIV long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_DIV],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_MOD",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_MOD],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_MOD long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_MOD],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_LSHIFT",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_LSHIFT],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_LSHIFT long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_LSHIFT],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_RSHIFT",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_RSHIFT],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_RSHIFT long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_RSHIFT],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_BOOLAND",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_BOOLAND],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_BOOLAND long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_BOOLAND],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_BOOLOR",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_BOOLOR],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_BOOLOR long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_BOOLOR],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NUMEQUAL",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NUMEQUAL],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NUMEQUAL long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NUMEQUAL],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NUMEQUALVERIFY",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NUMEQUALVERIFY],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NUMEQUALVERIFY long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NUMEQUALVERIFY],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NUMNOTEQUAL",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NUMNOTEQUAL],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NUMNOTEQUAL long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NUMNOTEQUAL],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_LESSTHAN",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_LESSTHAN],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_LESSTHAN long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_LESSTHAN],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_GREATERTHAN",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_GREATERTHAN],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_GREATERTHAN long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_GREATERTHAN],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_LESSTHANOREQUAL",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_LESSTHANOREQUAL],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_LESSTHANOREQUAL long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_LESSTHANOREQUAL],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_GREATERTHANOREQUAL",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_GREATERTHANOREQUAL],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_GREATERTHANOREQUAL long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_GREATERTHANOREQUAL],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_MIN",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_MIN],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_MIN long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_MIN],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_MAX",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_MAX],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_MAX long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_MAX],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_WITHIN",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_WITHIN],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_WITHIN long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_WITHIN],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_RIPEMD160",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_RIPEMD160],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_RIPEMD160 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_RIPEMD160],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_SHA1",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_SHA1],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_SHA1 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_SHA1],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_SHA256",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_SHA256],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_SHA256 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_SHA256],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_HASH160",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_HASH160],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_HASH160 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_HASH160],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_HASH256",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_HASH256],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_HASH256 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_HASH256],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_CODESAPERATOR",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_CODESEPARATOR],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_CODESEPARATOR long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_CODESEPARATOR],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_CHECKSIG",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_CHECKSIG],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_CHECKSIG long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_CHECKSIG],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_CHECKSIGVERIFY",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_CHECKSIGVERIFY],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_CHECKSIGVERIFY long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_CHECKSIGVERIFY],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_CHECKMULTISIG",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_CHECKMULTISIG],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_CHECKMULTISIG long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_CHECKMULTISIG],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_CHECKMULTISIGVERIFY",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_CHECKMULTISIGVERIFY],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_CHECKMULTISIGVERIFY long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_CHECKMULTISIGVERIFY],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NOP1",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP1],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NOP1 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP1],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NOP2",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP2],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NOP2 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP2],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NOP3",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP3],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NOP3 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP3],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NOP4",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP4],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NOP4 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP4],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NOP5",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP5],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NOP5 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP5],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NOP6",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP6],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NOP6 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP6],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NOP7",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP7],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NOP7 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP7],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NOP8",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP8],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NOP8 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP8],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NOP9",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP9],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NOP9 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP9],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_NOP10",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP10],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_NOP10 long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_NOP10],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_PUBKEYHASH",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_PUBKEYHASH],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_PUBKEYHASH long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_PUBKEYHASH],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_PUBKEY",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_PUBKEY],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_PUBKEY long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_PUBKEY],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
	{
		name: "OP_INVALIDOPCODE",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_INVALIDOPCODE],
			data:   nil,
		},
		expectedErr: nil,
	},
	{
		name: "OP_INVALIDOPCODE long",
		pop: &parsedOpcode{
			opcode: &opcodeArray[OP_INVALIDOPCODE],
			data:   make([]byte, 1),
		},
		expectedErr: ErrStackInvalidOpcode,
	},
}

func TestUnparsingInvalidOpcodes(t *testing.T) {
	for _, test := range popTests {
		_, err := test.pop.bytes()
		if err != test.expectedErr {
			t.Errorf("Parsed Opcode test '%s' failed", test.name)
			t.Error(err, test.expectedErr)
		}
	}
}
