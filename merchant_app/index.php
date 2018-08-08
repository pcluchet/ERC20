<!doctype html>

<html lang="en">
<head>
  <meta charset="utf-8">

  <title>Merchant App</title>
  <meta name="description" content="The HTML5 Herald">
  <meta name="author" content="SitePoint">

  <link rel="stylesheet" href="css/style.css">

</head>

<body>
<script src="js/script.js"></script>
<?php include "parts/header.php" ?>
<div id="maincontainer">

    <div id="ongoinginvoices">
          <ul>
                        <li>Banana - $9</li>
                        <li>Banana - $9</li>
                        <li>Banana - $9</li>
                    </ul>
    <button onclick="customer_interface()">Open customer interface</button>
    </div>

    <div id="addinvoice">
        <div id="invoiceeditor">
            <div id="invoice">
                <div id="products">
                    <ul id="invoicelist">
                    </ul>
                </div>
                <div id="total">
                    Grand Total : <span id="grand">0</span>€
                </div>
            </div>

        <div id="editorcontrol">
            Product : 
            <select name="food" id="selectedproduct" >
               <option price="2" value="banana">Banana</option>
                <option price="5" value="sandwich">Sandwich</option>
                 <option price="2" value="icecream">Ice cream</option>
                  <option price="3" value="salad">Salad</option>
            </select> 
            <br>
            <button onclick="dje()">Add to invoice</button>
            <button onclick="resetInvoice()">Reset invoice</button>
        </div>
        </div>
        <a href="./add.php">
            <div id="submit">
                Create invoice
            </div>
        </a>
    </div>

</div>
</body>
</html>