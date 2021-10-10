package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrTheSameFiles          = errors.New("the same files")
	blockToCopy              = 1024 * 1024 * 4 // MB
)

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
	if limit == 0 {
		toCopySize = fullInputSize - offset
	}

	if err = processCopy(dstFile, soursFile, toCopySize); err != nil {
		return err
	}

	return nil
}

func processCopy(dstFile, soursFile *os.File, toCopySize int64) error {
	var (
		hasCopped int64
		tmpCopy   int64
		err       error
	)

	bar := pb.Start64(toCopySize)
	bar.Set(pb.Bytes, true)
	defer bar.Finish()

	for ; toCopySize > 0; toCopySize -= hasCopped {
		if toCopySize < int64(blockToCopy) {
			tmpCopy = toCopySize
		} else {
			tmpCopy = int64(blockToCopy)
		}
		if hasCopped, err = io.CopyN(dstFile, soursFile, tmpCopy); err != nil {
			return err
		}
		bar.Add64(hasCopped)
	}
	if err = dstFile.Sync(); err != nil {
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
