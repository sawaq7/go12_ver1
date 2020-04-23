package html1000


   const Booklist = `
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
       <section>
          <form method="POST" enctype="multipart/form-data" action="/booklist2" >

            <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
            <h1  align="center">Register Your deliverly-number ,please</h1>
            <tr> <th>title</th> <th>image-file</th>  <th>access</th> </tr>

            <td >
              <input type="text" name="title" id="title" size="5" align="center"  />
            </td>
            <td >
              <input type="file" name="image" id="image" size="10" align="center">
            </td>

            <td >
              <input type="submit" size="2"  value="登録" />
            </td>
            </form>
       </section>
       <section id="main">


           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">List Of Deliverly Situation</h2>

           <tr> <th>title</th> <th>image-file</th> <th>access1</th> <th>access2</th>
                                                                             <th>access3</th> <th>access4</th> </tr>
           {{range .}}
             <tr>

             <form method="GET" enctype="multipart/form-data" action="/booklist2_url_select" >


               <td>
                  <input type="text" name="title" id="title" size="5" align="center" value="{{.Title|html}}" />
                  <input type="hidden" name="id"  value="{{.ID|html}}"/>

               </td>
               <td>
                  <input type="text" name="image" id="image" size="10" align="center"  value="{{.ImageURL}}" />


               </td>

               <td>
                  <input type="submit"  size="2" value="url-select"  />
               </td>
               <input type="hidden" name="imageURL" value="{{.ImageURL}}"/>

             </form>

             <form method="GET" action="/booklist2_file_select" >

               <td>
                  <input type="hidden" name="id"  value="{{.ID|html}}"/>
                  <input type="submit"  size="2" value="file_select"  />
               </td>
             </form>
             <form method="GET" action="/booklist2_file_make" >

               <td>
                  <input type="hidden" name="id"  value="{{.ID|html}}"/>
                  <input type="submit"  size="2" value="file_make"  />
               </td>
             </form>
             <form method="GET" action="/d_schedule_keyin" >
               <td>
                  <input type="hidden" name="id"  value="{{.ID|html}}"/>
                  <input type="submit"  size="2" value="スケジュール"  />
               </td>
             </form>
             <div>
             <form  >
               <td>
                 <img src="{{.ImageURL}}" width="100" height="100"      name="List" alt="png"/>
               </td>
             </form>


           </tr>
           {{end}}
           </table>

        </section>


     </body>
   </html>
   `

