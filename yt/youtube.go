package yt

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/kkdai/youtube/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"os"
	"time"
)

type DownloadConfig struct {
	Dir    string `json:"string"`
	DbFile string `json:"db"`
}

type DownloadRecord struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	Url      string
	Filename string `gorm:"uniqueIndex"`
	CreateAt time.Time
}

var config DownloadConfig
var db *gorm.DB

func InitConfig() (err error) {
	configFilename := "youtube.json"
	configData, err := os.ReadFile(configFilename)
	if err != nil {
		return
	}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		return
	}
	if config.Dir == "" || !pathExist(config.Dir) {
		return errors.New("youtube config download dir is not exist")
	}
	if config.DbFile == "" {
		return errors.New("youtube config youtube db is empty")
	}

	db, err = gorm.Open(sqlite.Open(config.DbFile), &gorm.Config{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&DownloadRecord{})
	return
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

func GenerateFilenameUnique(db *gorm.DB) (string, error) {
	maxGen := 10
	for i := 0; i < maxGen; i++ {
		filename := GenerateFilename()
		record := DownloadRecord{}
		result := db.Where("filename = ?", filename).First(&record)
		if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
			return filename, nil
		}
	}
	return "", errors.New("cannot generate filename, time == 10")
}

func DownloadWithRecord(url string) (err error) {
	filename, err := GenerateFilenameUnique(db)
	if err != nil {
		return
	}
	err = DownloadSingleVideo(url, filename)
	return
}

func pathExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
