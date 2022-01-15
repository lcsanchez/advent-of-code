package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed testdata/input.txt
var input []byte

const (
	LITERAL_GROUP_LEN  = 4
	HEADER_BIT_LEN     = 6
	LEN_TYPE_0_BIT_LEN = 16
)

type Header struct {
	version int
	typeID  int
}

type Packet struct {
	header  *Header
	literal uint64
	packets []*Packet
}

type PacketOp func(values ...uint64) uint64

func SumOp(values ...uint64) uint64 {
	result := uint64(0)

	for _, value := range values {
		result += value
	}

	return result
}

func ProductOp(values ...uint64) uint64 {
	result := uint64(1)

	for _, value := range values {
		result *= value
	}

	return result
}

func MinOp(values ...uint64) uint64 {
	result := uint64(math.MaxUint64)

	for _, value := range values {
		if value < result {
			result = value
		}
	}

	return result
}

func MaxOp(values ...uint64) uint64 {
	result := uint64(0)

	for _, value := range values {
		if value > result {
			result = value
		}
	}

	return result
}

func GreaterThanOp(values ...uint64) uint64 {
	if values[0] > values[1] {
		return 1
	}

	return 0
}

func LessThanOp(values ...uint64) uint64 {
	if values[0] < values[1] {
		return 1
	}

	return 0
}

func EqualToOp(values ...uint64) uint64 {
	if values[0] == values[1] {
		return 1
	}

	return 0
}

func main() {
	input, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	binary, err := hexToBinary(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(binary)

	packet, _, err := readPacket(binary)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumVersions(packet))
	fmt.Println(processPacket(packet))
}

func sumVersions(p *Packet) int {
	sum := p.header.version

	for _, sub := range p.packets {
		sum += sumVersions(sub)
	}

	return sum
}

func processPacket(p *Packet) uint64 {
	if p.header.typeID == 4 {
		return p.literal
	}

	values := []uint64{}
	for _, packet := range p.packets {
		values = append(values, processPacket(packet))
	}

	var op PacketOp
	switch p.header.typeID {
	case 0:
		op = SumOp
	case 1:
		op = ProductOp
	case 2:
		op = MinOp
	case 3:
		op = MaxOp
	case 5:
		op = GreaterThanOp
	case 6:
		op = LessThanOp
	case 7:
		op = EqualToOp
	}

	return op(values...)
}

func readInput(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return scanner.Text(), nil
}

func hexToBinary(hex string) (string, error) {
	var sb strings.Builder

	for _, r := range hex {
		ui, err := strconv.ParseUint(string(r), 16, 4)
		if err != nil {
			return "", err
		}
		sb.WriteString(fmt.Sprintf("%04b", ui))
	}

	return sb.String(), nil
}

func readPacket(b string) (*Packet, int, error) {
	header, err := readPacketHeader(b)
	if err != nil {
		return nil, 0, err
	}

	if header.typeID == 4 {
		literal, bitsRead, err := readLiteral(b[6:])
		if err != nil {
			return nil, 0, err
		}

		return &Packet{
			header:  header,
			literal: literal,
		}, HEADER_BIT_LEN + bitsRead, nil
	}

	packets, bitsRead, err := readOperator(header, b[6:])
	if err != nil {
		return nil, 0, err
	}

	return &Packet{
		header:  header,
		packets: packets,
	}, HEADER_BIT_LEN + bitsRead, nil
}

func readPacketHeader(b string) (*Header, error) {
	version, err := strconv.ParseUint(b[:3], 2, 32)
	if err != nil {
		return nil, err
	}

	typeID, err := strconv.ParseUint(b[3:6], 2, 32)
	if err != nil {
		return nil, err
	}

	return &Header{version: int(version), typeID: int(typeID)}, nil
}

func readLiteral(b string) (uint64, int, error) {
	var idx = 0
	var literalBin string
	var isEnd bool
	for !isEnd {

		if len(b) < idx+LITERAL_GROUP_LEN {
			return 0, 0, fmt.Errorf("unable to read literal: %v", b)
		}

		isEnd = b[idx] == '0'
		idx += 1

		literalBin += b[idx : idx+(LITERAL_GROUP_LEN)]
		idx += LITERAL_GROUP_LEN
	}

	num, err := strconv.ParseUint(literalBin, 2, 64)
	if err != nil {
		return 0, 0, err
	}

	return num, idx, nil
}

func readOperator(h *Header, b string) ([]*Packet, int, error) {
	lengthTypeID := b[0]

	if lengthTypeID == '0' {
		subPacketsBitLen, err := strconv.ParseUint(b[1:16], 2, 32)
		if err != nil {
			return nil, 0, err
		}

		packets := []*Packet{}
		subPacketBits := b[16:]
		for i := 0; i < int(subPacketsBitLen); {
			packet, bitsRead, err := readPacket(subPacketBits[i:])
			if err != nil {
				return nil, 0, err
			}

			packets = append(packets, packet)
			i += bitsRead
		}

		return packets, int(subPacketsBitLen) + LEN_TYPE_0_BIT_LEN, nil
	} else {
		numSubPackets, err := strconv.ParseUint(b[1:12], 2, 32)
		if err != nil {
			return nil, 0, err
		}

		var totalBitsRead = 12
		packets := []*Packet{}
		for i := 0; i < int(numSubPackets); i++ {
			packet, bitsRead, err := readPacket(b[totalBitsRead:])
			if err != nil {
				return nil, 0, err
			}

			packets = append(packets, packet)
			totalBitsRead += bitsRead
		}

		return packets, totalBitsRead, nil
	}
}
