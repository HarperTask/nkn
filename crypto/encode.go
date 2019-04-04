package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/nknorg/nkn/crypto/ed25519"
)

const (
	INFINITYLEN      = 1
	FLAGLEN          = 1
	XORYVALUELEN     = 32
	COMPRESSEDLEN    = 33
	NOCOMPRESSEDLEN  = 65
	COMPEVENFLAG     = 0x02
	COMPODDFLAG      = 0x03
	NOCOMPRESSEDFLAG = 0x04
	P256PARAMA       = -3
)

func isEven(k *big.Int) bool {
	z := big.NewInt(0)
	z.Mod(k, big.NewInt(2))
	if z.Int64() == 0 {
		return true
	}
	return false
}

func getLowestSetBit(k *big.Int) int {
	i := 0
	for i = 0; k.Bit(i) != 1; i++ {
	}
	return i
}

// fastLucasSequence refer to https://en.wikipedia.org/wiki/Lucas_sequence
func fastLucasSequence(curveP, lucasParamP, lucasParamQ, k *big.Int) (*big.Int, *big.Int) {
	n := k.BitLen()
	s := getLowestSetBit(k)

	uh := big.NewInt(1)
	vl := big.NewInt(2)
	ql := big.NewInt(1)
	qh := big.NewInt(1)
	vh := big.NewInt(0).Set(lucasParamP)
	tmp := big.NewInt(0)

	for j := n - 1; j >= s+1; j-- {
		ql.Mul(ql, qh)
		ql.Mod(ql, curveP)

		if k.Bit(j) == 1 {
			qh.Mul(ql, lucasParamQ)
			qh.Mod(qh, curveP)

			uh.Mul(uh, vh)
			uh.Mod(uh, curveP)

			vl.Mul(vh, vl)
			tmp.Mul(lucasParamP, ql)
			vl.Sub(vl, tmp)
			vl.Mod(vl, curveP)

			vh.Mul(vh, vh)
			tmp.Lsh(qh, 1)
			vh.Sub(vh, tmp)
			vh.Mod(vh, curveP)
		} else {
			qh.Set(ql)

			uh.Mul(uh, vl)
			uh.Sub(uh, ql)
			uh.Mod(uh, curveP)

			vh.Mul(vh, vl)
			tmp.Mul(lucasParamP, ql)
			vh.Sub(vh, tmp)
			vh.Mod(vh, curveP)

			vl.Mul(vl, vl)
			tmp.Lsh(ql, 1)
			vl.Sub(vl, tmp)
			vl.Mod(vl, curveP)
		}
	}

	ql.Mul(ql, qh)
	ql.Mod(ql, curveP)

	qh.Mul(ql, lucasParamQ)
	qh.Mod(qh, curveP)

	uh.Mul(uh, vl)
	uh.Sub(uh, ql)
	uh.Mod(uh, curveP)

	vl.Mul(vh, vl)
	tmp.Mul(lucasParamP, ql)
	vl.Sub(vl, tmp)
	vl.Mod(vl, curveP)

	ql.Mul(ql, qh)
	ql.Mod(ql, curveP)

	for j := 1; j <= s; j++ {
		uh.Mul(uh, vl)
		uh.Mul(uh, curveP)

		vl.Mul(vl, vl)
		tmp.Lsh(ql, 1)
		vl.Sub(vl, tmp)
		vl.Mod(vl, curveP)

		ql.Mul(ql, ql)
		ql.Mod(ql, curveP)
	}

	return uh, vl
}

