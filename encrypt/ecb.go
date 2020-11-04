package encrypt

import (
	"bytes"
	"crypto/aes"
	"errors"
)

type AesECB struct {
	key       []byte
	blockSize int
}

func NewEcb(key []byte, blocks ...int) *AesECB {
	size := aes.BlockSize
	if len(blocks) > 0 {
		size = blocks[0]
	}
	return &AesECB{
		key:       key,
		blockSize: size,
	}
}

// When isPad is false, use Nopadding mode
func (a *AesECB) Encrypt(src []byte, isPad ...bool) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}

	if len(src) == 0 {
		return nil, errors.New("content is empty")
	}

	if len(isPad) > 0 && isPad[0] == false {
		src = a.noPadding(src)
	} else {
		src = a.padding(src)
	}

	buf := make([]byte, a.blockSize)
	encrypted := make([]byte, 0)
	for i := 0; i < len(src); i += a.blockSize {
		block.Encrypt(buf, src[i:i+a.blockSize])
		encrypted = append(encrypted, buf...)
	}
	return encrypted, nil
}

// When isPad is false, use Nopadding mode
func (a *AesECB) Decrypt(src []byte, isPad ...bool) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}
	if len(src) == 0 {
		return nil, errors.New("content is empty")
	}
	buf := make([]byte, a.blockSize)
	decrypted := make([]byte, 0)
	for i := 0; i < len(src); i += a.blockSize {
		block.Decrypt(buf, src[i:i+a.blockSize])
		decrypted = append(decrypted, buf...)
	}

	if len(isPad) > 0 && isPad[0] == false {
		decrypted = a.unNoPadding(decrypted)
	} else {
		decrypted = a.unPadding(decrypted)
	}

	return decrypted, nil
}

// Nopadding mode
func (a *AesECB) noPadding(src []byte) []byte {
	count := a.blockSize - len(src)%a.blockSize
	if len(src)%a.blockSize == 0 {
		return src
	} else {
		return append(src, bytes.Repeat([]byte{byte(0)}, count)...)
	}
}

// Nopadding mode
func (a *AesECB) unNoPadding(src []byte) []byte {
	for i := len(src) - 1; ; i-- {
		if src[i] != 0 {
			return src[:i+1]
		}
	}
	return nil
}

// Padding Mode
func (a *AesECB) padding(src []byte) []byte {
	count := a.blockSize - len(src)%a.blockSize
	padding := bytes.Repeat([]byte{byte(0)}, count)
	padding[count-1] = byte(count)
	return append(src, padding...)
}

// Padding Mode
func (a *AesECB) unPadding(src []byte) []byte {
	l := len(src)
	p := int(src[l-1])
	return src[:l-p]
}
