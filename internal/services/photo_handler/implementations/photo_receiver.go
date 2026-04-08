package photo_handler

import (
	"net/http"
	"test-golang-project/internal/models"
)

func ReceivePhoto(photo_name string, data []byte) models.PhotoObject{
	mime := http.DetectContentType(data)

	ret := models.PhotoObject{
		PictureName: photo_name,
		MimeType: mime,
		Bytes: data,
	}
	return ret
}

func ReceivePhotoFromBytes(name string, bytes []byte) models.PhotoObject{
	mime := http.DetectContentType(bytes)
	ret := models.PhotoObject{
		PictureName: name,
		MimeType: mime,
		Bytes: bytes,
	}
	return ret
}