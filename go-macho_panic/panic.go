package main

import (
	"encoding/binary"
	"fmt"

	"github.com/blacktop/go-macho/pkg/codesign"
)

func main() {
	b := make([]byte, 28)

	binary.BigEndian.PutUint32(b[0:], 0xfade0cc0) // MAGIC_EMBEDDED_SIGNATURE
	binary.BigEndian.PutUint32(b[4:], 28)         // total length
	binary.BigEndian.PutUint32(b[8:], 1)          // count

	binary.BigEndian.PutUint32(b[12:], 5)  // CSSLOT_ENTITLEMENTS
	binary.BigEndian.PutUint32(b[16:], 20) // offset to blob header

	binary.BigEndian.PutUint32(b[20:], 0xfade7171) // MAGIC_EMBEDDED_ENTITLEMENTS
	binary.BigEndian.PutUint32(b[24:], 0)          // invalid length

	cs, err := codesign.ParseCodeSignature(b)
	fmt.Printf("cs=%v err=%v\n", cs, err)
}