package main

import (

	    "net/http"
	    "fmt"
	    "log"
	    "os"

///
///  ---------------------   application all index   ----------------------------
///

	    "web1/appri_index"

///
///  ---------------------   general's soft index   ----------------------------
///

	    "web1/general2/general_index"

        ///
        ///  ---------------------   storage appri menu   ----------------------------
        ///

	       "web1/general2/db/storage/storage_bucket_list"

	       "web1/general2/db/storage/storage_object_list"

	       "web1/general2/db/storage/storage_object_copy_keyin"

	       "web1/general2/db/storage/storage_object_copy_excute"

	       "web1/general2/db/storage/storage_object_rename_keyin"

	       "web1/general2/db/storage/storage_object_rename_excute"

	       "web1/general2/db/storage/storage_object_show"

	       "web1/general2/db/storage/storage_object_delete"

	       "web1/general2/db/storage/storage_object_upload"

           ///
           ///  ---------------------   csv appri menu   ----------------------------
           ///

	       "web1/general2/db/csv/csv_show"

	       "web1/general2/db/csv/csv_show_test1"

	       "web1/general2/db/csv/csv_update"

	       "web1/general2/db/csv/csv_copy"

	       "web1/general2/db/csv/csv_delete"

	       "web1/general2/db/csv/csv_make"

	       "web1/general2/db/csv/csv_sort"

	       "web1/general2/db/csv/csv_column_join"

	       "web1/general2/db/csv/csv_column_delete"

	       "web1/general2/db/csv/csv_column_exchange"

	       "web1/general2/db/csv/csv_match_exp"

	       "web1/general2/db/csv/csv_match_exp2"

	       "web1/general2/db/csv/csv_match_wording"

	       "web1/general2/db/csv/csv_match_wording2"

           ///
           ///  ---------------------   datastore appri menu   ----------------------------
           ///

	       "web1/general2/db/datastore/datastore_copy_list_keyin"

	       "web1/general2/db/datastore/datastore_copy_list_show"

	       "web1/general2/db/datastore/datastore_copy_list_delete"

	       "web1/general2/db/datastore/datastore_copy_excute"

           ///
           ///  ---------------------  db_access_list appri menu   ----------------------------
           ///

	       "web1/general2/db/db_access_list"

///
///  ---------------------  sgh appri menu index   ----------------------------
///

	    "web1/sgh2/sgh_index"

        ///
	    ///  ---------------------  sgh appri/district・area menu   ---------------------------
        ///

	    "web1/sgh2/top/district/d_district_showall1"

	    "web1/sgh2/top/district/d_district_showall2"

	    "web1/sgh2/top/district/d_district_update"

	    "web1/sgh2/top/district/d_district_delete"

	    "web1/sgh2/top/district/d_district_area"

	    "web1/sgh2/top/district/d_district_area_show"

	    "web1/sgh2/top/district/d_district_area_update"

        "web1/sgh2/top/district/d_district_area_delete"

        ///
        ///  -------- sgh appri/deliver menu   ---------------------------
        ///

        "web1/sgh2/top/deliver/deliver_showall2"

        "web1/sgh2/top/deliver/deliver_showall1"

        "web1/sgh2/top/deliver/deliver_update"

        "web1/sgh2/top/deliver/deliver_delete"

        "web1/sgh2/top/deliver/deliver_copy"

        ///  -------- sgh appri/schedule menu   ---------------------------

        "web1/sgh2/top/schedule/d_schedule_keyin"

        "web1/sgh2/top/schedule/d_schedule_showall"

        "web1/sgh2/top/schedule/d_schedule_update"

        "web1/sgh2/top/schedule/d_schedule_delete"

        "web1/sgh2/top/schedule/d_schedule_copy"

        ///
        ///  -------- sgh appri/car menu   ---------------------------
        ///

        "web1/sgh2/top/car/car_show"

        "web1/sgh2/top/car/car_show2"

        "web1/sgh2/top/car/car_delete"

        "web1/sgh2/top/car/car_update"

        ///
        ///  -------- sgh appri/private menu   ---------------------------
        ///

        "web1/sgh2/top/private/private_showall1"

        "web1/sgh2/top/private/private_showall2"

        "web1/sgh2/top/private/private_update"

        "web1/sgh2/top/private/private_delete"

///
///  ---------------------  tokura appri menu index   ----------------------------
///
	    "web1/tokura2/tokura_index_top/tokura_index"

        ///
        ///  ---------------------  tokura hp appri menu index   ----------------------------
        ///

	    "web1/tokura2/tokura_index_top/tokura_index_hp"

           ///
	       ///  ---------------------  tokura hp appri menu   ----------------------------
           ///

	       "web1/tokura2/top/hydrostatic_pressure/hydrostatic_pressure1/hydrostatic_pressure1_excute"

	       "web1/tokura2/top/hydrostatic_pressure/hydrostatic_pressure1/hydrostatic_pressure1_show"

	       "web1/tokura2/top/hydrostatic_pressure/hydrostatic_pressure2"

        ///
	    ///  ---------------------  tokura pl appri menu index  ---------------------------
        ///

	    "web1/tokura2/tokura_index_top/tokura_index_pl"

           ///
           ///  ---------------------  tokura pl/storage-free appri menu   ----------------------------
           ///

	       "web1/tokura2/top/pipe_line/pipe_line1/pipe_line1_excute_all"

	       "web1/tokura2/top/pipe_line/pipe_line1/pipe_line1_excute_single"

	       "web1/tokura2/top/pipe_line/pipe_line1/pipe_line1_delete"

	       "web1/tokura2/top/pipe_line/pipe_line1/pipe_line1_show"

           ///
           ///  ---------------------  tokura pl/storage appri menu   ----------------------------
           ///

           "web1/tokura2/top/pipe_line/pipe_line_st/pipe_line_st_keyin"

           "web1/tokura2/top/pipe_line/pipe_line_st/pipe_line_st_show"

           "web1/tokura2/top/pipe_line/pipe_line_st/pipe_line_st_update"

           "web1/tokura2/top/pipe_line/pipe_line_st/pipe_line_st_delete"

           "web1/tokura2/top/pipe_line/pipe_line_st/pipe_line_st_wl_keyin"

           "web1/tokura2/top/pipe_line/pipe_line_st/pipe_line_st_wl_show"

           "web1/tokura2/top/pipe_line/pipe_line_st/pipe_line_st_wl_delete"

           "web1/tokura2/top/pipe_line/pipe_line_st/pipe_line_st_wl_update"

           "web1/tokura2/top/pipe_line/pipe_line_st/pipe_line_st_cal"

           ///
	       ///  ---------------------  tokura pl/datastore appri menu   ----------------------------
           ///

	       "web1/tokura2/top/pipe_line/pipe_line_ds/pipe_line_ds_keyin"

	       "web1/tokura2/top/pipe_line/pipe_line_ds/pipe_line_ds_show"

	       "web1/tokura2/top/pipe_line/pipe_line_ds/pipe_line_ds_update"

	       "web1/tokura2/top/pipe_line/pipe_line_ds/pipe_line_ds_delete"

	       "web1/tokura2/top/pipe_line/pipe_line_ds/pipe_line_ds_cal"

	       "web1/tokura2/top/pipe_line/pipe_line_ds/pipe_line_ds_wl_keyin"

	       "web1/tokura2/top/pipe_line/pipe_line_ds/pipe_line_ds_wl_show"

	       "web1/tokura2/top/pipe_line/pipe_line_ds/pipe_line_ds_wl_delete"

	       "web1/tokura2/top/pipe_line/pipe_line_ds/pipe_line_ds_wl_update"

           ///
	       ///  ---------------------  tokura pl/water_slope appri menu   ----------------------------
           ///

	       "web1/tokura2/top/pipe_line/water_slope/water_slope_show"

	       "web1/tokura2/top/pipe_line/water_slope/water_slope_delete"

	       "web1/tokura2/top/pipe_line/water_slope/water_slope_graf"
///
///  ---------------------  medical appri menu index   ----------------------------
///

	    "web1/reserve2/reserve_index"

        ///
	    ///  ---------------------  medical guest appri menu   ----------------------------
        ///

	    "web1/reserve2/top/guest/guest_show"

	    "web1/reserve2/top/guest/guest_show2"

        "web1/reserve2/top/guest/guest_delete"

        ///
	    ///  ---------------------  medical reserve appri menu   ----------------------------
        ///

        "web1/reserve2/top/reserve/reserve_register"

        "web1/reserve2/top/reserve/reserve_register_delete"

        "web1/reserve2/top/reserve/reserve_register_excute"

        "web1/reserve2/top/reserve/reserve_situation"

        "web1/reserve2/top/reserve/reserve_situation2"

        "web1/reserve2/top/reserve/reserve_situation3"

        ///
	    ///  ---------------------  medical medical-record appri menu   ----------------------------
        ///

        "web1/reserve2/top/medical_record/medical_record_show"

        "web1/reserve2/top/medical_record/medical_record_show2"

        "web1/reserve2/top/medical_record/medical_record_delete"

        ///
	    ///  ---------------------  medical xray appri menu   ----------------------------
        ///

        "web1/reserve2/top/payment/payment_register"

        "web1/reserve2/top/payment/payment_register_excute"

        "web1/reserve2/top/payment/payment_delete"

        ///
	    ///  ---------------------  medical xray appri menu   ----------------------------
        ///

        "web1/reserve2/top/medical_record/medical_xray_show"

        "web1/reserve2/top/medical_record/medical_xray_upload"

        "web1/reserve2/top/medical_record/medical_xray_delete"

///
///  ---------------------  sample appri menu index   ----------------------------
///

      "web1/sample1/sample_index"

      ///
	  ///  ---------------------  sample multi-struct appri menu   ----------------------------
      ///

      "web1/sample1/top/d_district_showall2_sample"

      "web1/sample1/top/d_district_showall2_sample2"

      "web1/sample1/top/d_district_showall2_sample3"


                                                                                   )
