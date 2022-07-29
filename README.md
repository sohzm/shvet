![name](assets/name.png)

An application for changing HSV values of images.

# Available Themes

+ dracula
+ gruvbox
+ nord
+ solarized
+ tokyonight

NOTE: if the theme you like isn't shown here. Adding a theme to Shvet is really eazy, for more info [see here](https://example.com).

# Installation

There are two ways to install this application on your device:
+ Download precompiled binary
+ Build from source on your local PC

## Download

Click here to [download](https://example.com) the application

## Build from source

```
GOOS=linux CGO_ENABLED=0 go build -o shvet main.go     #replace linux with your target OS
```

# Getting started

```
go run main.go lighten <path to jpg or png>  # will increase the overall brightness of the image
go run main.go darken <path to jpg or png>   # will decrease the overall brightness of the image
```
