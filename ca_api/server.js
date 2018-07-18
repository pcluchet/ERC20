'use strict';

const express = require('express');

// Constants
const PORT = 8081;
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

  var caAddr = "http://localhost:7054";
  //console.log(req.body.param2);
  var query = require('./registerUser.1.js');
  query.ca_register(usr,caAddr).then(
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


  //var peerAddr = 'grpc://localhost:7051';
  var peerAddr = process.env.PEER_ADDR;
  console.log("peerAddr=",peerAddr);
  //var peerListenerAddr = 'grpc://localhost:7053';
  var peerListenerAddr = process.env.PEER_LISTENER_ADDR;
  console.log("peerListenerAddr=",peerListenerAddr);
  //var ordererAddr = 'grpc://localhost:7050';
  var ordererAddr = process.env.ORDERER_ADDR;
  console.log("ordererAddr=",ordererAddr);

  //console.log(req.body.param2);

  var query = require('./invoke.1.js');
  query.cc_invoke('user1',request, channel, peerAddr, ordererAddr, peerListenerAddr).then(
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
