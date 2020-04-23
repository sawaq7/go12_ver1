package trans

import (

//	    "google.golang.org/appengine/datastore"
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
///     make expression for each course
///                       　


func Sgh_ai( course_no int64 ,w http.ResponseWriter, r *http.Request )  (sgh_ai_view []type2.Sgh_Ai ) {

//     IN 　course_no　 : コースNo
//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//     OUT sgh_ai_view  : slice of struct ( Sgh_Ai )

//    fmt.Fprintf( w, "trans.sgh_ai start \n" )

//    fmt.Fprintf( w, "trans.sgh_ai course_no %v\n" ,course_no )

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

    query := datastore.NewQuery("Sgh_Ai").Order("Course_No")

    count, err := client.Count(ctx, query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	sgh_ai      := make([]type2.Sgh_Ai, 0, count)
	sgh_ai_view = make([]type2.Sgh_Ai, 0)

    keys, err := client.GetAll(ctx, query , &sgh_ai)
//	keys, err := q.GetAll(c, &sgh_ai)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)

        return	nil
	}

    keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

    count2 := 0

	for pos, sgh_aiw := range sgh_ai {

      if  course_no == sgh_aiw.Course_No &&
          count2     == 0                    {

          count2 ++

         sgh_ai_view = append(sgh_ai_view, type2.Sgh_Ai {           keys_wk[pos]          ,

                                                                    sgh_aiw.Course_No     ,
                                                                    sgh_aiw.District_No   ,
                                                                    sgh_aiw.District_Name ,
                                                                    sgh_aiw.Area_No       ,
                                                                    sgh_aiw.Area_Name     ,
                                                                    sgh_aiw.Date_Basic      ,
                                                                    sgh_aiw.Date_Basic_Real ,

                                                                    sgh_aiw.Ex_Type       ,
                                                                    sgh_aiw.Expression    ,
                                                                    sgh_aiw.Item_Num      ,
                                                                    sgh_aiw.Item1_Name    ,
                                                                    sgh_aiw.Item1_Factor  ,
                                                                    sgh_aiw.Item2_Name    ,
                                                                    sgh_aiw.Item2_Factor  ,
                                                                    sgh_aiw.Item3_Name    ,
                                                                    sgh_aiw.Item3_Factor  ,
                                                                    sgh_aiw.Item4_Name    ,
                                                                    sgh_aiw.Item4_Factor  ,
                                                                    sgh_aiw.Item5_Name    ,
                                                                    sgh_aiw.Item5_Factor           })

//        fmt.Fprintf( w, "trans.sgh_ai sgh_aiw.Expression %v\n" ,sgh_aiw.Expression )

      }
	}

//    fmt.Fprintf( w, "trans.sgh_ai normal end \n" )
    return	sgh_ai_view
}

