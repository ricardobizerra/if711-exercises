package udp

import (
	"fmt"
	"net"
	"os"
	"slices"
	"strconv"
)

// Formato: Tamanho <5 caracteres> JSON codificado <tamanho caracteres> \xFF
func ReceiveUDPMessage(conn *net.UDPConn) ([]byte, *net.UDPAddr) {
	buffer := make([]byte, 64000)
	_, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	msg_len, err := strconv.Atoi(string(buffer[:5]))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	msg := buffer[5 : 5+msg_len]
	return msg, addr
}

func SendUDPMessage(conn *net.UDPConn, msg []byte) {
	msg_len := len(msg)
	msg_len_bytes := []byte(fmt.Sprintf("%05d", msg_len))

	end_byte := []byte{255}
	final_msg := slices.Concat(slices.Concat(msg_len_bytes, msg), end_byte)
	_, err := conn.Write(final_msg)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func SendUDPMessageAddr(conn *net.UDPConn, addr *net.UDPAddr, msg []byte) {
	msg_len := len(msg)
	msg_len_bytes := []byte(fmt.Sprintf("%05d", msg_len))

	end_byte := []byte{255}
	final_msg := slices.Concat(slices.Concat(msg_len_bytes, msg), end_byte)
	_, err := conn.WriteTo(final_msg, addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
