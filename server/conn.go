package main

import (
	"context"
	"encoding/binary"
	"errors"
	"github.com/quic-go/quic-go"
)

type conn struct {
	Connection quic.Connection
	infoStream quic.Stream
	dataStream quic.Stream
}

// 自定义的Conn，方便操作
func newConn(sess quic.Connection, is_server bool) (*conn, error) {
	if is_server {
		dstream, err := sess.OpenStream()
		if err != nil {
			return nil, err
		}
		return &conn{
			Connection: sess,
			dataStream: dstream,
		}, nil
	} else {
		istream, err := sess.OpenStream()
		if err != nil {
			return nil, err
		}
		return &conn{
			Connection: sess,
			infoStream: istream,
		}, nil
	}
}

//	func (c *conn) DataStream() quic.Stream {
//		return c.dataStream
//	}
func (c *conn) ReadSeq(seq *uint16) (int, error) {
	if c.infoStream == nil {
		var err error
		c.infoStream, err = c.Connection.AcceptStream(context.Background())
		// TODO: check stream id
		if err != nil {
			return 0, err
		}
	}
	seq_b := make([]byte, 2)
	_, err := c.infoStream.Read(seq_b)
	if err != nil {
		return -1, err
	}
	*seq = binary.BigEndian.Uint16(seq_b)
	return 0, err

	//return io.ReadFull(c.dataStream,b)
}

func (c *conn) ReadSsrc(ssrc *uint32) error {
	if c.infoStream == nil {
		var err error
		c.infoStream, err = c.Connection.AcceptStream(context.Background())
		// TODO: check stream id
		if err != nil {
			return err
		}
	}
	ssrc_b := make([]byte, 4)
	_, err := c.infoStream.Read(ssrc_b)
	if err != nil {
		return errors.New("read ssrc failed")
	}
	*ssrc = binary.BigEndian.Uint32(ssrc_b)
	return err

	//return io.ReadFull(c.dataStream,b)
}

func (c *conn) SendLen(len uint16) (int, error) {
	len_b := make([]byte, 2)
	binary.BigEndian.PutUint16(len_b, len)
	return c.infoStream.Write(len_b)
}

func (c *conn) SendRtp(pkt []byte) (int, error) {
	if pkt == nil { //缓存中没找到该包
		_, err := c.SendLen(uint16(0))
		if err != nil {
			panic(err)
		}
		return 0, nil
	} else {
		_, err := c.SendLen(uint16(len(pkt)))
		if err != nil {
			panic(err)
		}
		return c.dataStream.Write(pkt)
	}
}
