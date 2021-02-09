package bmff

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
)

// bufReader adds some HEIF/BMFF-specific methods around a *bufio.Reader.
type bufReader struct {
	*bufio.Reader
	err    error
	remain int64
	// sticky error
}

func (b box) newReader(size int64) Reader {
	return Reader{br: bufReader{Reader: b.r.Reader, remain: size}}
}

func (br *bufReader) discard(n int) error {
	m, err := br.Discard(n)
	br.remain -= int64(m)
	if m != n || err != nil {
		br.err = errors.New("bufReader discard error")
		return br.err
	}
	return err
}

// ok reports whether all previous reads have been error-free.
func (br *bufReader) ok() bool { return br.err == nil }

func (br *bufReader) anyRemain() bool {
	if br.err != nil {
		return false
	}
	_, err := br.Peek(1)
	return err == nil && br.remain > 0
}

func (br *bufReader) readString() (string, error) {
	if br.err != nil {
		return "", br.err
	}
	s0, err := br.ReadString(0)
	if err != nil {
		br.err = err
		return "", err
	}
	br.remain -= int64(len(s0))
	if s0[len(s0)-1] == '\x00' {
		s0 = s0[:len(s0)-1]
		return string(s0), nil
	}
	s := strings.TrimSuffix(s0, "\x00")
	if len(s) == len(s0) {
		err = fmt.Errorf("unexpected non-null terminated string")
		br.err = err
		return "", err
	}
	return s, nil
}

func (br *bufReader) readUint16() (uint16, error) {
	if br.err != nil {
		return 0, br.err
	}
	buf, err := br.Peek(2)
	if err != nil {
		br.err = err
		return 0, err
	}
	v := binary.BigEndian.Uint16(buf[:2])
	return v, br.discard(2)
}

func (br *bufReader) readBrand() (b Brand, err error) {
	if br.err != nil {
		err = br.err
		return
	}

	if br.remain < 4 {
		err = errors.New("bufReader error: brand insufficient length")
		br.err = err
		return
	}
	var buf []byte
	if buf, err = br.Peek(4); err != nil {
		return
	}
	return brand(buf[:4]), br.discard(4)
}

func (br *bufReader) readItemType() (it ItemType, err error) {
	if br.remain < 4 {
		err = errors.New("bufReader error: itemType infufficient length")
		br.err = err
	}
	buf, err := br.Peek(5)
	if err != nil {
		return ItemTypeUnknown, err
	}

	it = itemType(buf[:4])
	if buf[4] != '\x00' {
		err = errors.New("bufReader error: itemType doesn't end on whitespace")
	}

	return it, br.discard(5)
}

// Errors
var (
	ErrBufReaderLength = errors.New("bufReader error: infufficient length")
)

func (br *bufReader) readFlags() (f Flags, err error) {
	if br.remain < 4 {
		err = ErrBufReaderLength
		br.err = err
	}

	// Parse FullBox header.
	buf, err := br.Peek(4)
	if err != nil {
		return f, fmt.Errorf("failed to read 4 bytes of Flags: %v", err)
	}
	f.Read(buf)

	err = br.discard(4)
	return f, err
}
