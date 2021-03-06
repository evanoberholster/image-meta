package exif

import (
	"bytes"
	"encoding/binary"
	"testing"
)

// TODO: Write tests for exifReader
func TestExifReader(t *testing.T) {
	exifOffset := uint32(0)
	byteOrder := binary.BigEndian
	reader := bytes.NewReader([]byte{0, 0, 0, 0})

	er := newExifReader(reader, byteOrder, exifOffset, 0)

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

	// TODO: test Reader
}
