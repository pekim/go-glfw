package glfw

import (
	"image"
	"image/draw"
	"runtime"
	"unsafe"
)

/*
Get returns the red, green, and blue values for the Gammaramp.
*/
func (gr *Gammaramp) Get() ([]UShort, []UShort, []UShort) {
	red := make([]UShort, gr.size)
	copy(red, unsafe.Slice(gr.red, gr.size))
	green := make([]UShort, gr.size)
	copy(green, unsafe.Slice(gr.green, gr.size))
	blue := make([]UShort, gr.size)
	copy(blue, unsafe.Slice(gr.blue, gr.size))

	return red, green, blue
}

/*
Set sets the red, green, and blue values for the ramp.

The slice arguments will be pinned to ensure that they are not moved or freed
while the struct may be referenced by glfw C code.
When the Gammaramp is no longer going to be passed to any glfw function,
the returned function can be called to unpin the slices.
*/
func (gr *Gammaramp) Set(red []UShort, green []UShort, blue []UShort) func() {
	gr.red = &red[0]
	gr.green = &green[0]
	gr.blue = &blue[0]
	gr.size = UInt(len(red))

	var pinner runtime.Pinner
	pinner.Pin(&red)
	pinner.Pin(&green)
	pinner.Pin(&blue)
	return func() {
		pinner.Unpin()
	}
}

/*
Set will set the Image's fields with the argument's size and pixels.

The image pixels will be pinned to ensure that they are not moved or freed
while the struct may be referenced by glfw C code.
When the Image is no longer going to be passed to any glfw function,
the returned function can be called to unpin the pixels.
*/
func (_image *Image) Set(img image.Image) func() {
	_image.width = Int(img.Bounds().Dx())
	_image.height = Int(img.Bounds().Dy())

	// Can only get pixels from an NRGBA image.
	// So if not an NRGBA image, create one and draw the image into it.
	var nrgba *image.NRGBA
	if i, isNRGBA := img.(*image.NRGBA); !isNRGBA {
		nrgba = i
	} else {
		nrgba = image.NewNRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
		draw.Draw(nrgba, nrgba.Bounds(), img, img.Bounds().Min, draw.Src)
	}

	pixels := nrgba.Pix
	_image.pixels = (*UChar)(unsafe.Pointer(&pixels[0]))

	var pinner runtime.Pinner
	pinner.Pin(&pixels)
	return func() {
		pinner.Unpin()
	}
}
