// SPDX-License-Identifier: MIT
pragma solidity ^0.6.8;

contract BaseERC20 {
    string public name; 
    string public symbol; 
    uint8 public decimals; 

    uint256 public totalSupply; 

    mapping (address => uint256) balances; 

    mapping (address => mapping (address => uint256)) allowances; 

    mapping (address => mapping (string => uint256)) public item; 
    

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);

    constructor() public{
        name = "MyToken"; 
        symbol = "MTK"; 
        decimals = 18; 
        totalSupply = 100000000 * 10 ** uint256(decimals);

        balances[msg.sender] = totalSupply;  
    }

    function balanceOf(address _owner) public view returns (uint256 balance) {
        return balances[_owner];    
    }

    function transfer(address _to, uint256 _value) public returns (bool success) {
        require(balances[msg.sender] >= _value, "ERC20: transfer amount exceeds balance");

        balances[msg.sender] -= _value;    
        balances[_to] += _value;   

        emit Transfer(msg.sender, _to, _value);  
        return true;   
    }

    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
        require(balances[_from] >= _value, "ERC20: transfer amount exceeds balance");
        require(allowances[_from][msg.sender] >= _value,"ERC20: transfer amount exceeds allowance");

        balances[_from] -= _value; 
        balances[_to] += _value; 

        allowances[_from][msg.sender] -= _value;
        
        emit Transfer(_from, _to, _value); 
        return true; 
    }

    function approve(address _spender, uint256 _value) public returns (bool success) {
        allowances[msg.sender][_spender] = _value; 
        emit Approval(msg.sender, _spender, _value); 
        return true; 
    }

    function allowance(address _owner, address _spender) public view returns (uint256 remaining) {
        return allowances[_owner][_spender];
    }

    function setitem(string calldata key,uint256 val) external  {
    
            item[msg.sender][key]=val;
     
    }

    function getitem(string memory key)public view returns(uint256){
        return item[msg.sender][key];
    }
  
    
}