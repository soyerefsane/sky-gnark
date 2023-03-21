package fields_bn254

import (
	"github.com/consensys/gnark-crypto/ecc/bn254"
)

type E12 struct {
	C0, C1 E6
}

type Ext12 struct {
	*Ext6
}

func NewExt12(baseField *curveF) *Ext12 {
	return &Ext12{Ext6: NewExt6(baseField)}
}

func (e Ext12) Add(x, y *E12) *E12 {
	z0 := e.Ext6.Add(&x.C0, &y.C0)
	z1 := e.Ext6.Add(&x.C1, &y.C1)
	return &E12{
		C0: *z0,
		C1: *z1,
	}
}

func (e Ext12) Sub(x, y *E12) *E12 {
	z0 := e.Ext6.Sub(&x.C0, &y.C0)
	z1 := e.Ext6.Sub(&x.C1, &y.C1)
	return &E12{
		C0: *z0,
		C1: *z1,
	}
}

func (e Ext12) Conjugate(x *E12) *E12 {
	z1 := e.Ext6.Neg(&x.C1)
	return &E12{
		C0: x.C0,
		C1: *z1,
	}
}

func (e Ext12) Mul(x, y *E12) *E12 {
	a := e.Ext6.Add(&x.C0, &x.C1)
	b := e.Ext6.Add(&y.C0, &y.C1)
	a = e.Ext6.Mul(a, b)
	b = e.Ext6.Mul(&x.C0, &y.C0)
	c := e.Ext6.Mul(&x.C1, &y.C1)
	z1 := e.Ext6.Sub(a, b)
	z1 = e.Ext6.Sub(z1, c)
	z0 := e.Ext6.MulByNonResidue(c)
	z0 = e.Ext6.Add(z0, b)
	return &E12{
		C0: *z0,
		C1: *z1,
	}
}

func (e Ext12) CyclotomicSquare(x *E12) *E12 {
	t0 := e.Ext2.Square(&x.C1.B1)
	t1 := e.Ext2.Square(&x.C0.B0)
	t6 := e.Ext2.Add(&x.C1.B1, &x.C0.B0)
	t6 = e.Ext2.Square(t6)
	t6 = e.Ext2.Sub(t6, t0)
	t6 = e.Ext2.Sub(t6, t1)
	t2 := e.Ext2.Square(&x.C0.B2)
	t3 := e.Ext2.Square(&x.C1.B0)
	t7 := e.Ext2.Add(&x.C0.B2, &x.C1.B0)
	t7 = e.Ext2.Square(t7)
	t7 = e.Ext2.Sub(t7, t2)
	t7 = e.Ext2.Sub(t7, t3)
	t4 := e.Ext2.Square(&x.C1.B2)
	t5 := e.Ext2.Square(&x.C0.B1)
	t8 := e.Ext2.Add(&x.C1.B2, &x.C0.B1)
	t8 = e.Ext2.Square(t8)
	t8 = e.Ext2.Sub(t8, t4)
	t8 = e.Ext2.Sub(t8, t5)
	t8 = e.Ext2.MulByNonResidue(t8)
	t0 = e.Ext2.MulByNonResidue(t0)
	t0 = e.Ext2.Add(t0, t1)
	t2 = e.Ext2.MulByNonResidue(t2)
	t2 = e.Ext2.Add(t2, t3)
	t4 = e.Ext2.MulByNonResidue(t4)
	t4 = e.Ext2.Add(t4, t5)
	z00 := e.Ext2.Sub(t0, &x.C0.B0)
	z00 = e.Ext2.Double(z00)
	z00 = e.Ext2.Add(z00, t0)
	z01 := e.Ext2.Sub(t2, &x.C0.B1)
	z01 = e.Ext2.Double(z01)
	z01 = e.Ext2.Add(z01, t2)
	z02 := e.Ext2.Sub(t4, &x.C0.B2)
	z02 = e.Ext2.Double(z02)
	z02 = e.Ext2.Add(z02, t4)
	z10 := e.Ext2.Add(t8, &x.C1.B0)
	z10 = e.Ext2.Double(z10)
	z10 = e.Ext2.Add(z10, t8)
	z11 := e.Ext2.Add(t6, &x.C1.B1)
	z11 = e.Ext2.Double(z11)
	z11 = e.Ext2.Add(z11, t6)
	z12 := e.Ext2.Add(t7, &x.C1.B2)
	z12 = e.Ext2.Double(z12)
	z12 = e.Ext2.Add(z12, t7)
	return &E12{
		C0: E6{
			B0: *z00,
			B1: *z01,
			B2: *z02,
		},
		C1: E6{
			B0: *z10,
			B1: *z11,
			B2: *z12,
		},
	}
}

