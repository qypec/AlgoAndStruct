package main

import (
	"container/list"
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type Packet struct {
	arrival []int
	duration []int
	result []int
	counter int
	number int
}

func (p *Packet) Init() {
	p.arrival = make([]int, p.number)
	p.duration = make([]int, p.number)
	p.result = make([]int, p.number)
}

func (p *Packet) Cancel(packetIndex int) {
	if p.result[packetIndex] == 0 {
		p.result[packetIndex] = -1
	}
}

type Buffer struct {
	size int
	usedSize int
}

func (b *Buffer) isFull() {
	if b.size == b.usedSize {
		return true
	}
	return false
}

func (b *Buffer) Inc() {
	if b.usedSize < b.size {
		b.usedSize++
	}
}

func (b *Buffer) Dec() {
	if b.usedSize != 0 {
		b.usedSize--
	}
}

func PacketProcessing(packet *Packet, buff *Buffer) []int {
	// queue = list.New()

	time := packet.arrival[0]
	for ; packet.counter < packet.number; packet.counter++ {
		for j := packet.counter; j < packet.number && time == packet.arrival[j]; j++ {
			if buff.isFull() {
				packet.Cancel(j)
			} else {
				buff.Inc()
			}
		}

		for {
			
		}
	}
}

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
