package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"github.com/disintegration/imaging"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ex, err := os.Getwd()
	checkErr(err)
	//fmt.Println(ex)

	width := flag.Int("width", 800, "width for resize")
	height := flag.Int("height", 600, "height for resize")

	flag.Parse()

	files, err := ioutil.ReadDir(ex + `/images/`)
	checkErr(err)

	for _, f := range files {
		fmt.Println("Processing: " + ex + `/images/` + f.Name())
		src, err := imaging.Open(ex + `/images/` + f.Name())
		checkErr(err)
		//imaging.Fit(src, 800, 600, imaging.Lanczos)
		destImg := imaging.Fill(src, *width, *height, imaging.Center, imaging.Lanczos)
		err = imaging.Save(destImg, ex+`/images/resized-`+f.Name())
		checkErr(err)
	}
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
