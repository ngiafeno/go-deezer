package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-deezer"
)

func main() {
	client := deezer.NewClient()
	res, _ := client.Artist.GetAlbums("27")
	top, _ := json.Marshal(res)
	fmt.Println(string(top))
}
