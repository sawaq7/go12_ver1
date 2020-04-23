package struct_set

import (

	    "net/http"
//	    "fmt"
	    "strings"
	    "strconv"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
                                                )

///                           　　　　　　　　　　　
///    expand water-line data ( from line data(string) to struct (type4.Water_Line)
///                       　　　　　　　　　　　

func Water_line( w http.ResponseWriter, water_line_string string )  ( type4.Water_Line ) {

//     IN         w             : レスポンスライター
//     IN   water_line_string   : water-line(in)

//     OUT  water_line_struct   : water-line(out)

//    fmt.Fprintf( w, "struct_set.water_line start \n" )
//    fmt.Fprintf(w, "struct_set.water_line : water_line_string %s\n", water_line_string )

    var water_line_struct type4.Water_Line                 //   allocate work area

    str := strings.Fields(water_line_string)  //   split line data by brank

    for ii := 0 ; ii < len(str) ; ii++ {

      switch ii {

        case 0 :     //    Id
          str2 := strings.Split(str[ii], "{"  )     // trim mark "{"

//          fmt.Printf("struct_set: str2 %v\n", str2 )
//          fmt.Printf("struct_set: len(str2) %v\n", len(str2) )

          water_line_struct.Id ,_ =strconv.ParseInt(str2[1],10 ,64)

//          fmt.Printf("struct_set: water_line_struct.Id %v\n", water_line_struct.Id )

        break;

        case 1 :      //　water-name
          water_line_struct.Name = str[ii]

//          fmt.Printf("struct_set: water_line_struct.Name %v\n", water_line_struct.Name )

        break;

        case 2 :      //　water-section
          water_line_struct.Section = str[ii]

//          fmt.Printf("struct_set: water_line_struct.Section %v\n", water_line_struct.Section )

        break;

        case 3 :      //　friction_factor
          water_line_struct.Friction_Factor,_ =strconv.ParseFloat(str[ii],64)

//          fmt.Printf("struct_set: water_line_struct.High %v\n", water_line_struct.Friction_Factor )

        break;

        case 4 :      //　Velocity
          water_line_struct.Velocity,_ =strconv.ParseFloat(str[ii],64)

//          fmt.Printf("struct_set: water_line_struct.High %v\n", water_line_struct.Velocity )

        break;

        case 5 :      //　Pipe_Diameter
          water_line_struct.Pipe_Diameter,_ =strconv.ParseFloat(str[ii],64)

//          fmt.Printf("struct_set: water_line_struct.High %v\n", water_line_struct.Pipe_Diameter )

        break;

        case 6 :      //  Pipe_Length

          str3 := strings.Split(str[ii], "}"  )      // trim mark "{"
          water_line_struct.Pipe_Length,_ =strconv.ParseFloat(str3[0],64)

//         fmt.Printf("struct_set: water_line_struct.Roughness_Factor %v\n", water_line_struct.Pipe_Length )

        break;

     }

   }

   return water_line_struct

}
