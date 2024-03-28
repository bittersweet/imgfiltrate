package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bittersweet/imgfiltrate/color"
)

type Result struct {
	ColorPercentage float64 `json:"color_percentage"`
	TotalColors     int     `json:"total_colors"`
	Prediction      string  `json:"prediction"`
	Score           string  `json:"score"`
}

type Prediction struct {
	Result struct {
		Prediction string `json:"prediction"`
		Score      string `json:"score"`
	} `json:"result"`
}

func getLocalFile(location string) *bytes.Buffer {
	content, err := ioutil.ReadFile(location)
	if err != nil {
		panic("Failed on ReadFile")
	}

	return bytes.NewBuffer(content)
}

func getRemoteFile(location string) *bytes.Buffer {
	resp, err := http.Get(location)
	if err != nil {
		panic("failed on download")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("failed on ReadAll")
	}
	resp.Body.Close()
	return bytes.NewBuffer(body)
}

func getPrediction(contents io.Reader) (string, string) {
	hostname := os.Getenv("IMGFILTRATE_HOSTNAME")
	path := fmt.Sprintf("%s/predict", hostname)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", path, contents)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("error in client do ", err)
	}
	defer res.Body.Close()

	var p Prediction
	json.NewDecoder(res.Body).Decode(&p)

	return p.Result.Prediction, p.Result.Score
}

func processImage(file *bytes.Buffer) Result {
	prediction, score := getPrediction(bytes.NewReader(file.Bytes()))
	pct, totalColors := color.ProcessImage(bytes.NewReader(file.Bytes()))

	r := Result{
		ColorPercentage: pct,
		TotalColors:     totalColors,
		Prediction:      prediction,
		Score:           score,
	}

	return r
}

func (r *Result) output() {
	output, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		log.Fatal("MarshalIndent", err)
	}

	fmt.Println(string(output))
}

func main() {
	urlPtr := flag.String("url", "none", "string to remote file location")
	filePtr := flag.String("file", "none", "string to file location")
	flag.Parse()

	// Check against default value, if not none, we have input
	if *urlPtr != "none" {
		remoteFile := getRemoteFile(*urlPtr)
		r := processImage(remoteFile)
		r.output()
	} else {
		localFile := getLocalFile(*filePtr)
		r := processImage(localFile)
		r.output()
	}
}
