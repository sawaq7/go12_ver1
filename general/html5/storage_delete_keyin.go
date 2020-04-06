package html5

   const Storage_delete_keyin = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>datastore's update/delete</title>
        <link rel="stylesheet" href="css/member1_show.css" type="text/css">
     </head>
     <body>
       <article>
       </article>

          <section id="main">
            <form method="GET" action="/storage_delete_excute" >

              <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
              <h1  align="center">key-in file-name & bucket-name which you want to delete ,please</h1>
              <tr> <th>bucket-name</th> <th>file-name</th>    <th>access</th> </tr>
              <td > <input type="text" name="bucket_name"  size="10" />               </td>
              <td > <input type="text" name="file_name" size="10" />            </td>
              <td > <input type="submit" size="4"  value="delete-start" />           </td>
            </form>
          <section>


     </body>
   </html>
   `
