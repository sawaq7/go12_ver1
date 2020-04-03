package type2

import (
          "time"
                  )
//
// the collection of struct information for sagawa express
//

//
// deliver district information縲・亥慍蛹ｺ諠・ｱ・・//

type D_District struct {               /// 繝・・繧ｿ繧ｹ繝医い逕ｨ

       Id               int64           //縲繝・・繧ｿid
       District_No      int64           // 驟埼＃蝨ｰ蝓櫻O.
	   District_Name    string          // 驟埼＃蝨ｰ蝓溷錐


   }

type D_District_Temp struct {         /// 荳譎ゅヵ繧｡繧､繝ｫ逕ｨ

       Id             int64           //縲繝・・繧ｿid
       District_No    int64           // 驟埼＃蝨ｰ蝓櫻O.
	   District_Name  string          // 驟埼＃蝨ｰ蝓溷錐


   }

type D_District_View struct {               /// 陦ｨ遉ｺ逕ｨ

       Id               int64           //縲繝・・繧ｿid
       District_No      int64           // 驟埼＃蝨ｰ蝓櫻O.
	   District_Name    string          // 驟埼＃蝨ｰ蝓溷錐
       D_Area_Slice     []D_Area        //縲繧ｨ繝ｪ繧｢諠・ｱ縺ｮ繧ｹ繝ｩ繧､繧ｹ

   }

//
// deliver area information縲・医お繝ｪ繧｢諠・ｱ・・//

