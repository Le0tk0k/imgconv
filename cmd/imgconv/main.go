package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func convert(dst, src string) error {
	sf, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("ファイルを開けませんでした。 %s", src)
	}
	defer sf.Close()

	df, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("ファイルを書き出せませんでした %s", dst)
	}
	defer df.Close()

	img, _, err := image.Decode(sf)
	if err != nil {
		return err
	}

	switch strings.ToLower(filepath.Ext(dst)) {
	case ".png":
		err = png.Encode(df, img)
	case ".jpen", ".jpg":
		err = jpeg.Encode(df, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	}

	if err != nil {
		return fmt.Errorf("画像ファイルを書き出せませんでした。 %s", dst)
	}

	return nil
}

func run() error {
	if len(os.Args) < 3 {
		return fmt.Errorf("画像ファイルを指定してください。")
	}

	return convert(os.Args[2], os.Args[1])
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
