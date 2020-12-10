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

function addResult(result) {
    axios({
        method: 'post',
        url: "http://localhost:42069/api/result",
        data: result,
        auth: Auth
    }).then( response => { 
        console.log(response) 
    }).catch(error => {
        console.log(error)
    })
}

function Results(props) {
    const newId          = useFormInput(0);
    const newUserId      = useFormInput(0);
    const newFillingDate = useFormInput('');
    const newScore       = useFormInput(0);

    const [open, setOpen] = useState(false);
    const [results, setResults] = useState([])
    const openNewUserWindow = () => {
        setOpen(true);
    };
    const closeNewUserWindow = () => {
        setOpen(false);
    };
    const closeAndSend = () => {
        addResult({
            id: parseInt(newId.value),
            user_id: parseInt(newUserId.value),
            score: parseInt(newScore.value)
        })
        closeNewUserWindow()
    }
    const listResults = () => {
        axios({
            method: 'get',
            url: "http://localhost:42069/api/result",
            auth: Auth
        }).then( response => { setResults(response.data); console.log(response) }).catch(error => {
            console.log(error)
        })
    }
    const columns = [
        { field: "id", headerName: 'ID', width: 70},
        { field: "user_id", headerName: 'User ID', width: 130 },
        { field: "score", headerName: 'Score', width: 130 },
        { field: "filling_date", headerName: 'Date', width: 200 }
    ]
    return  <div className="users">
                <Button onClick={ ev => { listResults() }} >Update</Button>
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
                        id="user_id"
                        label="User"
                        {...newUserId}
                        fullWidth
                        />
                        <TextField
                        id="filling_date"
                        label="Filling Date"
                        {...newFillingDate}
                        fullWidth
                        />
                        <TextField
                        id="score"
                        label="Score"
                        {...newScore}
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
                <DataGrid rows={results} columns={columns} pageSize={5} />
            </div>
}

export default Results;