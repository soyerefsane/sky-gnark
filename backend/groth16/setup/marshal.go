package setup

import (
	"io"

	"github.com/consensys/gnark-crypto/ecc/bn254"
)

// WriteTo implements io.WriterTo
func (phase1 *Phase1) WriteTo(writer io.Writer) (int64, error) {
	n, err := phase1.writeTo(writer)
	if err != nil {
		return n, err
	}
	nBytes, err := writer.Write(phase1.Hash)
	return int64(nBytes) + n, err
}

func (phase1 *Phase1) writeTo(writer io.Writer) (int64, error) {
	toEncode := []interface{}{
		&phase1.PublicKeys.Tau.SG,
		&phase1.PublicKeys.Tau.SXG,
		&phase1.PublicKeys.Tau.XR,
		&phase1.PublicKeys.Alpha.SG,
		&phase1.PublicKeys.Alpha.SXG,
		&phase1.PublicKeys.Alpha.XR,
		&phase1.PublicKeys.Beta.SG,
		&phase1.PublicKeys.Beta.SXG,
		&phase1.PublicKeys.Beta.XR,
		phase1.Parameters.G1.Tau,
		phase1.Parameters.G1.AlphaTau,
		phase1.Parameters.G1.BetaTau,
		phase1.Parameters.G2.Tau,
		&phase1.Parameters.G2.Beta,
	}

	enc := bn254.NewEncoder(writer)
	for _, v := range toEncode {
		if err := enc.Encode(v); err != nil {
			return enc.BytesWritten(), err
		}
	}
	return enc.BytesWritten(), nil
}

// ReadFrom implements io.ReaderFrom
func (phase1 *Phase1) ReadFrom(reader io.Reader) (int64, error) {
	toEncode := []interface{}{
		&phase1.PublicKeys.Tau.SG,
		&phase1.PublicKeys.Tau.SXG,
		&phase1.PublicKeys.Tau.XR,
		&phase1.PublicKeys.Alpha.SG,
		&phase1.PublicKeys.Alpha.SXG,
		&phase1.PublicKeys.Alpha.XR,
		&phase1.PublicKeys.Beta.SG,
		&phase1.PublicKeys.Beta.SXG,
		&phase1.PublicKeys.Beta.XR,
		&phase1.Parameters.G1.Tau,
		&phase1.Parameters.G1.AlphaTau,
		&phase1.Parameters.G1.BetaTau,
		&phase1.Parameters.G2.Tau,
		&phase1.Parameters.G2.Beta,
	}

	dec := bn254.NewDecoder(reader)
	for _, v := range toEncode {
		if err := dec.Decode(v); err != nil {
			return dec.BytesRead(), err
		}
	}
	phase1.Hash = make([]byte, 32)
	nBytes, err := reader.Read(phase1.Hash)
	return dec.BytesRead() + int64(nBytes), err
}

// WriteTo implements io.WriterTo
func (phase2 *Phase2) WriteTo(writer io.Writer) (int64, error) {
	n, err := phase2.writeTo(writer)
	if err != nil {
		return n, err
	}
	nBytes, err := writer.Write(phase2.Hash)
	return int64(nBytes) + n, err
}

func (c *Phase2) writeTo(writer io.Writer) (int64, error) {
	enc := bn254.NewEncoder(writer)
	toEncode := []interface{}{
		&c.PublicKey.SG,
		&c.PublicKey.SXG,
		&c.PublicKey.XR,
		&c.Parameters.G1.Delta,
		c.Parameters.G1.L,
		c.Parameters.G1.Z,
		&c.Parameters.G2.Delta,
	}

	for _, v := range toEncode {
		if err := enc.Encode(v); err != nil {
			return enc.BytesWritten(), err
		}
	}

	return enc.BytesWritten(), nil
}

// ReadFrom implements io.ReaderFrom
func (c *Phase2) ReadFrom(reader io.Reader) (int64, error) {
	dec := bn254.NewDecoder(reader)
	toEncode := []interface{}{
		&c.PublicKey.SG,
		&c.PublicKey.SXG,
		&c.PublicKey.XR,
		&c.Parameters.G1.Delta,
		&c.Parameters.G1.L,
		&c.Parameters.G1.Z,
		&c.Parameters.G2.Delta,
	}

	for _, v := range toEncode {
		if err := dec.Decode(v); err != nil {
			return dec.BytesRead(), err
		}
	}

	c.Hash = make([]byte, 32)
	n, err := reader.Read(c.Hash)
	return int64(n) + dec.BytesRead(), err

}

// WriteTo implements io.WriterTo
func (c *Phase2Evaluations) WriteTo(writer io.Writer) (int64, error) {
	enc := bn254.NewEncoder(writer)
	toEncode := []interface{}{
		c.G1.A,
		c.G1.B,
		c.G2.B,
	}

	for _, v := range toEncode {
		if err := enc.Encode(v); err != nil {
			return enc.BytesWritten(), err
		}
	}

	return enc.BytesWritten(), nil
}

// ReadFrom implements io.ReaderFrom
func (c *Phase2Evaluations) ReadFrom(reader io.Reader) (int64, error) {
	dec := bn254.NewDecoder(reader)
	toEncode := []interface{}{
		&c.G1.A,
		&c.G1.B,
		&c.G2.B,
	}

	for _, v := range toEncode {
		if err := dec.Decode(v); err != nil {
			return dec.BytesRead(), err
		}
	}

	return dec.BytesRead(), nil
}
