package trans3

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/general/type5"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )

///
///    csv諠・ｱ繧偵ご繝・ヨ縺吶ｋ
///

func Csv_inf ( w http.ResponseWriter, r *http.Request )  ( []type5.Csv_Inf ) {

//     IN    w      縲縲縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ

//     OUT csv_inf_view  : csv諠・ｱ窶昴・繧ｹ繝ｩ繧､繧ｹ

//    fmt.Fprintf( w, "trans.csv_inf start \n" )  // 繝・ヰ繝・け

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()

	query := datastore.NewQuery("Csv_Inf").Order("Line_No")

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return  nil
    }

    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return  nil
	}

	csv_inf      := make([]type5.Csv_Inf, 0, count)
	csv_inf_view := make([]type5.Csv_Inf, 0)

    keys, err := client.GetAll(ctx, query , &csv_inf)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, csv_infw := range csv_inf {

///  讖溯・縺ｫ繧医ｊ繝√ぉ繝・け鬆・岼繧偵そ繝・ヨ

         csv_inf_view = append(csv_inf_view, type5.Csv_Inf {        keys_wk[pos]           ,
                                                                    csv_infw.Line_No    ,
                                                                    csv_infw.File_Name  ,
                                                                    csv_infw.Column_Num ,
                                                                    csv_infw.Column1    ,
                                                                    csv_infw.Column2    ,
                                                                    csv_infw.Column3    ,
                                                                    csv_infw.Column4    ,
                                                                    csv_infw.Column5    ,
                                                                    csv_infw.Column6    ,
                                                                    csv_infw.Column7    ,
                                                                    csv_infw.Column8    ,
                                                                    csv_infw.Column9    ,
                                                                    csv_infw.Column10     })




	}

//    fmt.Fprintf( w, "trans.csv_inf : csv_inf_view %v\n", csv_inf_view )  // 繝・ヰ繝・け

    return	csv_inf_view
}

