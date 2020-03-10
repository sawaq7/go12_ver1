package html5

   const Csv_match_exp = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>データストアの更新・削除</title>

        <link rel="stylesheet" href="css/general/csv_show.css" type="text/css">
     </head>
     <body>
       <header>


       </header>
       <nav>

       </nav>
       <article>
         <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="left">column's command list</h2>

           {{range .}}
             {{ if eq .Line_No 1}}
              <tr>  <th>column-no</th> <th>expression</th> </tr>
              <form method="GET" action="/csv_match_exp2" >
              <tr>

               <th>
                  <input type="text" name="column_no1" size="3" align="center"   />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>
               <th>
                  <input type="text" name="expression1" size="15" align="center"  />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>
              </tr>

              <tr>

               <th>
                 <input type="text" name="column_no2" size="3" align="center"   />
                 <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>

               <th>
                 <input type="text" name="expression2" size="15" align="center"  />
                 <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>
              </tr>

              <tr>
                <th>
                 <input type="text" name="column_no3" size="3" align="center"   />
                 <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>

               <th>
                 <input type="text" name="expression3" size="15" align="center"  />
                 <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>
              </tr>

              <tr>
                <th>
                 <input type="text" name="column_no4" size="3" align="center"   />
                 <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>

               <th>
                 <input type="text" name="expression4" size="15" align="center"  />
                 <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>

              </tr>

              <tr>
                <th>
                 <input type="text" name="column_no5" size="3" align="center"   />
                 <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>

               <th>
                 <input type="text" name="expression5" size="15" align="center"  />
                 <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>

              </tr>

              <tr>
                <th>
                 <input type="text" name="column_no6" size="3" align="center"   />
                 <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>

               <th>
                 <input type="text" name="expression6" size="15" align="center"  />
                 <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>


              </tr>

              <tr>
                <th>
                 <input type="text" name="column_no7" size="3" align="center"   />
                 <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>

               <th>
                 <input type="text" name="expression7" size="15" align="center"  />
                 <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>


              </tr>

              <tr>
                <th>
                 <input type="text" name="column_no8" size="3" align="center"   />
                 <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>

               <th>
                 <input type="text" name="expression8" size="15" align="center"  />
                 <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </th>
              <tr>
              </tr>
               <th>
                <input type="radio" name="logical" value="and" checked="checked" >and
                <input type="radio" name="logical" value="or" >or
               </th>

               <th>
                <input type="submit"  size="2" value="excute"  />
               </th>
              </tr>
             </form>
             {{end}}
           {{end}}

          </table>

       </article>

       <section id="main">

           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">Csv's List Information</h2>

           <tr> <th>line-no</th>  <th>column1</th> <th>column2</th> <th>column3</th><th>column4</th> <th>column5</th></tr>
           {{range .}}
             <tr>

             <form method="GET" action="/csv_update" >
               <td>
                  <input type="text"  size="5" align="center" value="{{.Line_No|html}}" />

                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>
               <td>
                  <input type="text" name="column1" size="10" align="center" value="{{.Column1|html}}" />

                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>
               <td>
                  <input type="text" name="column2" size="10" align="center" value="{{.Column2|html}}" />

                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>
               <td>
                  <input type="text" name="column3" size="10" align="center" value="{{.Column3|html}}" />

                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>
               <td>
                  <input type="text" name="column4" size="10" align="center" value="{{.Column4|html}}" />

                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>
               <td>
                  <input type="text" name="column5" size="10" align="center" value="{{.Column5|html}}" />

                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>

               <td>
                  <input type="submit"  size="2" value="update"  />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>


             </form>

             <form method="GET" action="/csv_delete" >
               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="delete"  />
               </td>
             </form>

             <form method="GET" action="/csv_copy" >
               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="copy"  />
               </td>
             </form>

             <form method="GET" action="/storage_object_rename_keyin" >
               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="rename"  />
               </td>
             </form>

             <form method="GET" action="/csv_show" >
               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="csv access"  />
               </td>
             </form>

           </tr>
           {{end}}
           </table>

        </section>
        <aside>

        </aside>

        <footer>

        </footer>

     </body>
   </html>
   `
