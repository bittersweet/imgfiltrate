package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func getPrediction(file string) (string, string) {
	fContent, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("ReadFile ", err)
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://localhost:5000/predict", bytes.NewBuffer(fContent))
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("error in client do ", err)
	}
	defer res.Body.Close()

	var p Prediction
	json.NewDecoder(res.Body).Decode(&p)

	return p.Result.Prediction, p.Result.Score
}

func processImage(file string) Result {
	pct, totalColors := color.ProcessImage(file)
	prediction, score := getPrediction(file)

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
	file := os.Args[1]
	r := processImage(file)
	r.output()
}
