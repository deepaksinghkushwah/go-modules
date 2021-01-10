package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func main() {
	var mainFile, outputFile string
	flag.StringVar(&mainFile, "main", "movie.mp4", "The main movie file name which need to be converted")
	flag.StringVar(&outputFile, "output", "output.mp4", "Output video file name")
	flag.Parse()
	currentPath, _ := os.Getwd()
	ps := string(os.PathSeparator)
	cmd := newCmd(currentPath+ps+mainFile, currentPath+ps+outputFile)
	//go infmt.Println(cmd.String())
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

func newCmd(mainFile, outputFile string) *exec.Cmd {
	return exec.Command("ffmpeg",
		"-i", mainFile,
		"-vcodec", "h264",
		"-acodec", "aac",
		"-strict", "-2",
		"-vf", "scale=1280:720",
		"-crf", "24",
		outputFile,
	)
}
