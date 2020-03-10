package html2

   const Deliver_keyin_private = `
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
          <section>
            <form method="GET" action="/deliver_showprivate" >

              <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
              <h1  align="center">Register Your deliverly-number ,please</h1>
              <tr> <th>private_no</th> <th>access</th> </tr>

              <td > <input type="text" name="private_no" size="5" />            </td>
              <td > <input type="submit" size="2"  value="表示" />           </td>
            </form>
           </section>


        </section>
     </body>
   </html>
   `
