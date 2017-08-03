package main

import (
	"testing"
)

func TestSetImage(t *testing.T) {
	imageController := New()

	srcPath := "fa5.jpg"

	img, err := imageController.SetImage(srcPath)
	if err != nil {
		t.Error(err)
	}

	if img.SourceFileName != srcPath {
		t.Errorf("expected %s, got %v", srcPath, img.SourcePath)
	}

}

func TestSaveImage(t *testing.T) {
	imageController := New()

	srcPath := "fa5.jpg"

	img, err := imageController.SetImage(srcPath)
	if err != nil {
		t.Error(err)
	}

	if img.SourceFileName != srcPath {
		t.Errorf("expected %s, got %v", srcPath, img.SourcePath)
	}

	err = imageController.SaveImage("hombreng.jpg")
	if err != nil {
		t.Error(err)
	}
}
