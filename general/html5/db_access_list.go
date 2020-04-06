package html5

   const Db_access_list = `
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

              <tr> <th>basic-name</th> <th>new-file</th>  <th>access</th> </tr>

              <td > <input type="text" name="basic_name"  size="15" />               </td>

              <td > <input type="text" name="new_file" size="15"    />               </td>
              <td > <input type="submit" size="2"  value="登録"     />               </td>

            </form>

          </section>
          <section>
           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">copy list</h2>

           <tr> <th>db-type</th><th>access-type</th><th>basic-file-name</th> <th>new_file_name</th>
                                                    <th>access1</th><th>access2</th> </tr>
           {{range .}}

             <tr>

             <form method="GET" >

               <td>
                  <input type="text" name="db_type" size="5" align="center" value="{{.Db_Type|html}}" />
               </td>
               <td>
                  <input type="text" name="access_type" size="5" align="center" value="{{.Access_Type|html}}" />
               </td>

               <td>
                  <input type="text" name="basic_name" size="15" align="center" value="{{.Basic_File_Name|html}}" />
               </td>

               <td>
                  <input type="text" name="new_file" size="15" align="center" value="{{.New_File_Name|html}}" />
               </td>

             </form>
             <form method="GET" action="/datastore_copy_excute" >

               <td>
                  <input type="hidden" name="line_no"  value="{{.Line_No|html}}"/>
                  <input type="submit"  size="2" value="excute"  />
               </td>
             </form>

             <form method="GET" action="/datastore_copy_list_delete" >

               <td>
                  <input type="hidden" name="line_no"  value="{{.Line_No|html}}"/>
                  <input type="submit"  size="2" value="delete in list"  />
               </td>
             </form>


           </tr>

           {{end}}
           </table>




          </section>



     </body>
   </html>
   `
