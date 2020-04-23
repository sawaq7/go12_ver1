package trans

import (

	    "net/http"
//	    "fmt"
//	    "html/template"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
//	    "time"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///
///     get area inf.
///

func D_area(funct int64 ,some_no int64 ,w http.ResponseWriter, r *http.Request )  ([]type2.D_Area ) {

//     IN  funct  　　　: function
//     　　　　　1:  area no
//     　　　　　2:  car no
//     　　　　　3:  private no
//     IN  some_no      : if flag
//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//     OUT d_area_view  : slice of struct ( D_Area )

//    fmt.Fprintf( w, "trans.d_area start \n" )
//    fmt.Fprintf( w, "trans.d_area funct \n" ,funct )
//    fmt.Fprintf( w, "trans.d_area some_no \n" ,some_no)

    var check_no int64

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

    query := datastore.NewQuery("D_Area").Order("Area_No")

	count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	d_area      := make([]type2.D_Area, 0, count)
	d_area_view := make([]type2.D_Area, 0)

	keys, err := client.GetAll(ctx, query , &d_area)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)
		return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, d_areaw := range d_area {

//	  fmt.Fprintf( w, "trans.d_area d_areaw %v\n" ,d_areaw)

///  branch flag set from function key

	  if funct == 0 {   //  when func. eq. area
	     check_no = d_areaw.District_No

	  }else if funct == 1 {   //  when func. eq. car
	     check_no = 1

	  }else if funct == 2 {   //  when func. eq. private
	     check_no = 2

	  }
      if  some_no == 0 {

         d_area_view = append(d_area_view, type2.D_Area { keys_wk[pos]      ,
                                                          d_areaw.Course_No      ,
                                                          d_areaw.District_No    ,
                                                          d_areaw.District_Name  ,
                                                          d_areaw.Area_No        ,
                                                          d_areaw.Area_Name      ,
                                                          d_areaw.Area_Detail    ,
                                                          d_areaw.Number_Total   ,
                                                          d_areaw.Time_Total     ,
                                                          d_areaw.Productivity   ,
                                                          d_areaw.Car_No           })


      }else if some_no == check_no {

         d_area_view = append(d_area_view, type2.D_Area { keys_wk[pos]      ,
                                                          d_areaw.Course_No      ,
                                                          d_areaw.District_No    ,
                                                          d_areaw.District_Name  ,
                                                          d_areaw.Area_No        ,
                                                          d_areaw.Area_Name      ,
                                                          d_areaw.Area_Detail    ,
                                                          d_areaw.Number_Total   ,
                                                          d_areaw.Time_Total     ,
                                                          d_areaw.Productivity   ,
                                                          d_areaw.Car_No           })

      }
	}

    return	d_area_view
}

