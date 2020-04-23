package html1000

   const Queries = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>update&delete in datastore</title>
        <link rel="stylesheet" href="css/member1_show.css" type="text/css">
     </head>
     <body>
       <article>
       </article>
       <section id="main">


           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">List Of Deliverly Situation</h2>

           <tr> <th>district-no</th> <th>district-name</th> <th>area-no</th> <th>area-name</th><th>area-detail</th><th>access1</th> <th>access2</th> </tr>
           {{range .}}
             <tr>

             <form method="GET" action="/d_district_area_update" >


               <td>
                  <input type="text" name="district_no" size="5" align="center" value="{{.FirstName|html}}" />


               </td>
               <td>
                  <input type="text" name="district_no" size="5" align="center" value="{{.LastName|html}}" />


               </td>
               <td>
                  <input type="text" name="district_no" size="5" align="center" value="{{.Height|html}}" />


               </td>

             </form>

           </tr>
           {{end}}
           </table>

        </section>
        <section >
          <h2 align="center">Input New Data ,please</h2>
          <form method="GET" action="/queries_show" >

              <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
    /*          <h1  align="center">Register Your deliverly-number ,please</h1>   */
              <tr>   <th>area_name</th> <th>area-detail</th>  <th>access</th> </tr>
                                                                                                 </td>
              <td > <input type="text" name="first_name" size="10" />            </td>
               <td > <input type="text" name="last_name" size="10" />            </td>
              <td > <input type="text" name="height" size="10" />            </td>
              <td > <input type="submit" size="2"  value="登録" />           </td>
          </form>

        </section>

     </body>
   </html>
   `
