package main

import (
	"golang-design-pattern/proxy-youtube/proxy"
	"log"
)

func main() {
	//var downloader proxy.ThirdPartyYoutubeLib = proxy.ThirdPartyYoutubeDownloader{}
	var downloader proxy.ThirdPartyYoutubeLib = proxy.NewCachedYoutubeDownloader(proxy.ThirdPartyYoutubeDownloader{})

	videoId := "X-350W_gwjY"
	// naive download
	var id = downloader.DownloadVideo(videoId)
	log.Println("video downloaded :", id)

	// cached download
	id = downloader.DownloadVideo(videoId)
	log.Println("video downloaded :", id)

}
