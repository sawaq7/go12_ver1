package put1

import (

	    "net/http"
//	    "fmt"
	    "storage2"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "cloud.google.com/go/storage"

                                                )

///                           　　　　　　　　　　　
///   ストレチE��ファイルに水路ファイル惁E��を書ぁE水路ファイルがなぁE��合！E///                          　　　　　　　　　　　

func Water2_new( w http.ResponseWriter, r *http.Request ,water_inf type4.Water2 ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     IN   water_inf   : 水路惁E��の構造体　　struct : Water2

//    fmt.Fprintf( w, "put1.water2_new start \n" )  // チE��チE��

    bucket := "sample-7777"
    filename1 := "Water2.txt"

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,filename1 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)  // インターフェイス型を型変換

//    writer := storage2.File_Create( w ,r ,bucket ,filename1 )   // "Water2.txt"を新規作�E

    defer writer.Close()

    lf_flag := int64( 0 )

    water_inf.Id = int64( 1 )

    storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_inf )

    return

}


