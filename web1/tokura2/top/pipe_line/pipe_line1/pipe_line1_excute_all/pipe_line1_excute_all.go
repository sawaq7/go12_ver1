////////////////////////////////////////////////////////////////////////
///                                                                  ///
///   動水勾配線データを作�Eする　　　　　　　　　　　　　　　　　　　  ///
///                                                                  ///
////////////////////////////////////////////////////////////////////////

package pipe_line1_excute_all

import (
//	      "fmt"
	      "io"
	      "strings"
	      "net/http"
	      "bufio"
	      "github.com/sawaq7/go12_ver1/client/tokura/suiri"
	      "storage2"
	      "cloud.google.com/go/storage"
                                         )



///  main process ///

func Pipe_line1_excute_all(w http.ResponseWriter, r *http.Request) {

//    var gcs_gae storage2.Gcs_Gae  // struct for accessing Google Cloud Storage get from storage/srorage2/storage_access.go

    bucket := "sample-7777"
    filename1 := "water_inf.txt"
    filename2 := "grade_line.txt"

// 導水勾配線ファイル�E�Erade_line.txt�E�　�E�Erite file�E�EクリエイチE

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,filename2 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)  // インターフェイス型を型変換

//	writer := storage2.File_Create( w ,r ,bucket ,filename2 )

	defer writer.Close()

// 水路惁E��ファイル　�E�Eead file�E�Eオープン

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // インターフェイス型を型変換

//    reader := storage2.File_Open(w ,r ,bucket ,filename1)

    defer reader.Close()

// ファイルリーダー(string用�E�を�E��E��E�

    sreader := bufio.NewReaderSize(reader, 4096)

    index := 0

    for {

      index ++     // レコードカウンターをカウンチE

//      fmt.Fprintf(w, "pipe_line1_excute_all : lndex %v\n", index )  // チE��チE��

// ファイルを１行read

      line ,_  := sreader.ReadString('\n')

//斁E��単位にスペ�Eスで刁E��

      str := strings.Fields(line)

      num := len(str)

//      fmt.Fprintf(w, "pipe_line1_excute_all : num %v\n", num )  // チE��チE��

      if num != 0 {
         if index == 1{

// ヘッダーは、スルーする

//             fmt.Fprintf(w, "pipe_line1_excute_all (header) : line %s\n", line )  // チE��チE��

          }else{
//             fmt.Fprintf(w, "pipe_line1_excute_all (the other): line %s\n", line )  // チE��チE��

/// 動水勾配線データを作�E

             ad_hp ,ad_hl ,ad_vhead ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown := suiri.Kansui1( line  )

//             fmt.Fprintf(w, "pipe_line1_excute_all : ad_hp %s\n", ad_hp )  // チE��チE��
//             fmt.Fprintf(w, "pipe_line1_excute_all : ad_hl %s\n", ad_hl )  // チE��チE��
//             fmt.Fprintf(w, "pipe_line1_excute_all : ad_vhead %s\n", ad_vhead )  // チE��チE��
//             fmt.Fprintf(w, "pipe_line1_excute_all : ad_eneup %s\n", ad_eneup )  // チE��チE��
//             fmt.Fprintf(w, "pipe_line1_excute_all : ad_enedown %s\n", ad_enedown )  // チE��チE��
//             fmt.Fprintf(w, "pipe_line1_excute_all : ad_glineup %s\n", ad_glineup )  // チE��チE��
//             fmt.Fprintf(w, "pipe_line1_excute_all : ad_glinedown %s\n", ad_glinedown )  // チE��チE��

// ポイント損失惁E��をwrite

//                 storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_hp )
                 storage2.File_write ( w ,writer ,ad_hp )

// ライン損失惁E��をwrite

//                 storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_hl )
                 storage2.File_write ( w ,writer ,ad_hl )

// 速度水頭惁E��をwrite

//                 storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_vhead )
                 storage2.File_write ( w ,writer ,ad_vhead )

// エネルギー線upをwrite

//                 storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_eneup )
                 storage2.File_write ( w ,writer ,ad_eneup )

// エネルギー線downをwrite

//                 storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_enedown )
                 storage2.File_write ( w ,writer ,ad_enedown )

// 動水勾配線upをwrite

//                 storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_glineup )
                 storage2.File_write ( w ,writer ,ad_glineup )

// 動水勾配線downをwrite

//                 storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_glinedown )
                 storage2.File_write ( w ,writer ,ad_glinedown )

          }

      } else if num == 0 {

//          io.WriteString(w, "\n pipe_line1_excute_all : data end \n")   //チE��チE��

         break

      }
   }

// end process

//    fmt.Fprintf(w, "\n pipe_line1_excute_all : Calculate succeeded.\n" )

	return
}







