// Copyright 2012 Google, Inc. All rights reserved.

package gopacket

import (
	"encoding/binary"
	"errors"
)

// PPP is the layer for PPP encapsulation headers.
type PPP struct {
	PPPType PPPType
}

// LayerType returns LayerTypePPP
func (p *PPP) LayerType() LayerType { return LayerTypePPP }
func (p *PPP) LinkFlow() Flow       { return PPPFlow }

func decodePPP(data []byte) (out DecodeResult, err error) {
	ppp := &PPP{}
	if data[0]&0x1 == 0 {
		if data[1]&0x1 == 0 {
			err = errors.New("PPP has invalid type")
			return
		}
		ppp.PPPType = PPPType(binary.BigEndian.Uint16(data[:2]))
		out.RemainingBytes = data[2:]
	} else {
		ppp.PPPType = PPPType(data[0])
		out.RemainingBytes = data[1:]
	}
	out.DecodedLayer = ppp
	out.NextDecoder = ppp.PPPType
	out.LinkLayer = ppp
	return
}
