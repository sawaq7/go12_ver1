package csv_show_test1

import (

	    "net/http"
	    "fmt"


//	    "github.com/sawaq7/go12_ver1/storage2/get"
	    "github.com/sawaq7/go12_ver1/storage2"
	    "strconv"
	    "strings"
	    "io"
	    "bufio"


	    "github.com/sawaq7/go12_ver1/general/type5"

	    "github.com/sawaq7/go12_ver1/general/datastore5/initialize"


        "cloud.google.com/go/datastore"
        "context"
        "os"


                                                  )


func Csv_show_test1(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf( w, "csv_show_test start \n" )  // チE��チE��

    var bucket ,filename string

    var index   int64

    var column int


    line_no := r.FormValue("line_no")

//    fmt.Fprintf( w, "csv_show : line_no %v\n", line_no )  // チE��チE��

	select_id ,_ := strconv.Atoi(line_no)

//    fmt.Fprintf( w, "csv_show : select_id %v\n", select_id )  // チE��チE��


///
///   バケチE��名をゲチE��
///
     projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // チE��チE��

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
           bucket    = storage_b_o_tempw.Bucket_Name    // バケチE��名をゲチE��

        }
	  }
	}

//    fmt.Fprintf( w, "csv_show : bucket2: %v\n", bucket2 )

///
///   ファイル名をゲチE��
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
///      csvファイル　惁E��をゲチE��して表示
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // インターフェイス型を型変換

    defer reader.Close()

    csv_reader := bufio.NewReaderSize(reader, 4096)

    index = 0 // レコードカウンターをイニシャライズ

    for {

        index ++     // レコードカウンターをカウンチE
	    record ,err  := csv_reader.ReadString('\n')

	    fmt.Fprintf( w, "csv_show : record %v\n", record )  // チE��チE��

	    if err == io.EOF {

	      break

		}
		if err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
	      return

		}

		record = strings.Replace( record, ",", " ", -1)     /// 区刁E��斁E��を変更

		if index == 1 {   // 列数をゲチE��

		  column = strings.Count( record ," ") + 1

		  fmt.Fprintf( w, "csv_show : column %v\n", column )  // チE��チE��

		}




///    ワークエリア(チE�Eタストア�E�にcsv惁E��をセチE��


	}

//	fmt.Fprintf( w, "csv_show : normal end \n" )  // チE��チE��

}

