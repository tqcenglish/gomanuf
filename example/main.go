package main

import (
	manuf "github.com/tqcenglish/gomanuf"
)

func main() {
	print("mac: 80:e6:50:07:8a:f8")
	print(manuf.Search("80:e6:50:07:8a:f8"))
}
