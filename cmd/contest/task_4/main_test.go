package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"ozon/pkg/config"
	"reflect"
	"testing"
)

const (
	confPath   = "../../../config/main.json"
	testFolder = "compact-boxes"
)

func TestCompactBoxes(t *testing.T) {
	variants := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		14, 16, 17, 18, 19, 20,
		22, 24, 25, 29, 30,
		31, 32, 33, 34, 35, 36, 37, 41,
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
			compactBoxes(inFile, outFile)
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

	ans := make([]map[string]any, 0)
	err = json.Unmarshal(ansBytes, &ans)
	require.Nil(t, err)

	out := make([]map[string]any, 0)
	err = json.Unmarshal(resBytes, &out)
	require.Nil(t, err)

	require.Equal(t, true, reflect.DeepEqual(ans, out))
}
