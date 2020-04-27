package initialize3

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/storage2"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "cloud.google.com/go/storage"

                                        )

///                           　　　　　　　　　　　
///   write current water-inf. in temp.-file
///                          　　　　　　　　　　　

func Water2_temp( w http.ResponseWriter, r *http.Request ,water_inf type4.Water2_Temp ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     IN   water_inf   : slice of struct's format which is type4.Water2_Temp

//    fmt.Fprintf( w, "initialize3.water2_temp start \n" )

    bucket := "sample-7777"

    filename1 := "Water2_Temp.txt"

///
///    　delete existing "Water2_Temp.txt"
///

    storage2.File_Delete ( w , r  ,bucket  ,filename1  )

///
///    　　make　new "Water2_Temp.txt"
///

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,filename1 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)

//    writer := storage2.File_Create( w ,r ,bucket ,filename1 )

    defer writer.Close()

///
///    　　　write Water2_temp by format which is struct
///

    lf_flag := int64( 0 )
    water_inf.Id = int64( 1 )

    storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_inf )

    return

}
