package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Swop/advent-of-code/pkg/color"
	"github.com/charmbracelet/lipgloss"
)

var (
	failureStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ff0000"))
	defaultStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#9c9c9c"))

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color("#d0ff4f")).
			Padding(1, 2).
			Margin(1, 1)
	dayTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#4f92ff")).
			Padding(0, 2).
			MarginTop(1)
	partStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#d0ff4f")).MarginLeft(2)
	dotStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#333333"))
	resultStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#4f92ff")).MarginLeft(1)

	defaultTimingStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#9c9c9c")).MarginLeft(1)
	lowTimingStyle     = defaultTimingStyle.Copy().Foreground(lipgloss.Color("#00ff00"))
	medLowTimingStyle  = defaultTimingStyle.Copy().Foreground(lipgloss.Color("#d0ff4f"))
	medHighTimingStyle = defaultTimingStyle.Copy().Foreground(lipgloss.Color("#eb9534"))
	highTimingStyle    = defaultTimingStyle.Copy().Foreground(lipgloss.Color("#ff0000"))

	compilationFailureStyle = failureStyle.Copy().MarginLeft(1)
	runFailureStyle         = failureStyle.Copy().MarginLeft(1)

	width = 80
)

func main() {
	var year string
	var rootPath string
	flag.StringVar(&year, "year", "", "year")
	flag.StringVar(&rootPath, "path", "", "root path")
	flag.Parse()

	if year == "" || rootPath == "" {
		fatal(usage())
	}

	if !strings.HasPrefix(rootPath, "/") {
		cwd, err := os.Getwd()
		if err != nil {
			fatal(err.Error())
		}
		rootPath, err = filepath.Abs(path.Join(cwd, rootPath))
		if err != nil {
			fatal(err.Error())
		}
	}

	fmt.Print(titleStyle.Render(fmt.Sprintf("AoC %s", year)) + "\n\n")

	days, err := getDays(rootPath, year)
	if err != nil {
		fatal(err.Error())
	}
	if err := compileAll(rootPath, year, days); err != nil {
		fatal(err.Error())
	}
	totalDuration, durations := runAll(rootPath, year, days)
	sort.Slice(durations, func(i, j int) bool { return durations[i] < durations[j] })

	fmt.Print("\n\n")

	totalPartsCount := len(days) * 2
	var medianDuration time.Duration
	if totalPartsCount%2 == 0 {
		medianDuration = (durations[totalPartsCount/2-1] + durations[totalPartsCount/2]) / 2
	} else {
		medianDuration = durations[(totalPartsCount+1)/2-1]
	}
	durStyle := durationStyle(medianDuration)
	fmt.Println(createGroup(
		defaultStyle.Render("Median duration: "),
		durStyle.Render(medianDuration.String()),
	))
	fmt.Println(createGroup(
		defaultStyle.Render("Total cumulated duration: "),
		durStyle.Render(totalDuration.String()),
	))
}

func usage() string {
	return "Usage: go run bench.go -path=<path> -year=<year>"
}

func getDays(rootPath string, year string) ([]string, error) {
	var days []string
	dayFolders, err := os.ReadDir(path.Join(rootPath, year))
	if err != nil {
		if os.IsNotExist(err) {
			return days, errors.New("year not found")
		}
		return days, err
	}

	for _, e := range dayFolders {
		_, err := strconv.Atoi(e.Name())
		if err != nil {
			continue
		}
		days = append(days, e.Name())
	}
	slices.Sort(days)
	return days, nil
}

func createGroup(strs ...string) string {
	left := strs[:len(strs)-1]
	right := strs[len(strs)-1]
	leftGroup := lipgloss.JoinHorizontal(lipgloss.Center, left...)
	dotSize := width - lipgloss.Width(leftGroup) - lipgloss.Width(right)
	dots := dotStyle.Render(strings.Repeat(".", dotSize))
	return lipgloss.JoinHorizontal(lipgloss.Center, leftGroup, dots, right)
}

