/*
 * @Author: F1
 * @Date: 2021-08-28 19:40:43
 * @LastEditTime: 2021-09-18 23:00:06
 * @LastEditors: F1
 * @Description:
 *  *
 *  *				test
 *  *
 * @FilePath: /im-base-protocols/utils/utils_test.go
 *
 */
package utils

import "testing"

func Test_Header(t *testing.T) {
	println(2 << 31)
	println(2 << 30)
	b := UintToBytes(2 << 30)
	println("Test BytesToUInt:...", BytesToUInt(b))
}
