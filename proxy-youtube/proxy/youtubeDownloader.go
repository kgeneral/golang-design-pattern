package proxy

import (
	"github.com/kkdai/youtube/v2"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type ThirdPartyYoutubeDownloader struct {
}

func (yd ThirdPartyYoutubeDownloader) ListVideos() []string {
	log.Println("this is a video list for u")
	return nil
}

func (yd ThirdPartyYoutubeDownloader) GetVideoInfo(id string) {
	log.Println("this is a video info for :", id)
}

func (yd ThirdPartyYoutubeDownloader) DownloadVideo(id string) string {
	client := youtube.Client{}

	video, err := client.GetVideo(id)
	if err != nil {
		panic(err)
	}

	log.Println("video downloaded. title :", video.Title)
	log.Println("video id :", id)

	formats := video.Formats.WithAudioChannels() // only get videos with audio

	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}

	filePath := "./" + id + "_" + strconv.FormatInt(time.Now().UnixMilli(), 10) + ".mp4"
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	//defer file.Close()

	fullPath, err := filepath.Abs(filePath)
	if err != nil {
		panic(err)
	}

	// stream read 중 hang이 걸림, EOF 가 오지 않는 듯
	// google content server 를 찾아서 임의 다운로드 받는 방식이라 정상적인 termination signal 을 기대할 수 없음

	// youtube downloader 의 stream writer 가 정상 동작하지 않아서 강제 timeout 적용
	downloadChannel := make(chan int64, 1)
	go func() {
		defer file.Close()

		written, err := io.Copy(file, stream)
		if err != nil {
			panic(err)
		}
		downloadChannel <- written
	}()

	log.Println("download started for :", fullPath)

	select {
	case written := <-downloadChannel:
		log.Printf("video file has been created. path : %s, size : %d", fullPath, written)
	case <-time.After(time.Second * 30):
		log.Printf("timeout! video file has been created. path : %s", fullPath)
	}

	return fullPath
}
