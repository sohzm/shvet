![Shvet](assets/name.png)
background image source [deviantart/leikoi](https://www.deviantart.com/leikoi/art/The-Neon-Shallows-823330548)

An application for changing HSV values of images.

## Available Themes

+ dracula
+ gruvbox
+ nord
+ solarized
+ tokyonight

NOTE: if the theme you like isn't shown here. Adding a theme to Shvet is really eazy, for more info [see here](https://example.com).

## Installation

There are two ways to install this application on your device:
+ Download precompiled binary
+ Build from source on your local PC

### Download

Click here to [download](https://example.com) the application

### Build from source

```
git clone https://github.com/sz47/shvet.git            # clone this repo
cd shvet                                               # goto the shvet directory
GOOS=linux CGO_ENABLED=0 go build -o shvet main.go     # replace linux with your target OS
```

## Usage

There are two different ways a photo can be themed. You should try both and decide the better one. In some cases, first method will be better, in others the latter one.
1. This will try to replace the color space of the image with the one from theme.

```
        $ shvet <theme> path/to/image.png
example $ shvet gruvbox ~/Wallpapers/linux.png
```
2. This will try to change the average colors of the image to the one matching with the theme. Also you can use an optional parameter intensity, if intensity is 1 the image will only contain the colors from theme itself. You should generally check by putting the intensities 0, 1, 2, 10, 25 and come to conclusion which one fits your need better.
```
        $ shvet <theme>2 <intensity number (optional)> path/to/image.png
example $ shvet gruvbox2 1 ~/Wallpapers/linux.png
```

For list of themes supported do `shvet list`
