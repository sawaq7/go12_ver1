package cal1000

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


//	     "github.com/gonum/plot/internal/cmpimg"

                                                   )

///
/// 導水勾配線群よりグラフを作り、ストレッジに保存する。
///

func Make_graf_sample( w http.ResponseWriter ,r *http.Request ,p_number int ,ad_eneup []type3.Point ,
                           ad_enedown []type3.Point ,ad_glineup []type3.Point ,ad_glinedown []type3.Point ) (f_name string) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     IN  p_number 　　: 点の数
//     IN  ad_eneup  　 : エネルギー線（up）のスライス   (pointの構造体）
//     IN  ad_enedown   : エネルギー線（down）のスライス (pointの構造体）
//     IN  ad_glineup   : 導水勾配線（up）のスライス     (pointの構造体）
//     IN  ad_glinedown : 導水勾配線（down）のスライス   (pointの構造体）


   fmt.Fprintf( w, "make_graf_sample start \n" )  // デバック

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
 	p.X.Label.Text = "time"
 	p.Y.Label.Text = "guest"

// 	p.X.Min = 0
//    p.X.Max = 200
//    p.Y.Min = 15
//    p.Y.Max = 20

    p.BackgroundColor = color.RGBA{R: 102,G: 204, B: 255}

    name1 := "sawamoto"

//    p.NominalX( name1, "Two", "Three", "Four", "Five") // 各値のラベル(X軸)

    p.NominalY( name1, "two", "Three", "Four", "Five") // 各値のラベル(X軸)

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

	wide := vg.Points(10) // 棒グラフの幅

  	// groupの各値について棒グラフを生成
  	group := plotter.Values{20, 21, 22, 23, 24} // 描画対象

  	bars, _ := plotter.NewBarChart(group, wide)

//  	bars, _ := plotter.NewBarChart_test( w ,group, wide)

  	bars.LineStyle.Width = vg.Length(0) // 棒グラフの枠線の太さ

  	bars.Color = plotutil.Color(3)      // 棒グラフの色。0から6まででplotutilにSoftColorsとして定義されている。

//  	bars.Offset = wide * 2                    // 棒グラフを表示する位置のオフセット(X方向)。複数のグループを並べたいときは-wなどで位置を調整する。

    bars.Offset = 0
//    bars.XMin = 2

// 	bars.Horizontal = false             // trueにすると横向きの棒グラフになる。

 	bars.Horizontal = true

//	p.Add(bars)

	p.Add_test( w ,bars)

// 	if err := plotutil.AddLinePoints(p, "エネルギー線(up)", ad_eneup_xys); err != nil {
// 	   http.Error(w, err.Error(), http.StatusInternalServerError)
//	   return " "
// 	}

//    if err := plotutil.AddLinePoints(p, "エネルギー線(down)", ad_enedown_xys ); err != nil {
// 	   http.Error(w, err.Error(), http.StatusInternalServerError)
//	   return " "
// 	}

// 	if err := plotutil.AddLinePoints(p, "動水勾配線(up)", ad_glineup_xys ); err != nil {
// 	   http.Error(w, err.Error(), http.StatusInternalServerError)
//	   return " "
// 	}

// 	if err := plotutil.AddLinePoints(p, "動水勾配線(down)", ad_glinedown_xys ); err != nil {
// 	   http.Error(w, err.Error(), http.StatusInternalServerError)
//	   return " "
// 	}

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

    f_name = "bar_graph_" + unique_no + ".png"
    fmt.Fprintf( w, "deliver_showall1 : f_name %v\n", f_name )  // デバック

    if err := p.Save_Storage(w ,r ,5*vg.Inch, 5*vg.Inch, bucket , f_name ); err != nil {  // 新ファイルを保存

       http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "

 	}
    return f_name
}
