// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonum

// This list is duplicated in netlib/lapack/netlib. Keep in sync.
const (
	absIncNotOne  = "lapack: increment not one or negative one"
	badAlpha      = "lapack: bad alpha length"
	badApplyOrtho = "lapack: bad ApplyOrtho"
	badAuxv       = "lapack: auxv has insufficient length"
	badBeta       = "lapack: bad beta length"
	badD          = "lapack: d has insufficient length"
	badDiag       = "lapack: bad diag"
	badDims       = "lapack: bad input dimensions"
	badDirect     = "lapack: bad direct"
	badE          = "lapack: e has insufficient length"
	badEVComp     = "lapack: bad EVComp"
	badEVHowMany  = "lapack: bad EVHowMany"
	badEVJob      = "lapack: bad EVJob"
	badEVSide     = "lapack: bad EVSide"
	badGenOrtho   = "lapack: bad GenOrtho"
	badGSVDJob    = "lapack: bad GSVDJob"
	badIlo        = "lapack: ilo out of range"
	badIhi        = "lapack: ihi out of range"
	badIpiv       = "lapack: bad permutation length"
	badBalanceJob = "lapack: bad BalanceJob"
	badJ1         = "lapack: j1 out of range"
	badK1         = "lapack: k1 out of range"
	badK2         = "lapack: k2 out of range"
	badKperm      = "lapack: incorrect permutation length"
	badLdA        = "lapack: bad leading dimension of A"
	badLdB        = "lapack: bad leading dimension of B"
	badLdC        = "lapack: bad leading dimension of C"
	badLdF        = "lapack: bad leading dimension of F"
	badLdH        = "lapack: bad leading dimension of H"
	badLdQ        = "lapack: bad leading dimension of Q"
	badLdT        = "lapack: bad leading dimension of T"
	badLdU        = "lapack: bad leading dimension of U"
	badLdV        = "lapack: bad leading dimension of V"
	badLdVL       = "lapack: bad leading dimension of VL"
	badLdVR       = "lapack: bad leading dimension of VR"
	badLdVT       = "lapack: bad leading dimension of VT"
	badLdX        = "lapack: bad leading dimension of X"
	badLdY        = "lapack: bad leading dimension of Y"
	badLdZ        = "lapack: bad leading dimension of Z"
	badNb         = "lapack: nb out of range"
	badNorm       = "lapack: bad norm"
	badPivot      = "lapack: bad pivot"
	badS          = "lapack: s has insufficient length"
	badSchurComp  = "lapack: bad SchurComp"
	badSchurJob   = "lapack: bad SchurJob"
	badShifts     = "lapack: bad shifts"
	badSide       = "lapack: bad side"
	badSlice      = "lapack: bad input slice length"
	badSort       = "lapack: bad Sort"
	badStore      = "lapack: bad store"
	badTau        = "lapack: tau has insufficient length"
	badTauQ       = "lapack: tauQ has insufficient length"
	badTauP       = "lapack: tauP has insufficient length"
	badTrans      = "lapack: bad trans"
	badVn1        = "lapack: vn1 has insufficient length"
	badVn2        = "lapack: vn2 has insufficient length"
	badUplo       = "lapack: illegal triangle"
	badWork       = "lapack: insufficient working memory"
	badZ          = "lapack: insufficient z length"
	badIncX       = "lapack: incX <= 0"
	badIncY       = "lapack: incY <= 0"
	kGTM          = "lapack: k > m"
	kGTN          = "lapack: k > n"
	kLT0          = "lapack: k < 0"
	mLT0          = "lapack: m < 0"
	mLTN          = "lapack: m < n"
	nanScale      = "lapack: NaN scale factor"
	negDimension  = "lapack: negative matrix dimension"
	negZ          = "lapack: negative z value"
	nLT0          = "lapack: n < 0"
	nLT1          = "lapack: n < 1"
	nbLT0         = "lapack: nb < 0"
	nLTM          = "lapack: n < m"
	nrhsLT0       = "lapack: nrhs < 0"
	ncvtLT0       = "lapack: ncvt < 0"
	nruLT0        = "lapack: nru < 0"
	nccLT0        = "lapack: ncc < 0"
	pLT0          = "lapack: p < 0"
	offsetLT0     = "lapack: offset < 0"
	offsetGTM     = "lapack: offset > m"
	shortA        = "lapack: insufficient length of a"
	shortB        = "lapack: insufficient length of b"
	shortC        = "lapack: insufficient length of c"
	shortF        = "lapack: insufficient length of f"
	shortH        = "lapack: insufficient length of h"
	shortIsgn     = "lapack: insufficient length of isgn"
	shortQ        = "lapack: insufficient length of q"
	shortScale    = "lapack: insufficient length of scale"
	shortT        = "lapack: insufficient length of t"
	shortU        = "lapack: insufficient length of u"
	shortV        = "lapack: insufficient length of v"
	shortVL       = "lapack: insufficient length of vl"
	shortVR       = "lapack: insufficient length of vr"
	shortVT       = "lapack: insufficient length of vt"
	shortWR       = "lapack: insufficient length of wr"
	shortWI       = "lapack: insufficient length of wi"
	shortX        = "lapack: insufficient length of x"
	shortY        = "lapack: insufficient length of y"
	shortZ        = "lapack: insufficient length of z"
	shortWork     = "lapack: working array shorter than declared"
	zeroDiv       = "lapack: zero divisor"
)
