package proxy

import (
	"log"
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
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		log.Printf("downloading...%d\n", i+1)
	}

	return id
}
