package main

import "fmt"

type model struct {
	src     string
	version string
}

func main() {
	m := model{"http://path_to_model", "1.0"}
	fmt.Println(m)
}
