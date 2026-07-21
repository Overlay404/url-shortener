package pkg

import (
	"log/slog"
	"os"
	"strings"
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

func ReadFromFile(log *slog.Logger) ([]string, error) {
	data, err := os.ReadFile("short-urls.bin")
	if err != nil {
		log.Debug("cannot read from file", slog.String("err", err.Error()))
		return nil, err
	}

	log.Debug("Read file short-urls.bin", slog.String("value", string(data)))
	return strings.Fields(string(data)), nil
}
