package csv_show

import (

	    "net/http"
//	    "fmt"

	    "storage2/vaccine1"
	    "storage2"
	    "strconv"
	    "strings"
	    "io"
	    "bufio"
	    "html/template"

	    "general/type5"
	    "general/html5"
	    "general/datastore5/initialize"
	    "general/datastore5/trans3"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                  )

func Csv_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "csv_show start \n" )  // デバック

    var bucket ,filename string

    var index   int64

    var column  int

    var str_work[10] string

    var csv_inf type5.Csv_Inf

    var storage_b_o_temp2 type5.Storage_B_O_Temp

    line_no := r.FormValue("line_no")

//    fmt.Fprintf( w, "csv_show : line_no %v\n", line_no )  // デバック

	select_id ,_ := strconv.Atoi(line_no)

//    fmt.Fprintf( w, "csv_show : select_id %v\n", select_id )  // デバック

///
///   バケット名をゲット
///

     projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // デバック

      projectID = "sample-7777"

	}

    ctx := context.Background()

    query := datastore.NewQuery("Storage_B_O_Temp")

    client, err := datastore.NewClient(ctx, projectID)
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

    if _, err = client.GetAll(ctx, query , &storage_b_o_temp ) ;  err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{

      for pos, storage_b_o_tempw := range storage_b_o_temp {

        if pos == 0 {
           bucket    = storage_b_o_tempw.Bucket_Name    // バケット名をゲット

        }
	  }
	}

//    fmt.Fprintf( w, "csv_show : bucket2: %v\n", bucket2 )

///
///   ファイル名をゲット
///

	objects :=  storage2.Object_List ( w  ,r , bucket )

//    fmt.Fprintf( w, "csv_show : objects: %v\n", objects )

    for pos2 , objectsw := range objects {

      if pos2 == select_id-1 {

        filename = objectsw

      }

    }

///
///      csvファイルをイニシャライズ
///

    initialize.Csv_inf (w , r )

///
///      csvファイル　情報をゲットして表示
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // インターフェイス型を型変換

    defer reader.Close()

    csv_reader := bufio.NewReaderSize(reader, 4096)

    index = 0 // レコードカウンターをイニシャライズ

    warn_flag := 0  // ワーニングフラグのイニシャライズ

    for {

	    record ,err  := csv_reader.ReadString('\n')

//	    fmt.Fprintf( w, "csv_show : record %v\n", record )  // デバック

	    if err == io.EOF {

	      break

		}
		if err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
	      return

		}

        record = strings.Replace( record, ",", " ", -1)     /// 区切り文字を変更

        column = strings.Count( record ," ") + 1

//        fmt.Fprintf( w, "csv_show : column %v\n", strings.Count( record ," ") + 1 )  // デバック

        if  column > 1 {  //   レコードがスペースでない場合処理

          index ++     // レコードカウンターをカウント

          str := strings.Fields(record)

//		fmt.Fprintf( w, "csv_show : str %v\n", str )  // デバック

		  for ii := 0 ; ii < column ; ii++ {

             str_work[ii] = str[ii]

          }

///    ワークエリア(データストア）にcsv情報をセット

          csv_inf.Line_No    = index
          csv_inf.File_Name  = filename
          csv_inf.Column_Num = int64(column)
	      csv_inf.Column1    = str_work[0]
	      csv_inf.Column2    = str_work[1]
	      csv_inf.Column3    = str_work[2]
	      csv_inf.Column4    = str_work[3]
	      csv_inf.Column5    = str_work[4]
	      csv_inf.Column6    = str_work[5]
	      csv_inf.Column7    = str_work[6]
	      csv_inf.Column8    = str_work[7]
	      csv_inf.Column9    = str_work[8]
	      csv_inf.Column10   = str_work[9]

          new_key := datastore.IncompleteKey("Csv_Inf", nil)

//    fmt.Fprintf(w, "storage_object_copy_excute: new_key %v\n", new_key )  // デバック

         _, err = client.Put(ctx, new_key, &csv_inf )
	     if err != nil {
		   http.Error(w,err.Error(), http.StatusInternalServerError)
		   return
	     }

       } else {

      	 warn_flag = 1

       }

	}

    initialize.Storage_b_o_temp (w , r ) //  既存の　Storage_B_O_Temp コモン用のtemporary-fileをクリアー

    storage_b_o_temp2.Line_No =  1
    storage_b_o_temp2.Project_Name = projectID
    storage_b_o_temp2.Bucket_Name = bucket
    storage_b_o_temp2.Object_Name = filename

/// コモン用のtemporary-fileにバケット名を再セット

    new_key := datastore.IncompleteKey("Storage_B_O_Temp", nil)

    if _, err := client.Put(ctx, new_key, &storage_b_o_temp2 ); err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError )
    }

///
///　　 ファイルの補正する
///

    if warn_flag == 1 {

       vaccine1.File_Pack ( w , r ,bucket  ,filename )

	}

///
///　　web にcsv情報を表示
///

     csv_inf_view := trans3.Csv_inf ( w ,r )  ///      csv情報をゲット

     monitor := template.Must( template.New("html").Parse( html5.Csv_show )) // テンプレートのヘッダーをGET

     err = monitor.Execute ( w, csv_inf_view )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

//	fmt.Fprintf( w, "csv_show : normal end \n" )  // デバック

}

