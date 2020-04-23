package type1000

import (

//	    "google.golang.org/appengine/datastore"
//        "time"
//        "github.com/sawaq7/go12_ver1/client/sgh/type2"

                                                )
//
// struct information for sagawa express
//


// Book holds metadata about a book.

type Book_Test struct {
	ID            int64
	Title         string
	Author        string
	PublishedDate string
	ImageURL      string
	Description   string
	CreatedBy     string
	CreatedByID   string
	File_Name     string
}

//
// deliver district information　�E�地区惁E��・マルチ構造体バーション�E�E//

type D_District struct {               /// チE�Eタストア用

       Id               int64           //　
       District_No      int64           // 配達地域NO.
	   District_Name    string          // 配達地域名
//       D_Area_Slice   []type2.D_Area    //
       D_Area_Slice     []D_Area    //
     D_Area_Small_Slice []D_Area_Small  //
   }

//
// deliver area information　�E�エリア惁E���E�E//

type D_Area struct {               /// 常用ファイル用

       Id              int64           //
       Course_No       int64           // コースNO.
       District_No     int64           // 配達地域NO.
       District_Name   string          // 配達地域名
       Area_No         int64           // 配達エリアNO.
	   Area_Name       string          // 配達エリア
	   Area_Detail     string          // 配達エリアの詳細
	   Number_Total    int64           //
	   Time_Total      float64         //
	   Productivity    float64         //
       Car_No          int64            // レギュラー号

   }

type D_Area_Small struct {               /// 常用ファイル用

	   Area_Name       string          // 配達エリア吁E
	   Area_Small_Name string          // 配達エリアの詳細

   }

//
// deliver district information　

type D_District_2 struct {               ///

       Id               int64           //
       District_No      int64           // 配達地域NO.
	   District_Name    string          // 配達地域名
       D_Area_Slice     []D_Area_2    // エリア
//     D_Area_Small_Slice []D_Area_Small  // スモールエリア

   }

//
// deliver area information　

type D_Area_2 struct {               /// 常用ファイル用

       Id              int64           //
       Course_No       int64           // コースNO.
       District_No     int64           // 配達地域NO.
       District_Name   string          // 配達地域名
       Area_No         int64           // 配達エリアNO.
	   Area_Name       string          // 配達エリア吁E
	   Area_Detail     string          // 配達エリアの詳細
	   Number_Total    int64           //
	   Time_Total      float64         //
	   Productivity    float64         //
       Car_No          int64           // レギュラー号軁E
       D_Area_Small_Slice []D_Area_Small  // スモールエリア

   }
