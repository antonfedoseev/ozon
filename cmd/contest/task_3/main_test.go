package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"ozon/pkg/config"
	"strconv"
	"strings"
	"testing"
)

const (
	confPath   = "../../../config/main.json"
	testFolder = "even-strings"
)

func TestEvenStrings(t *testing.T) {
	variants := []int{
		1, 2, 3, 8, 9, 10, 21,
	}

	conf := config.New()
	err := conf.Load(confPath)
	require.Nil(t, err)

	for _, v := range variants {
		t.Run(fmt.Sprintf("%d", v), func(t *testing.T) {
			inPath := fmt.Sprintf("%s\\%s\\%d", conf.TestsPath, testFolder, v)
			ansPath := fmt.Sprintf("%s\\%s\\%d.a", conf.TestsPath, testFolder, v)
			outPath := fmt.Sprintf("%s\\%s\\%d.o", conf.TestsPath, testFolder, v)

			inFile, err := os.Open(inPath)
			if err != nil {
				panic(err)
			}
			outFile, err := os.Create(outPath)
			if err != nil {
				panic(err)
			}
			doTask(inFile, outFile)
			compareFiles(t, ansPath, outPath)
		})
	}
}

func BenchmarkEvenStrings(b *testing.B) {
	variants := []int{
		1, 2, 3, 8, 9, 10, 21,
	}

	conf := config.New()
	err := conf.Load(confPath)
	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		for _, v := range variants {
			b.Run(fmt.Sprintf("%d", v), func(b *testing.B) {
				inPath := fmt.Sprintf("%s\\%s\\%d", conf.TestsPath, testFolder, v)
				ansPath := fmt.Sprintf("%s\\%s\\%d.a", conf.TestsPath, testFolder, v)
				outPath := fmt.Sprintf("%s\\%s\\%d.o", conf.TestsPath, testFolder, v)

				inFile, err := os.Open(inPath)
				if err != nil {
					panic(err)
				}
				outFile, err := os.Create(outPath)
				if err != nil {
					panic(err)
				}
				doTask(inFile, outFile)
				compareFiles(b, ansPath, outPath)
			})
		}
	}
}

func compareFiles(t require.TestingT, ansPath string, outPath string) {
	ansBytes, err := os.ReadFile(ansPath)
	if err != nil {
		panic(err)
	}
	resBytes, err := os.ReadFile(outPath)
	if err != nil {
		panic(err)
	}

	ansLines := strings.Split(string(ansBytes), "\n")
	resLines := strings.Split(string(resBytes), "\n")

	require.Equal(t, len(ansLines), len(resLines))

	for i := 0; i < len(ansLines); i++ {
		if ansLines[i] == "" {
			require.Equal(t, i, len(ansLines)-1)
			require.Equal(t, resLines[i], "")
			return
		}

		ans, err := strconv.Atoi(ansLines[i])
		require.Nil(t, err)
		res, err := strconv.Atoi(resLines[i])
		require.Nil(t, err)

		assert.Equal(t, ans, res)
	}
}
