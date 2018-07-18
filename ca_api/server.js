'use strict';

const express = require('express');

// Constants
const PORT = 8084;
const HOST = '0.0.0.0';

const bodyParser = require("body-parser");


const app = express();


app.use(bodyParser.urlencoded({
  extended: true
}));

app.use(bodyParser.json());
// App


app.post('/register/', (req, res) => {


  var usr = req.body.username;
  if (!usr)
  { 
    //TODO : handle error
  }

//  var caAddr = "http://localhost:7054";

  var caAddr = process.env.CA_ADDR;
  console.log("caAddr=", caAddr);
  //console.log(req.body.param2);
  var query = require('./registerUser.1.js');
  query.ca_register(usr, caAddr).then(
    (result) => {

  console.log(result);

  res.send(result);
    }
  );

});



app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
