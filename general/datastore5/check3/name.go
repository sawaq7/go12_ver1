package check3

import (

	    "net/http"
	    "fmt"
	    "errors"

                     )

///                          ///
/// ćEEćæć¹ćć¢ćć³ććEćć  ///
///                         ///

func Name( w http.ResponseWriter, basic_name string  ) (err error ){

//     IN    w        : ć¬ć¹ćć³ć¹ć©ć¤ćæć¼
//     IN  basic_Name : åŗę¬ć®ćEEćæć¹ćć¢åE

    fmt.Fprintf( w, "check3.name start \n" )  // ćEćEÆ
    fmt.Fprintf( w, "check3.name basic_name %v\n" ,basic_name)  // ćEćEÆ

///
///  ćØć©ć¼ć”ćE»ć¼ćøćć»ćE
///

var (

	  Err1 = errors.New("can't find this datastore's file (check3.name)")

	                                                                        )
///
/// ćEEćæć¹ćć¢åććć§ćEÆ
///
    ok_flag := 0

    switch basic_name {

      case "Deliver" :

        ok_flag = 1


      break;

    }

    if ok_flag == 0 {

		return Err1
	}

    fmt.Fprintf( w, "check3.name ok_flag %v\n" ,ok_flag)  // ćEćEÆ
	fmt.Fprintf( w, "check3.name normal end \n" )  // ćEćEÆ

    return nil
}

