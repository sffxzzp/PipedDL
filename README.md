# PipedDL

一个从 Piped 下载视频的工具。

支持任意 Piped 播放页地址，支持油管地址。

出于兼容性考虑，当前仅下载 mp4 以及 m4a 格式的视频和音频

* [x] 下载视频
* [x] 下载播放列表
* [x] 指定 pipedapi 地址
* [x] 支持高画质（实验性，需要 ffmpeg）

## 使用方法

不支持双击运行，在 cmd/shell 中使用。

`PipedDL -url https://piped.video/watch?v=dQw4w9WgXcQ -api https://pipedapi.kavin.rocks`

参数列表：

* -Q
  * 实验性高画质
* -url
  * 要下载的视频/播放列表地址
* -api
  * 指定的 pipedapi 地址
* -help
  * 显示帮助

## EN

A tool that downloads video from Piped Instances.

Support links from any piped frontend or youtube.

current it only downloads videos in mp4 codec, and audio in m4a codec

Features:

* [x] download video
* [x] download playlist
* [x] download from specified pipedapi address
* [x] high quality video (experimental, needs ffmpeg)

## Usage

Run in cmd/shell only.

`PipedDL -url https://piped.video/watch?v=dQw4w9WgXcQ -api https://pipedapi.kavin.rocks`

arguments:

* -Q
  * High quality
* -url
  * video/playlist link
* -api
  * specified pipedapi address
* -help
  * show help
