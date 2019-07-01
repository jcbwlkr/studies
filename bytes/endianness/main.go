package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	num := uint8(7)
	fmt.Printf("%x\n", num)
}

func main2() {

	rand.Seed(time.Now().Unix())

	buf := make([]byte, 8)

	// Store the 8 bytes of a uint64 into buf Big End first
	n := rand.Uint64()
	binary.BigEndian.PutUint64(buf, n)

	fmt.Println("Original", n, "==", buf)

	// Make a uint32 and stick its 4 bytes into the first part of buf, little end first
	n32 := rand.Uint32()
	binary.LittleEndian.PutUint32(buf[0:4], n32)

	fmt.Println("uint32", n32, "==", buf)

	// Read the 8 bytes of buf back into an int64
	x := binary.BigEndian.Uint64(buf)

	fmt.Println("x", x, "==", buf)

	m := uint64(math.MaxUint64)
	binary.BigEndian.PutUint64(buf, m)
	fmt.Println("m", m, "==", buf)

	buf2 := make([]byte, 4)
	m2 := uint32(math.MaxUint32)
	binary.BigEndian.PutUint32(buf2, m2)
	fmt.Println("m", m2, "==", buf2)
}
