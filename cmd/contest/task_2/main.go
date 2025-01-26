package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	doTask(os.Stdin, os.Stdout)
}

func doTask(reader io.Reader, writer io.Writer) {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(reader)
	out = bufio.NewWriter(writer)
	defer out.Flush()

	var setsAmount int
	fmt.Fscanln(in, &setsAmount)

	for i := 0; i < setsAmount; i++ {
		var lines []string
		var output string
		err := parseSet(in, &lines, &output)
		if err != nil {
			printResult(out, false)
		}

		res := validate(lines, &output)
		printResult(out, res)
	}
}

func validate(lines []string, output *string) bool {
	expectedPrices := make(map[int][]string, len(lines))

	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")

		if len(parts) < 2 {
			return false
		}

		item := parts[0]
		price, err := strconv.Atoi(parts[1])
		if err != nil {
			return false
		}

		expectedPrices[price] = append(expectedPrices[price], item)
	}

	pairs := strings.Split(*output, ",")
	if len(expectedPrices) != len(pairs) {
		return false
	}

	gotItems := make(map[string]int, len(pairs))

	for i := 0; i < len(pairs); i++ {
		if !validPairExp.MatchString(pairs[i]) {
			return false
		}

		parts := strings.Split(pairs[i], ":")
		if len(parts) < 2 {
			return false
		}
		title := parts[0]
		priceStr := parts[1]
		if len(priceStr) > 1 && priceStr[0] == '0' {
			return false
		}

		price, err := strconv.Atoi(priceStr)
		if err != nil {
			return false
		}
		if price < 1 || price > 1_000_000_000 {
			return false
		}
		gotItems[title] = price
	}

	if len(expectedPrices) != len(gotItems) {
		return false
	}

	for expectedPrice, expectedNames := range expectedPrices {
		var contains bool
		var gotPrice int

		for i := 0; i < len(expectedNames); i++ {
			gotPrice, contains = gotItems[expectedNames[i]]
			if contains {
				break
			}
		}

		if !contains {
			return false
		}

		if expectedPrice != gotPrice {
			return false
		}
	}

	return true
}

var validPairExp = regexp.MustCompile("^[a-z]{1,10}:[0-9]{1,10}$")

const (
	yes = "YES"
	no  = "NO"
)

func printResult(out *bufio.Writer, res bool) {
	if res {
		fmt.Fprintln(out, yes)
	} else {
		fmt.Fprintln(out, no)
	}
}

func parseSet(in *bufio.Reader, lines *[]string, output *string) (err error) {
	var linesAmount int
	fmt.Fscanln(in, &linesAmount)

	*lines = make([]string, linesAmount)

	for i := 0; i < linesAmount; i++ {
		line, err := readFullLine(in)
		if err != nil {
			return err
		}

		(*lines)[i] = line
	}

	fmt.Fscanln(in, output)

	return nil
}

func readFullLine(in *bufio.Reader) (string, error) {
	builder := strings.Builder{}
	isPrefix := true
	var line []byte
	var err error

	for isPrefix {
		line, isPrefix, err = in.ReadLine()
		if err != nil {
			return "", err
		}

		builder.Write(line)
	}

	return builder.String(), nil
}
