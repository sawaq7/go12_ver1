package graph1

import (

	    "net/http"
	    "fmt"
//	    "github.com/sawaq7/go12_ver1/client/sgh/process"
        "github.com/sawaq7/go12_ver1/basic/type3"
        "math/rand"
        "image/color"
//        "storage2"
        "time"

// 	     "github.com/gonum/plot"
// 	     "github.com/gonum/plot/plotter"
// 	     "github.com/gonum/plot/plotutil"
// 	     "github.com/gonum/plot/vg"

 	     "gonum.org/v1/plot/vg"
	     "gonum.org/v1/plot"
         "gonum.org/v1/plot/plotter"
 	     "gonum.org/v1/plot/plotutil"
                                                   )

///
///       make graf and reserve it in storage from water-slope-lines
///

func Line_graph1( w http.ResponseWriter ,r *http.Request ,p_number int ,ad_eneup []type3.Point ,
                           ad_enedown []type3.Point ,ad_glineup []type3.Point ,ad_glinedown []type3.Point ) (f_name string) {

//     IN     w         : response-writer
//     IN     r         : request-parameter
//     IN  p_number 　　: point number
//     IN  ad_eneup  　 : energy-line-up slice
//     IN  ad_enedown   : energy-line-down slice
//     IN  ad_glineup   : water-slope-line-up slice
//     IN  ad_glinedown : water-slope-line-down slice

//   fmt.Fprintf( w, "pipe_line1_make_graf start \n" )

   rand.Seed(int64(0))

///
///     make graf
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
///   allocate work-area from points　
///

    ad_eneup_xys     := make(plotter.XYs, p_number)
    ad_enedown_xys   := make(plotter.XYs, p_number)
    ad_glineup_xys   := make(plotter.XYs, p_number)
    ad_glinedown_xys := make(plotter.XYs, p_number)

///
///    set points data for line　
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
///    make data for graf　
///

 	if err := plotutil.AddLinePoints(p, "energy-line-up)", ad_eneup_xys); err != nil {
 	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "
 	}

    if err := plotutil.AddLinePoints(p, "energy-line-down)", ad_enedown_xys ); err != nil {
 	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "
 	}

 	if err := plotutil.AddLinePoints(p, "water-slope-line-up)", ad_glineup_xys ); err != nil {
 	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "
 	}

 	if err := plotutil.AddLinePoints(p, "water-slope-line-down)", ad_glinedown_xys ); err != nil {
 	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "
 	}

 	bucket := "sample-7777"

///
///     make file-name
///

 	date_w := time.Now()
    unique_no := fmt.Sprintf("%04d%02d%02d%02d%02d%02d",
		date_w.Year(), date_w.Month(),date_w.Day(), date_w.Hour(), date_w.Minute(), date_w.Second())

    f_name = "water_slope_" + unique_no + ".png"

//    fmt.Fprintf( w, "deliver_showall1 : f_name %v\n", f_name )

///     save graf data in storage

    if err := p.Save_Storage(w ,r ,5*vg.Inch, 5*vg.Inch, bucket , f_name ); err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
	   return " "

 	}
    return f_name
}
