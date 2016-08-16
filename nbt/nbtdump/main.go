package main

import (
	"strconv"
	"fmt"
	"os"
	"github.com/silvasur/gonbt/nbt"
)

func main() {
	tag, name, err := nbt.ReadNamedTag(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read NBT data: %s", err)
		os.Exit(1)
	}
	
	fmt.Printf("Tag Name:\n%s\n\nData:\n%s\n", strconv.Quote(name), tag)
}
