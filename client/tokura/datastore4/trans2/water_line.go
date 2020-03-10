package trans2

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

//	    "html/template"
//	    "web/htmls/sgh"
	    "client/tokura/suiri/type4"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )

///                           　　　　　　　　　　
/// データストアーから水路ライン情報をGETする（水路ラインファイル）
///                          　　　　　　　　　　　


func Water_line( funct int64 ,wname string ,w http.ResponseWriter, r *http.Request )  ([]type4.Water_Line ) {

//     IN  funct : ファンクション　0:すべての水路ラインを表示
//               　　　　　　　　　1:指定した水路名の水路ラインを表示
//     IN  wname : 水路名
//     IN    w   : レスポンスライター
//     IN    r   : リクエストパラメータ
//     OUT  one  : 水路ラインのスライス

//    fmt.Fprintf( w, "trans.water_line start \n" )  // デバック
//    fmt.Fprintf( w, "trans.water_line funct %v   \n" , funct  )  // デバック
//    fmt.Fprintf( w, "trans.water_line wname %v   \n" , wname  )  // デバック

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

    query := datastore.NewQuery("Water_Line").Order("Section")
//	q := datastore.NewQuery("Water_Line").Order("Section")

//	count, err := q.Count(c)
    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	water_line      := make([]type4.Water_Line, 0, count)

	water_line_view := make([]type4.Water_Line, 0)


//	keys, err := q.GetAll(c, &water_line)
	keys, err := client.GetAll(ctx, query , &water_line)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "water_line err \n" ,err)  // デバック
//		return	water_line_view
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, water_linew := range water_line {

///  機能によりチェック項目をセット
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

