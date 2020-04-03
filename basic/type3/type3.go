package type3


///
///    poly-line
///

type Poly_Line struct {
        Number  int64
        X[10]  float64
        Y[10]  float64
        Z[10]  float64

}

///
///    point
///

type Point struct {

        X  float64
        Y  float64
        Z  float64

}

///
///    最小二乗法！Eeast squares method�E�E///

type Ls_Method struct {   // 使用するか検討中

        X   float64
        Y   float64
        XX  float64
        YY  float64
        XY  float64

}

