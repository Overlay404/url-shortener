package pkg

import (
	"io"
	"log/slog"
	"os"
)

func WriteInFile(text string, log *slog.Logger) error {
	data := []byte(text + "\n")
	file, err := os.OpenFile("short-urls.bin", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Debug("cannot write in file", slog.String("err", err.Error()))
		return err
	}
	defer file.Close()
	file.Write(data)
	s, _ := file.Stat()
	log.Debug("Write in file", slog.String("text", text), slog.Any("file", s))
	return nil
}

func ReadFromFile(log *slog.Logger) (res string, err error) {
	var n int
	file, err := os.Open("short-urls.bin")
	if err != nil {
		log.Debug("cannot read from file", slog.String("err", err.Error()))
		return "", err
	}
	defer file.Close()

	data := make([]byte, 64)
	for {
		n, err = file.Read(data)
		if err == io.EOF {
			break
		}
	}
	log.Debug("Read file short-urls.bin")
	return string(data[:n]), nil
}
