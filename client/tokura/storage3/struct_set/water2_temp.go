package struct_set

import (

	    "net/http"
//	    "fmt"
	    "strings"
	    "strconv"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
                                                )

///                           　　　　　　　　　　　
///    expand water-line data ( from line data(string) to struct (type4.Water2_Temp)
///                      　　　　　　　　　　　

func Water2_temp( w http.ResponseWriter, water2_temp_string string )  ( type4.Water2_Temp ) {

//     IN         w             : レスポンスライター
//     IN   water2_string   : water(in)

//     OUT  water2_struct   : water(out)

//    fmt.Fprintf( w, "struct_set.water2_temp start \n" )
//    fmt.Fprintf(w, "struct_set.water2_temp : water2_temp_string %s\n", water2_temp_string )

    var water2_temp_struct type4.Water2_Temp   //   allocate work area

    str := strings.Fields(water2_temp_string)  //   split line data by brank

    for ii := 0 ; ii < len(str) ; ii++ {

      switch ii {

        case 0 :     //    Id
          str2 := strings.Split(str[ii], "{"  )     // trim mark "{"

//          fmt.Printf("struct_set: str2 %v\n", str2 )
//          fmt.Printf("struct_set: len(str2) %v\n", len(str2) )

          water2_temp_struct.Id ,_ =strconv.ParseInt(str2[1],10 ,64)

//          fmt.Printf("struct_set: water2_temp_struct.Id %v\n", water2_temp_struct.Id )

        break;

        case 1 :      //　water-name
          water2_temp_struct.Name = str[ii]

//          fmt.Printf("struct_set: water2_temp_struct.Name %v\n", water2_temp_struct.Name )

        break;

        case 2 :      //　water-high
          water2_temp_struct.High,_ =strconv.ParseFloat(str[ii],64)

//          fmt.Printf("struct_set: water2_temp_struct.High %v\n", water2_temp_struct.High )

        break;

        case 3 :      //　roughness_factor
          str3 := strings.Split(str[ii], "}"  )        // trim mark "{"
          water2_temp_struct.Roughness_Factor,_ =strconv.ParseFloat(str3[0],64)

//          fmt.Printf("struct_set: water2_temp_struct.Roughness_Factor %v\n", water2_temp_struct.Roughness_Factor )

        break;

     }

   }

   return water2_temp_struct

}

