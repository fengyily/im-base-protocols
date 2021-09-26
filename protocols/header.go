/*
 * @Author: F1
 * @Date: 2020-07-14 21:16:18
 * @LastEditors: F1
 * @LastEditTime: 2021-09-20 20:11:06
 * @Description: 协议包中的头部相关定义
 */
package protocols

import (
	"fmt"

	"im-base-protocols/utils"
)

/**
 * @Title:协议头部，包含指令，标识及包体的长度
 * @Description:
 * 		一般来讲一个完整的包是由Header|Body组成，特殊的包只有包头：
 *　	心跳包及心跳回复包，只有1字节，其它包允许只有包头，也就是Body为空，此时Header.Length为0
 *
 * @Author: F1
 * @Date: 2020-07-21 10:01:05
 * @LastEditors: F1
 * @Param:
 * @Return:
 *		[]byte 长度为：protocols.HEADER_LENGTH
 */
type Header struct {
	Cmd    Command
	Flag   Flag
	Length uint32
}

/**
 * @Title: ToBytes
 * @Description: 将头部转换为字节数组，将返回1-4字节
 *
 * @Author: F1
 * @Date: 2020-07-21 10:48:07
 * @Return:[]byte 1-4 byte
 */
func (header *Header) ToBytes() []byte {
	var data []byte
	if header.Cmd == HEARTBEAT_REQUEST || header.Cmd == HEARTBEAT_RESPONSE {
		data = make([]byte, 1)
		data[0] = byte(header.Cmd)
	} else {
		data = make([]byte, HEADER_LENGTH)
		data[0] = byte(header.Cmd)
		data[1] = byte(header.Flag)

		//copy(&data, 2, utils.Uint16ToBytes(header.Length), 0, 2)
		copy(&data, 2, utils.UintToBytes(header.Length), 0, BODY_LENGTH)
		fmt.Println("cmd", header.Cmd, "length", header.Length)
	}
	return data
}

/**
 * @Title: LoadHeader
 * @Description: 从字节流中加载头部，如果字节流不是以头部开始，将被丢弃，直致找到下一个头部为止。
 *
 * @Author: F1
 * @Date: 2020-07-21 10:49:49
 * @Param: buffer *[]byte 字节流
 * @Return:ok bool	是否成功的加载出头部
 * @Return:header Header　头部对象
 */
func LoadHeader(buffer *[]byte) (ok bool, header Header) {
	ok = true
	for total := len(*buffer); total > 0; total = len(*buffer) {
		i := 0
		header.Cmd = Command((*buffer)[i])
		// 心跳包相关的是1字节长度，正常包为>=4字节，接受body为0的包
		if header.Cmd == HEARTBEAT_REQUEST ||
			header.Cmd == HEARTBEAT_RESPONSE {
			*buffer = (*buffer)[i+1:]
			return ok, header
		}

		//var length int32
		// 正常包的处理逻辑，如果遇到不识别的包，重新定位头部位置。
		if !header.Cmd.IsCommandType((*buffer)[i]) {
			rIndex := 0
			for ; rIndex < total-i; rIndex++ {
				if header.Cmd.IsCommandType((*buffer)[i+rIndex]) {
					fmt.Println("重新定位成功。丢弃　", i+rIndex, "字节。")
					break
				}
			}
			*buffer = (*buffer)[i+rIndex:]
			fmt.Println("未识别的指令", rIndex, len(*buffer))
			continue
		}

		// 是否满足正常包的一个头部
		if total < i+HEADER_LENGTH {
			fmt.Println(fmt.Sprintf("len(cs.Buffer) < i+%d", HEADER_LENGTH))
			return false, header
		}
		header.Flag = Flag((*buffer)[1])
		// body length = HEADER_LENGTH - BODY_LENGTH 到 HEADER_LENGTH
		header.Length = utils.BytesToUInt((*buffer)[HEADER_LENGTH-BODY_LENGTH : HEADER_LENGTH])

		break
	}
	return true, header
}

func copy(desc *[]byte, index int, src []byte, startIndex int, length int) {
	for i := startIndex; i < startIndex+length; i++ {
		(*desc)[index] = src[i]
		index++
	}
}
