package main

import (
	"os"
	"image"
	"strings"
	"path/filepath"
	"image/jpeg"
	"image/png"
	"github.com/nfnt/resize"
	"errors"
	"io/ioutil"
)

func Arq(path, out string, w int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	var (
		img, m image.Image
	)
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".jpg":
		img, err = jpeg.Decode(file)
		if err != nil {
			return err
		}
		m = resize.Resize(uint(w), 0, img, resize.Lanczos3)
		out, err := os.Create(filepath.Join(out, filepath.Base(path)))
		if err != nil {
			return err
		}
		defer out.Close()
		err = jpeg.Encode(out, m, nil)
		if err != nil {
			return err
		}
	case ".png":
		img, err = png.Decode(file)
		if err != nil {
			return err
		}
		m = resize.Resize(uint(500), 0, img, resize.Lanczos3)
		out, err := os.Create(filepath.Join(out, filepath.Base(path)))
		if err != nil {
			return err
		}
		defer out.Close()
		err = png.Encode(out, m)
		if err != nil {
			return err
		}
	default:
		return errors.New("resize: extensão inválida")
	}
	return nil
}

func Dir(in, out string, w int) error {
	fi, err := ioutil.ReadDir(in)
	if err != nil {
		return err
	}
	for _, f := range fi {
		err = Arq(filepath.Join(in, f.Name()), out, w)
		if err != nil {
			return err
		}
	}
	return nil
}
