// SHA1 hashes are frequently used to compute short identities for binary or text blobs.
// For example, the git revision control system uses SHA1s extensively to identify versioned
// files and directories.
// Here’s how to compute SHA1 hashes in Go.

package main

// Go implements several hash functions in various crypto/* packages.

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	// The pattern for generating a hash is sha1.New(), sha1.Write(bytes), then sha1.Sum([]byte{}).
	// Here we start with a new hash.

	fmt.Println("-------- SHA1 --------")
	h := sha1.New()

	// Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes.
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte slice.
	// The argument to Sum can be used to append to an existing byte slice: it usually isn’t needed.
	bs := h.Sum(nil)

	// SHA1 values are often printed in hex, for example in git commits.
	// Use the %x format verb to convert a hash results to a hex string.
	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	// using MD5 - to compute MD5 hashes import crypto/md5 and use md5.New()
	fmt.Println("-------- MD5 --------")
	md5 := md5.New()
	md5.Write([]byte(s))
	ms := md5.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", ms)

}
