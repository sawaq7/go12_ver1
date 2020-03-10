package type4

///
///    水路ファイル１(ストレッジ・フリー）
///

type  Struct_Colle    struct  {

	      Water2_Slice   []Water2
	   Water_Line_Slice  []Water_Line

}

///
///    水路ファイル１(ストレッジ・フリー）
///

type  Water    struct           {
	      No             string   // データID
	      Name           string   // 水路名
	      High           string   // 水路高
	    Roughness_factor string   // 粗粒係数
}

///
///    水路ファイル２
///

type  Water2    struct           {
          Id             int64      //　データid  * ストレッジの場合はレコードNO
	      Name           string     // 水路名
	      High           float64    // 水路高
	    Roughness_Factor float64    // 粗粒係数
}
///
///    水路ファイル２(データストア・テンポラリー）
///

type  Water2_Temp    struct           {       //  データは1レコード
          Id             int64      //　データid
	      Name           string     // 水路名
	      High           float64    // 水路高
	    Roughness_Factor float64    // 粗粒係数
}
///
///    水路ラインファイル(データストア）
///

type  Water_Line    struct           {
          Id              int64    // データid　　* ストレッジの場合は、水路単位のレコードNO
	      Name            string   // 水路名
	      Section         string   // 区間名
	      Friction_Factor float64  // 摩擦係数
	      Velocity        float64  // 速度
	      Pipe_Diameter   float64  // 管径
	      Pipe_Length     float64  // 管長
}

///
///    導水勾配線ファイル(グラフ）
///
///  *** 画像ファイル表示用　struct "Image_Show" と　同フォーマット

type  Water_Slope    struct           {

          Id              int64    // データid
	      File_Name       string   // ファイル名
	      Url             string   // url

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
