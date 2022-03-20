package proxy

type ThirdPartyYoutubeLib interface {
	ListVideos() []string
	GetVideoInfo(id string)
	DownloadVideo(id string) string
}
