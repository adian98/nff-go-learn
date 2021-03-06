package ipfix

import (
	"fmt"
	"net"

	"golang.org/x/net/ipv4"
)

//GenerateIPHeader 生成简易ip头
func GenerateIPHeader(srcIP, dstIP net.IP, dataLen int) (*ipv4.Header, error) {
	iph := &ipv4.Header{
		Version: ipv4.Version,
		//IP头长一般是20
		Len: ipv4.HeaderLen,
		TOS: 0x00,
		//buff为数据
		TotalLen: ipv4.HeaderLen + dataLen,
		TTL:      64,
		Flags:    ipv4.DontFragment,
		FragOff:  0,
		Protocol: 17,
		Checksum: 0,
		Src:      srcIP,
		Dst:      dstIP,
	}

	h, err := iph.Marshal()
	if err != nil {
		return nil, fmt.Errorf("Generate IP datagram error:%s", err.Error())
	}
	//计算IP头部校验值
	iph.Checksum = int(checkSum(h))

	return iph, nil
}

