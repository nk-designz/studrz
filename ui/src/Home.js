import Auth from './Auth';
import React from 'react'
import PropTypes from 'prop-types';
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';
import Users from './Users';
import Results from './Results';
import './Home.css'
import Consume from './Consume';

function a11yProps(index) {
    return {
      id: `simple-tab-${index}`,
      'aria-controls': `simple-tabpanel-${index}`,
    };
}

  function TabPanel(props) {
    const { children, value, index, ...other } = props;
    if(Auth.username == "" && Auth.password == "") {
      
    }
    return (
      <div
        role="tabpanel"
        hidden={value !== index}
        id={`simple-tabpanel-${index}`}
        aria-labelledby={`simple-tab-${index}`}
        {...other}
      >
        {value === index && (
          <Box p={3}>
            <Typography>{children}</Typography>
          </Box>
        )}
      </div>
    );
  }
  
  TabPanel.propTypes = {
    children: PropTypes.node,
    index: PropTypes.any.isRequired,
    value: PropTypes.any.isRequired,
  };
  

function Home(props) {
    const [value, setValue] = React.useState(0);

    const handleChange = (event, newValue) => {
        setValue(newValue);
    };
    return <div>
            <AppBar position="static">
                <Tabs value={value} onChange={handleChange} aria-label="simple tabs example">
                    <Tab label="Users" {...a11yProps(0)} />
                    <Tab label="Results" {...a11yProps(1)} />
                    <Tab label="Consumes" {...a11yProps(2)} />
                </Tabs>
            </AppBar>
            <TabPanel value={value} index={0}>
                <Users></Users>
            </TabPanel>
            <TabPanel value={value} index={1}>
                <Results></Results>
            </TabPanel>
            <TabPanel value={value} index={2}>
                <Consume></Consume>
            </TabPanel>
        </div>
}

export default Home;