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

func calcDuplicates(words []string) uint {
	amount := uint(0)
	evenGroups := make(map[uint]uint32)
	oddGroups := make(map[uint]uint32)
	groups := make(map[uint]uint32)

	for i := 0; i < len(words); i++ {
		evenAmount := (len(words[i]) + 1) / 2
		oddAmount := len(words[i]) / 2
		evenKey := buildKey(i, words, evenAmount, 0, 2)
		oddKey := buildKey(i, words, oddAmount, 1, 2)
		key := buildKey(i, words, oddAmount, 0, 1)

		if evenKey > 0 {
			evenGroups[evenKey]++
		}

		if oddKey > 0 {
			oddGroups[oddKey]++
		}

		if key > 0 {
			groups[key]++
		}
	}

	for _, count := range evenGroups {
		amount += uint(calcCombinations(count, 2))
	}

	for _, count := range oddGroups {
		amount += uint(calcCombinations(count, 2))
	}

	for _, count := range groups {
		amount -= uint(calcCombinations(count, 2))
	}

	return amount
}

func calcCombinations(n uint32, k uint32) uint32 {
	if n < k {
		return 0
	}

	return n * (n - 1) / k
}

const (
	firstChar   = 'a'
	charsAmount = 26
)

func calcMaskAndSum(s *string, offset int, step int) (mask uint, sum uint) {
	for i := offset; i < len(*s); i += step {
		setBit(&mask, (*s)[i]-firstChar)
		val := uint((*s)[i]) + uint(i)
		sum += val * val
	}
	return
}

func setBit(x *uint, index uint8) {
	*x |= 1 << index
}

func buildKey(wordIndex int, words []string, amount int, offset int, step int) uint {
	if amount == 0 {
		return 0
	}

	mask, sum := calcMaskAndSum(&words[wordIndex], offset, step)
	key := sum << charsAmount
	key |= mask
	return key
}

func printResult(out *bufio.Writer, res uint) {
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
