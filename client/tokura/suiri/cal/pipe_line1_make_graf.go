package cal

import (

	    "net/http"
	    "fmt"
//	    "client/sgh/process"
        "basic/type3"
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
/// 導水勾配線群よりグラフを作り、ストレッジに保存する。
///

func Pipe_line1_make_graf( w http.ResponseWriter ,r *http.Request ,p_number int ,ad_eneup []type3.Point ,
                           ad_enedown []type3.Point ,ad_glineup []type3.Point ,ad_glinedown []type3.Point ) (f_name string) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     IN  p_number 　　: 点の数
//     IN  ad_eneup  　 : エネルギー線（up）のスライス   (pointの構造体）
//     IN  ad_enedown   : エネルギー線（down）のスライス (pointの構造体）
//     IN  ad_glineup   : 導水勾配線（up）のスライス     (pointの構造体）
//     IN  ad_glinedown : 導水勾配線（down）のスライス   (pointの構造体）


//   fmt.Fprintf( w, "pipe_line1_make_graf start \n" )  // デバック

   rand.Seed(int64(0))

///
/// グラフの枠を作成　
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
/// 各種ラインの　点のワークエリアをＧＥＴ　
///

    ad_eneup_xys     := make(plotter.XYs, p_number)
    ad_enedown_xys   := make(plotter.XYs, p_number)
    ad_glineup_xys   := make(plotter.XYs, p_number)
    ad_glinedown_xys := make(plotter.XYs, p_number)

///
/// 各種ラインの点のデータをSET　
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
/// 各種ラインのグラフのデータをSET　
///

 	if err := plotutil.AddLinePoints(p, "エネルギー線(up)", ad_eneup_xys); err != nil {
 	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "
 	}

    if err := plotutil.AddLinePoints(p, "エネルギー線(down)", ad_enedown_xys ); err != nil {
 	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "
 	}

 	if err := plotutil.AddLinePoints(p, "動水勾配線(up)", ad_glineup_xys ); err != nil {
 	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "
 	}

 	if err := plotutil.AddLinePoints(p, "動水勾配線(down)", ad_glinedown_xys ); err != nil {
 	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "
 	}

///
///   グラフ-ファイル(画像ファイル）をストレッジに保存
///

 	bucket := "sample-7777"     // バケット名セット

///
/// ファイル名を作成
///

 	date_w := time.Now()        // 日付をセット
    unique_no := fmt.Sprintf("%04d%02d%02d%02d%02d%02d",
		date_w.Year(), date_w.Month(),date_w.Day(), date_w.Hour(), date_w.Minute(), date_w.Second())

    f_name = "water_slope_" + unique_no + ".png"

//    fmt.Fprintf( w, "deliver_showall1 : f_name %v\n", f_name )  // デバック

//    storage2.File_Delete ( w , r  ,bucket ,f_name  ) // 旧ファイルを削除

    if err := p.Save_Storage(w ,r ,5*vg.Inch, 5*vg.Inch, bucket , f_name ); err != nil {  // 新ファイルを保存

       http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "

 	}
    return f_name
}
