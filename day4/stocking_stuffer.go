package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

const PuzzleInput = "yzbqklnj"

func main() {
	found := false
	num := 1
	for !found {
		sum := md5HashCombo(PuzzleInput, strconv.Itoa(num))
		fmt.Printf("MD5 sum for %v and num %d is %s, first 5 bytes are %v\n", PuzzleInput, num, sum, sum[0:5])
		num += 1
		if sum[0:5] == "00000" || num == 1000000 {
			found = true
			fmt.Println("Found it")
		}
	}
	fmt.Printf("Output for pqrstuv1048970 is %s", md5HashCombo("pqrstuv", "1048970"))
	fmt.Printf("Otuput for abcdef609043 is %s", md5HashCombo("abcdef", "609043"))
}

func md5HashCombo(code, num string) string {
	hasher := md5.New()
	data := []byte(code + num)
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}
