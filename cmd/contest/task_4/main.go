package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	compactBoxes(os.Stdin, os.Stdout)
}

func compactBoxes(reader io.Reader, writer io.Writer) {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(reader)
	out = bufio.NewWriter(writer)
	defer out.Flush()

	var setsAmount int
	fmt.Fscanln(in, &setsAmount)
	s := make([]map[string]any, setsAmount)

	for i := 0; i < setsAmount; i++ {
		var field [][]byte
		err := parseSet(in, &field)
		if err != nil {
			panic(err)
		}

		rows := len(field)
		cols := len(field[0])
		m := make(map[string]any)
		scanArea(0, 0, rows-1, cols-1, field, m)
		s[i] = m
	}

	printResult(out, s)
}

func scanArea(beginRow int, beginCol int, endRow, endCol int, field [][]byte, m map[string]any) {
	for i := beginRow; i <= endRow; i++ {
		for j := beginCol; j <= endCol; j++ {
			if field[i][j] == '+' {
				ok, val, _, boxEndCol := scanBox(i, j, field)
				if !ok {
					continue
				}

				name := readBoxName(i, j, boxEndCol, field)
				m[name] = val
				/*scanArea(i, boxEndCol+1, boxEndRow, endCol, field, m)
				scanArea(i+1, beginCol, boxEndRow, j-1, field, m)
				i = boxEndRow
				break*/
			}
		}
	}
}

func scanBox(beginBoxRow int, beginBoxCol int, field [][]byte) (bool, any, int, int) {
	found, endBoxRow, endBoxCol := defineBorders(beginBoxRow, beginBoxCol, field)
	if !found {
		return false, -1, -1, -1
	}

	field[beginBoxRow][beginBoxCol] = ' '
	field[endBoxRow][beginBoxCol] = ' '
	field[beginBoxRow][endBoxCol] = ' '
	field[endBoxRow][endBoxCol] = ' '

	rowsAmount := endBoxRow - beginBoxRow + 1
	colsAmount := endBoxCol - beginBoxCol + 1
	area := (rowsAmount - 2) * (colsAmount - 2)
	m := make(map[string]any)

	scanArea(beginBoxRow+1, beginBoxCol+1, endBoxRow-1, endBoxCol-1, field, m)

	if len(m) > 0 {
		return true, m, endBoxRow, endBoxCol
	}

	return true, area, endBoxRow, endBoxCol
}

func defineBorders(beginBoxRow int, beginBoxCol int, field [][]byte) (ok bool, endBoxRow int, endBoxCol int) {
	i := beginBoxRow + 1
	j := beginBoxCol + 1

	for ; ; i++ {
		if field[i][beginBoxCol] == '+' {
			break
		}
	}

	for ; ; j++ {
		if field[beginBoxRow][j] == '+' {
			break
		}
	}

	endBoxRow = i
	endBoxCol = j
	return true, endBoxRow, endBoxCol
}

func readBoxName(beginRow int, beginCol int, endCol int, field [][]byte) string {
	beginRow++
	beginCol++
	endCol--

	b := strings.Builder{}
	for j := beginCol; j <= endCol; j++ {
		if field[beginRow][j] == '.' || field[beginRow][j] == '|' || field[beginRow][j] == '+' || field[beginRow][j] == ' ' {
			break
		}

		b.WriteByte(field[beginRow][j])
	}
	return b.String()
}

func printResult(out *bufio.Writer, m []map[string]any) {
	str, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(out, string(str))
}

func parseSet(in *bufio.Reader, field *[][]byte) (err error) {
	sizeBytes, _, err := in.ReadLine()
	if err != nil {
		return err
	}

	sizeArr := strings.Split(string(sizeBytes), " ")

	rows, err := strconv.Atoi(sizeArr[0])
	if err != nil {
		return err
	}

	columns, err := strconv.Atoi(sizeArr[1])
	if err != nil {
		return err
	}

	*field = make([][]byte, rows)

	for i := 0; i < rows; i++ {
		(*field)[i] = make([]byte, columns)
	}

	err = fillField(in, field, rows, columns)
	if err != nil {
		return err
	}

	return nil
}

func fillField(in *bufio.Reader, field *[][]byte, rows, columns int) error {
	for i := 0; i < rows; i++ {
		rowStr, err := readFullLine(in)
		if err != nil {
			return err
		}

		if len(rowStr) != columns {
			return fmt.Errorf("wrong columns amount. expected: %d, but got: %d", columns, len(rowStr))
		}

		for j, cell := range rowStr {
			(*field)[i][j] = byte(cell)
		}
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
