package html2

   const D_district_keyin_all = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>datastore update/delete</title>
        <link rel="stylesheet" href="css/member1_show.css" type="text/css">
     </head>
     <body>
       <article>
       </article>
       <section id="main">

          <section>
            <form method="GET" action="/d_district_showall1" >

              <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
              <h1  align="center">Register Your deliverly-number ,please</h1>
              <tr> <th>district-no</th> <th>district-name</th>  <th>access</th> </tr>
              <td > <input type="text" name="district_no"  size="5" />               </td>
              <td > <input type="text" name="district_name" size="10" />            </td>

              <td > <input type="submit" size="2"  value="登録" />           </td>
            </form>
           </section>


        </section>
     </body>
   </html>
   `
