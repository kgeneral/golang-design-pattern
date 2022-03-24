package proxy

import "log"

type CachedYoutubeDownloader struct {
	service    ThirdPartyYoutubeLib
	listCache  []string
	videoCache map[string]bool
	needReset  bool
}

func NewCachedYoutubeDownloader(service ThirdPartyYoutubeLib) CachedYoutubeDownloader {
	return CachedYoutubeDownloader{
		service:    service,
		videoCache: make(map[string]bool),
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
		yd.videoCache[id] = true
	} else {
		log.Println("cached video :", id)
	}

	return id
}

func (yd CachedYoutubeDownloader) downloadExists(id string) bool {
	return yd.videoCache[id]
}
