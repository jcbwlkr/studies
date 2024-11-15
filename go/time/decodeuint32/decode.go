package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

func main0() {
	//a := []byte{10, 161, 171, 28}                  // This came from the first 4 bytes of an OC generated traceID
	a := []byte{61, 212, 250, 90}                    // This came from the first 4 bytes of an OC generated traceID
	fmt.Printf("%v\n", a)                            //
	epoch := int64(binary.LittleEndian.Uint32(a))    // This turns those 4 bytes into a uint32 which we convert to an int64
	fmt.Println("that time is", time.Unix(epoch, 0)) // This shows that time as a time.Time
}

func main() {
	ids := []string{
		"5af99bcf",
		"5af99bd4",
		"5af99bd9",
	}
	for _, id := range ids {
		src := []byte(id)

		b := make([]byte, 12)

		n, err := hex.Decode(b, src)
		if err != nil {
			log.Fatal(err)
		}

		i := int64(binary.BigEndian.Uint32(b[:n]))

		t := time.Unix(i, 0)

		fmt.Println("id  ", id)
		fmt.Println("unix", i)
		fmt.Println("time", t)
		fmt.Println("----------------------------------------")
	}
}
