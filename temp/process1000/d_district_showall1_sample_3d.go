package process1000

import (
	    "google.golang.org/appengine"
	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/trans"
	    "github.com/sawaq7/go12_ver1/temp/datastore1000/trans1000"
// 	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "github.com/sawaq7/go12_ver1/temp/type1000"
	    "github.com/sawaq7/go12_ver1/temp/html1000"

//	    "github.com/sawaq7/go12_ver1/client/sgh/type2"

                                                )


func D_district_showall1_sample_3d(w http.ResponseWriter, r *http.Request) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

    fmt.Fprintf( w, "d_district_showall1_sample_3d start \n" )

    d_district_view := make([]type1000.D_District_2, 0)


    c := appengine.NewContext(r)

    d_district := trans.D_district ( 0 ,0 ,w ,r )

    for _, d_districtw := range d_district {

      d_area_slice := trans1000.D_area_district ( w ,r ,d_districtw.District_No )

      d_area_slice_3d := trans1000.D_area_district_3d ( w ,r ,d_districtw.District_No )

      d_area_small_slice := make([]type1000.D_Area_Small, 1)

      fmt.Fprintf( w, "d_district_showall1_sample_3d len( d_area_slice ) %v\n" ,len( d_area_slice ))
      if len( d_area_slice ) != 0 {

        d_area_small_slice[0].Area_Name = d_area_slice[0].Area_Name

	  }

      d_area_small_slice[0].Area_Small_Name = "area_small_name_test"

      d_district_view = append(d_district_view, type1000.D_District_2 { d_districtw.Id            ,
                                                                      d_districtw.District_No   ,
                                                                      d_districtw.District_Name ,
                                                                      d_area_slice_3d               })


///
      var d_district_work type1000.D_District

      d_district_work.District_No   = d_districtw.District_No
      d_district_work.District_Name = d_districtw.District_Name
      d_district_work.D_Area_Slice  = d_area_slice
      d_district_work.D_Area_Small_Slice = d_area_small_slice

	  if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "D_District_Sample", nil), &d_district_work); err != nil {
		 http.Error(w,err.Error(), http.StatusInternalServerError)
		 return
	  }

	}

//

    monitor := template.Must(template.New("html").Parse(html1000.D_district_showall1_sample_3d))

	err := monitor.Execute(w, d_district_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

