import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';

import { DefaultApi } from '../../api/apis';
import { EntBook} from '../../api/models/EntBook';

const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});

export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [Books, setBooks] = useState<EntBook[]>(Array);
 const [loading, setLoading] = useState(true);
 
console.log(Books)

useEffect(() => {
    const getBooks = async () => {
      const res = await api.listBook({ limit: 10, offset: 0 });
      setLoading(false);
      setBooks(res);
      console.log(res);
    };
    getBooks();
  }, [loading]);
 
 const deleteBook = async (id: number) => {
    const res = await api.deleteBook({ id: id });
    setLoading(true);
   };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No.</TableCell>
           <TableCell align="center">เลขห้อง</TableCell>
           <TableCell align="center">วันเช็คอิน</TableCell>
           <TableCell align="center"></TableCell>         
          </TableRow>
       </TableHead>
       <TableBody>
       {Books === undefined 
          ? null
          : Books.map((item :any)=> (
           <TableRow>               
             <TableCell align="center">{item.id}</TableCell>                 
             <TableCell align="center">{item.edges?.bookroom?.rOOMNAME}</TableCell>             
             <TableCell align="center">{item.rESERVATIONS}</TableCell>     
             <TableCell align="center">
             <Button
                 onClick={() => {
                  deleteBook(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 ยกเลิก
               </Button>
                 </TableCell>                        
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
 
