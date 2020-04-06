package html5

   const Datastore_copy_list_keyin = `
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
            <form method="GET" action="/datastore_copy_list_show" >

              <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
              <h1  align="center">Register Your deliverly-number ,please</h1>

              <tr> <th>basic-name</th> <th>copy-file</th>  <th>new-file</th>  <th>access</th> </tr>

              <td > <input type="text" name="basic_name"  size="15" />               </td>
              <td > <input type="text" name="copy_file" size="15"   />               </td>
              <td > <input type="text" name="new_file" size="15"    />               </td>
              <td > <input type="submit" size="2"  value="登録"     />               </td>

            </form>

          </section>
          <section>
           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">copy list</h2>

           <tr> <th>basic-name</th><th>copy-file</th> <th>new-file</th> <th>access1</th> <th>access2</th> </tr>
           {{range .}}

             <tr>

             <form method="GET" >

               <td>
                  <input type="text" name="basic_name" size="15" align="center" value="{{.Basic_Name|html}}" />
               </td>
               <td>
                  <input type="text" name="copy_file" size="15" align="center" value="{{.Copy_Name|html}}" />
               </td>
               <td>
                  <input type="text" name="new_file" size="15" align="center" value="{{.New_Name|html}}" />
               </td>

             </form>

             <form method="GET" action="/datastore_copy_excute" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="copy"  />
               </td>
             </form>

             <form method="GET" action="/datastore_copy_list_delete" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="delete"  />
               </td>
             </form>


           </tr>

           {{end}}
           </table>




          </section>



     </body>
   </html>
   `