func (e Ext12) Frobenius(x *E12) *E12 {
	t0 := e.Ext2.Conjugate(&x.C0.B0)
	t1 := e.Ext2.Conjugate(&x.C0.B1)
	t2 := e.Ext2.Conjugate(&x.C0.B2)
	t3 := e.Ext2.Conjugate(&x.C1.B0)
	t4 := e.Ext2.Conjugate(&x.C1.B1)
	t5 := e.Ext2.Conjugate(&x.C1.B2)
	t1 = e.Ext2.MulByNonResidue1Power2(t1)
	t2 = e.Ext2.MulByNonResidue1Power4(t2)
	t3 = e.Ext2.MulByNonResidue1Power1(t3)
	t4 = e.Ext2.MulByNonResidue1Power3(t4)
	t5 = e.Ext2.MulByNonResidue1Power5(t5)
	return &E12{
		C0: E6{
			B0: *t0,
			B1: *t1,
			B2: *t2,
		},
		C1: E6{
			B0: *t3,
			B1: *t4,
			B2: *t5,
		},
	}
}

func (e Ext12) FrobeniusSquare(x *E12) *E12 {
	z00 := &x.C0.B0
	z01 := e.Ext2.MulByNonResidue2Power2(&x.C0.B1)
	z02 := e.Ext2.MulByNonResidue2Power4(&x.C0.B2)
	z10 := e.Ext2.MulByNonResidue2Power1(&x.C1.B0)
	z11 := e.Ext2.MulByNonResidue2Power3(&x.C1.B1)
	z12 := e.Ext2.MulByNonResidue2Power5(&x.C1.B2)
	return &E12{
		C0: E6{B0: *z00, B1: *z01, B2: *z02},
		C1: E6{B0: *z10, B1: *z11, B2: *z12},
	}
}

func (e Ext12) FrobeniusCube(x *E12) *E12 {
	t0 := e.Ext2.Conjugate(&x.C0.B0)
	t1 := e.Ext2.Conjugate(&x.C0.B1)
	t2 := e.Ext2.Conjugate(&x.C0.B2)
	t3 := e.Ext2.Conjugate(&x.C1.B0)
	t4 := e.Ext2.Conjugate(&x.C1.B1)
	t5 := e.Ext2.Conjugate(&x.C1.B2)
	t1 = e.Ext2.MulByNonResidue3Power2(t1)
	t2 = e.Ext2.MulByNonResidue3Power4(t2)
	t3 = e.Ext2.MulByNonResidue3Power1(t3)
	t4 = e.Ext2.MulByNonResidue3Power3(t4)
	t5 = e.Ext2.MulByNonResidue3Power5(t5)
	return &E12{
		C0: E6{
			B0: *t0,
			B1: *t1,
			B2: *t2,
		},
		C1: E6{
			B0: *t3,
			B1: *t4,
			B2: *t5,
		},
	}
}

func (e Ext12) Zero() *E12 {
	zero := e.fp.Zero()
	return &E12{
		C0: E6{
			B0: E2{A0: *zero, A1: *zero},
			B1: E2{A0: *zero, A1: *zero},
			B2: E2{A0: *zero, A1: *zero},
		},
		C1: E6{
			B0: E2{A0: *zero, A1: *zero},
			B1: E2{A0: *zero, A1: *zero},
			B2: E2{A0: *zero, A1: *zero},
		},
	}
}

