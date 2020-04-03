package vaccine1

import (
//	      "fmt"
	      "io"
	      "strings"
	      "net/http"
	      "bufio"

	      "storage2"

	      "cloud.google.com/go/storage"

                                         )
///
///   StoragePack : pack a file in Google Cloud Storage.
///

func File_Pack ( w http.ResponseWriter , r *http.Request ,bucket_name string ,file_name string ){

    file_name2 := file_name + "temp"

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket_name ,file_name2 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

	defer writer.Close()

// 豌ｴ霍ｯ諠・ｱ繝輔ぃ繧､繝ｫ縲・・ead file・・繧ｪ繝ｼ繝励Φ

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket_name ,file_name , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

    defer reader.Close()

// 繝輔ぃ繧､繝ｫ繝ｪ繝ｼ繝繝ｼ(string逕ｨ・峨ｒ・ｧ・･・ｴ

    sreader := bufio.NewReaderSize(reader, 4096)

    index := 0

    for {

      index ++     // 繝ｬ繧ｳ繝ｼ繝峨き繧ｦ繝ｳ繧ｿ繝ｼ繧偵き繧ｦ繝ｳ繝・

//      fmt.Fprintf(w, "File_Pack : lndex %v\n", index )  // 繝・ヰ繝・け

// 繝輔ぃ繧､繝ｫ繧抵ｼ題｡罫ead

      line ,err  := sreader.ReadString('\n')

      if err == io.EOF {

	      break

	  }

	  if err != nil {

		 http.Error(w, err.Error(), http.StatusInternalServerError)
	     return

	  }

//	  line = strings.Replace( line, ",", " ", -1)     /// 蛹ｺ蛻・ｊ譁・ｭ励ｒ螟画峩

      column := strings.Count( line ,",") + 1

//      fmt.Fprintf(w, "File_Pack : column %v\n", column )  // 繝・ヰ繝・け

      if  column > 1 {      //   繝ｬ繧ｳ繝ｼ繝峨′繧ｹ繝壹・繧ｹ縺ｧ縺ｪ縺・ヵ繧｡繧､繝ｫ縺ｫ譖ｸ縺崎ｾｼ縺ｿ

          line2 := strings.Trim(line, " ")           ///   荳｡遶ｯ繧ｹ繝壹・繧ｹ繧偵ヨ繝ｪ繝
//          fmt.Fprintf(w, "File_Pack :line2 [%s]\n", line2 )  // 繝・ヰ繝・け

          storage2.File_Write_Line ( w ,writer ,line2 )

      }

   }

///
/// 縲縲縲縲繝輔ぃ繧､繝ｫ蜷阪・螟画峩
///

   storage2.File_Delete ( w , r ,bucket_name ,file_name  )    //  譌ｧ繝輔ぃ繧､繝ｫ繧貞炎髯､

   storage2.File_Rename ( w , r  ,bucket_name ,file_name2 ,file_name ) //  譁ｰ繝輔ぃ繧､繝ｫ繧偵Μ繝阪・繝


//	fmt.Fprintf(w, " File_Pack : Calculate succeeded.\n" )

	return
}
