package trans

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"                                      )

///                           ///
/// å°åŗęE ±ćć²ćEćć ///
///                          ///

func D_district2( w http.ResponseWriter, r *http.Request )  ([]type2.D_District_View ) {

//     IN    w      ćććć: ć¬ć¹ćć³ć¹ć©ć¤ćæć¼
//     IN    r      ćććć: ćŖćÆćØć¹ććć©ć”ć¼ćæ

//     OUT d_district_view  : ę§é ä½ćāå°åŗęE ±āćEć¹ć©ć¤ć¹

//    fmt.Fprintf( w, "trans.d_district2 start \n" )  // ćEćEÆ

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "trans.d_district2 :  projectID unset \n"  )  // ćEćEÆ

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
	d_district_view := make([]type2.D_District_View, 0)

    keys, err := client.GetAll(ctx, query , &d_district)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district2 err \n" ,err)  // ćEćEÆ
		return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, d_districtw := range d_district {

///  ę©čEć«ćććć§ćEÆé E®ćć»ćE

        d_area_slice :=   D_area_district ( w ,r ,d_districtw.District_No )

        d_district_view = append(d_district_view, type2.D_District_View { keys_wk[pos]            ,
                                                                          d_districtw.District_No   ,
                                                                          d_districtw.District_Name ,
                                                                          d_area_slice                })
	}

    return	d_district_view
}

