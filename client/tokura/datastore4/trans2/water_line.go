package trans2

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )

///                           　　　　　　　　　　
///    get water-line                         　　　　　　　　　　　
///

func Water_line( funct int64 ,wname string ,w http.ResponseWriter, r *http.Request )  ([]type4.Water_Line ) {

//     IN  funct : ファンクション　0:   show all water-line
//               　　　　　　　　　1:   show water-line which was selected
//     IN  wname :  water-name
//     IN    w   : レスポンスライター
//     IN    r   : リクエストパラメータ
//     OUT  one  : water-line slice

//    fmt.Fprintf( w, "trans.water_line start \n" )
//    fmt.Fprintf( w, "trans.water_line funct %v   \n" , funct  )
//    fmt.Fprintf( w, "trans.water_line wname %v   \n" , wname  )

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

    query := datastore.NewQuery("Water_Line").Order("Section")

    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	water_line      := make([]type4.Water_Line, 0, count)

	water_line_view := make([]type4.Water_Line, 0)

	keys, err := client.GetAll(ctx, query , &water_line)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "water_line err \n" ,err)
//		return	water_line_view
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, water_linew := range water_line {

      if  funct == 0 {

         water_line_view = append(water_line_view, type4.Water_Line {  keys_wk[pos]         ,
                                                                    water_linew.Name              ,
                                                                    water_linew.Section              ,
                                                                    water_linew.Friction_Factor   ,
                                                                    water_linew.Velocity          ,
                                                                    water_linew.Pipe_Diameter     ,
                                                                    water_linew.Pipe_Length         })

      }else if wname == water_linew.Name  {

         water_line_view = append(water_line_view, type4.Water_Line { keys_wk[pos]           ,
                                                                    water_linew.Name              ,
                                                                    water_linew.Section              ,
                                                                    water_linew.Friction_Factor   ,
                                                                    water_linew.Velocity          ,
                                                                    water_linew.Pipe_Diameter     ,
                                                                    water_linew.Pipe_Length         })

      }

	}
    return	water_line_view
}

