package cal

import (

	    "net/http"
	    "fmt"
//	    "github.com/sawaq7/go12_ver1/client/sgh/process"
        "github.com/sawaq7/go12_ver1/basic/type3"
        "math/rand"
        "image/color"
//        "storage2"
        "time"

 	     "github.com/gonum/plot"
 	     "github.com/gonum/plot/plotter"
 	     "github.com/gonum/plot/plotutil"
 	     "github.com/gonum/plot/vg"
                                                   )

///
/// 導水勾配線群よりグラフを作り、ストレチE��に保存する、E///

func Pipe_line1_make_graf( w http.ResponseWriter ,r *http.Request ,p_number int ,ad_eneup []type3.Point ,
                           ad_enedown []type3.Point ,ad_glineup []type3.Point ,ad_glinedown []type3.Point ) (f_name string) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     IN  p_number 　　: 点の数
//     IN  ad_eneup  　 : エネルギー線！Ep�E��Eスライス   (pointの構造体！E//     IN  ad_enedown   : エネルギー線！Eown�E��Eスライス (pointの構造体！E//     IN  ad_glineup   : 導水勾配線！Ep�E��Eスライス     (pointの構造体！E//     IN  ad_glinedown : 導水勾配線！Eown�E��Eスライス   (pointの構造体！E

//   fmt.Fprintf( w, "pipe_line1_make_graf start \n" )  // チE��チE��

   rand.Seed(int64(0))

///
/// グラフ�E枠を作�E　
///

 	p, err := plot.New()
 	if err != nil {
 	    http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "

 	}

 	p.Title.Text = "water-slope-line"
 	p.X.Label.Text = "x"
 	p.Y.Label.Text = "y"

 	p.X.Min = 0
    p.X.Max = 200
    p.Y.Min = 15
    p.Y.Max = 20

    p.BackgroundColor = color.RGBA{R: 102,G: 204, B: 255}

///
/// 吁E��ラインの　点のワークエリアをＧ�E��E�　
///

    ad_eneup_xys     := make(plotter.XYs, p_number)
    ad_enedown_xys   := make(plotter.XYs, p_number)
    ad_glineup_xys   := make(plotter.XYs, p_number)
    ad_glinedown_xys := make(plotter.XYs, p_number)

///
/// 吁E��ラインの点のチE�EタをSET　
///

 	for i := 0; i < p_number; i++ {

 		ad_eneup_xys[i].X = ad_eneup[i].X
 		ad_eneup_xys[i].Y = ad_eneup[i].Y

 		ad_enedown_xys[i].X = ad_enedown[i].X
 		ad_enedown_xys[i].Y = ad_enedown[i].Y

 		ad_glineup_xys[i].X = ad_glineup[i].X
 		ad_glineup_xys[i].Y = ad_glineup[i].Y

 		ad_glinedown_xys[i].X = ad_glinedown[i].X
 		ad_glinedown_xys[i].Y = ad_glinedown[i].Y

 	}

///
/// 吁E��ラインのグラフ�EチE�EタをSET　
///

 	if err := plotutil.AddLinePoints(p, "エネルギー緁Eup)", ad_eneup_xys); err != nil {
 	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "
 	}

    if err := plotutil.AddLinePoints(p, "エネルギー緁Edown)", ad_enedown_xys ); err != nil {
 	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "
 	}

 	if err := plotutil.AddLinePoints(p, "動水勾配緁Eup)", ad_glineup_xys ); err != nil {
 	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "
 	}

 	if err := plotutil.AddLinePoints(p, "動水勾配緁Edown)", ad_glinedown_xys ); err != nil {
 	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "
 	}

///
///   グラチEファイル(画像ファイル�E�をストレチE��に保孁E///

 	bucket := "sample-7777"     // バケチE��名セチE��

///
/// ファイル名を作�E
///

 	date_w := time.Now()        // 日付をセチE��
    unique_no := fmt.Sprintf("%04d%02d%02d%02d%02d%02d",
		date_w.Year(), date_w.Month(),date_w.Day(), date_w.Hour(), date_w.Minute(), date_w.Second())

    f_name = "water_slope_" + unique_no + ".png"

//    fmt.Fprintf( w, "deliver_showall1 : f_name %v\n", f_name )  // チE��チE��

//    storage2.File_Delete ( w , r  ,bucket ,f_name  ) // 旧ファイルを削除

    if err := p.Save_Storage(w ,r ,5*vg.Inch, 5*vg.Inch, bucket , f_name ); err != nil {  // 新ファイルを保孁E
       http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "

 	}
    return f_name
}
