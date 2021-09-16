package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var reader *bufio.Reader = bufio.NewReader(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprint(write, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main(){
	defer writer.Flush
}