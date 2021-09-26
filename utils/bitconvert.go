/*
 * @Author: F1
 * @Date: 2020-07-14 21:16:18
 * @LastEditors: F1
 * @LastEditTime: 2021-09-20 20:10:30
 * @Description:
 *
 *				协议中通用的一些函数，通常为字节流转常用类型以及常用类型转字节流
 *
 */
package utils

import (
	"bytes"
	"encoding/binary"
)

/**
 * @Title:Uint16ToBytes
 * @Description:
 *
 * 				将uint16位的数字转成byte
 *
 * @Author: F1
 * @Date: 2020-07-21 11:09:34
 * @Param:uint16
 * @Return:[]byte
 */
func Uint16ToBytes(data uint16) []byte {
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func UintToBytes(data uint32) []byte {
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

/**
 * @Title:BytesToUInt16
 * @Description:
 *
 *				字节流转uint16,默认为大端
 *
 * @Author: F1
 * @Date: 2020-07-21 11:11:38
 * @Param: []byte
 * @Return: uint16
 */
func BytesToUInt16(bys []byte) uint16 {
	bytebuff := bytes.NewBuffer(bys)
	var data uint16
	binary.Read(bytebuff, binary.BigEndian, &data)
	return uint16(data)
}

func BytesToUInt(bys []byte) uint32 {
	bytebuff := bytes.NewBuffer(bys)
	var data uint32
	binary.Read(bytebuff, binary.BigEndian, &data)
	return uint32(data)
}
