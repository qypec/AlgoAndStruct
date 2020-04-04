package main

import (
	"container/list"
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var packet Packet
	var buff Buffer

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	buff.size, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	packet.number, _ := strconv.Atoi(scanner.Text())

	for i := 0; scanner.Scan(); i++ {
		packet.arrival[i], _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		packet.duration[i], _ := strconv.Atoi(scanner.Text())
	}

	result := PacketProcessing(packet, buff)
}
