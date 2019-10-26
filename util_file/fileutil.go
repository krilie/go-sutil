package util_file

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"strings"
)

func IsFileAImage(ctx context.Context, file io.ReadSeeker) (bool, error) {
	content, err := GetFileContent(ctx, file)
	if err != nil {
		return false, err
	}
	return strings.Contains(content, "image"), nil
}

func GetFileContent(ctx context.Context, file io.ReadSeeker) (content string, err error) {
	decByte := make([]byte, 512)
	if _, err := file.Read(decByte); err != nil {
		return "", err
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", err
	}
	contentType := http.DetectContentType(decByte)
	return contentType, nil
}

func GetFileSha256(ctx context.Context, file io.ReadSeeker) (sha string, err error) {
	defer func() {
		_, err = file.Seek(0, io.SeekStart)
	}()
	hash := sha256.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum), nil
}

// 后缀名，不带点 没有返回 ""
func GetFileSuffix(fileName string) string {
	if fileName == "" {
		return ""
	}
	indexByte := strings.LastIndexByte(fileName, '.')
	if indexByte == -1 {
		return ""
	}
	return fileName[indexByte+1:]
}
