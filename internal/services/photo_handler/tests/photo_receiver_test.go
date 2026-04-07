package photo_receiver

import (
	"fmt"
	photo_receiver "test-golang-project/internal/services/photo_handler/implementations"
	"test-golang-project/internal/models"
	"testing"
) 

func TestSimple(t *testing.T){
	fmt.Println("hello there")
	var vit models.PhotoObject = photo_receiver.ReceivePhoto("mignle",[]byte("hello world"))
	if vit.PictureName != "mignle"{
		t.Errorf("man there's an error here")
	}
}

func TestProcessImage(t *testing.T){
	
}
