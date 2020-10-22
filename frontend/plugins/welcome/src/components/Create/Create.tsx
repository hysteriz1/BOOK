import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import {InputLabel,MenuItem,Button,TextField,FormControl,Select} from '@material-ui/core';
import { Alert } from '@material-ui/lab';
import React, { useState, useEffect } from 'react';

import Link from '@material-ui/core/Link';
import { Link as RouterLink } from 'react-router-dom';

import { DefaultApi } from'../../api/apis';

import { EntRoom } from '../../api/models/EntRoom';
const useStyles = makeStyles((theme: Theme) =>
    createStyles({
      button: {
        display: 'block',
        marginTop: theme.spacing(2),
      },
      formControl: {
          width: 500,
        },
        selectEmpty: {
          marginTop: theme.spacing(2),
        },
    }),
  );

export  default  function Create() {
    const classes = useStyles();
    const api = new DefaultApi();

    const [status, setStatus] = useState(false);
    const [alert, setAlert] = useState(true);
    const [rooms, setRooms] = useState<EntRoom[]>([]);    
    const [reservations, setRESERVATIONS] = useState(String);
    const [room, setRoom] = useState(Number);    
    const [loading, setLoading] = useState(true);

    useEffect(() => {
    
        const getRoom = async () => {
          const res = await api.listRoom({ limit: 3, offset: 0 });
          setLoading(false);
          setRooms(res);
        };
        getRoom();   
         
      }, [loading]);

//handle
const RESERVATIONShandleChange = (event: any) => {
    setRESERVATIONS(event.target.value as string);
  };

  const RoomhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setRoom(event.target.value as number);
  };  

  const createBooks = async () => {
    const book = {        
      reservations: reservations + ":00+07:00",
      bookStatus: 1,
      room: room,
      user: 1,

    };
    const r = {    
        id: room,   
        roomstatus: 2,  
      }
    console.log(book);    
    const res1: any = await api.updateRoom({ id:room,room: r});
    const res: any = await api.createBook({ book: book });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
  };

  return(
    <div>
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  บันทึกการจองสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    ไม่สามารถบันทึกการจองได้
                  </Alert>
                )}
            </div>
          ) : null}
           <FormControl className={classes.formControl}>
            <InputLabel>หมายเลขห้อง</InputLabel>
            <Select
                  name="room"
                  id="roomid"
                  value={room}
                  onChange={RoomhandleChange}
                  style={{ width: 200 }}
                >
                {rooms.map((item: any) => (
                  <MenuItem value={item.id}>{item.rOOMNAME}</MenuItem>
                ))}
                </Select>
  </FormControl> 
  <br/>
  <TextField
              id="reservations"
              label="วันเช็คอิน"
              type="datetime-local"
              value={reservations}
              onChange={RESERVATIONShandleChange}
              InputLabelProps={{
                shrink: true,
              }}
            />
  <FormControl >
  <br/>
  <br/>
  <br/>
  <br/>
  <Button 
            onClick={() => {
              createBooks();
            }}
            variant="contained"            
          >
            ยืนยัน
         </Button>
         <br/>
         <Link component={RouterLink} to="/">
           <Button variant="contained" color="primary">
          ย้อนกลับ     
           </Button></Link>
    </FormControl>
    </div>
);
}

