package tools

// 系统工具
import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

// 将int64转为字节数组
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// 反转数组
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// 获取本地ip
func GetLocalIp1() string {
	var ip_str string
	addrSlice, err := net.InterfaceAddrs()
	if nil != err {
		log.Panic("获取本地IP失败!")
		ip_str = "localhost"
	}
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				ip_str = ipnet.IP.String()
				return ip_str
			}
		}
	}
	ip_str = "localhost"
	return ip_str
}

// 获取本地ip
func GetLocalIp() string {
	var ip_str string
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
		ip_str = "localhost"
		log.Panic("获取本地IP失败!")
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						ip_str = ipnet.IP.String()
						return ip_str
					}
				}
			}
		}
	}
	ip_str = "localhost"
	return ip_str
}