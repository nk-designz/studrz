import { Button, Select } from '@material-ui/core';
import { DataGrid } from '@material-ui/data-grid';
import axios from 'axios';
import { useState } from 'react';
import Auth from './Auth';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '@material-ui/core/DialogTitle';

const useFormInput = initialValue => {
    const [value, setValue] = useState(initialValue);
   
    const handleChange = e => {
      setValue(e.target.value);
    }
    return {
      value,
      onChange: handleChange
    }
  }

function addConsume(consume) {
    axios({
        method: 'post',
        url: "http://localhost:42069/api/consume",
        data: consume,
        auth: Auth
    }).then( response => { 
        console.log(response) 
    }).catch(error => {
        console.log(error)
    })
}

function Consume(props) {
    const newId          = useFormInput(0);
    const newDrug      = useFormInput('');
    const newRate = useFormInput(0);
    const newConsumeId = useFormInput(0);

    const [open, setOpen] = useState(false);
    const [consumes, setConsumes] = useState([])
    const openNewUserWindow = () => {
        setOpen(true);
    };
    const closeNewUserWindow = () => {
        setOpen(false);
    };
    const closeAndSend = () => {
        addConsume({
            id: parseInt(newId.value),
            name: newDrug.value,
            rate: parseInt(newRate.value),
            questionair_result_id:  parseInt(newConsumeId.value)
        })
        closeNewUserWindow()
    }
    const listConsumes = () => {
        axios({
            method: 'get',
            url: "http://localhost:42069/api/consume",
            auth: Auth
        }).then( response => { setConsumes(response.data); console.log(response) }).catch(error => {
            console.log(error)
        })
    }
    const columns = [
        { field: "id", headerName: 'ID', width: 70},
        { field: "name", headerName: 'User ID', width: 130 },
        { field: "rate", headerName: 'Score', width: 130 },
        { field: "questionair_result_id", headerName: 'Result ID', width: 130 }
    ]
    return  <div className="users">
                <Button onClick={ ev => { listConsumes() }} >Update</Button>
                <Button onClick={ ev => { openNewUserWindow() }} >Add</Button>
                <Dialog open={open} onClose={closeNewUserWindow} aria-labelledby="form-dialog-title">
                    <DialogTitle id="form-dialog-title">Add new user</DialogTitle>
                    <DialogContent>
                        <TextField
                        id="id"
                        label="ID"
                        fullWidth
                        {...newId}
                        />
                        <TextField
                        id="drug"
                        label="Drug"
                        {...newDrug}
                        fullWidth
                        />
                        <TextField
                        id="rate"
                        label="Rate per Week"
                        {...newRate}
                        fullWidth
                        />
                        <TextField
                        id="result"
                        label="Result ID"
                        {...newConsumeId}
                        fullWidth
                        />
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={closeNewUserWindow} color="secondary">
                        Cancel
                        </Button>
                        <Button onClick={closeAndSend} color="primary">
                            Add
                        </Button>
                    </DialogActions>
                </Dialog>
                <DataGrid rows={consumes} columns={columns} pageSize={5} />
            </div>
}

export default Consume;