package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gammazero/workerpool"
)

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".mp4" || filepath.Ext(path) == ".mpeg" || filepath.Ext(path) == ".mov" {
			*files = append(*files, path)
		}

		return nil
	}
}

func main() {

	root, _ := os.Getwd()
	var files []string
	err := filepath.Walk(root, visit(&files))
	if err != nil {
		panic(err)
	}

	wp := workerpool.New(len(files))
	for _, file := range files {
		source := file
		dest := strings.TrimSuffix(source, filepath.Ext(source)) + "-out.mp4"
		cmd := newCmd(source, dest)
		fmt.Println(cmd.String())
		wp.Submit(func() {
			if err := cmd.Run(); err != nil {
				log.Fatalln(err)
			}
		})

	}
	wp.StopWait()
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