type D_Area struct {               /// 蟶ｸ逕ｨ繝輔ぃ繧､繝ｫ逕ｨ

       Id             int64           // 繝・・繧ｿid
       Course_No      int64           // 繧ｳ繝ｼ繧ｹNO.
       District_No    int64           // 驟埼＃蝨ｰ蝓櫻O.
       District_Name  string          // 驟埼＃蝨ｰ蝓溷錐
       Area_No        int64           // 驟埼＃繧ｨ繝ｪ繧｢NO.
	   Area_Name      string          // 驟埼＃繧ｨ繝ｪ繧｢蜷・	   Area_Detail    string          // 驟埼＃繧ｨ繝ｪ繧｢縺ｮ隧ｳ邏ｰ
	   Number_Total   int64           // 螳・・驟埼＃邱乗焚
	   Time_Total     float64         // 螳・・驟埼＃邱乗凾髢・	   Productivity   float64         // 螳・・逕溽肇諤ｧ
       Car_No         int64           // 繝ｬ繧ｮ繝･繝ｩ繝ｼ蜿ｷ霆・   }

//
// deliver area information縲・医お繝ｪ繧｢諠・ｱ・・//

type D_Area_Temp struct {           /// 荳譎ゅヵ繧｡繧､繝ｫ逕ｨ

       Id           int64           // 繝・・繧ｿid
       Course_No    int64           // 繧ｳ繝ｼ繧ｹNO.
       District_No    int64         // 驟埼＃蝨ｰ蝓櫻O.
       District_Name  string          // 驟埼＃蝨ｰ蝓溷錐
       Area_No      int64           // 驟埼＃繧ｨ繝ｪ繧｢NO.
	   Area_Name    string          // 驟埼＃繧ｨ繝ｪ繧｢蜷・	   Area_Detail  string          // 驟埼＃繧ｨ繝ｪ繧｢縺ｮ隧ｳ邏ｰ
	   Number_Total   int64           // 螳・・驟埼＃邱乗焚
	   Time_Total     float64         // 螳・・驟埼＃邱乗凾髢・	   Productivity   float64         // 螳・・逕溽肇諤ｧ
       Car_No       int64           // 繝ｬ繧ｮ繝･繝ｩ繝ｼ蜿ｷ霆・   }

//
// deliver information縲・磯・驕費ｼ・//

type Deliver struct {

       Id           int64           // 繝・・繧ｿid
       Line_No      int64           // 陦君O.
       Course_No    int64           // 繧ｳ繝ｼ繧ｹNO.
       District_No    int64         // 驟埼＃蝨ｰ蝓櫻O.
       District_Name  string          // 驟埼＃蝨ｰ蝓溷錐
       Area_No      int64           // 驟埼＃繧ｨ繝ｪ繧｢NO.
	   Area_Name    string          // 驟埼＃繧ｨ繝ｪ繧｢蜷・       Date         string         // 驟埼＃譌･
       Date_Real    time.Time      // 驟埼＃譌･(繧ｿ繧､繝繝・・繧ｿ・・	   Car_No       int64           // 蜿ｷ霆・	   Private_No   int64           // 蛟倶ｺｺ逡ｪ蜿ｷ
	   Car_Division int64           // 霆贋ｸ｡蛹ｺ蛻・・撰ｼ夊ｻｽ閾ｪ蜍戊ｻ翫・托ｼ壹Ρ繧ｴ繝ｳ縲・抵ｼ壹ヨ繝ｩ繝・け
	   Number       int64           // 螳・・驟埼＃謨ｰ

   }


//
// deliver's Schedule information縲・磯・驕・繧ｹ繧ｱ繧ｸ繝･繝ｼ繝ｫ・・//

type Private struct {

       Id             int64           //縲繝・・繧ｿid
       Worker_No      int64           // 繝ｯ繝ｼ繧ｫ繝ｼNO.
       Worker_Name    string          // 繝ｯ繝ｼ繧ｫ繝ｼ蜷・       Worker_Type    string          // 繝ｯ繝ｼ繧ｫ繝ｼ繧ｿ繧､繝励縲SD : 繧ｻ繝ｼ繝ｫ繧ｹ繝峨Λ繧､繝舌・縲DD : 螳・・繝峨Λ繧､繝舌・
                                      // 縲縲縲縲縲縲縲縲縲OD : 螟匁ｳｨ繝峨Λ繧､繝舌・
       Worker_Salary  float64         // 繝ｯ繝ｼ繧ｫ繝ｼ繧ｵ繝ｩ繝ｪ繝ｼ・亥ｹｴ蜿趣ｼ・       Worker_Twh     float64         // 繝ｯ繝ｼ繧ｫ繝ｼ邱丞感蜒肴凾髢難ｼ亥ｹｴ髢難ｼ欝wh : total working hours
       Worker_H_Pay   float64         // 繝ｯ繝ｼ繧ｫ繝ｼ譎らｵｦ縲H_Pay : hourly縲pay
       Number_Total   int64           // 螳・・驟埼＃邱乗焚
	   Time_Total     float64         // 螳・・驟埼＃邱乗凾髢・	   Productivity   float64         // 蜉ｴ蜒咲函逕｣諤ｧ

   }


//
//   car information縲・亥捷霆頑ュ蝣ｱ・・//

type Car struct {

       Id                int64           //縲繝・・繧ｿid
       District_No       int64           // 驟埼＃蝨ｰ蝓櫻O.
       District_Name     string          // 驟埼＃蝨ｰ蝓溷錐
       Car_No            int64           // 蜿ｷ霆劾O.(繧ｷ繝ｼ繧ｱ繝ｳ繧ｷ繝｣繝ｫ縺ｪNO)
	   Car_Name          string          // 蜿ｷ霆雁錐
	   Car_Explain       string          // 蜿ｷ霆翫・襍ｰ陦後Ν繝ｼ繝育ｭ峨・隱ｬ譏・	   Number_Total      int64           // 蜿ｷ霆翫・螳・・驟埼＃邱乗焚
	   Time_Total        float64         // 螳・・驟埼＃邱乗凾髢・	   Productivity      float64         // 蜉ｴ蜒咲函逕｣諤ｧ

   }



//
// deliver's Schedule information縲・磯・驕・繧ｹ繧ｱ繧ｸ繝･繝ｼ繝ｫ・・//

type D_Schedule struct {

       Id             int64           //縲繝・・繧ｿid
       District_No    int64           // 驟埼＃蝨ｰ蝓櫻O.
       District_Name  string          // 驟埼＃蝨ｰ蝓溷錐
       Date           string          // 驟埼＃譌･
       Date_Real    time.Time         // 驟埼＃譌･(繧ｿ繧､繝繝・・繧ｿ・・       Expected_Num   float64         // 闕ｷ迚ｩ縺ｮ莠域Φ蛟区焚
       Judge          string          // 驟埼＃縺ｮ蛻､螳・       Course_Num     int64           // 繧ｳ繝ｼ繧ｹ謨ｰ
	   Course_01      string          // 繧ｳ繝ｼ繧ｹ1繧剃ｹ苓ｻ翫☆繧九ラ繝ｩ繧､繝舌・
	   Car_Name_01     string          // 繧ｳ繝ｼ繧ｹ1縺ｮ蜿ｷ霆雁錐
	   Course_02      string          // 繧ｳ繝ｼ繧ｹ2繧剃ｹ苓ｻ翫☆繧九ラ繝ｩ繧､繝舌・
	   Car_Name_02     string          // 繧ｳ繝ｼ繧ｹ1縺ｮ蜿ｷ霆雁錐
	   Course_03      string          // 繧ｳ繝ｼ繧ｹ3繧剃ｹ苓ｻ翫☆繧九ラ繝ｩ繧､繝舌・
	   Car_Name_03     string          // 繧ｳ繝ｼ繧ｹ1縺ｮ蜿ｷ霆雁錐
	   Course_04      string          // 繧ｳ繝ｼ繧ｹ4繧剃ｹ苓ｻ翫☆繧九ラ繝ｩ繧､繝舌・
	   Car_Name_04     string          // 繧ｳ繝ｼ繧ｹ1縺ｮ蜿ｷ霆雁錐
       Course_05      string          // 繧ｳ繝ｼ繧ｹ5繧剃ｹ苓ｻ翫☆繧九ラ繝ｩ繧､繝舌・
       Car_Name_05     string          // 繧ｳ繝ｼ繧ｹ1縺ｮ蜿ｷ霆雁錐
	   Course_06      string          // 繧ｳ繝ｼ繧ｹ6繧剃ｹ苓ｻ翫☆繧九ラ繝ｩ繧､繝舌・
	   Car_Name_06     string          // 繧ｳ繝ｼ繧ｹ1縺ｮ蜿ｷ霆雁錐
	   Course_07      string          // 繧ｳ繝ｼ繧ｹ7繧剃ｹ苓ｻ翫☆繧九ラ繝ｩ繧､繝舌・
	   Car_Name_07     string          // 繧ｳ繝ｼ繧ｹ1縺ｮ蜿ｷ霆雁錐
	   Course_08      string          // 繧ｳ繝ｼ繧ｹ8繧剃ｹ苓ｻ翫☆繧九ラ繝ｩ繧､繝舌・
	   Car_Name_08     string          // 繧ｳ繝ｼ繧ｹ1縺ｮ蜿ｷ霆雁錐
	   Course_09      string          // 繧ｳ繝ｼ繧ｹ9繧剃ｹ苓ｻ翫☆繧九ラ繝ｩ繧､繝舌・
	   Car_Name_09     string          // 繧ｳ繝ｼ繧ｹ1縺ｮ蜿ｷ霆雁錐
	   Course_10      string          // 繧ｳ繝ｼ繧ｹ10繧剃ｹ苓ｻ翫☆繧九ラ繝ｩ繧､繝舌・
	   Car_Name_10     string          // 繧ｳ繝ｼ繧ｹ1縺ｮ蜿ｷ霆雁錐

   }
//
// the collection of some condition expression縲・域擅莉ｶ蠑擾ｼ・//

type Sgh_Ai struct {

       Id              int64          // 繝・・繧ｿid
       Course_No       int64          // 繧ｳ繝ｼ繧ｹNO.
       District_No     int64          // 驟埼＃蝨ｰ蝓櫻O.
       District_Name   string         // 驟埼＃蝨ｰ蝓溷錐
       Area_No         int64          // 驟埼＃繧ｨ繝ｪ繧｢NO.
	   Area_Name       string         // 驟埼＃繧ｨ繝ｪ繧｢蜷・       Date_Basic      string         // 蝓ｺ貅匁律(髢｢謨ｰ縺ｮ蟋狗せ・・       Date_Basic_Real time.Time      // 蝓ｺ貅匁律(繧ｿ繧､繝繝・・繧ｿ・・       Ex_Type         string         // 譚｡莉ｶ蠑上・繧ｿ繧､繝・                                         // 1. function : 髢｢謨ｰ
       Expression    string         // 譚｡莉ｶ蠑・	   Item_Num      int64          // 鬆・焚
	   Item1_Name    string         // 鬆・錐1
	   Item1_Factor  float64        // 鬆・・菫よ焚1
	   Item2_Name    string         // 鬆・錐1
	   Item2_Factor  float64        // 鬆・・菫よ焚1
	   Item3_Name    string         // 鬆・錐1
	   Item3_Factor  float64        // 鬆・・菫よ焚1
	   Item4_Name    string         // 鬆・錐1
	   Item4_Factor  float64        // 鬆・・菫よ焚1
	   Item5_Name    string         // 鬆・錐1
	   Item5_Factor  float64        // 鬆・・菫よ焚1


   }



