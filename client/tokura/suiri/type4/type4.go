package type4

///
///    水路ファイル�E�EストレチE��・フリー�E�E///

type  Struct_Colle    struct  {

	      Water2_Slice   []Water2
	   Water_Line_Slice  []Water_Line

}

///
///    水路ファイル�E�EストレチE��・フリー�E�E///

type  Water    struct           {
	      No             string   // チE�EタID
	      Name           string   // 水路吁E	      High           string   // 水路髁E	    Roughness_factor string   // 粗粒係数
}

///
///    水路ファイル�E�E///

type  Water2    struct           {
          Id             int64      //　チE�Eタid  * ストレチE��の場合�EレコードNO
	      Name           string     // 水路吁E	      High           float64    // 水路髁E	    Roughness_Factor float64    // 粗粒係数
}
///
///    水路ファイル�E�EチE�Eタストア・チE��ポラリー�E�E///

type  Water2_Temp    struct           {       //  チE�Eタは1レコーチE          Id             int64      //　チE�Eタid
	      Name           string     // 水路吁E	      High           float64    // 水路髁E	    Roughness_Factor float64    // 粗粒係数
}
///
///    水路ラインファイル(チE�Eタストア�E�E///

type  Water_Line    struct           {
          Id              int64    // チE�Eタid　　* ストレチE��の場合�E、水路単位�EレコードNO
	      Name            string   // 水路吁E	      Section         string   // 区間名
	      Friction_Factor float64  // 摩擦係数
	      Velocity        float64  // 速度
	      Pipe_Diameter   float64  // 管征E	      Pipe_Length     float64  // 管長
}

///
///    導水勾配線ファイル(グラフ！E///
///  *** 画像ファイル表示用　struct "Image_Show" と　同フォーマッチE
type  Water_Slope    struct           {

          Id              int64    // チE�Eタid
	      File_Name       string   // ファイル吁E	      Url             string   // url

}
///
///    hydro-static information
///

type Seisui struct {

      Omega string
      D1    string
      D2    string
      P     string
      H     string
     P2    float64

}
