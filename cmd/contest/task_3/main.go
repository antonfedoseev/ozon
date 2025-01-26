package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
		err := parseSet(in, &lines)
		if err != nil {
			panic(err)
		}

		res := calcDuplicates(lines)
		printResult(out, res)
	}

}

func calcDuplicates(words []string) int {
	amount := 0

	for i := 0; i < len(words); i++ {
		evenChars := getChars(words[i], 0, 2)
		oddChars := getChars(words[i], 1, 2)

		for j := i + 1; j < len(words); j++ {
			evenCharsOther := getChars(words[j], 0, 2)
			if len(evenChars) == len(evenCharsOther) && evenChars == evenCharsOther {
				amount++
				continue
			}

			oddCharsOther := getChars(words[j], 1, 2)
			if len(oddChars) > 0 && len(oddChars) == len(oddCharsOther) && oddChars == oddCharsOther {
				amount++
			}
		}
	}

	return amount
}

func getChars(s string, offset int, delta int) string {
	b := strings.Builder{}
	b.Grow(len(s) / 2)

	for i := offset; i < len(s); i += delta {
		b.WriteByte(s[i])
	}

	return b.String()
}

func printResult(out *bufio.Writer, res int) {
	fmt.Fprintln(out, fmt.Sprintf("%d", res))
}

func parseSet(in *bufio.Reader, lines *[]string) (err error) {
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
