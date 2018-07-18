'use strict';

const express = require('express');

// Constants
const PORT = 8080;
const HOST = '0.0.0.0';

// App

const request = {
  //targets : --- letting this default to the peers assigned to the channel
  chaincodeId: 'fabcar',
  fcn: 'get',
  args: ['a']
};


const app = express();
app.post('/query/', (req, res) => {


  var query = require('./query.1.js');
  query.cc_query('user1',request).then(
    (result) => {

  console.log(result);

  res.send(result);
    }
  );

});

//TODO
app.post('/invoke/', (req, res) => {

  res.send("{}");

});

//TODO
app.post('/listpeers/', (req, res) => {

  res.send("{}");

});



app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
