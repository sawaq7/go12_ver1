    ///////////////////////////////////////////////////////
   ///    pipe_line_show                               ///
  ///     水路ヘッダー情報を　表示                      ///
 ///                                                  ///
////////////////////////////////////////////////////////

package pipe_line1_delete

import (
//	     "fmt"
	     "strings"
	     "bufio"
	     "io"
	     "net/http"
	     "strconv"
	     "client/tokura/suiri"
	     "client/tokura/suiri/type4"
	     "storage2"
	     "cloud.google.com/go/storage"

	                                     )

///  main process ///

func Pipe_line1_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "pipe_line1_show delete \n" )  // デバック

    water := make([]type4.Water,100 )

// 水路情報カウンターをinitialize

    pos := 0

// バケット名・ファイル名　セット

    bucket := "sample-7777"
    filename1 := "water_inf.txt"
    filename2 := "work.txt"    // test test test test

// 実行する水路を判定する

    water_id , err := strconv.Atoi(r.FormValue("water_id"))
//    fmt.Fprintf( w, "pipe_line1_excute_delete water_id %v\n", water_id )  // デバック

	if err  != nil {

//	   fmt.Fprintf( w, "pipe_line1_excute_delete :error water_id"  )  // デバック

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

// ファイルの rename

    storage2.File_Rename ( w ,r ,bucket ,filename1 ,filename2 )

// 差し替えた、水路情報ファイルを　（read file） オープン

//    reader := storage2.File_Open(w ,r ,bucket ,filename2)

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename2 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // インターフェイス型を型変換

//    defer reader.Close()

// ファイルリーダー(string用）をＧＥＴ

    sreader := bufio.NewReaderSize(reader, 4096)

// 新しく水路情報ファイルを作成

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,filename1 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)  // インターフェイス型を型変換

//    writer := storage2.File_Create( w ,r ,bucket ,filename1 )

    defer writer.Close()


// 導水勾配線ファイル、オープン

   index := 0        // レコードカウンターをinitialize


   for {
      index ++     // レコードカウンターをカウント

//      fmt.Fprintf(w, "pipe_line1_delete : lndex %v\n", index )  // デバック

// ファイルを１行read

      line ,_  := sreader.ReadString('\n')

      str := strings.Fields(line) // ブランクで分割

      num := len(str)

//      fmt.Fprintf(w, "pipe_line1_delete : num %v\n", num )  // デバック

      if num != 0 && index != water_id {

          storage2.File_write ( w ,writer ,str )

         if index == 1 {

// ヘッダーは、スルーする

//             fmt.Fprintf(w, "pipe_line1_delete (header) : line %s\n", line )  // デバック

          }else{

/// 水路ヘッダー情報を　GET
             pos ++     // 水路情報カウンターをカウント

//             fmt.Fprintf(w, "pipe_line1_delete (the other): line %s\n", line )  // デバック

             water[pos-1].No = strconv.Itoa(index) //　整数を文字に変換
             water[pos-1].Name ,water[pos-1].High ,water[pos-1].Roughness_factor = suiri.Kansui1_2( line  )

//             fmt.Fprintf(w, "pipe_line1_delete : 水路ナンバー %v\n", water[pos-1].No )  // デバック
//             fmt.Fprintf(w, "pipe_line1_delete : 水路名 %s\n", water[pos-1] .Name)  // デバック
//             fmt.Fprintf(w, "pipe_line1_delete : 水路高 %s\n", water[pos-1].High )  // デバック
//             fmt.Fprintf(w, "pipe_line1_delete : 粗度係数 %s\n", water[pos-1].Roughness_factor )  // デバック

          }

      } else if num == 0 {                                         // End check

//          io.WriteString(w, "\n pipe_line1_delete : data end \n")   //デバック

         break

      }
   }

// ワークファイルを削除

   storage2.Storage_basic( "delete" ,bucket ,filename2 , w , r  )

//   storage2.File_Delete ( w , r ,bucket ,filename2 )

// スライスを圧縮

//   fmt.Fprintf(w, "pipe_line1_delete : len(water) cap(water) %v\n", len(water)  ,cap(water))  // デバック

   water2 := make([]type4.Water, pos )
   copy ( water2 ,water[0:pos] ) // 注：データは、0　から　pos－1　まで

/// 水路情報　表示

   suiri.Pipe_line1_show( w ,pos , water2 )

// end process

//	fmt.Fprintf(w, "\n pipe_line1_delete : Calculate succeeded.\n" )

    return

}
