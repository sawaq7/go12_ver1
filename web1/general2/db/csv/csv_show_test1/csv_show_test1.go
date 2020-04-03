package csv_show_test1

import (

	    "net/http"
	    "fmt"


//	    "github.com/sawaq7/go12_ver1/storage2/get"
	    "storage2"
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

    fmt.Fprintf( w, "csv_show_test start \n" )  // 繝・ヰ繝・け

    var bucket ,filename string

    var index   int64

    var column int


    line_no := r.FormValue("line_no")

//    fmt.Fprintf( w, "csv_show : line_no %v\n", line_no )  // 繝・ヰ繝・け

	select_id ,_ := strconv.Atoi(line_no)

//    fmt.Fprintf( w, "csv_show : select_id %v\n", select_id )  // 繝・ヰ繝・け


///
///   繝舌こ繝・ヨ蜷阪ｒ繧ｲ繝・ヨ
///
     projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // 繝・ヰ繝・け

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
           bucket    = storage_b_o_tempw.Bucket_Name    // 繝舌こ繝・ヨ蜷阪ｒ繧ｲ繝・ヨ

        }
	  }
	}

//    fmt.Fprintf( w, "csv_show : bucket2: %v\n", bucket2 )

///
///   繝輔ぃ繧､繝ｫ蜷阪ｒ繧ｲ繝・ヨ
///

	objects :=  storage2.Object_List ( w  ,r , bucket )

//    fmt.Fprintf( w, "csv_show : objects: %v\n", objects )

    for pos2 , objectsw := range objects {

      if pos2 == select_id-1 {

        filename = objectsw


      }

    }

///
///      csv繝輔ぃ繧､繝ｫ繧偵う繝九す繝｣繝ｩ繧､繧ｺ
///

    initialize.Csv_inf (w , r )

///
///      csv繝輔ぃ繧､繝ｫ縲諠・ｱ繧偵ご繝・ヨ縺励※陦ｨ遉ｺ
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

    defer reader.Close()

    csv_reader := bufio.NewReaderSize(reader, 4096)

    index = 0 // 繝ｬ繧ｳ繝ｼ繝峨き繧ｦ繝ｳ繧ｿ繝ｼ繧偵う繝九す繝｣繝ｩ繧､繧ｺ

    for {

        index ++     // 繝ｬ繧ｳ繝ｼ繝峨き繧ｦ繝ｳ繧ｿ繝ｼ繧偵き繧ｦ繝ｳ繝・
	    record ,err  := csv_reader.ReadString('\n')

	    fmt.Fprintf( w, "csv_show : record %v\n", record )  // 繝・ヰ繝・け

	    if err == io.EOF {

	      break

		}
		if err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
	      return

		}

		record = strings.Replace( record, ",", " ", -1)     /// 蛹ｺ蛻・ｊ譁・ｭ励ｒ螟画峩

		if index == 1 {   // 蛻玲焚繧偵ご繝・ヨ

		  column = strings.Count( record ," ") + 1

		  fmt.Fprintf( w, "csv_show : column %v\n", column )  // 繝・ヰ繝・け

		}




///    繝ｯ繝ｼ繧ｯ繧ｨ繝ｪ繧｢(繝・・繧ｿ繧ｹ繝医い・峨↓csv諠・ｱ繧偵そ繝・ヨ


	}

//	fmt.Fprintf( w, "csv_show : normal end \n" )  // 繝・ヰ繝・け

}