// compute the coordinate of Y from Y**2
func curveSqrt(ySquare *big.Int, curve *elliptic.CurveParams) *big.Int {
	if curve.P.Bit(1) == 1 {
		tmp1 := big.NewInt(0)
		tmp1.Rsh(curve.P, 2)
		tmp1.Add(tmp1, big.NewInt(1))

		tmp2 := big.NewInt(0)
		tmp2.Exp(ySquare, tmp1, curve.P)

		tmp3 := big.NewInt(0)
		tmp3.Exp(tmp2, big.NewInt(2), curve.P)

		if 0 == tmp3.Cmp(ySquare) {
			return tmp2
		}
		return nil
	}

	qMinusOne := big.NewInt(0)
	qMinusOne.Sub(curve.P, big.NewInt(1))

	legendExponent := big.NewInt(0)
	legendExponent.Rsh(qMinusOne, 1)

	tmp4 := big.NewInt(0)
	tmp4.Exp(ySquare, legendExponent, curve.P)
	if 0 != tmp4.Cmp(big.NewInt(1)) {
		return nil
	}

	k := big.NewInt(0)
	k.Rsh(qMinusOne, 2)
	k.Lsh(k, 1)
	k.Add(k, big.NewInt(1))

	lucasParamQ := big.NewInt(0)
	lucasParamQ.Set(ySquare)
	fourQ := big.NewInt(0)
	fourQ.Lsh(lucasParamQ, 2)
	fourQ.Mod(fourQ, curve.P)

	seqU := big.NewInt(0)
	seqV := big.NewInt(0)

	for {
		lucasParamP := big.NewInt(0)
		for {
			tmp5 := big.NewInt(0)
			lucasParamP, _ = rand.Prime(rand.Reader, curve.P.BitLen())

			if lucasParamP.Cmp(curve.P) < 0 {
				tmp5.Mul(lucasParamP, lucasParamP)
				tmp5.Sub(tmp5, fourQ)
				tmp5.Exp(tmp5, legendExponent, curve.P)

				if 0 == tmp5.Cmp(qMinusOne) {
					break
				}
			}
		}

		seqU, seqV = fastLucasSequence(curve.P, lucasParamP, lucasParamQ, k)

		tmp6 := big.NewInt(0)
		tmp6.Mul(seqV, seqV)
		tmp6.Mod(tmp6, curve.P)
		if 0 == tmp6.Cmp(fourQ) {
			if 1 == seqV.Bit(0) {
				seqV.Add(seqV, curve.P)
			}
			seqV.Rsh(seqV, 1)
			return seqV
		}
		if (0 == seqU.Cmp(big.NewInt(1))) || (0 == seqU.Cmp(qMinusOne)) {
			break
		}
	}
	return nil
}

// deCompress is for computing the coordinate of Y based the coordinate of X
func deCompress(yTilde int, xValue []byte, curve *elliptic.CurveParams) (*PubKey, error) {
	xCoord := big.NewInt(0)
	xCoord.SetBytes(xValue)

	//y**2 = x**3 + A*x +B, A = -3, there is no A's clear definition in the realization of p256.
	paramA := big.NewInt(P256PARAMA)
	//compute x**3 + A*x +B
	ySqare := big.NewInt(0)
	ySqare.Exp(xCoord, big.NewInt(2), curve.P)
	ySqare.Add(ySqare, paramA)
	ySqare.Mod(ySqare, curve.P)
	ySqare.Mul(ySqare, xCoord)
	ySqare.Mod(ySqare, curve.P)
	ySqare.Add(ySqare, curve.B)
	ySqare.Mod(ySqare, curve.P)

	yValue := curveSqrt(ySqare, curve)
	if nil == yValue {
		return nil, errors.New("Invalid point compression")
	}

	yCoord := big.NewInt(0)
	if (isEven(yValue) && 0 != yTilde) || (!isEven(yValue) && 1 != yTilde) {
		yCoord.Sub(curve.P, yValue)
	} else {
		yCoord.Set(yValue)
	}
	return &PubKey{xCoord, yCoord}, nil
}

