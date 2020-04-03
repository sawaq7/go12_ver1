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
/// 髱呎ｰｴ蝨ｧ縲U蟄礼ｮ｡縺ｮ險育ｮ・   type2縲・医ヵ繧｡繧､繝ｫ蜈･蜉幢ｼ・///

///  main process ///

func Hydrostatic_pressure2(w http.ResponseWriter, r *http.Request) {

// 蜊倅ｽ榊ｮｹ遨埼㍾驥上・委会ｼ峨ｒ繧ｻ繝・ヨ

   var omega ,drad1 ,drad2 ,press1 ,press2,high ,area1 ,area2 float64
   var fname ,fname2  string
   var index ,num int

// 繝輔ぃ繧､繝ｫ繧ｪ繝ｼ繝励Φ
   bucket := "sample-7777"
   fname = "seisui_inf.txt"
   fname2 = "seisui.txt"

   ad_fdata := make([]string ,6)        // keep work data for etc float data

// 髱呎ｰｴ諠・ｱ繝輔ぃ繧､繝ｫ縲√が繝ｼ繝励Φ

//   reader  := storage2.File_Open(w ,r ,bucket ,fname)

   reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,fname , w , r  )

   reader, _ := reader_minor.(io.ReadCloser)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

   defer reader.Close()

// 繝輔ぃ繧､繝ｫ繝ｪ繝ｼ繝繝ｼ繧抵ｼｧ・･・ｴ

   sreader := bufio.NewReaderSize(reader, 4096)

// 髱呎ｰｴ繝輔ぃ繧､繝ｫ縲√が繝ｼ繝励Φ

   writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,fname2 , w , r  )

   writer, _ := writer_minor.(*storage.Writer)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤


//   writer := storage2.File_Create ( w ,r ,bucket ,fname2 )

   index = 0        // 繝ｬ繧ｳ繝ｼ繝峨き繧ｦ繝ｳ繧ｿ繝ｼ繧段nitialize

   for {

      index ++     // 繝ｬ繧ｳ繝ｼ繝峨き繧ｦ繝ｳ繧ｿ繝ｼ繧偵き繧ｦ繝ｳ繝・
//      fmt.Fprintf (w, "hydrostatic-pressure2縲index " ,index)  // 繝・ヰ繝・け

// 繝輔ぃ繧､繝ｫ繧抵ｼ題｡罫ead

      line ,_  := sreader.ReadString('\n')

//譁・ｭ怜腰菴阪↓繧ｹ繝壹・繧ｹ縺ｧ蛻・牡

      str := strings.Fields(line)

      num = len(str)

//      fmt.Fprintf (w,"hydrostatic-pressure2縲num " ,num)  // 繝・ヰ繝・け

      if num == 0 {  //縲END縲繝√ぉ繝・け

//         fmt.Fprintf (w,"hydrostatic-pressure2 normal end")  // 繝・ヰ繝・け
         goto END
      }

      if index != 1{   // 隕句・縺嶺ｻ･螟悶ｒmake

         omega ,_ =strconv.ParseFloat(str[0],64)
         drad1 ,_ =strconv.ParseFloat(str[1],64)
         drad2 ,_ =strconv.ParseFloat(str[2],64)
         press1 ,_ =strconv.ParseFloat(str[3],64)

         high ,_ =strconv.ParseFloat(str[5],64)

//         fmt.Fprintf ( w,"hydrostatic-pressure2 file data " ,omega, drad1 , drad2 ,press1  ,high ) // 繝・ヰ繝・け

// U蟄礼ｮ｡縺ｮ髱｢遨阪ｒ險育ｮ励☆繧・
         area1 = sum.Circle_Area(drad1/2)
         area2 = sum.Circle_Area(drad2/2)

         press2 =  suiri.Seisui1( area1 ,area2  ,press1  ,omega  ,high  )

//         fmt.Fprintf(w,"hydrostatic-pressure2 蝨ｧ蜉幢ｼ偵",press2,"・・)   //繝・ヰ繝・け

// U蟄礼ｮ｡縺ｮ蜷・ｨｮ諠・ｱ繧呈枚蟄怜・縺ｫ螟画鋤

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






