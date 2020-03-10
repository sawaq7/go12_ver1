package process3

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"
	    "general/datastore5/reformat"

//	    "general/html5"
	    "general/type5"
	    "general/datastore5/trans3"
	    "general/datastore5/set1"
	    "storage2"
	    "io"
	    "strings"
	    "bufio"
	    "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )

///
///     指定したファイルの1列のデータを追加する
///

func Csv_column_join(w http.ResponseWriter, r *http.Request , filename string ,column_no int ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  filename 　  : ファイル名
//     IN  column_no　  : 追加する行NO

//    fmt.Fprintf( w, "process3.csv_column_join start \n" )  // デバック

    var index   int64

    var column   int

    var str_work[10] string

    var bucket string

    csv_inf_join := make( []string , 0)

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // デバック

      project_name = "sample-7777"

	}
    ctx := context.Background()

	query := datastore.NewQuery("Storage_B_O_Temp")

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storage_b_o_temp      := make([]type5.Storage_B_O_Temp, 0, count)

    if _, err := client.GetAll(ctx, query , &storage_b_o_temp); err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{

      for pos, storage_b_o_tempw := range storage_b_o_temp {

        if pos == 0 {

//           project_name       = storage_b_o_tempw.Project_Name    // プロジェクト名をゲット
           bucket   = storage_b_o_tempw.Bucket_Name    // バケット名をゲット
//           filename = storage_b_o_tempw.Object_Name    // ベーシックファイル名をゲット

        }
	  }
	}

///
/// 　　　csv情報を修正する　
///

	csv_inf := trans3.Csv_inf ( w ,r )  ///      csv情報をゲット

    csv_inf2 := reformat.Csv_inf ( 1,csv_inf[0].Column_Num+1 ,csv_inf ,w ,r )
                                                           /// フォーマットを1列拡張する
///
///      追加するcsv情報をゲット
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // インターフェイス型を型変換

    defer reader.Close()

    csv_reader := bufio.NewReaderSize(reader, 4096)

    index = 0 // レコードカウンターをイニシャライズ

    for {

        index ++     // レコードカウンターをカウント

	    record ,err  := csv_reader.ReadString('\n')

//	    fmt.Fprintf( w, "csv_show : record %v\n", record )  // デバック

	    if err == io.EOF {

	      break

		}
		if err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
	      return

		}
		if index == 1 {   // 列数をゲット

		  column = strings.Count( record ,",") + 1

//		  fmt.Fprintf( w, "csv_show : column %v\n", column )  // デバック

		}

		str := strings.Split ( record, ","  )

//		fmt.Fprintf( w, "csv_show : str %v\n", str )  // デバック

		for ii := 0 ; ii < column ; ii++ {

           str_work[ii] = str[ii]

        }

        csv_inf_join = append( csv_inf_join ,str_work[column_no-1] )
                                                      ///    追加するcsv情報をセット

    }

    csv_inf_new := set1.Csv_inf (  csv_inf2 ,csv_inf_join ,int(csv_inf2[0].Column_Num) , w ,r )
                                                        /// 追加するデータ1列をセット

///
/// 　　　データストアに、csv情報を再セットする　
///

    for _, csv_inf_neww := range csv_inf_new {

//   	  fmt.Fprintf( w, "process3.csv_column_join csv_inf_neww %v\n", csv_inf_neww )  // デバック

      key := datastore.IDKey("Csv_Inf", csv_inf_neww.Id, nil)

      if _, err := client.Put(ctx, key, &csv_inf_neww ); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	  }

    }

//	fmt.Fprintf( w, "process3.csv_column_join normal end \n" )  // デバック


}

