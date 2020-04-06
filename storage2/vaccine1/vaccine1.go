package vaccine1

import (
//	      "fmt"
	      "io"
	      "strings"
	      "net/http"
	      "bufio"

	      "github.com/sawaq7/go12_ver1/storage2"

	      "cloud.google.com/go/storage"

                                         )
///
///   StoragePack : pack a file in Google Cloud Storage.
///

func File_Pack ( w http.ResponseWriter , r *http.Request ,bucket_name string ,file_name string ){

    file_name2 := file_name + "temp"

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket_name ,file_name2 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)  // インターフェイス型を型変換

	defer writer.Close()

// 水路惁E��ファイル　�E�Eead file�E�Eオープン

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket_name ,file_name , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // インターフェイス型を型変換

    defer reader.Close()

// ファイルリーダー(string用�E�を�E��E��E�

    sreader := bufio.NewReaderSize(reader, 4096)

    index := 0

    for {

      index ++     // レコードカウンターをカウンチE

//      fmt.Fprintf(w, "File_Pack : lndex %v\n", index )  // チE��チE��

// ファイルを１行read

      line ,err  := sreader.ReadString('\n')

      if err == io.EOF {

	      break

	  }

	  if err != nil {

		 http.Error(w, err.Error(), http.StatusInternalServerError)
	     return

	  }

//	  line = strings.Replace( line, ",", " ", -1)     /// 区刁E��斁E��を変更

      column := strings.Count( line ,",") + 1

//      fmt.Fprintf(w, "File_Pack : column %v\n", column )  // チE��チE��

      if  column > 1 {      //   レコードがスペ�EスでなぁE��ァイルに書き込み

          line2 := strings.Trim(line, " ")           ///   両端スペ�Eスをトリム
//          fmt.Fprintf(w, "File_Pack :line2 [%s]\n", line2 )  // チE��チE��

          storage2.File_Write_Line ( w ,writer ,line2 )

      }

   }

///
/// 　　　　ファイル名�E変更
///

   storage2.File_Delete ( w , r ,bucket_name ,file_name  )    //  旧ファイルを削除

   storage2.File_Rename ( w , r  ,bucket_name ,file_name2 ,file_name ) //  新ファイルをリネ�Eム


//	fmt.Fprintf(w, " File_Pack : Calculate succeeded.\n" )

	return
}
