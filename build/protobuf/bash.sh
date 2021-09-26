###
 # @Author: F1
 # @Date: 2020-07-20 16:19:50
 # @LastEditors: F1
 # @LastEditTime: 2021-09-18 22:47:58
 # @Description: 用于生成protocbuf .pb.go文件
### 
export PATH=$PATH;
#v3.17.3
protoc register.proto --go_out=.
protoc sendto.proto --go_out=.