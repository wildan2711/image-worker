package main

import (
	"image"
	"log"
	"strings"

	"github.com/disintegration/imaging"
)

const PathSeperator = "/"

type ImageData struct {
	ImageObject    image.Image
	SourceFileName string
	SourcePath     string
	TargetFileName string
	TargetFormat   string
}

type ImageController interface {
	SetImage(srcPath string) (*ImageData, error)
	SaveImage(targetPath string) error
	Resize(width, height int)
}

// Resize resizes the image to the specified width and height
func (i *ImageData) Resize(width, height int) {
	src := imaging.Resize(i.ImageObject, width, height, imaging.Lanczos)
	i.ImageObject = src
}

// SetImage opens the image from a path
// then sets the Image object
func (i *ImageData) SetImage(srcPath string) (*ImageData, error) {
	src, err := imaging.Open(srcPath)
	if err != nil {
		log.Fatalf("Setting image failed from path %s: %v", srcPath, err)
		return nil, err
	}

	dir := strings.Split(srcPath, PathSeperator)

	i.ImageObject = src
	i.SourcePath = srcPath
	i.SourceFileName = dir[len(dir)-1]

	return i, nil
}

// SaveImage saves the image object to the specified
// destination path
func (i *ImageData) SaveImage(dstPath string) error {
	err := imaging.Save(i.ImageObject, dstPath)
	if err != nil {
		log.Fatalf("Saving image at path %s failed: %v", dstPath, err)
		return err
	}
	return nil
}

func New() ImageController {
	return new(ImageData)
}
