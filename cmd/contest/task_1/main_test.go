package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"ozon/pkg/config"
	"strconv"
	"strings"
	"testing"
)

const (
	confPath   = "../../../config/main.json"
	testFolder = "dark-room"
)

func TestDarkRoom(t *testing.T) {
	variants := []int{
		1,
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
			darkRoom(inFile, outFile)
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

		if ansLines[i] == "" {
			require.Equal(t, i, len(ansLines)-1)
			continue
		}

		amount, err := strconv.Atoi(ansLines[i])
		require.Nil(t, err)

		ansParts := strings.Split(ansLines[i+1], " ")
		resParts := strings.Split(resLines[i+1], " ")
		require.Equal(t, len(ansParts), 3)
		require.Equal(t, len(ansParts), len(resParts))

		require.Equal(t, ansParts[0], resParts[0])
		require.Equal(t, ansParts[1], resParts[1])

		if amount == 1 {
			require.Equal(t, ansParts[2], resParts[2])
		} else if amount == 2 {
			ansParts2 := strings.Split(ansLines[i+2], " ")
			resParts2 := strings.Split(resLines[i+2], " ")
			require.Equal(t, len(ansParts2), 3)
			require.Equal(t, len(ansParts2), len(resParts2))

			require.Equal(t, ansParts2[0], resParts2[0])
			require.Equal(t, ansParts2[1], resParts2[1])

			rows, err := strconv.Atoi(resParts2[0])
			require.Nil(t, err)
			cols, err := strconv.Atoi(resParts2[1])
			require.Nil(t, err)

			if rows < cols {
				require.Equal(t, resParts[2], string(flashlightDirectionRight))
				require.Equal(t, resParts2[2], string(flashlightDirectionLeft))
			} else {
				require.Equal(t, resParts[2], string(flashlightDirectionDown))
				require.Equal(t, resParts2[2], string(flashlightDirectionUp))
			}
		}

		i += amount
	}
}
