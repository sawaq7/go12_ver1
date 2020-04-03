package hydrostatic_pressure1_excute

import (
	     "net/http"
	     "html/template"
	     "github.com/sawaq7/go12_ver1/client/tokura/suiri"
	     "github.com/sawaq7/go12_ver1/basic/maths/sum"
	     "strconv"
	     "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"

//        "github.com/sawaq7/go12_ver1/client/tokura/html4"
         "github.com/sawaq7/go12_ver1/general/html5"
	                                  )

func Hydrostatic_pressure1_excute(w http.ResponseWriter, r *http.Request) {

   var seisui type4.Seisui

// 入力データをGET

	seisui.Omega = r.FormValue("omega")
	seisui.D1    = r.FormValue("d1")
	seisui.D2    = r.FormValue("d2")
	seisui.P     = r.FormValue("p")
	seisui.H     = r.FormValue("h")

// string を　float に変換

	r_omega,_ :=strconv.ParseFloat(seisui.Omega,64)
	r_d1,_ :=strconv.ParseFloat(seisui.D1,64)
	r_d2,_ :=strconv.ParseFloat(seisui.D2,64)
	r_p,_ :=strconv.ParseFloat(seisui.P ,64)
	r_h,_ :=strconv.ParseFloat(seisui.H ,64)

	a1 := sum.Circle_Area(r_d1/2)
    a2 := sum.Circle_Area(r_d2/2)

    seisui.P2 =  suiri.Seisui1( a1 ,a2  ,r_p  ,r_omega  ,r_h  )

// チE��プレート�EヘッダーをGET

    monitor := template.Must(template.New("html").Parse(html5.Storage_bucket_list))
//    monitor := template.Must( template.New("html").Parse( html4.Hydrostatic_pressure2_show) )

// 吁E��入力データを表示

//    err := monitor.Execute(w, seisui)
    var str_dmy string

	err := monitor.Execute( w, str_dmy )
    if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
