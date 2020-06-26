package exiftool

import (
	"bufio"
	"bytes"
	"io"

	"github.com/evanoberholster/exiftool/exif"
	"github.com/evanoberholster/exiftool/ifds"
	"github.com/evanoberholster/exiftool/imagetype"
	"github.com/evanoberholster/exiftool/meta"
	"github.com/evanoberholster/exiftool/meta/tiffmeta"
)

// Errors
var (
	// Alias to tiffmeta ErrInvalidHeader
	ErrInvalidHeader = tiffmeta.ErrInvalidHeader
	ErrNoExif        = meta.ErrNoExif
)

// ScanExif identifies the imageType based on magic bytes and
// searches for exif headers, then it parses the io.ReaderAt for exif
// information and returns it.
// Sets exif imagetype from magicbytes, if not found sets imagetype
// Unknown.
//
// If no exif information is found ScanExif will return ErrNoExif.
func ScanExif(r io.ReaderAt) (e *exif.Exif, err error) {
	er := NewExifReader(r, nil, 0)
	br := bufio.NewReader(er)

	// Identify Image Type
	t, err := imagetype.ScanBuf(br)
	if err != nil {
		return
	}

	// Search Image for Metadata Header using
	// Imagetype information
	m, err := meta.ScanBuf(br, t)
	if err != nil {
		if err != ErrNoExif {
			return
		}
	}

	// NewExif with an ExifReader attached
	e = exif.NewExif(er, t)
	e.SetMetadata(m)

	if err == nil {
		header := m.TiffHeader()
		// Set TiffHeader sets the ExifReader and checks
		// the header validity.
		// Returns ErrInvalidHeader if header is not valid.
		if err = er.SetHeader(header); err != nil {
			return
		}

		// Scan the RootIFD with the FirstIfdOffset from the ExifReader
		err = scan(er, e, ifds.RootIFD, header.FirstIfdOffset)
	}
	return
}

// ParseExif parses a tiff header from the io.ReaderAt and
// returns exif and an error.
// Sets exif imagetype as imageTypeUnknown
//
// If the header is invalid ParseExif will return ErrInvalidHeader.
func ParseExif(r io.ReaderAt) (e *exif.Exif, err error) {
	er := NewExifReader(r, nil, 0)
	br := bufio.NewReader(er)

	// Search Image for Metadata Header using
	// Imagetype information
	m, err := meta.ScanBuf(br, imagetype.ImageUnknown)
	if err != nil {
		if err != ErrNoExif {
			return
		}
	}

	// NewExif with an ExifReader attached
	e = exif.NewExif(er, imagetype.ImageUnknown)
	e.SetMetadata(m)

	if err == nil {
		header := m.TiffHeader()
		// Set TiffHeader sets the ExifReader and checks
		// the header validity.
		// Returns ErrInvalidHeader if header is not valid.
		if err = er.SetHeader(header); err != nil {
			return
		}

		// Scan the RootIFD with the FirstIfdOffset from the ExifReader
		err = scan(er, e, ifds.RootIFD, header.FirstIfdOffset)
	}
	return
}

// ParseExifBytes parses exif information from the tiff header and a byte slice
// that contains raw exif data and returns exif and an error.
// Sets Exif imagetype to ImageTypeUnknown.
//
// If the header is invalid ParseExif will return ErrInvalidHeader.
func ParseExifBytes(rawexif []byte, header tiffmeta.Header) (e *exif.Exif, err error) {
	if !header.IsValid() {
		err = ErrInvalidHeader
		return
	}

	// Creates NewReader
	r := bytes.NewReader(rawexif)

	// NewExifReader from io.ReaderAt with the ByteOrder and TiffHeaderOffset
	er := NewExifReader(r, header.ByteOrder, header.TiffHeaderOffset)

	// NewExif with an ExifReader attached
	e = exif.NewExif(er, imagetype.ImageUnknown)

	// Scan the RootIFD with the FirstIfdOffset from the ExifReader
	err = scan(er, e, ifds.RootIFD, header.FirstIfdOffset)
	return
}

//
// OLD - Will be removed

// ParseExif parses an io.ReaderAt for exif informationan and returns it
func (eh ExifHeader) ParseExif(r io.ReaderAt) (e *exif.Exif, err error) {
	er := NewExifReader(r, eh.byteOrder, eh.tiffHeaderOffset)

	e = exif.NewExif(er, imagetype.ImageUnknown)
	if err = scan(er, e, ifds.RootIFD, eh.firstIfdOffset); err != nil {
		return
	}
	return
}
