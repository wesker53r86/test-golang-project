package photo_receiver

import (
	"test-golang-project/internal/models"
)

func ReceivePhoto(photo_name string, data []byte) models.PhotoObject{
	ret := models.PhotoObject{
		PictureName: photo_name,
		MimeType: "image/png",
		Bytes: []byte("hi there"),
	}
	return ret
}