func durationStyle(d time.Duration) lipgloss.Style {
	switch {
	case d > 500*time.Millisecond:
		return highTimingStyle
	case d > 50*time.Millisecond:
		return medHighTimingStyle
	case d > 1*time.Millisecond:
		return medLowTimingStyle
	default:
		return lowTimingStyle
	}
}

func createBuildDir(rootPath string) error {
	dir := path.Join(rootPath, "build")
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getBuildTarget(rootPath string, year string, day string) string {
	return path.Join(rootPath, "build", fmt.Sprintf("%s-%s", year, day))
}

func compileAll(rootPath string, year string, days []string) error {
	fmt.Println(defaultStyle.Render("Compiling solvers..."))

	if err := createBuildDir(rootPath); err != nil {
		return fmt.Errorf("error when creating build dir: %w", err)
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(days))
	for _, day := range days {
		wg.Add(1)
		go func(day string) {
			defer wg.Done()
			err := compile(rootPath, year, day)
			if err != nil {
				errCh <- err
			}
		}(day)
	}
	go func() {
		wg.Wait()
		close(errCh)
	}()

	shouldExit := false
	for err := range errCh {
		shouldExit = true
		_, _ = fmt.Fprintln(os.Stderr, compilationFailureStyle.Render("❌ ", err.Error()))
	}
	if shouldExit {
		return errors.New("compilation failed")
	}
	return nil
}

func compile(rootPath string, year string, day string) error {
	cmd := exec.Command("/opt/homebrew/bin/go", "build", "-o", getBuildTarget(rootPath, year, day), "main.go") //nolint:gosec
	cmd.Dir = path.Join(rootPath, year, day)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error when compiling puzzle: %w", err)
	}
	return nil
}

func runAll(rootPath string, year string, days []string) (time.Duration, []time.Duration) {
	fmt.Println(defaultStyle.Render("Running solvers..."))

	totalDuration := time.Duration(0)
	var durations []time.Duration
	for _, day := range days {
		fmt.Println(dayTitleStyle.Render(fmt.Sprintf("Day %s", day)))
		for _, part := range []int{1, 2} {
			partStr := partStyle.Render(fmt.Sprintf("Part %d:", part))
			result, duration, err := run(rootPath, year, day, part)
			if err != nil {
				fmt.Println(lipgloss.JoinHorizontal(lipgloss.Center, partStr, runFailureStyle.Render("❌ ", err.Error())))
				continue
			}
			timingStyle := defaultTimingStyle
			d, err := time.ParseDuration(duration)
			if err == nil {
				totalDuration += d
				durations = append(durations, d)
				timingStyle = durationStyle(d)
			}

			fmt.Println(createGroup(partStr, resultStyle.Render(result), timingStyle.Render(duration)))
		}
	}
	return totalDuration, durations
}

func run(rootPath string, year string, day string, part int) (string, string, error) {
	var result, duration string
	cmd := exec.Command(getBuildTarget(rootPath, year, day), "-part="+strconv.Itoa(part)) //nolint:gosec
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	file, err := os.Open(path.Join(rootPath, year, day, "input.txt"))
	if err != nil {
		return result, duration, fmt.Errorf("error when opening input file: %w", err)
	}
	cmd.Stdin = file
	err = cmd.Run()
	out := b.String()
	outLines := strings.Split(out, "\n")
	if err != nil {
		return result, duration, fmt.Errorf("error when running part: %w", err)
	}
	var lastTwo [2]string
	lastTwoCurs := 1
	for i := len(outLines) - 1; i >= 0 && lastTwoCurs >= 0; i-- {
		line := outLines[i]
		if line == "" {
			continue
		}
		lastTwo[lastTwoCurs] = line
		lastTwoCurs--
	}
	result = strings.Fields(lastTwo[0])[1]
	duration = strings.Fields(lastTwo[1])[1]
	return result, duration, nil
}

func fatal(err string) {
	_, _ = fmt.Fprintln(os.Stderr, color.Format(color.Red, err))
	os.Exit(1)
}
