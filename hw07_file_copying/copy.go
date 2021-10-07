package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrTheSameFiles          = errors.New("the same files")
	fromInput, toInput       string
	offsetInput, limitInput  int64
)

func init() {
	flag.StringVar(&fromInput, "from", "", "file to read from")
	flag.StringVar(&toInput, "to", "", "file to write to")
	flag.Int64Var(&limitInput, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offsetInput, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()
	err := Copy(fromInput, toInput, offsetInput, limitInput)
	if err != nil {
		fmt.Println("err = ", err)
	}
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	var toCopySize int64
	fullInputSize, err := checkInputParams(fromPath, toPath, offset)
	if err != nil {
		return err
	}
	soursFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer soursFile.Close()
	soursFile.Seek(offset, 0)

	dstFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if fullInputSize-offset >= limit {
		toCopySize = limit
	} else {
		toCopySize = fullInputSize - offset
	}
	fmt.Println(toCopySize)

	if _, err = io.CopyN(dstFile, soursFile, toCopySize); err != nil {
		return err
	}

	return nil
}

func checkInputParams(soursFile, outFile string, offset int64) (int64, error) {
	soursStat, err := os.Stat(soursFile)
	if err != nil {
		return 0, err
	}

	if soursStat.IsDir() || !soursStat.Mode().IsRegular() {
		return 0, ErrUnsupportedFile
	}

	if soursStat.Size() <= offset {
		return 0, ErrOffsetExceedsFileSize
	}

	outPutStat, err := os.Stat(outFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return 0, err
		}
	} else {
		if outPutStat.IsDir() || !outPutStat.Mode().IsRegular() {
			return 0, ErrUnsupportedFile
		}
		if os.SameFile(soursStat, outPutStat) {
			return 0, ErrTheSameFiles
		}
	}

	return soursStat.Size(), nil
}
