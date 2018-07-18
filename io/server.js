'use strict';

const express = require('express');

// Constants
const PORT = 8080;
const HOST = '0.0.0.0';

const bodyParser = require("body-parser");


const app = express();


app.use(bodyParser.urlencoded({
  extended: true
}));

app.use(bodyParser.json());
// App


app.post('/query/', (req, res) => {


  var cc = req.body.chaincode;
  if (!cc)
  { 
    //TODO : handle error
  }

  var channel = req.body.channel;
  if (!channel)
  { 
    //TODO : handle error
  }

  var func = req.body.func;
  if (!func)
  { 
    //TODO : handle error
  }

  var args = req.body.args;
  if (!args)
  { 
    //TODO : handle error
  }




const request = {
  //targets : --- letting this default to the peers assigned to the channel
  chaincodeId: cc,
  fcn: func,
  args: JSON.parse(args)
};

  //console.log(req.body.param2);
  var query = require('./query.1.js');
  query.cc_query('user1',request, channel).then(
    (result) => {

  console.log(result);

  res.send(result);
    }
  );

});

//TODO
app.post('/invoke/', (req, res) => {

  var cc = req.body.chaincode;
  if (!cc)
  { 
    //TODO : handle error
  }

  var channel = req.body.channel;
  if (!channel)
  { 
    //TODO : handle error
  }

  var func = req.body.func;
  if (!func)
  { 
    //TODO : handle error
  }

  var args = req.body.args;
  if (!args)
  { 
    //TODO : handle error
  }

var request = {
  //targets : --- letting this default to the peers assigned to the channel
  chaincodeId: cc,
  fcn: func,
  args: JSON.parse(args),
  chainId: channel,
};

  //console.log(req.body.param2);
  var query = require('./invoke.1.js');
  query.cc_invoke('user1',request, channel).then(
    (result) => {

  console.log(result);

  res.send(result);
    }
  );


//  res.send("{}");

});

//TODO
app.post('/listpeers/', (req, res) => {

  res.send("{}");

});



app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
