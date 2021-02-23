package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

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
		var stdout, stderr []byte

		stdoutIn, _ := cmd.StdoutPipe()
		stderrIn, _ := cmd.StderrPipe()

		//fmt.Println(cmd.String())
		wp.Submit(func() {
			cmd.Start()
			var wg sync.WaitGroup
			wg.Add(1)
			var errStdout, errStderr error
			go func(errStdout error, errStderr error) {
				stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
				wg.Done()
			}(errStdout, errStderr)
			stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)

			wg.Wait()

			cmd.Wait()
			outStr, errStr := string(stdout), string(stderr)
			fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
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

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}
