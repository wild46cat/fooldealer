package socket

import (
	"bytes"
	"encoding/binary"
	"errors"
)

const BUFFER_SIZE = 512
const HEAD_SIZE = 4

/**
基本消息结构 消息头+body
4 + n
 */
type BaseMessageInfo struct {
	head int
	body string
}

type BaseBytesMessageInfo struct {
	head []byte
	body []byte
}

func ConvertToBytes(str string) []byte {
	temp := BaseMessageInfo{
		head: len(str),
		body: str,
	}
	bytesMessageInfo := convertToByteInfo(temp)
	res := make([]byte, 0)
	res = append(res, bytesMessageInfo.head...)
	res = append(res, bytesMessageInfo.body...)
	return res
}

//转换成head结构
func ConvertFromBytes(bytes []byte) (info []BaseMessageInfo, lastBytes []byte, err error) {
	byteSize := len(bytes)
	//fmt.Print("bytes content:")
	//fmt.Println(bytes)
	if byteSize < HEAD_SIZE {
		return nil, nil, errors.New("消息格式不正确")
	} else {
		res := make([]BaseMessageInfo, 0)
		headLength := 0
		lastOffset := 0
		for offset := 0; offset <= byteSize && offset+HEAD_SIZE <= byteSize; offset = offset + HEAD_SIZE + headLength {
			head := bytes[offset : offset+HEAD_SIZE]
			headLength = bytesToInt32(head)
			if offset+HEAD_SIZE+headLength <= byteSize {
				body := bytes[offset+HEAD_SIZE : offset+HEAD_SIZE+headLength]
				res = append(res, BaseMessageInfo{
					head: headLength,
					body: string(body),
				})
				lastOffset = offset + HEAD_SIZE + headLength
			} else {
				//存在半包
				lastOffset = offset
			}
		}
		lastBytes := make([]byte, 0)
		//fmt.Print("lastOffset bytes:")
		//fmt.Println(bytes[lastOffset:])
		lastBytes = append(lastBytes, bytes[lastOffset:]...)
		return res, lastBytes, nil
	}
}

func convertToByteInfo(info BaseMessageInfo) BaseBytesMessageInfo {
	return BaseBytesMessageInfo{
		head: int32ToBytes(info.head),
		body: []byte(info.body),
	}
}

func convertFromByteInfo(byteInfo BaseBytesMessageInfo) BaseMessageInfo {
	return BaseMessageInfo{
		head: bytesToInt32(byteInfo.head),
		body: string(byteInfo.body),
	}
}

//整形转换成字节
func int32ToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func bytesToInt32(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
