package tools

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/labstack/gommon/bytes"
	"io"
	"os"
)

var (
	hash_block_size = int64(bytes.MB * 1)
)

func ReadFileInLine(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	lines := []string{}
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}

func WriteFileInLine(path string, data []string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	for idx, v := range data {
		if idx != len(data)-1 {
			if _, err := fmt.Fprintln(file, v); err != nil {
				return err
			}
		} else {
			if _, err := fmt.Fprint(file, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func SetHashBlock(size int64) {
	hash_block_size = size
}

func GetFileSHA1(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return ``, err
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		return ``, err
	}
	file_block_num := fi.Size()/hash_block_size + 1
	if fi.Size()%hash_block_size == 0 {
		file_block_num--
	}
	hSHA1 := sha1.New()
	for i := int64(0); i < file_block_num; i++ {
		buf := make([]byte, hash_block_size)
		if _, err := file.Seek(i*hash_block_size, os.SEEK_SET); err != nil {
			return ``, err
		}
		offset, err := file.Read(buf)
		if err != nil {
			return ``, err
		}
		if _, err := hSHA1.Write(buf[:offset]); err != nil {
			return ``, err
		}
	}
	return hex.EncodeToString(hSHA1.Sum(nil)), nil
}

func GetFileMD5(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return ``, err
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		return ``, err
	}
	file_block_num := fi.Size()/hash_block_size + 1
	if fi.Size()%hash_block_size == 0 {
		file_block_num--
	}
	hMD5 := md5.New()
	for i := int64(0); i < file_block_num; i++ {
		buf := make([]byte, hash_block_size)
		if _, err := file.Seek(i*hash_block_size, os.SEEK_SET); err != nil {
			return ``, err
		}
		offset, err := file.Read(buf)
		if err != nil {
			return ``, err
		}
		if _, err := hMD5.Write(buf[:offset]); err != nil {
			return ``, err
		}
	}
	return hex.EncodeToString(hMD5.Sum(nil)), nil
}