///
///     process list
///

func main() {

///
///  ---------------------   application main process  ----------------------------
///

  http.HandleFunc("/appri_index", appri_index.Appri_index)

   ///
   ///  ---------------------   general's soft index process   ----------------------------
   ///

   http.HandleFunc("/general_index", general_index.General_index)

     ///
     ///  ---------------------   storage appri process   ----------------------------
     ///

     http.HandleFunc("/storage_bucket_list", storage_bucket_list.Storage_bucket_list)

     http.HandleFunc("/storage_object_list", storage_object_list.Storage_object_list)

     http.HandleFunc("/storage_object_show", storage_object_show.Storage_object_show)

     http.HandleFunc("/storage_object_copy_keyin", storage_object_copy_keyin.Storage_object_copy_keyin)

     http.HandleFunc("/storage_object_copy_excute", storage_object_copy_excute.Storage_object_copy_excute)

     http.HandleFunc("/storage_object_rename_keyin", storage_object_rename_keyin.Storage_object_rename_keyin)

     http.HandleFunc("/storage_object_rename_excute", storage_object_rename_excute.Storage_object_rename_excute)

     http.HandleFunc("/storage_object_upload", storage_object_upload.Storage_object_upload)

     http.HandleFunc("/storage_object_delete", storage_object_delete.Storage_object_delete)

     ///
     ///  ---------------------   datastore appri process   ----------------------------
     ///

     http.HandleFunc("/datastore_copy_list_keyin", datastore_copy_list_keyin.Datastore_copy_list_keyin)

     http.HandleFunc("/datastore_copy_list_show", datastore_copy_list_show.Datastore_copy_list_show)

     http.HandleFunc("/datastore_copy_list_delete", datastore_copy_list_delete.Datastore_copy_list_delete)

     http.HandleFunc("/datastore_copy_excute", datastore_copy_excute.Datastore_copy_excute)

     ///
     ///  ---------------------   csv appri process   ----------------------------
     ///

     http.HandleFunc("/csv_show", csv_show.Csv_show)

     http.HandleFunc("/csv_show_test1", csv_show_test1.Csv_show_test1)

     http.HandleFunc("/csv_update", csv_update.Csv_update)

     http.HandleFunc("/csv_copy", csv_copy.Csv_copy)

     http.HandleFunc("/csv_delete", csv_delete.Csv_delete)

     http.HandleFunc("/csv_make", csv_make.Csv_make)

     http.HandleFunc("/csv_sort", csv_sort.Csv_sort)

     http.HandleFunc("/csv_column_join", csv_column_join.Csv_column_join)

     http.HandleFunc("/csv_column_delete", csv_column_delete.Csv_column_delete)

     http.HandleFunc("/csv_column_exchange", csv_column_exchange.Csv_column_exchange)

     http.HandleFunc("/csv_match_exp", csv_match_exp.Csv_match_exp)

     http.HandleFunc("/csv_match_exp2", csv_match_exp2.Csv_match_exp2)

     http.HandleFunc("/csv_match_wording", csv_match_wording.Csv_match_wording)

     http.HandleFunc("/csv_match_wording2", csv_match_wording2.Csv_match_wording2)

     ///
     ///  ---------------------  db_access_list appri process   ----------------------------
     ///

     http.HandleFunc("/db_access_list", db_access_list.Db_access_list)

  ///
  ///  ---------------------  tokura appri index process   ----------------------------
  ///

  http.HandleFunc("/tokura_index", tokura_index.Tokura_index)

    ///
    ///  ---------------------  tokura hp appri index process   ----------------------------
    ///

    http.HandleFunc("/tokura_index_hp", tokura_index_hp.Tokura_index_hp)

      ///
      ///  ---------------------  tokura hp appri index process   ----------------------------
      ///

      http.HandleFunc("/hydrostatic_pressure1_excute", hydrostatic_pressure1_excute.Hydrostatic_pressure1_excute)

      http.HandleFunc("/hydrostatic_pressure1_show", hydrostatic_pressure1_show.Hydrostatic_pressure1_show)

      http.HandleFunc("/hydrostatic_pressure2", hydrostatic_pressure2.Hydrostatic_pressure2)

   ///
   ///  ---------------------  tokura pl appri index  process   ----------------------------
   ///

    http.HandleFunc("/tokura_index_pl", tokura_index_pl.Tokura_index_pl)

      ///
      ///  ---------------------  tokura pl/datastore appri process   ----------------------------
      ///

      http.HandleFunc("/pipe_line_ds_keyin", pipe_line_ds_keyin.Pipe_line_ds_keyin)

      http.HandleFunc("/pipe_line_ds_show", pipe_line_ds_show.Pipe_line_ds_show)

      http.HandleFunc("/pipe_line_ds_update", pipe_line_ds_update.Pipe_line_ds_update)

      http.HandleFunc("/pipe_line_ds_delete", pipe_line_ds_delete.Pipe_line_ds_delete)

      http.HandleFunc("/pipe_line_ds_cal", pipe_line_ds_cal.Pipe_line_ds_cal)

      http.HandleFunc("/pipe_line_ds_wl_keyin", pipe_line_ds_wl_keyin.Pipe_line_ds_wl_keyin)

      http.HandleFunc("/pipe_line_ds_wl_show", pipe_line_ds_wl_show.Pipe_line_ds_wl_show)

      http.HandleFunc("/pipe_line_ds_wl_delete", pipe_line_ds_wl_delete.Pipe_line_ds_wl_delete)

      http.HandleFunc("/pipe_line_ds_wl_update", pipe_line_ds_wl_update.Pipe_line_ds_wl_update)

      ///
      ///  ---------------------  tokura pl/storage appri process   ----------------------------
      ///

      http.HandleFunc("/pipe_line_st_keyin", pipe_line_st_keyin.Pipe_line_st_keyin)

      http.HandleFunc("/pipe_line_st_show", pipe_line_st_show.Pipe_line_st_show)

      http.HandleFunc("/pipe_line_st_update", pipe_line_st_update.Pipe_line_st_update)

      http.HandleFunc("/pipe_line_st_delete", pipe_line_st_delete.Pipe_line_st_delete)

      http.HandleFunc("/pipe_line_st_wl_keyin", pipe_line_st_wl_keyin.Pipe_line_st_wl_keyin)

      http.HandleFunc("/pipe_line_st_wl_show", pipe_line_st_wl_show.Pipe_line_st_wl_show)

      http.HandleFunc("/pipe_line_st_wl_update", pipe_line_st_wl_update.Pipe_line_st_wl_update)

      http.HandleFunc("/pipe_line_st_wl_delete", pipe_line_st_wl_delete.Pipe_line_st_wl_delete)

      http.HandleFunc("/pipe_line_st_cal", pipe_line_st_cal.Pipe_line_st_cal)

      ///
      ///  ---------------------  tokura pl/storage-free appri process   ----------------------------
      ///

      http.HandleFunc("/pipe_line1_excute_all", pipe_line1_excute_all.Pipe_line1_excute_all)

      http.HandleFunc("/pipe_line1_excute_single", pipe_line1_excute_single.Pipe_line1_excute_single)

      http.HandleFunc("/pipe_line1_delete", pipe_line1_delete.Pipe_line1_delete)

      http.HandleFunc("/pipe_line1_show", pipe_line1_show.Pipe_line1_show)

      ///
      ///  ---------------------  tokura pl/water_slope appri process   ----------------------------
      ///

      http.HandleFunc("/water_slope_show", water_slope_show.Water_slope_show)

      http.HandleFunc("/water_slope_delete", water_slope_delete.Water_slope_delete)

      http.HandleFunc("/water_slope_graf", water_slope_graf.Water_slope_graf)

///
///  ---------------------  sgh appri index process   ----------------------------
///

  http.HandleFunc("/sgh_index", sgh_index.Sgh_index)

  ///
  ///  ---------------------  sgh appri/district・area process   ---------------------------
  ///

  http.HandleFunc("/d_district_showall1", d_district_showall1.D_district_showall1)

  http.HandleFunc("/d_district_showall2", d_district_showall2.D_district_showall2)

  http.HandleFunc("/d_district_update", d_district_update.D_district_update)

  http.HandleFunc("/d_district_delete", d_district_delete.D_district_delete)

  http.HandleFunc("/d_district_area", d_district_area.D_district_area)

  http.HandleFunc("/d_district_area_show", d_district_area_show.D_district_area_show)

  http.HandleFunc("/d_district_area_update", d_district_area_update.D_district_area_update)

  ///
  ///  -------- sgh appri/deliver process   ---------------------------
  ///

  http.HandleFunc("/d_district_area_delete", d_district_area_delete.D_district_area_delete)

  http.HandleFunc("/deliver_showall1", deliver_showall1.Deliver_showall1)

  http.HandleFunc("/deliver_showall2", deliver_showall2.Deliver_showall2)

  http.HandleFunc("/deliver_update", deliver_update.Deliver_update)

  http.HandleFunc("/deliver_delete", deliver_delete.Deliver_delete)

  http.HandleFunc("/deliver_copy", deliver_copy.Deliver_copy)

  ///
  ///  -------- sgh appri/schedule process   ---------------------------
  ///

  http.HandleFunc("/d_schedule_keyin", d_schedule_keyin.D_schedule_keyin)

  http.HandleFunc("/d_schedule_showall", d_schedule_showall.D_schedule_showall)

  http.HandleFunc("/d_schedule_update", d_schedule_update.D_schedule_update)

  http.HandleFunc("/d_schedule_delete", d_schedule_delete.D_schedule_delete)

  http.HandleFunc("/d_schedule_copy", d_schedule_copy.D_schedule_copy)

  ///
  ///  -------- sgh appri/car process   ---------------------------
  ///

  http.HandleFunc("/car_show", car_show.Car_show)

  http.HandleFunc("/car_show2", car_show2.Car_show2)

  http.HandleFunc("/car_delete", car_delete.Car_delete)

  http.HandleFunc("/car_update", car_update.Car_update)

  http.HandleFunc("/private_showall1", private_showall1.Private_showall1)

  http.HandleFunc("/private_showall2", private_showall2.Private_showall2)

  http.HandleFunc("/private_update", private_update.Private_update)

  http.HandleFunc("/private_delete", private_delete.Private_delete)

///
///  ---------------------   application medical process  ----------------------------
///

  http.HandleFunc("/reserve_index", reserve_index.Reserve_index)

  ///
  ///  ---------------------   application medical/guest process  ----------------------------
  ///

  http.HandleFunc("/guest_show", guest_show.Guest_show)

  http.HandleFunc("/guest_show2", guest_show2.Guest_show2)

  http.HandleFunc("/guest_delete", guest_delete.Guest_delete)

  ///
  ///  ---------------------   application medical/reserve process  ----------------------------
  ///

  http.HandleFunc("/reserve_register", reserve_register.Reserve_register)

  http.HandleFunc("/reserve_register_excute", reserve_register_excute.Reserve_register_excute)

  http.HandleFunc("/reserve_register_delete", reserve_register_delete.Reserve_register_delete)

  http.HandleFunc("/reserve_situation", reserve_situation.Reserve_situation)

  http.HandleFunc("/reserve_situation2", reserve_situation2.Reserve_situation2)

  http.HandleFunc("/reserve_situation3", reserve_situation3.Reserve_situation3)

  ///
  ///  ---------------------   application medical/medical_record process  ----------------------------
  ///

  http.HandleFunc("/medical_record_show", medical_record_show.Medical_record_show)

  http.HandleFunc("/medical_record_show2", medical_record_show2.Medical_record_show2)

  http.HandleFunc("/medical_record_delete", medical_record_delete.Medical_record_delete)

  ///
  ///  ---------------------   application medical/xray process  ----------------------------
  ///

  http.HandleFunc("/medical_xray_show", medical_xray_show.Medical_xray_show)

  http.HandleFunc("/medical_xray_delete", medical_xray_delete.Medical_xray_delete)

  http.HandleFunc("/medical_xray_upload", medical_xray_upload.Medical_xray_upload)

  ///
  ///  ---------------------   application medical/payment process  ----------------------------
  ///

  http.HandleFunc("/payment_register", payment_register.Payment_register)

  http.HandleFunc("/payment_register_excute", payment_register_excute.Payment_register_excute)

  http.HandleFunc("/payment_delete", payment_delete.Payment_delete)

///
///  ---------------------   application sample process  ----------------------------
///

   http.HandleFunc("/sample_index", sample_index.Sample_index)

   ///
   ///  ---------------------   application medical/xray process  ----------------------------
   ///

   http.HandleFunc("/d_district_showall2_sample", d_district_showall2_sample.D_district_showall2_sample)

   http.HandleFunc("/d_district_showall2_sample2", d_district_showall2_sample2.D_district_showall2_sample2)

   http.HandleFunc("/d_district_showall2_sample3", d_district_showall2_sample3.D_district_showall2_sample3)

///
///       ポートをセット
///

  port := os.Getenv("PORT")

  if port == "" {
        port = "8080"
        log.Printf("Defaulting to port %s", port)
  }

  log.Printf("Listening on port %s", port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}
