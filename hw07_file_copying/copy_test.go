package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Какой смысл писать тесты если уже все протестировано через bash скрипт.
func TestCopy(t *testing.T) {
	testInputString := "test date"
	tFile, errTmp := ioutil.TempFile("", "soursFile")
	require.NoError(t, errTmp)
	defer tFile.Close()

	outFile, errOut := ioutil.TempFile("", "outFile")
	require.NoError(t, errOut)
	defer outFile.Close()

	tFile.WriteString(testInputString)

	t.Run("sucsec test", func(t *testing.T) {
		err := Copy(tFile.Name(), outFile.Name(), 0, 0)
		require.NoError(t, err)
		tFile.Seek(0, 0)

		inputBuf, err := io.ReadAll(tFile)
		require.NoError(t, err)
		outBuf, err := io.ReadAll(outFile)
		require.NoError(t, err)
		require.Equal(t, inputBuf, outBuf)
	})
}

func TestCheckInputParams(t *testing.T) {
	testInputString := "test date"
	tFile, errTmp := ioutil.TempFile("", "soursFile")
	require.NoError(t, errTmp)
	defer tFile.Close()

	tFile.WriteString(testInputString)

	t.Run("no input file", func(t *testing.T) {
		inputSize, err := checkInputParams("LOLOLO", "out", 1)
		require.Equal(t, inputSize, int64(0))
		require.True(t, os.IsNotExist(err))
	})

	t.Run("there is input file", func(t *testing.T) {
		inputSize, err := checkInputParams(tFile.Name(), "out", 0)
		require.Equal(t, inputSize, int64(len(testInputString)))
		require.NoError(t, err)
	})

	t.Run("offset exceeds file size", func(t *testing.T) {
		inputSize, err := checkInputParams(tFile.Name(), "out", int64(len(testInputString))+10)
		require.Equal(t, inputSize, int64(0))
		require.Equal(t, err, ErrOffsetExceedsFileSize)
	})

	t.Run("the same files", func(t *testing.T) {
		inputSize, err := checkInputParams(tFile.Name(), tFile.Name(), 0)
		require.Equal(t, inputSize, int64(0))
		require.Equal(t, err, ErrTheSameFiles)
	})
}
