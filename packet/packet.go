package packet

import "encoding/binary"

const (
	ACK       = 0
	DATA      = 1
	EOT       = 2
	MaxData   = 500
	SeqNumMod = 32
)

type Packet struct {
	Type   int    `json:"type"`    // 0: ACK, 1: Data, 2: EOT
	SeqNum int    `json:"seq_num"` // Modulo 32
	Length int    `json:"length"`  // Length of the String variable ‘data’
	Data   []byte `json:"data"`    // String with Max Length 500
}

func (p Packet) Bytes() []byte {
	b := make([]byte, 12)
	binary.BigEndian.PutUint32(b, uint32(p.Type))
	binary.BigEndian.PutUint32(b[4:], uint32(p.SeqNum))
	binary.BigEndian.PutUint32(b[8:], uint32(p.Length))
	b = append(b, []byte(p.Data)...)
	return b
}

func DecodePacket(b []byte) Packet {
	p := Packet{}
	p.Type = int(binary.BigEndian.Uint32(b))
	p.SeqNum = int(binary.BigEndian.Uint32(b[4:]))
	p.Length = int(binary.BigEndian.Uint32(b[8:]))
	p.Data = b[12 : p.Length+12]
	return p
}
