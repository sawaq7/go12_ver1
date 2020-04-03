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
// deliver district information縲・亥慍蛹ｺ諠・ｱ繝ｻ繝槭Ν繝∵ｧ矩菴薙ヰ繝ｼ繧ｷ繝ｧ繝ｳ・・//

type D_District struct {               /// 繝・・繧ｿ繧ｹ繝医い逕ｨ

       Id               int64           //縲繝・・繧ｿid
       District_No      int64           // 驟埼＃蝨ｰ蝓櫻O.
	   District_Name    string          // 驟埼＃蝨ｰ蝓溷錐
//       D_Area_Slice   []type2.D_Area    // 繧ｨ繝ｪ繧｢諠・ｱ
       D_Area_Slice     []D_Area    // 繧ｨ繝ｪ繧｢諠・ｱ
     D_Area_Small_Slice []D_Area_Small  // 繧ｹ繝｢繝ｼ繝ｫ繧ｨ繝ｪ繧｢諠・ｱ

   }

//
// deliver area information縲・医お繝ｪ繧｢諠・ｱ・・//

type D_Area struct {               /// 蟶ｸ逕ｨ繝輔ぃ繧､繝ｫ逕ｨ

       Id              int64           // 繝・・繧ｿid
       Course_No       int64           // 繧ｳ繝ｼ繧ｹNO.
       District_No     int64           // 驟埼＃蝨ｰ蝓櫻O.
       District_Name   string          // 驟埼＃蝨ｰ蝓溷錐
       Area_No         int64           // 驟埼＃繧ｨ繝ｪ繧｢NO.
	   Area_Name       string          // 驟埼＃繧ｨ繝ｪ繧｢蜷・	   Area_Detail     string          // 驟埼＃繧ｨ繝ｪ繧｢縺ｮ隧ｳ邏ｰ
	   Number_Total    int64           // 螳・・驟埼＃邱乗焚
	   Time_Total      float64         // 螳・・驟埼＃邱乗凾髢・	   Productivity    float64         // 螳・・逕溽肇諤ｧ
       Car_No          int64           // 繝ｬ繧ｮ繝･繝ｩ繝ｼ蜿ｷ霆・//    D_Area_Small_Slice []D_Area_Small  // 繧ｹ繝｢繝ｼ繝ｫ繧ｨ繝ｪ繧｢諠・ｱ

   }

type D_Area_Small struct {               /// 蟶ｸ逕ｨ繝輔ぃ繧､繝ｫ逕ｨ

	   Area_Name       string          // 驟埼＃繧ｨ繝ｪ繧｢蜷・	   Area_Small_Name string          // 驟埼＃繧ｨ繝ｪ繧｢縺ｮ隧ｳ邏ｰ

   }

//
// deliver district information縲・亥慍蛹ｺ諠・ｱ繝ｻ繝槭Ν繝∵ｧ矩菴薙ヰ繝ｼ繧ｷ繝ｧ繝ｳ・・//

type D_District_2 struct {               /// 繝・・繧ｿ繧ｹ繝医い逕ｨ

       Id               int64           //縲繝・・繧ｿid
       District_No      int64           // 驟埼＃蝨ｰ蝓櫻O.
	   District_Name    string          // 驟埼＃蝨ｰ蝓溷錐
       D_Area_Slice     []D_Area_2    // 繧ｨ繝ｪ繧｢諠・ｱ
//     D_Area_Small_Slice []D_Area_Small  // 繧ｹ繝｢繝ｼ繝ｫ繧ｨ繝ｪ繧｢諠・ｱ

   }

//
// deliver area information縲・医お繝ｪ繧｢諠・ｱ・・//

type D_Area_2 struct {               /// 蟶ｸ逕ｨ繝輔ぃ繧､繝ｫ逕ｨ

       Id              int64           // 繝・・繧ｿid
       Course_No       int64           // 繧ｳ繝ｼ繧ｹNO.
       District_No     int64           // 驟埼＃蝨ｰ蝓櫻O.
       District_Name   string          // 驟埼＃蝨ｰ蝓溷錐
       Area_No         int64           // 驟埼＃繧ｨ繝ｪ繧｢NO.
	   Area_Name       string          // 驟埼＃繧ｨ繝ｪ繧｢蜷・	   Area_Detail     string          // 驟埼＃繧ｨ繝ｪ繧｢縺ｮ隧ｳ邏ｰ
	   Number_Total    int64           // 螳・・驟埼＃邱乗焚
	   Time_Total      float64         // 螳・・驟埼＃邱乗凾髢・	   Productivity    float64         // 螳・・逕溽肇諤ｧ
       Car_No          int64           // 繝ｬ繧ｮ繝･繝ｩ繝ｼ蜿ｷ霆・    D_Area_Small_Slice []D_Area_Small  // 繧ｹ繝｢繝ｼ繝ｫ繧ｨ繝ｪ繧｢諠・ｱ

   }
