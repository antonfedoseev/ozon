package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"ozon/pkg/config"
	"strings"
	"testing"
)

const (
	confPath   = "../../../config/main.json"
	testFolder = "validate-result"
)

func TestValidateResult(t *testing.T) {
	variants := []int{
		1, 2, 3, 4, 5, 7, 8, 9, 10,
		12, 13, 14, 15, 17, 18, 19, 20,
		21, 23, 24, 25, 26, 28, 29, 30,
		31, 33, 34, 35, 36, 38, 39, 40, 41,
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

func compareFiles(t *testing.T, ansPath string, outPath string) {
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
		require.Equal(t, ansLines[i], resLines[i])
	}
}
