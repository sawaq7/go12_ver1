package process3

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"
	    "github.com/sawaq7/go12_ver1/general/datastore5/reformat"

//	    "github.com/sawaq7/go12_ver1/general/html5"
	    "github.com/sawaq7/go12_ver1/general/type5"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
	    "github.com/sawaq7/go12_ver1/general/datastore5/set1"
	    "storage2"
	    "io"
	    "strings"
	    "bufio"
	    "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )

///
///     持E��したファイルの1列�EチE�Eタを追加する
///

func Csv_column_join(w http.ResponseWriter, r *http.Request , filename string ,column_no int ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  filename 　  : ファイル吁E//     IN  column_no　  : 追加する行NO

//    fmt.Fprintf( w, "process3.csv_column_join start \n" )  // チE��チE��

    var index   int64

    var column   int

    var str_work[10] string

    var bucket string

    csv_inf_join := make( []string , 0)

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // チE��チE��

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

//           project_name       = storage_b_o_tempw.Project_Name    // プロジェクト名をゲチE��
           bucket   = storage_b_o_tempw.Bucket_Name    // バケチE��名をゲチE��
//           filename = storage_b_o_tempw.Object_Name    // ベ�EシチE��ファイル名をゲチE��

        }
	  }
	}

///
/// 　　　csv惁E��を修正する　
///

	csv_inf := trans3.Csv_inf ( w ,r )  ///      csv惁E��をゲチE��

    csv_inf2 := reformat.Csv_inf ( 1,csv_inf[0].Column_Num+1 ,csv_inf ,w ,r )
                                                           /// フォーマットを1列拡張する
///
///      追加するcsv惁E��をゲチE��
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // インターフェイス型を型変換

    defer reader.Close()

    csv_reader := bufio.NewReaderSize(reader, 4096)

    index = 0 // レコードカウンターをイニシャライズ

    for {

        index ++     // レコードカウンターをカウンチE
	    record ,err  := csv_reader.ReadString('\n')

//	    fmt.Fprintf( w, "csv_show : record %v\n", record )  // チE��チE��

	    if err == io.EOF {

	      break

		}
		if err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
	      return

		}
		if index == 1 {   // 列数をゲチE��

		  column = strings.Count( record ,",") + 1

//		  fmt.Fprintf( w, "csv_show : column %v\n", column )  // チE��チE��

		}

		str := strings.Split ( record, ","  )

//		fmt.Fprintf( w, "csv_show : str %v\n", str )  // チE��チE��

		for ii := 0 ; ii < column ; ii++ {

           str_work[ii] = str[ii]

        }

        csv_inf_join = append( csv_inf_join ,str_work[column_no-1] )
                                                      ///    追加するcsv惁E��をセチE��

    }

    csv_inf_new := set1.Csv_inf (  csv_inf2 ,csv_inf_join ,int(csv_inf2[0].Column_Num) , w ,r )
                                                        /// 追加するチE�Eタ1列をセチE��

///
/// 　　　チE�Eタストアに、csv惁E��を�EセチE��する　
///

    for _, csv_inf_neww := range csv_inf_new {

//   	  fmt.Fprintf( w, "process3.csv_column_join csv_inf_neww %v\n", csv_inf_neww )  // チE��チE��

      key := datastore.IDKey("Csv_Inf", csv_inf_neww.Id, nil)

      if _, err := client.Put(ctx, key, &csv_inf_neww ); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	  }

    }

//	fmt.Fprintf( w, "process3.csv_column_join normal end \n" )  // チE��チE��


}

