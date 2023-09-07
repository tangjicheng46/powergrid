package yt

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/kkdai/youtube/v2"
	"io"
	"os"
)

type DownloadConfig struct {
	Dir string `json:"string"`
}

var config DownloadConfig

func InitConfig() error {
	configFilename := "youtube.json"
	configData, err := os.ReadFile(configFilename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		return err
	}
	if config.Dir == "" || !pathExist(config.Dir) {
		return errors.New("youtube download config dir error")
	}
	return nil
}

func DownloadSingleVideo(url, filename string) (err error) {
	client := youtube.Client{}
	video, err := client.GetVideo(url)
	if err != nil {
		return err
	}

	format := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &format[0])
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	defer func(file *os.File) {
		errClose := file.Close()
		if errClose != nil {
			err = fmt.Errorf("origin error: %s, close error: %s", err, errClose)
		}
	}(file)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, stream)
	return
}

func GenerateFilename() string {
	u := uuid.New()
	return u.String() + ".mp4"
}

func RecordDownload(url string) (err error) {
	// TODO
	return nil
}

func pathExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
