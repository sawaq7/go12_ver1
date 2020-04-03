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
/// 蝨ｰ蛹ｺ諠・ｱ繧偵ご繝・ヨ縺吶ｋ ///
///                          ///

func D_district2( w http.ResponseWriter, r *http.Request )  ([]type2.D_District_View ) {

//     IN    w      縲縲縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ

//     OUT d_district_view  : 讒矩菴薙窶晏慍蛹ｺ諠・ｱ窶昴・繧ｹ繝ｩ繧､繧ｹ

//    fmt.Fprintf( w, "trans.d_district2 start \n" )  // 繝・ヰ繝・け

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "trans.d_district2 :  projectID unset \n"  )  // 繝・ヰ繝・け

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
//		fmt.Fprintf( w, "d_district2 err \n" ,err)  // 繝・ヰ繝・け
		return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, d_districtw := range d_district {

///  讖溯・縺ｫ繧医ｊ繝√ぉ繝・け鬆・岼繧偵そ繝・ヨ

        d_area_slice :=   D_area_district ( w ,r ,d_districtw.District_No )

        d_district_view = append(d_district_view, type2.D_District_View { keys_wk[pos]            ,
                                                                          d_districtw.District_No   ,
                                                                          d_districtw.District_Name ,
                                                                          d_area_slice                })
	}

    return	d_district_view
}

