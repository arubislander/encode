package main

import "flag"
import "os"
import "io"
import "encoding/base64"

var encoding = flag.String("encoding", "base64", "encoding to use")
var filepath = flag.String("file", "", "file to encode")

func main() {
	flag.Parse()
	if *filepath == "" {
		flag.Usage()
		return
	}

	f, err := os.Open(*filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	enc := base64.NewEncoder(base64.RawStdEncoding, os.Stdout)
	defer enc.Close()

	var buffer = make([]byte, 512)
	for true {
		count, err := f.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if _, err := enc.Write(buffer[:count]); err != nil {
			panic(err)
		}
	}
}
