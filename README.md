# image converter written in Go
This CLI convert extension of images.
.png, .jpg(.jpeg), and .gif are only supported.

## Install

```
go get github.com/Le0tk0k/imgpro
```

## Option

|  Flag  |  Description  | Default |
| ---- | ---- | --- |
|  -f  |  Extension before conversion  | .jpeg |
|  -t  |  Extension after conversion  | .png |
|  -r  |  Remove file before conversion | false |

## Usage

```
$ # jpeg -> png（default, Do not delete the file before conversion）
$ ./main sample.jpeg

$ # png -> jpg（Specify extension before and after conversion）
$ ./main -f .png -t .jpg sample.png

$ # jpg -> png（Delete the file before conversion）
$ ./main -r sample.jpg
```
