package process1000

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/trans"
	    "github.com/sawaq7/go12_ver1/temp/datastore1000/trans1000"
// 	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "github.com/sawaq7/go12_ver1/temp/type1000"
	    "github.com/sawaq7/go12_ver1/temp/html1000"

//	    "github.com/sawaq7/go12_ver1/client/sgh/type2"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"

                                                )


func D_district_showall1_sample(w http.ResponseWriter, r *http.Request) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//    fmt.Fprintf( w, "d_district_showall1_sample start \n" )

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

    d_district_view := make([]type1000.D_District, 0)

    d_district := trans.D_district ( 0 ,0 ,w ,r )

    for _, d_districtw := range d_district {

      d_area_slice := trans1000.D_area_district ( w ,r ,d_districtw.District_No )

      d_area_small_slice := make([]type1000.D_Area_Small, 1)

//      fmt.Fprintf( w, "d_district_showall1_sample len( d_area_slice ) %v\n" ,len( d_area_slice ))
      if len( d_area_slice ) != 0 {

        d_area_small_slice[0].Area_Name = d_area_slice[0].Area_Name

	  }

      d_area_small_slice[0].Area_Small_Name = "area_small_name_test"

      d_district_view = append(d_district_view, type1000.D_District { d_districtw.Id            ,
                                                                      d_districtw.District_No   ,
                                                                      d_districtw.District_Name ,
                                                                      d_area_slice              ,
                                                                      d_area_small_slice })


      var d_district_work type1000.D_District

      d_district_work.District_No   = d_districtw.District_No
      d_district_work.District_Name = d_districtw.District_Name
      d_district_work.D_Area_Slice  = d_area_slice
      d_district_work.D_Area_Small_Slice = d_area_small_slice

      new_key := datastore.IncompleteKey("D_District_Sample", nil)

      if _, err = client.Put(ctx, new_key, &d_district_work ); err != nil {

		 http.Error(w,err.Error(), http.StatusInternalServerError)
		 return
	  }

	}

    monitor := template.Must(template.New("html").Parse(html1000.D_district_showall1_sample))

	err = monitor.Execute(w, d_district_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

