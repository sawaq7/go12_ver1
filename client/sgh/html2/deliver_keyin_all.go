package html2

   const Deliver_keyin_all = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>データストアの更新・削除</title>
        <link rel="stylesheet" href="css/member1_show.css" type="text/css">
     </head>
     <body>
       <article>
       </article>

          <section id="main">
            <form method="GET" action="/deliver_showall1" >

              <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
              <h1  align="center">Register Your deliverly-number ,please</h1>
              <tr> <th>deliverly-area</th> <th>car-no</th>  <th>private-no</th> <th>deliverly-no</th>  <th>access</th> </tr>
              <td > <input type="text" name="area_name"  size="10" />               </td>
              <td > <input type="text" name="car_no" size="5" />            </td>
              <td > <input type="text" name="private_no" size="5" />            </td>
              <td > <input type="text" name="number" size="5" />            </td>
              <td > <input type="submit" size="2"  value="登録" />           </td>
            </form>
          <section>
          <section>
          <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
              <h1  align="center">Register Your deliverly-number ,please</h1>
              <tr> <th>all</th> <th>car</th>  <th>private</th> </tr>

            <td >
              <form method="GET"  align="center" action="/deliver_showall2">
                <input type="submit" value="配達履歴（all)" />
              </form>
                                                                             </td>
            <td >
              <form method="GET" align="center" action="/deliver_keyin_car">
                <input type="submit" value="配達履歴（car)" />
              </form>
                                                                             </td>
            <td >
              <form method="GET" align="center" action="/deliver_keyin_private">
                <input type="submit" value="配達履歴（private)" />
              /form>
                                                                             </td>
           </section>



     </body>
   </html>
   `
