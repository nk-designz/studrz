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

function addStatistic(statistic) {
    axios({
        method: 'get',
        url: "http://localhost:42069/api/statistic/" + statistic.function + "/" + statistic.name,
        auth: Auth
    }).then( response => { 
        console.log(response) 
    }).catch(error => {
        console.log(error)
    })
}

function Statistic(props) {
    const newFunction  = useFormInput('');
    const newDrug      = useFormInput('');

    const [open, setOpen] = useState(false);
    const [statistics, setStatistics] = useState([])
    const openNewUserWindow = () => {
        setOpen(true);
    };
    const closeNewUserWindow = () => {
        setOpen(false);
    };
    const closeAndSend = () => {
        addStatistic({
            function: newFunction.value,
            name: newDrug.value,
        })
        closeNewUserWindow()
    }
    const listStatistics = () => {
        axios({
            method: 'get',
            url: "http://localhost:42069/api/statistic",
            auth: Auth
        }).then( response => { setStatistics(response.data); console.log(response) }).catch(error => {
            console.log(error)
        })
    }
    const columns = [
        { field: "id", headerName: 'ID', width: 70},
        { field: "key", headerName: 'Substance', width: 200 },
        { field: "calc", headerName: 'Function', width: 130 },
        { field: "value", headerName: 'Result', width: 70 },
        { field: "findings", headerName: 'Findings', width: 130 }
    ]
    return  <div className="users">
                <Button onClick={ ev => { listStatistics() }} >Update</Button>
                <Button onClick={ ev => { openNewUserWindow() }} >Add</Button>
                <Dialog open={open} onClose={closeNewUserWindow} aria-labelledby="form-dialog-title">
                    <DialogTitle id="form-dialog-title">Add new user</DialogTitle>
                    <DialogContent>
                        <TextField
                        id="function"
                        label="Function"
                        fullWidth
                        {...newFunction}
                        />
                        <TextField
                        id="drug"
                        label="Drug"
                        {...newDrug}
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
                <DataGrid rows={statistics} columns={columns} pageSize={20} />
            </div>
}

export default Statistic;