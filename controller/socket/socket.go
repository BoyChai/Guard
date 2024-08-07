package socket

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/BoyChai/Guard/dao"
	"github.com/BoyChai/Guard/utils"
	"github.com/spf13/viper"
)

func StartSocket() {
	// 监听本地端口
	listen, err := net.Listen("tcp", ":"+viper.GetString("Settings.Socket_Port"))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listen.Close()
	// 等待客户端链接
	for {
		// listen.Accept如果没有链接会一直阻塞
		conn, err := listen.Accept() //等待建立链接
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 启动一个goroutine去处理链接
		go process(conn)
	}
}
func process(conn net.Conn) {
	defer conn.Close() // 处理完关闭链接

	var recvData string // 保存接收到的数据的变量

	reader := bufio.NewReader(conn)
	delimiter := viper.GetString("Settings.Socket_EndSymbol")
	buf := make([]byte, 1024)

	for {
		// 设置读取操作的超时时间10秒
		err := conn.SetReadDeadline(time.Now().Add(time.Duration(viper.GetInt64("Settings.Socket_Timeout")) * time.Second))
		if err != nil {
			fmt.Println("[Socket] " + conn.RemoteAddr().String() + " connection timed out: " + err.Error())
			return
		}
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		recvData += string(buf[:n])
		// 检查接收的数据是否包含分隔符
		if strings.Contains(recvData, delimiter) {
			// 找到分隔符，停止读取
			break
		}
	}

	// 处理完整的消息
	data, err := utils.DecryptWithPrivateKey(recvData)
	if err != nil {
		conn.Write([]byte("No, that's not right: " + err.Error() + delimiter))
		fmt.Println("[Socket] " + conn.RemoteAddr().String() + " No, that's not right: " + err.Error())
		return
	}
	isTrue, err := dao.Dao.CheckCard(data)
	if !isTrue {
		conn.Write([]byte("No, that's not right: " + err.Error() + delimiter))
		fmt.Println("[Socket] " + conn.RemoteAddr().String() + " No, that's not right: " + err.Error())
		return
	}
	log.Println(recvData)
	signData, err := utils.SignData(recvData)
	if !isTrue {
		conn.Write([]byte("server error:" + err.Error() + delimiter))
		fmt.Println("[Socket] " + conn.RemoteAddr().String() + " server error:" + err.Error())
		return
	}
	conn.Write([]byte(signData + delimiter))
}
