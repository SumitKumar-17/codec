# Codec from Scratch Golang Video Transcoder

This is a simple video transcoder written in Golang. It uses the `ffmpeg` command line tool to transcode videos.

Make sure to put the video file you want to transcode in the `input` folder.
and update the name of the file in the `cmd/main.go` file.


```bash
docker build -t video_transcoder .
docker run -it -v /home/sumitkumar/Desktop/codec-from-scratch:/home/app  video_transcoder

apt-get install golang

go run cmd/main.go 
```

- The index.m3u8 file will be generated in the output folder.
- Use the index.html file to paly the video 
