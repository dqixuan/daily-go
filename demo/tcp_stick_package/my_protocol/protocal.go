package my_protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

type MyProtocol struct {
}

// Encode 加密要传输的信息
func (mp *MyProtocol) Encode(s string) ([]byte, error) {
	length := int32(len(s))
	pkg := new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, fmt.Errorf("encode write length: %w", err)
	}
	err = binary.Write(pkg, binary.LittleEndian, []byte(s)) // 必须是基础类型或是基础类型切片
	if err != nil {
		return nil, fmt.Errorf("encode write body: %w", err)
	}
	return pkg.Bytes(), nil
}

func (mp *MyProtocol) Decode(reader *bufio.Reader) (string, error) {
	lengthByte, err := reader.Peek(4)
	if err != nil {
		return "", fmt.Errorf("decode peek: %w", err)
	}
	lengthBuf := bytes.NewBuffer(lengthByte)
	var dataLength int32
	err = binary.Read(lengthBuf, binary.LittleEndian, &dataLength)
	if err != nil {
		return "", fmt.Errorf("decode binary read: %w", err)
	}
	if int32(reader.Buffered()) < dataLength {
		return "", fmt.Errorf("wrong data")
	}
	pkg := make([]byte, 4+dataLength)
	_, err = reader.Read(pkg)
	if err != nil {
		return "", fmt.Errorf("decode reader Read: %w", err)
	}
	return string(pkg[4:]), nil
}
