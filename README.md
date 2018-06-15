# ERC20

## Descriptiom
This project is an implementation of an ERC20 compliant tokens.  
It respect the technical standard ERC20 on Hyperledger Fabric and gives you the following method-related functions:
```
totalSupply [Get the total token supply]
balanceOf(address _owner) [Get the account balance of another account with address _owner]
transfer(address _to, uint256 _value) [Send _value amount of tokens to address _to]
transferFrom(address _from, address _to, uint256 _value) [Send _value amount of tokens from address _from to address _to]
approve(address _spender, uint256 _value) [Allow _spender to withdraw from your account, multiple times, up to the _value amount. If this function is called again it overwrites the current allowance with _value]
allowance(address *_owner*, address *_spender*) [Returns the amount which _spender is still allowed to withdraw from _owner]
```

## Author
Pierre Cluchet [pcluchet](https://github.com/pcluchet) 🐝

Sebastien Huertas [cactusfluo](https://gitlab.com/cactusfluo) 🦍

Jefferson Le Quellec [jle-quel](https://github.com/jle-quel) 🐜
