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
import MenuItem from '@material-ui/core/MenuItem';

const getRoleText = role => {
    switch(role) {
        case 0:
            return "Admin";
        case 1:
            return "Reseacher";
        case 2:
            return "Subject";
        default:
            return "None";
    }
}

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

function addUser(user) {
    axios({
        method: 'post',
        url: "http://localhost:42069/api/user",
        data: user,
        auth: Auth
    }).then( response => { 
        console.log(response) 
    }).catch(error => {
        console.log(error)
    })
}

function Users(props) {
    const newUsername   = useFormInput('');
    const newPassword   = useFormInput('');
    const newRole       = useFormInput(2);
    const newBirthday   = useFormInput('');

    const [open, setOpen] = useState(false);
    const [users, setUsers] = useState([])
    const openNewUserWindow = () => {
        setOpen(true);
    };
    const closeNewUserWindow = () => {
        setOpen(false);
    };
    const closeAndSend = () => {
        addUser({
            username: newUsername.value,
            password: newPassword.value,
            birthday: newBirthday.value,
            role: newRole.value
        })
        closeNewUserWindow()
    }
    const listUsers = () => {
        axios({
            method: 'get',
            url: "http://localhost:42069/api/user",
            auth: Auth
        }).then( response => { setUsers(response.data); console.log(response) }).catch(error => {
            console.log(error)
        })
    }
    setInterval(listUsers, 10000);
    const columns = [
        { field: "id", headerName: 'ID', width: 70},
        { field: "username", headerName: 'Username', width: 130 },
        { field: "role", headerName: 'Role', width: 130 },
        { field: "birthday", headerName: 'Birthday', width: 200 }
    ]
    const roleUsers = users.map(u => {
        return { 
            id: u.id,
            username: u.username,
            role: getRoleText(u.role),
            birthday: u.birthday
        }})
    console.log(roleUsers)
    return  <div className="users">
                <Button onClick={ ev => { openNewUserWindow() }} >Add</Button>
                <Dialog open={open} onClose={closeNewUserWindow} aria-labelledby="form-dialog-title">
                    <DialogTitle id="form-dialog-title">Add new user</DialogTitle>
                    <DialogContent>
                        <TextField
                        id="name"
                        label="Username"
                        fullWidth
                        {...newUsername}
                        />
                        <TextField
                        id="password"
                        label="Password"
                        type="password"
                        {...newPassword}
                        fullWidth
                        />
                        <TextField
                        id="birthday"
                        label="Bithday"
                        {...newBirthday}
                        fullWidth
                        />
                        <Select id="role" labelId="Role" {...newRole}>
                            <MenuItem value={0}>Admin</MenuItem>
                            <MenuItem value={1}>Researcher</MenuItem>
                            <MenuItem value={2}>Subject</MenuItem>
                        </Select>
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
                <DataGrid rows={roleUsers} columns={columns} pageSize={5} />
            </div>
}

export default Users;