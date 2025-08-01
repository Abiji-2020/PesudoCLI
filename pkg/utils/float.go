/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/

package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func Float32SliceToBytes(floats []float32) ([]byte, error) {
	buf := new(bytes.Buffer)
	for _, f := range floats {
		if err := binary.Write(buf, binary.LittleEndian, f); err != nil {
			return nil, fmt.Errorf("failed to write float32 to bytes: %w", err)
		}
	}
	return buf.Bytes(), nil
}

func GetID(command string, os string) string {
	return fmt.Sprintf("%s_%s", command, os)
}
