package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	dir, err := os.MkdirTemp("", "example")
	require.Nil(t, err)

	defer func() {
		errRemove := os.RemoveAll(dir)
		require.Nil(t, errRemove)
	}()

	testParams := []struct {
		testName string
		dirName  string
		res      Environment
		err      bool
	}{
		{
			testName: "empty die",
			dirName:  "",
			res:      nil,
			err:      true,
		},
	}
	for _, params := range testParams {
		t.Run(params.testName, func(t *testing.T) {
			res, errRead := ReadDir(params.dirName)
			if params.err {
				require.NotNil(t, errRead)
			} else {
				require.Equal(t, params.res, res)
				require.Equal(t, errRead, nil)
			}
		})
	}
}

func TestCheckFileName(t *testing.T) {
	testParams := []struct {
		inputData string
		result    bool
	}{
		{
			inputData: "lolo",
			result:    false,
		},
		{
			inputData: "lo=lo",
			result:    true,
		},
		{
			inputData: "lo=;lo",
			result:    true,
		},
		{
			inputData: "lo;lo",
			result:    true,
		},
	}
	for _, params := range testParams {
		t.Run(params.inputData, func(t *testing.T) {
			require.Equal(t, params.result, checkFileName(params.inputData))
		})
	}
}

func TestCheckEnvValue(t *testing.T) {
	testParams := []struct {
		testName string
		fileName string
		res      string
		err      bool
	}{
		{
			testName: "BAR",
			fileName: "./testdata/env/BAR",
			res:      "bar",
			err:      false,
		},
		{
			testName: "FOO",
			fileName: "./testdata/env/FOO",
			res:      "   foo\nwith new line",
			err:      false,
		},
		{
			testName: "No file",
			fileName: "./testdata/env/FOO1234",
			res:      "",
			err:      true,
		},
	}
	for _, params := range testParams {
		t.Run(params.testName, func(t *testing.T) {
			res, err := checkEnvValue(params.fileName)
			if params.err {
				require.Equal(t, "", res)
				require.NotNil(t, err)
			} else {
				require.Equal(t, params.res, res)
				require.Equal(t, err, nil)
			}
		})
	}
}
