package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gocv.io/x/gocv"
)

var (
	enableOptical             = false
	initializationReservation = false
)

func api() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		if enableOptical {
			enableOptical = false
		} else {
			enableOptical = true
		}
	})

	r.GET("/init", func(ctx *gin.Context) {
		initializationReservation = true
	})

	r.Run()
}

func main() {
	go api()
	cam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatalln(err)
	}

	window := gocv.NewWindow("window")
	background := gocv.NewMat()
	cam.Read(&background)

	before := gocv.NewMat()
	after := gocv.NewMat()
	for {
		if initializationReservation {
			cam.Read(&background)
			initializationReservation = false
			continue
		}
		cam.Read(&after)
		if before.Empty() {
			after.CopyTo(&before)
			continue
		}

		diff := gocv.NewMat()
		after.CopyTo(&diff)

		if enableOptical {
			gocv.AbsDiff(after, before, &diff)
			gocv.AddWeighted(background, 1, diff, 0.5, 1, &diff)
		}

		window.IMShow(diff)
		window.WaitKey(1)

		bin, err := gocv.IMEncode(gocv.JPEGFileExt, diff)
		if err != nil {
			log.Fatal(err)
		}

		os.Stdout.Write(bin)
		os.Stdout.Sync()

		after.CopyTo(&before)
	}
}
