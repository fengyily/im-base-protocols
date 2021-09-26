/*
 * @Author: F1
 * @Date: 2021-09-18 22:36:00
 * @LastEditTime: 2021-09-18 23:07:00
 * @LastEditors: F1
 * @Description:
 *  *
 *  *				Description
 *  *
 * @FilePath: /im-base-protocols/protocols/flag.go
 *
 */
package protocols

/**
 * @Title: Flag 长度８位，占１字节
 * @Description:
 *
 *				DataType 前３位用来标识传输的数据类型，所以最多支持８种，
 *				Encrytion 第４位表示数据包是否加密
 *				COMPRESS 第５位表示数据包是否开启了压缩
 * 				第６到８位为扩展备用位
 * 				  ____________________________________________
 * 				 | DataType  | Encrytion |COMPRESS|ext|ext|ext|
 * 				 |-----------|-----------|--------|---|---|---|
 * 				 | 3 bit     | 1 bit     | 1 bit  | 1 | 1 | 1 |
 * 				 |-----------|-----------|--------|---|---|---|
 * 				 |[0] [1] [2]|    [3]    |   [4]  |[5]|[6]|[7]|
 * 				 |____________________________________________|
 *
 * @Author: F1
 * @Date: 2020-07-21 10:55:39
 * @Param:
 * @Return:
 */
type Flag byte

const (
	HEADER_LENGTH                     = 6      // 头部的长度
	BODY_LENGTH                       = 4      // 包长所占长度
	HEADER_FLAG_DATA_TYPE_JSON   Flag = 1      // 0000 0001
	HEADER_FLAG_DATA_TYPE_PB     Flag = 1 << 1 // 0000 0010
	HEADER_FLAG_DATA_TYPE_STRING Flag = 1 << 2 // 0000 0100
	HEADER_FLAG_IS_ENCRYTION     Flag = 1 << 3 // 0001 0000 是否加密
	HEADER_FLAG_IS_COMPRESS      Flag = 1 << 4 // 0010 0000 是否开启了压缩
	HEADER_FLAG_EXT1             Flag = 1 << 5 // 0000 1000
	HEADER_FLAG_EXT2             Flag = 1 << 6 // 0100 0000
	HEADER_FLAG_EXT3             Flag = 1 << 7 // 1000 0000

)
