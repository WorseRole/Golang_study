// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract InToRoman {

    function intToRoman(uint256 num) external pure returns(string memory) {
        // uint[] memory values = new uint[](13);
        // values[0] = 1000;
        // values[1] = 900;
        // values[2] = 500;
        // values[3] = 400;
        // values[4] = 100;
        // values[5] = 90;
        // values[6] = 50;
        // values[7] = 40;
        // values[8] = 10;
        // values[9] = 9;
        // values[10] = 5;
        // values[11] = 4;
        // values[12] = 1;
        // string[] memory symbols = new string[](13);
        // symbols[0] = "M";
        // symbols[1] = "CM";
        // symbols[2] = "D";
        // symbols[3] = "CD";
        // symbols[4] = "C";
        // symbols[5] = "XC";
        // symbols[6] = "L";
        // symbols[7] = "XL";
        // symbols[8] = "X";
        // symbols[9] = "IX";
        // symbols[10] = "V";
        // symbols[11] = "IV";
        // symbols[12] = "I";
        uint256[13] memory values = [uint(1000), 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
        string[13] memory symbols = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];
        bytes memory result;
        for (uint i=0; i<values.length; i++) {
            uint256 value = values[i];
            string memory symbol = symbols[i];
            while(num >= value) {
                num -= value;
                result = abi.encodePacked(result, symbol);
            }
            if(num == 0) {
                break;
            }
        }
        return string(result);
    }
}