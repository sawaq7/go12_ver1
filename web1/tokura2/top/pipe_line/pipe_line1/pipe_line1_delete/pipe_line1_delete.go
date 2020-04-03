    ///////////////////////////////////////////////////////
   ///    pipe_line_show                               ///
  ///     水路ヘッダー惁E��を　表示                      ///
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
	     "github.com/sawaq7/go12_ver1/client/tokura/suiri"
	     "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	     "storage2"
	     "cloud.google.com/go/storage"

	                                     )

///  main process ///

func Pipe_line1_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "pipe_line1_show delete \n" )  // チE��チE��

    water := make([]type4.Water,100 )

// 水路惁E��カウンターをinitialize

    pos := 0

// バケチE��名�Eファイル名　セチE��

    bucket := "sample-7777"
    filename1 := "water_inf.txt"
    filename2 := "work.txt"    // test test test test

// 実行する水路を判定すめE
    water_id , err := strconv.Atoi(r.FormValue("water_id"))
//    fmt.Fprintf( w, "pipe_line1_excute_delete water_id %v\n", water_id )  // チE��チE��

	if err  != nil {

//	   fmt.Fprintf( w, "pipe_line1_excute_delete :error water_id"  )  // チE��チE��

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

// ファイルの rename

    storage2.File_Rename ( w ,r ,bucket ,filename1 ,filename2 )

// 差し替えた、水路惁E��ファイルを　�E�Eead file�E�Eオープン

//    reader := storage2.File_Open(w ,r ,bucket ,filename2)

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename2 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // インターフェイス型を型変換

//    defer reader.Close()

// ファイルリーダー(string用�E�を�E��E��E�

    sreader := bufio.NewReaderSize(reader, 4096)

// 新しく水路惁E��ファイルを作�E

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,filename1 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)  // インターフェイス型を型変換

//    writer := storage2.File_Create( w ,r ,bucket ,filename1 )

    defer writer.Close()


// 導水勾配線ファイル、オープン

   index := 0        // レコードカウンターをinitialize


   for {
      index ++     // レコードカウンターをカウンチE
//      fmt.Fprintf(w, "pipe_line1_delete : lndex %v\n", index )  // チE��チE��

// ファイルを１行read

      line ,_  := sreader.ReadString('\n')

      str := strings.Fields(line) // ブランクで刁E��

      num := len(str)

//      fmt.Fprintf(w, "pipe_line1_delete : num %v\n", num )  // チE��チE��

      if num != 0 && index != water_id {

          storage2.File_write ( w ,writer ,str )

         if index == 1 {

// ヘッダーは、スルーする

//             fmt.Fprintf(w, "pipe_line1_delete (header) : line %s\n", line )  // チE��チE��

          }else{

/// 水路ヘッダー惁E��を　GET
             pos ++     // 水路惁E��カウンターをカウンチE
//             fmt.Fprintf(w, "pipe_line1_delete (the other): line %s\n", line )  // チE��チE��

             water[pos-1].No = strconv.Itoa(index) //　整数を文字に変換
             water[pos-1].Name ,water[pos-1].High ,water[pos-1].Roughness_factor = suiri.Kansui1_2( line  )

//             fmt.Fprintf(w, "pipe_line1_delete : 水路ナンバ�E %v\n", water[pos-1].No )  // チE��チE��
//             fmt.Fprintf(w, "pipe_line1_delete : 水路吁E%s\n", water[pos-1] .Name)  // チE��チE��
//             fmt.Fprintf(w, "pipe_line1_delete : 水路髁E%s\n", water[pos-1].High )  // チE��チE��
//             fmt.Fprintf(w, "pipe_line1_delete : 粗度係数 %s\n", water[pos-1].Roughness_factor )  // チE��チE��

          }

      } else if num == 0 {                                         // End check

//          io.WriteString(w, "\n pipe_line1_delete : data end \n")   //チE��チE��

         break

      }
   }

// ワークファイルを削除

   storage2.Storage_basic( "delete" ,bucket ,filename2 , w , r  )

//   storage2.File_Delete ( w , r ,bucket ,filename2 )

// スライスを圧縮

//   fmt.Fprintf(w, "pipe_line1_delete : len(water) cap(water) %v\n", len(water)  ,cap(water))  // チE��チE��

   water2 := make([]type4.Water, pos )
   copy ( water2 ,water[0:pos] ) // 注�E�データは、E　から　pos�E�E　まで

/// 水路惁E��　表示

   suiri.Pipe_line1_show( w ,pos , water2 )

// end process

//	fmt.Fprintf(w, "\n pipe_line1_delete : Calculate succeeded.\n" )

    return

}
