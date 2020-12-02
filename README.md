# optical-cam
webカメラを光学迷彩にする
OpenCV, GoCV, v4l2を利用する必要が有る

## api

### GET: /
光学迷彩モードのtoggle

### GET: /init
背景画像の初期化

## 起動

```
$ sudo modprobe v4l2loopback card_label="optical_cam" video_nr=20 exclusive_caps=1
$ go run main.go | ffmpeg -i - -vcodec rawvideo -pix_fmt yuv420p -f v4l2 /dev/video20
```