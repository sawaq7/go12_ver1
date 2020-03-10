package pipe_line1_show

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
	     "html/template"
	     "client/tokura/html4"
	                               )

///  main process ///

func Pipe_line1_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "Pipe_line1_show start \n" )  // デバック

    water := make([]type4.Water,100 )
//    water := make([]type4.Water,4 )

// 水路情報カウンターをinitialize

    pos := 0

// バケット名・ファイル名　セット

    bucket := "sample-7777"
    filename1 := "water_inf.txt"

// 水路情報ファイル　（read file） オープン

//	reader  := storage2.File_Open(w ,r ,bucket ,filename1)

	reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // インターフェイス型を型変換

    defer reader.Close()

// ファイルリーダー(string用）をＧＥＴ

    sreader := bufio.NewReaderSize(reader, 4096)

// 導水勾配線ファイル、オープン

   index := 0        // レコードカウンターをinitialize


   for {
      index ++     // レコードカウンターをカウント

//      fmt.Fprintf(w, "pipe_line1_show : lndex %v\n", index )  // デバック

// ファイルを１行read

      line ,_  := sreader.ReadString('\n')

      str := strings.Fields(line)

      num := len(str)

//      fmt.Fprintf(w, "pipe_line1_show : num %v\n", num )  // デバック

      if num != 0 {
         if index == 1 {

// ヘッダーは、スルーする

//             fmt.Fprintf(w, "pipe_line1_show (header) : line %s\n", line )  // デバック

          }else{

/// 水路ヘッダー情報を　GET
             pos ++     // 水路情報カウンターをカウント

//             fmt.Fprintf(w, "pipe_line1_show (the other): line %s\n", line )  // デバック

             water[pos-1].No = strconv.Itoa(index) //　整数を文字に変換
             water[pos-1].Name ,water[pos-1].High ,water[pos-1].Roughness_factor = suiri.Kansui1_2( line  )

//             fmt.Fprintf(w, "pipe_line1_show : 水路ナンバー %v\n", water[pos-1].No )  // デバック
//             fmt.Fprintf(w, "pipe_line1_show : 水路名 %s\n", water[pos-1] .Name)  // デバック
//             fmt.Fprintf(w, "pipe_line1_show : 水路高 %s\n", water[pos-1].High )  // デバック
//             fmt.Fprintf(w, "pipe_line1_show : 粗度係数 %s\n", water[pos-1].Roughness_factor )  // デバック

          }

      } else if num == 0 {                                         // End check

//          io.WriteString(w, "\n pipe_line1_show : data end \n")   //デバック

         break

      }
   }
//   fmt.Fprintf(w, "pipe_line1_show : len(water) cap(water) %v\n", len(water)  ,cap(water))  // デバック

// スライスを圧縮

   water2 := make([]type4.Water, pos )
   copy ( water2 ,water[0:pos] ) // 注：データは、0　から　pos－1　まで

/// 水路情報　表示

   monitor := template.Must(template.New("html").Parse( html4.Pipe_line1_show))
    err := monitor.Execute(w, water2)
    if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	}

//   suiri.Pipe_line1_show( w ,pos , water2 )

// end process

//    fmt.Fprintf(w, "\n pipe_line1_show : Calculate succeeded.\n" )

    return

}
