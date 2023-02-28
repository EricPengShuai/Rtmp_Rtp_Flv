package main

var (
	HEADER_BYTES = []byte{'F', 'L', 'V', 0x01, 0x05, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x00,
		0x12, 0x00, 0x00, 0x28, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // 11
		0x02, 0x00, 0x0a, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, // 13
		0x08, 0x00, 0x00, 0x00, 0x01, // 5
		0x00, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6F, 0x6E, // 10
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // 9
		0x00, 0x00, 0x09, // 3
		0x00, 0x00, 0x00, 0x33}
)

const (
	AUDIO_TAG           = byte(0x08)
	VIDEO_TAG           = byte(0x09)
	SCRIPT_DATA_TAG     = byte(0x12)
	DURATION_OFFSET     = 53
	HEADER_LEN          = 13
	MAX_RTP_PAYLOAD_LEN = 1000
	PACKET_LOSS_RATE    = 0.001
	MULTICAST_ADDRASS   = "239.0.0.0:5222"
	QUIC_ADDR           = "localhost:4242"
	USE_MULTICAST       = true
	RTP_INITIAL_SEQ     = uint16(65000)
)
