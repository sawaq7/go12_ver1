package trans

import (

	    "net/http"
//	    "fmt"
//	    "html/template"
//	    "web/htmls/sgh"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"                                      )

///
///     get district inf.
///

func D_district(funct int64 ,some_no int64 ,w http.ResponseWriter, r *http.Request )  ([]type2.D_District ) {

//     IN  funct  　　　　　: function
//     IN  some_no  　　　　: if flag
//     IN    w      　　　　: レスポンスライター
//     IN    r      　　　　: リクエストパラメータ

//     OUT d_district_view  : slice of struct ( D_District )

//    fmt.Fprintf( w, "trans.d_district start \n" )
//    fmt.Fprintf( w, "trans.d_district some_no \n" ,some_no)

//    var check_no int64

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )

      project_name = "sample-7777"

	}
    ctx := context.Background()

	query := datastore.NewQuery("D_District").Order("District_No")

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return  nil
    }

	count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	d_district      := make([]type2.D_District, 0, count)
	d_district_view := make([]type2.D_District, 0)

    keys, err := client.GetAll(ctx, query , &d_district)

    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)

//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)
		return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, d_districtw := range d_district {

         d_district_view = append(d_district_view, type2.D_District { keys_wk[pos]            ,
                                                                    d_districtw.District_No   ,
                                                                    d_districtw.District_Name            })

	}

    return	d_district_view
}

