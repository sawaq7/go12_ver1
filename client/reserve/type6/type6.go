package type6

import (

//          "time"

                  )
//
// the collection of struct information for reservation
//

//
// guest information　�E�地区惁E���E�E//

type Guest struct {            /// ゲスト情報

       Id            int64           //　チE�Eタid
       Guest_No      int64           // ゲスチEO.
	   Guest_Name    string          // ゲスト名


   }

type Guest_Temp struct {         /// 一時ファイル用

       Id          int64           //　チE�Eタid
       Guest_No    int64           // ゲスチEO.
	   Guest_Name  string          // ゲスト名


   }
//
// reservation information　
//

type Guest_Reserve_Minor struct {               ///  予紁E��報

       Id             int64          // チE�Eタid
       Line_No        int64          // 行NO.
       Date           string         // 予紁E��
       Guest_No       int64          // ゲスチEO.
	   Guest_Name     string         // ゲスト名
	   Start_Time     string         // 開始時閁E	   End_Time       string         // 終亁E��閁E
   }
//
// reservation information　
//

type Guest_Reserve_View struct {               ///  表示用予紁E��報

       Id             int64          // チE�Eタid
       Line_No        int64          // 行NO.
       Date           string         // 予紁E��
       Guest_No       int64          // ゲスチEO.
	   Guest_Name     string         // ゲスト名
	   Start_Time     string         // 開始時閁E	   End_Time       string         // 終亁E��閁E	   File_Name      string        // ファイル吁E	   Url            string        // url

   }
//
// reservation information　
//

type Guest_Medical_Record struct {   ///  カルチE��報

       Id             int64          // チE�Eタid
       Line_No        int64          // 行NO.
       Date           string         // 予紁E��
       Guest_No       int64          // ゲスチEO.
	   Guest_Name     string         // ゲスト名
	   Text           string         // 本斁E
   }

//
// reservation information　
//

type Guest_Medical_Xray struct {   ///  レントゲン写真

       Id             int64          // チE�Eタid
       Line_No        int64          // 行NO.
       Date           string         // 撮影日
       Guest_No       int64          // ゲスチEO.
	   Guest_Name     string         // ゲスト名
	   File_Name      string         // ファイル吁E	   Url            string         // url

	 }

//
// reservation information　
//

type Guest_Payment struct {               ///  予紁E��報

       Id             int64          // チE�Eタid
       Line_No        int64          // 行NO.
       Date           string         // 予紁E��
       Guest_No       int64          // ゲスチEO.
	   Guest_Name     string         // ゲスト名
	   Item           string         // 支払頁E��
	   Amount         int64          // 支払��顁E
   }
