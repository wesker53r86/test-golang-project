package photo_handler

import (
	"fmt"
	"os"
	"test-golang-project/internal/models"
	photo_receiver "test-golang-project/internal/services/photo_handler/implementations"
	"testing"
) 

func TestSimple(t *testing.T){
	var vit models.PhotoObject = photo_receiver.ReceivePhoto("mignle",[]byte("hello world"))
	if vit.PictureName != "mignle"{
		t.Errorf("The name of the Photo Object is incorrect")
	}
}

func TestProcessImage(t *testing.T){
	file_path := "./sample_files/sample_picture.png"
	//load the picture
	content, err := os.ReadFile(file_path)
	if err != nil{
		fmt.Println("There was an issue reading the file:",err)
		return
	}
	//call the method
	comp := photo_receiver.ReceivePhotoFromBytes("sample_picture",content)

	//check the values
	if comp.PictureName != "sample_picture"{
		t.Errorf("The name of the Photo Object is incorrect")
	}

	if comp.MimeType != "image/png"{
		t.Errorf("the mimetype of the image is incorrect")
	}
	fmt.Println(len(comp.Bytes))
	if len(comp.Bytes) == 0{
		t.Errorf("There is no data in the photo object")
	}

}
