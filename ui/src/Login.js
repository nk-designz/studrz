import React, { useState } from 'react';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import axios from 'axios';
import Auth from './Auth';
import './Login.css';
 
function Login(props) {
  const username = useFormInput('');
  const password = useFormInput('');
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(false);
 
  // handle button click of login form
  const handleLogin = () => {
    setError(null);
    setLoading(true);
    axios({
      method:'get',
      url: "http://localhost:42069/",
      auth: {
          username: username.value,
          password: password.value
      }}).then(response => {
        Auth.username = username.value
        Auth.password = password.value
        setLoading(false);
        props.history.push('/home');
      }).catch(error => {
        setLoading(false);
        console.log(error)
        if (error.response.status === 401) setError(error.response.data.message);
        else setError("Something went wrong. Please try again later.");
      });
  }
 
  return (
    <Card className="login">
      <h1>Login</h1><br />
      <div>
        <TextField label="Username" {...username} label="Standard" autoComplete="new-password"/>
      </div>
      <div style={{ marginTop: 10 }}>
        <TextField label="Password" type="password" {...password} autoComplete="new-password" />
      </div>
      {error && <><small style={{ color: 'red' }}>{error}</small><br /></>}<br />
      <Button value={loading ? 'Loading...' : 'Login'} onClick={handleLogin} disabled={loading} >Login</Button><br />
    </Card>
  );
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
 
export default Login;