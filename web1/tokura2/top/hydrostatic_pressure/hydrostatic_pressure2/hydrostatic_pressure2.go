package hydrostatic_pressure2

import (
//	    "fmt"
	    "strconv"
	    "strings"
	    "bufio"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri"
	    "github.com/sawaq7/go12_ver1/basic/maths/sum"
	    "storage2"
	    "net/http"
	    "io"
	    "cloud.google.com/go/storage"
    	                 )
///
/// 静水圧　U字管の計箁E   type2　�E�ファイル入力！E///

///  main process ///

func Hydrostatic_pressure2(w http.ResponseWriter, r *http.Request) {

// 単位容積重量　�E�ω）をセチE��

   var omega ,drad1 ,drad2 ,press1 ,press2,high ,area1 ,area2 float64
   var fname ,fname2  string
   var index ,num int

// ファイルオープン
   bucket := "sample-7777"
   fname = "seisui_inf.txt"
   fname2 = "seisui.txt"

   ad_fdata := make([]string ,6)        // keep work data for etc float data

// 静水惁E��ファイル、オープン

//   reader  := storage2.File_Open(w ,r ,bucket ,fname)

   reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,fname , w , r  )

   reader, _ := reader_minor.(io.ReadCloser)  // インターフェイス型を型変換

   defer reader.Close()

// ファイルリーダーをＧ�E��E�

   sreader := bufio.NewReaderSize(reader, 4096)

// 静水ファイル、オープン

   writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,fname2 , w , r  )

   writer, _ := writer_minor.(*storage.Writer)  // インターフェイス型を型変換


//   writer := storage2.File_Create ( w ,r ,bucket ,fname2 )

   index = 0        // レコードカウンターをinitialize

   for {

      index ++     // レコードカウンターをカウンチE
//      fmt.Fprintf (w, "hydrostatic-pressure2　index " ,index)  // チE��チE��

// ファイルを１行read

      line ,_  := sreader.ReadString('\n')

//斁E��単位にスペ�Eスで刁E��

      str := strings.Fields(line)

      num = len(str)

//      fmt.Fprintf (w,"hydrostatic-pressure2　num " ,num)  // チE��チE��

      if num == 0 {  //　END　チェチE��

//         fmt.Fprintf (w,"hydrostatic-pressure2 normal end")  // チE��チE��
         goto END
      }

      if index != 1{   // 見�Eし以外をmake

         omega ,_ =strconv.ParseFloat(str[0],64)
         drad1 ,_ =strconv.ParseFloat(str[1],64)
         drad2 ,_ =strconv.ParseFloat(str[2],64)
         press1 ,_ =strconv.ParseFloat(str[3],64)

         high ,_ =strconv.ParseFloat(str[5],64)

//         fmt.Fprintf ( w,"hydrostatic-pressure2 file data " ,omega, drad1 , drad2 ,press1  ,high ) // チE��チE��

// U字管の面積を計算すめE
         area1 = sum.Circle_Area(drad1/2)
         area2 = sum.Circle_Area(drad2/2)

         press2 =  suiri.Seisui1( area1 ,area2  ,press1  ,omega  ,high  )

//         fmt.Fprintf(w,"hydrostatic-pressure2 圧力２　",press2,"�E�E)   //チE��チE��

// U字管の吁E��惁E��を文字�Eに変換

         ad_fdata[0] = strconv.FormatFloat( omega ,  'f' ,8 ,64 )
         ad_fdata[1] = strconv.FormatFloat( drad1 ,  'f' ,8 ,64 )
         ad_fdata[2] = strconv.FormatFloat( drad2 ,  'f' ,8 ,64 )
         ad_fdata[3] = strconv.FormatFloat( press1 , 'f' ,8 ,64 )
         ad_fdata[4] = strconv.FormatFloat( press2 , 'f' ,8 ,64 )
         ad_fdata[5] = strconv.FormatFloat( high ,   'f' ,8 ,64 )

//         storage2.File_Write ( w ,bucket ,fname2 ,writer ,ad_fdata )
         storage2.File_write ( w ,writer ,ad_fdata )

      }
   }
   //  final process ,"example file delete ,rename ,close"

   END :

   defer writer.Close()

   return
}






