package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/test-input.txt
var testInput []byte

func TestReadInput(t *testing.T) {
	packet, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, "8A004A801A8002F478", packet)
}

func TestHexToBinary(t *testing.T) {
	binary, err := hexToBinary("8A004A801A8002F478")
	require.NoError(t, err)
	require.Equal(t, "100010100000000001001010100000000001101010000000000000101111010001111000", binary)
}

func TestReadLiteral(t *testing.T) {
	literal, bitsRead, err := readLiteral("101111111000101000")
	require.NoError(t, err)
	require.Equal(t, 15, bitsRead)
	require.Equal(t, uint64(2021), literal)
}

func TestReadOperatorPacket_TypeLenZero(t *testing.T) {
	packets, bitsRead, err := readPacket("00111000000000000110111101000101001010010001001000000000")
	require.NoError(t, err)
	require.Equal(t, 49, bitsRead)
	require.NotNil(t, packets)
	require.Len(t, packets.packets, 2)
}

func TestReadOperatorPacket_TypeLenOne(t *testing.T) {
	packets, bitsRead, err := readPacket("11101110000000001101010000001100100000100011000001100000")
	require.NoError(t, err)
	require.Equal(t, 51, bitsRead)
	require.NotNil(t, packets)
	require.Len(t, packets.packets, 3)
}
