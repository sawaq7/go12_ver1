package html4

   const Tokura_index_pl = `
   <!DOCTYPE html>
   <html>
	 <head>
		<meta charset="utf-8">
		<title>menu list</title>
		<link rel="stylesheet" href="css/sgh/tokura_index2.css" type="text/css">
	 </head>

	 <body>

	   <form method="GET"  action="/pipe_line1_excute_all" >
         <input type="submit" value="管水路(st-free-all)" />
       </form>

       <form method="GET"  action="/pipe_line1_show" >
         <input type="submit" value="管水路(st-free-single)" />
       </form>

       <form method="GET"  action="/pipe_line_ds_keyin" >
         <input type="submit" value="管水路(ds)" />
       </form>

       <form method="GET"  action="/pipe_line_st_keyin" >
         <input type="submit" value="管水路(st)" />
       </form>

       <form method="GET"  action="/water_slope_show" >
         <input type="submit" value="導水勾配線" />
       </form>

	</body>
  </html>
`

