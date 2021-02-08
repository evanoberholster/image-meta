package exif

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/evanoberholster/imagemeta/metadata"
)

// TODO: Write tests for exifReader
func TestExifReader(t *testing.T) {
	exifOffset := uint32(0)
	byteOrder := binary.BigEndian
	reader := bytes.NewReader([]byte{0, 0, 0, 0})

	er := newExifReader(reader, byteOrder, exifOffset)

	// Error ExifReader
	tempbuf := make([]byte, 0)
	if n, err := er.Read(tempbuf); err != nil && n != 0 {
		t.Errorf("Wanted Exif Read Error %s", err)
	}
	if _, err := er.ReadAt(tempbuf, -1); err != ErrReadNegativeOffset {
		t.Errorf("Error reader.ReadAt negative offset %s", err)
	}

	// ByteOrder
	if er.ByteOrder() != binary.BigEndian {
		t.Errorf("Error with ByteOrder")
	}

	// SetHeader
	th := metadata.NewTiffHeader(byteOrder, exifOffset, exifOffset, 0)
	if err := er.SetHeader(th); err != nil {
		t.Errorf("Error with reader.SetHeader expected no error")
	}

	// SetHeader (Invalid Header)
	if err := er.SetHeader(metadata.TiffHeader{}); err != metadata.ErrInvalidHeader {
		t.Errorf("Error with reader.SetHeader %s, expected %s", err.Error(), metadata.ErrInvalidHeader.Error())
	}

	// TODO: test Reader
}