func (e Ext12) One() *E12 {
	z000 := e.fp.One()
	zero := e.fp.Zero()
	return &E12{
		C0: E6{
			B0: E2{A0: *z000, A1: *zero},
			B1: E2{A0: *zero, A1: *zero},
			B2: E2{A0: *zero, A1: *zero},
		},
		C1: E6{
			B0: E2{A0: *zero, A1: *zero},
			B1: E2{A0: *zero, A1: *zero},
			B2: E2{A0: *zero, A1: *zero},
		},
	}
}

func (e Ext12) Square(x *E12) *E12 {
	c0 := e.Ext6.Sub(&x.C0, &x.C1)
	c3 := e.Ext6.MulByNonResidue(&x.C1)
	c3 = e.Ext6.Neg(c3)
	c3 = e.Ext6.Add(&x.C0, c3)
	c2 := e.Ext6.Mul(&x.C0, &x.C1)
	c0 = e.Ext6.Mul(c0, c3)
	c0 = e.Ext6.Add(c0, c2)
	z1 := e.Ext6.double(c2)
	c2 = e.Ext6.MulByNonResidue(c2)
	z0 := e.Ext6.Add(c0, c2)
	return &E12{
		C0: *z0,
		C1: *z1,
	}
}

func (e Ext12) AssertIsEqual(x, y *E12) {
	e.Ext6.AssertIsEqual(&x.C0, &y.C0)
	e.Ext6.AssertIsEqual(&x.C1, &y.C1)
}

func FromE12(y *bn254.E12) E12 {
	return E12{
		C0: FromE6(&y.C0),
		C1: FromE6(&y.C1),
	}

}

func (e Ext12) Inverse(x *E12) *E12 {
	res, err := e.fp.NewHint(inverseE12Hint, 12, &x.C0.B0.A0, &x.C0.B0.A1, &x.C0.B1.A0, &x.C0.B1.A1, &x.C0.B2.A0, &x.C0.B2.A1, &x.C1.B0.A0, &x.C1.B0.A1, &x.C1.B1.A0, &x.C1.B1.A1, &x.C1.B2.A0, &x.C1.B2.A1)
	if err != nil {
		// err is non-nil only for invalid number of inputs
		panic(err)
	}

	inv := E12{
		C0: E6{
			B0: E2{A0: *res[0], A1: *res[1]},
			B1: E2{A0: *res[2], A1: *res[3]},
			B2: E2{A0: *res[4], A1: *res[5]},
		},
		C1: E6{
			B0: E2{A0: *res[6], A1: *res[7]},
			B1: E2{A0: *res[8], A1: *res[9]},
			B2: E2{A0: *res[10], A1: *res[11]},
		},
	}

	one := e.One()

	// 1 == inv * x
	_one := e.Mul(&inv, x)
	e.AssertIsEqual(one, _one)

	return &inv

}

// DivUnchecked e2 elmts
func (e Ext12) DivUnchecked(x, y *E12) *E12 {
	res, err := e.fp.NewHint(divE12Hint, 12, &x.C0.B0.A0, &x.C0.B0.A1, &x.C0.B1.A0, &x.C0.B1.A1, &x.C0.B2.A0, &x.C0.B2.A1, &x.C1.B0.A0, &x.C1.B0.A1, &x.C1.B1.A0, &x.C1.B1.A1, &x.C1.B2.A0, &x.C1.B2.A1, &y.C0.B0.A0, &y.C0.B0.A1, &y.C0.B1.A0, &y.C0.B1.A1, &y.C0.B2.A0, &y.C0.B2.A1, &y.C1.B0.A0, &y.C1.B0.A1, &y.C1.B1.A0, &y.C1.B1.A1, &y.C1.B2.A0, &y.C1.B2.A1)

	if err != nil {
		// err is non-nil only for invalid number of inputs
		panic(err)
	}

	div := E12{
		C0: E6{
			B0: E2{A0: *res[0], A1: *res[1]},
			B1: E2{A0: *res[2], A1: *res[3]},
			B2: E2{A0: *res[4], A1: *res[5]},
		},
		C1: E6{
			B0: E2{A0: *res[6], A1: *res[7]},
			B1: E2{A0: *res[8], A1: *res[9]},
			B2: E2{A0: *res[10], A1: *res[11]},
		},
	}

	// x == div * y
	_x := e.Mul(&div, y)
	e.AssertIsEqual(x, _x)

	return &div
}
