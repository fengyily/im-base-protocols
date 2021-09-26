/*
 * protocols 协议包
 * @Author: F1
 * @Date: 2020-07-14 21:16:18
 * @LastEditTime: 2021-09-18 23:08:33
 * @LastEditors: F1
 * @Description: 协议包中指令部份，目前支持0-255的指令定义
 *
 * @FilePath: /im-base-protocols/protocols/command.go
 */
package protocols

import (
	"reflect"
	"strconv"

	"google.golang.org/protobuf/reflect/protoreflect"
)

/**
 * @Title: Command 消息指令定义
 * @Description:
 *
 * 	    消息指令位于头部（Header）中第一个字节，最多支持255种指令
 *  __________________________________________________________________
 * | 1 byte  | 1 byte  | 4 byte          | length byte                |
 * | ------  | ------  | --------------- | ----------------           |
 * | command | flag    | length          |  body                      |
 * | ------  | ------  | --------------- | ----------------           |
 * | [0]     | [1]     | [2][3][4][5]    | [4][5][6][7][]	 		  |
 * | 0-255   | 0-255   | 0-2^32          | length                     |
 * |__________________________________________________________________|
 *
 * @Author: F1
 * @Date: 2020-07-21 11:02:42
 */
type Command byte

const (
	HEARTBEAT_REQUEST         Command = 0 // 心跳包
	HEARTBEAT_RESPONSE        Command = 1 // 心跳包响应
	REGISTER_REQUEST          Command = 2 // 边缘端向服务端注册
	REGISTER_RESPONSE_SUCCESS Command = 3 // 注册响应 成功
	REGISTER_RESPONS_FAILED   Command = 4 // 注册响应 成功
	SENDTO_REQUEST            Command = 5 // 1 - 1 发消息
	SENDTO_RESPONSE           Command = 6 // 1 - 1 消息响应
	CAST_MSG_REQUEST          Command = 7 // 广播消息
	CAST_MSG_RESPONSE         Command = 8 // 广播响应
	TIMEOUT                   Command = 9 // 超时啦
)

// Enum value maps for Command.
var (
	Command_name = map[int32]string{
		0: "HEARTBEAT_REQUEST",
		1: "HEARTBEAT_RESPONSE",
		2: "REGISTER_REQUEST",
		3: "REGISTER_RESPONSE_SUCCESS",
		4: "REGISTER_RESPONS_FAILED",
		5: "SENDTO_REQUEST",
		6: "SENDTO_RESPONSE",
		7: "CAST_MSG_REQUEST",
		8: "CAST_MSG_RESPONSE",
		9: "TIMEOUT",
	}
)

func (x Command) String() string {
	return Command_name[int32(x.Number())]
}

func (x Command) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

/**
 * @Title: IsCommandType
 * @Description:
 *
 * 			说明：判断字节流中的当前位是否是定义的指令类型
 *
 * @Author: F1
 * @Date: 2020-07-21 11:05:14
 * @Param:
 * 		c byte 字节流中的第一位
 * @Return:
 *		ok bool 是否是指令类型
 */
func (cmd Command) IsCommandType(c byte) bool {
	tmp := Command(c)
	return reflect.TypeOf(tmp) == reflect.TypeOf(HEARTBEAT_RESPONSE)
}

/**
* @Title: ToString
* @Description:
* @Author: F1
* @Date: 2020-07-21 11:07:06
 * @Return: string
*/
func (cmd Command) ToString() string {
	return strconv.Itoa(int(cmd))
}