func DecodePoint(encodeData []byte) (*PubKey, error) {
	if len(encodeData) == 0 {
		return nil, errors.New("The encodeData cann't be nil")
	}

	if AlgChoice == Ed25519 {
		return &PubKey{X: new(big.Int).SetBytes(encodeData[1:]), Y: big.NewInt(0)}, nil
	} else {
		switch encodeData[0] {
		case 0x00:
			return &PubKey{nil, nil}, nil

		case 0x02, 0x03: //compressed
			if len(encodeData) != COMPRESSEDLEN {
				return nil, errors.New("encoded compressed public key length error")
			}
			yTilde := int(encodeData[0] & 1)
			pubKey, err := deCompress(yTilde, encodeData[FLAGLEN:FLAGLEN+XORYVALUELEN],
				&algSet.EccParams)
			if nil != err {
				return nil, fmt.Errorf("Invalid point encoding: (%v)", err)
			}
			return pubKey, nil

		case 0x04, 0x06, 0x07: //uncompressed
			if len(encodeData) != NOCOMPRESSEDLEN {
				return nil, errors.New("encoded uncompressed public key length error")
			}
			pubKeyX := new(big.Int).SetBytes(encodeData[FLAGLEN : FLAGLEN+XORYVALUELEN])
			pubKeyY := new(big.Int).SetBytes(encodeData[FLAGLEN+XORYVALUELEN : NOCOMPRESSEDLEN])
			return &PubKey{pubKeyX, pubKeyY}, nil

		default:
			return nil, errors.New("The encodeData format is error")
		}
	}
}

func (e *PubKey) EncodePoint(isCommpressed bool) ([]byte, error) {
	if AlgChoice == Ed25519 {
		encodedData := make([]byte, COMPRESSEDLEN)
		copy(encodedData[1:], e.X.Bytes())
		encodedData[0] = 0x04
		return encodedData, nil
	} else {
		//if X is infinity, then Y cann't be computed, so here used "||"
		if nil == e.X || nil == e.Y {
			infinity := make([]byte, INFINITYLEN)
			return infinity, nil
		}

		var encodedData []byte

		if isCommpressed {
			encodedData = make([]byte, COMPRESSEDLEN)
		} else {
			encodedData = make([]byte, NOCOMPRESSEDLEN)

			yBytes := e.Y.Bytes()
			copy(encodedData[NOCOMPRESSEDLEN-len(yBytes):], yBytes)
		}
		xBytes := e.X.Bytes()
		copy(encodedData[COMPRESSEDLEN-len(xBytes):COMPRESSEDLEN], xBytes)

		if isCommpressed {
			if isEven(e.Y) {
				encodedData[0] = COMPEVENFLAG
			} else {
				encodedData[0] = COMPODDFLAG
			}
		} else {
			encodedData[0] = NOCOMPRESSEDFLAG
		}

		return encodedData, nil
	}
}

func NewPubKeyFromBytes(pubkey []byte) (*PubKey, error) {
	if AlgChoice == Ed25519 {
		ed25519PubkeySize := ed25519.GetPublicKeySize()
		switch len(pubkey) {
		case COMPRESSEDLEN:
			return DecodePoint(pubkey)
		case ed25519PubkeySize:
			publicKeyX := new(big.Int).SetBytes(pubkey)
			return &PubKey{X: publicKeyX, Y: big.NewInt(0)}, nil
		default:
			return nil, errors.New("the size of ed25519 pubkey is incorrect")
		}
	} else {
		switch len(pubkey) {
		case COMPRESSEDLEN:
			fallthrough
		case NOCOMPRESSEDLEN:
			return DecodePoint(pubkey)
		case XORYVALUELEN * 2:
			publicKeyX := new(big.Int).SetBytes(pubkey[:XORYVALUELEN])
			publicKeyY := new(big.Int).SetBytes(pubkey[XORYVALUELEN:])
			return &PubKey{X: publicKeyX, Y: publicKeyY}, nil
		default:
			return nil, errors.New("the size of ecdsa pubkey is incorrect")
		}

	}
}

func NewPubKey(priKey []byte) *PubKey {
	if AlgChoice == Ed25519 {
		X := ed25519.NewKeyFromPrivkey(priKey)
		return &PubKey{X: X, Y: big.NewInt(0)}
	} else {
		privateKey := new(ecdsa.PrivateKey)
		privateKey.PublicKey.Curve = algSet.Curve

		k := new(big.Int)
		k.SetBytes(priKey)
		privateKey.D = k

		privateKey.PublicKey.X, privateKey.PublicKey.Y = algSet.Curve.ScalarBaseMult(k.Bytes())

		return &PubKey{X: privateKey.PublicKey.X, Y: privateKey.PublicKey.Y}
	}
}
