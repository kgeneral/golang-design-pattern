package main

import (
	"golang-design-pattern/proxy-youtube/proxy"
)

func main() {
	//var downloader proxy.ThirdPartyYoutubeLib = proxy.ThirdPartyYoutubeDownloader{}
	var downloader proxy.ThirdPartyYoutubeLib = proxy.NewCachedYoutubeDownloader(proxy.ThirdPartyYoutubeDownloader{})

	videoId := "X-350W_gwjY"
	// naive download
	downloader.DownloadVideo(videoId)
	// cached download
	downloader.DownloadVideo(videoId)

}
