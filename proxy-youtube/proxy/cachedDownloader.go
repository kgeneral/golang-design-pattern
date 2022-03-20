package proxy

import "log"

type CachedYoutubeDownloader struct {
	service    ThirdPartyYoutubeLib
	listCache  []string
	videoCache [][]byte
	needReset  bool
}

func NewCachedYoutubeDownloader(service ThirdPartyYoutubeLib) CachedYoutubeDownloader {
	return CachedYoutubeDownloader{
		service: service,
	}
}

func (yd CachedYoutubeDownloader) ListVideos() []string {
	if yd.listCache == nil || yd.needReset {
		yd.listCache = yd.ListVideos()
	}
	return yd.listCache
}

func (yd CachedYoutubeDownloader) GetVideoInfo(id string) {
	log.Println("this is a video info for :", id)
}

func (yd CachedYoutubeDownloader) DownloadVideo(id string) string {
	if !yd.downloadExists(id) || yd.needReset {
		yd.service.DownloadVideo(id)
	}

	return ""
}

func (yd CachedYoutubeDownloader) downloadExists(id string) bool {
	return false
}
