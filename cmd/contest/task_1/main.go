package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	darkRoom(os.Stdin, os.Stdout)
}

func darkRoom(reader io.Reader, writer io.Writer) {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(reader)
	out = bufio.NewWriter(writer)
	defer out.Flush()

	var setsAmount int
	fmt.Fscanln(in, &setsAmount)

	var field [][]byte

	for i := 0; i < setsAmount; i++ {
		rows, columns, err := parseSet(in, &field)
		if err != nil {
			panic(err)
		}

		res := doTaskLogic(rows, columns)
		printResult(out, res)
	}
}

func printResult(out *bufio.Writer, info []flashlight) {
	fmt.Fprintln(out, fmt.Sprintf("%d", len(info)))

	for i := 0; i < len(info); i++ {
		fmt.Fprintln(out, fmt.Sprintf("%d %d %s", info[i].row, info[i].col, info[i].direction))
	}
}

func parseSet(in *bufio.Reader, field *[][]byte) (rows int, columns int, err error) {
	sizeBytes, _, err := in.ReadLine()
	if err != nil {
		return -1, -1, err
	}

	sizeArr := strings.Split(string(sizeBytes), " ")

	rows, err = strconv.Atoi(sizeArr[0])
	if err != nil {
		return -1, -1, err
	}

	columns, err = strconv.Atoi(sizeArr[1])
	if err != nil {
		return -1, -1, err
	}

	return rows, columns, nil
}

type flashlightDirection string

const (
	flashlightDirectionUp    flashlightDirection = "U"
	flashlightDirectionDown  flashlightDirection = "D"
	flashlightDirectionLeft  flashlightDirection = "L"
	flashlightDirectionRight flashlightDirection = "R"
)

type flashlight struct {
	row       int
	col       int
	direction flashlightDirection
}

func doTaskLogic(rows int, columns int) []flashlight {
	var res []flashlight
	amount := 1
	direction := flashlightDirectionDown

	if rows == 1 && columns > 1 {
		amount = 1
		direction = flashlightDirectionRight
	} else if rows > 1 && columns == 1 {
		amount = 1
		direction = flashlightDirectionDown
	} else if rows == 1 && columns == 1 {
		amount = 1
		direction = flashlightDirectionRight
	} else if columns > rows {
		amount = 2
		direction = flashlightDirectionRight
	} else {
		amount = 2
		direction = flashlightDirectionDown
	}

	res = make([]flashlight, amount)

	for i := 0; i < amount; i++ {
		if i == 0 {
			res[i] = flashlight{row: 1, col: 1, direction: direction}
		} else {
			res[i] = flashlight{row: rows, col: columns, direction: direction}
		}

		if direction == flashlightDirectionRight {
			direction = flashlightDirectionLeft
		} else if direction == flashlightDirectionDown {
			direction = flashlightDirectionUp
		}
	}

	return res
}
