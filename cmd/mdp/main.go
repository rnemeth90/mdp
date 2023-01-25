package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"github.com/yosssi/gohtml"
)

const (
	header = `
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>mdp</title>
</head>

<body>
	`
	footer = `
</body>

</html>
	`
)

var (
	file    string
	preview bool
)

func init() {
	flag.StringVar(&file, "f", "", "the markdown file to preview")
	flag.BoolVar(&preview, "p", false, "preview the file")
}

func usage() {
	fmt.Println(os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if file == "" {
		usage()
		os.Exit(1)
	}

	if err := run(file, preview, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(fileName string, preview bool, w io.Writer) error {
	var outName string

	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	htmlData := parseContent(input)

	// this is ugly...
	if !preview {
		outName = fmt.Sprintf("%s.html", filepath.Base(fileName))
	} else {
		tempName, err := ioutil.TempFile(os.TempDir(), "mdp-*.html")
		if err != nil {
			return err
		}
		outName = tempName.Name()
	}

	fmt.Fprintln(w, outName)
	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}
	if preview {
		if err := saveHTML(outName, htmlData); err != nil {
			return err
		}
		if err != nil {
			return err
		}
		if err := previewFile(outName); err != nil {
			return err
		}
		fmt.Println("Press enter to exit.")
		fmt.Scanln()
		os.Remove(outName)
	}
	return nil
}

func previewFile(fileName string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", fileName).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", fileName).Start()
	case "darwin":
		return exec.Command("open", fileName).Start()
	default:
		return fmt.Errorf("unsupported platform")
	}
}

func parseContent(content []byte) []byte {
	html := blackfriday.Run(content)
	body := bluemonday.UGCPolicy().SanitizeBytes(html)

	var b bytes.Buffer

	b.WriteString(header)
	b.Write(body)
	b.WriteString(footer)

	formatted := gohtml.Format(b.String())

	return []byte(formatted)
}

func saveHTML(fileName string, htmlData []byte) error {
	return ioutil.WriteFile(fileName, htmlData, 0644)
}
