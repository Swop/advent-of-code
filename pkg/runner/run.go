package runner

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Swop/advent-of-code/pkg/hashmatrix"
)

const reportBaseURL = "http://localhost:1337"

var (
	part         int
	report       bool
	vis          bool
	reportHandle string
)

func init() {
	flag.IntVar(&part, "part", -1, "Puzzle part to execute")
	flag.BoolVar(&report, "report", false, "Report result to AoC report server")
	flag.BoolVar(&vis, "vis", false, "Send visualisations to AoC report server")
	flag.StringVar(&reportHandle, "report-handle", "", "AoC report server report handle")
}

func printUsage() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage: %s [-part=PART] [-report] [-vis]\n", os.Args[0])
	flag.PrintDefaults()
}

type PartFunc func([]string) any

func Run(part1 PartFunc, part2 PartFunc) {
	flag.Parse()
	if report && reportHandle == "" {
		_, _ = fmt.Fprint(os.Stderr, "Missing report handle\n")
		printUsage()
		os.Exit(2)
	}

	var cb PartFunc
	switch part {
	case 1:
		cb = part1
	case 2:
		cb = part2
	default:
		_, _ = fmt.Fprint(os.Stderr, "Invalid part. Should be 1 or 2\n")
		printUsage()
		os.Exit(2)
	}

	lines, err := ScanStdIn()
	if err != nil {
		panic(err)
	}

	t1 := time.Now()
	res := cb(lines)
	dur := time.Since(t1)

	var resStr string
	switch resTyped := res.(type) {
	case string:
		resStr = resTyped
	case uint, int, uint8, int8, uint16, int16, uint32, int32, uint64, int64:
		resStr = fmt.Sprintf("%d", resTyped)
	case float32, float64:
		resStr = fmt.Sprintf("%.0f", resTyped)
	}

	fmt.Println("RESULT\t", resStr)
	fmt.Println("TIME\t", dur)
	sendResult(resStr, dur)
}

func ScanStdIn() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error when scanning input: %w", err)
	}
	return lines, nil
}

func PushMatrixVisualisation[K hashmatrix.Pos3D | hashmatrix.Pos2D, V any](m *hashmatrix.Matrix[K, V], title string) {
	if !vis {
		return
	}
	err := send("/visualisations", map[string]any{
		"type":  "matrix",
		"title": title,
		"data":  m,
	})
	if err != nil {
		panic(err)
	}
}

func sendResult(res string, dur time.Duration) {
	if !report {
		return
	}
	err := send("/result", map[string]any{
		"result":   res,
		"duration": dur,
		"handle":   reportHandle,
	})
	if err != nil {
		panic(err)
	}
}

func send(url string, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error when marshalling data to send to AoC report server: %w", err)
	}
	req, err := http.NewRequest("POST", reportBaseURL+url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error when creating request to AoC report server: %w", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{Transport: &http.Transport{}}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error when sending data to AoC report server: %w", err)
	}
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error when reading response body received from AoC report server: %w", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("received invalid status code from AoC report server: [%d] %s", resp.StatusCode, respData)
	}
	return nil
}

type MatrixVisualisationElem struct {
	Color string `json:"color"`
	Value string `json:"value"`
}

func (e MatrixVisualisationElem) JSON() []byte {
	b, _ := json.Marshal(e)
	return b
}
