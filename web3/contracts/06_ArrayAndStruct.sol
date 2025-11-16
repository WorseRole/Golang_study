// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract ArrayAndStruct{
    // 数组array 和 结构体struct
    // 数组 用来存储一组数据（整数，字节，地址等等）。数组分为固定长度数组和可变长度数组 两种：
    
    // 固定长度数组：在声明时制定数组的长度。用T[k]的格式声明，其中T时元素的类型，k时长度，例如：
    uint[8] array1;
    bytes1[5] array2;
    address[100] array3;

    // 可变长度数组（动态数组）：在声明时不指定数组的长度。用T[]的格式声明，其中T是元素的类型，例如：
    uint[] array4;
    bytes1[] array5;
    address[] array6;
    bytes array7;
    // 注意bytes 比较特殊，是数组，但是不用加[]。另外，不能用byte[] 声明单字节数组，可以使用bytes 或bytes1[],bytes 比bytes1[] 省gas。


    // 创建数组的规则
    // 对于memory修饰的动态数组，可以用new操作符来创建，但是必须声明长度，并且声明后长度不能改变。例子：
    // memory动态数组
    // uint[] memory array8 = new uint[](5);
    // bytes memory array9 = new bytes(9);
    uint[] array8 = new uint[](5);
    bytes array9 = new bytes(9);

    // 数组字面常数(Array Literals)是写作表达式形式的数组，
    // 用方括号包着来初始化array的一种方式，并且里面每一个元素的type是以第一个元素为准的，
    // 例如[1,2,3]里面所有的元素都是uint8类型，因为在Solidity中，如果一个值没有指定type的话，
    // 会根据上下文推断出元素的类型，默认就是最小单位的type，这里默认最小单位类型是uint8。
    // 而[uint(1),2,3]里面的元素都是uint类型，因为第一个元素指定了是uint类型了，里面每一个元素的type都以第一个元素为准。

    // 给可变长度数组赋值
    function initArray() external pure returns(uint[] memory) {
        uint[] memory x = new uint[](3);
        x[0] = 1;
        x[1] = 2;
        x[2] = 3;
        return x;
    }
    // 数组成员
    // length: 数组有一个包含元素数量的length成员，memory数组的长度在创建后是固定的。
    // push(): 动态数组拥有push()成员，可以在数组最后添加一个0元素，并返回该元素的引用。
    // push(x): 动态数组拥有push(x)成员，可以在数组最后添加一个x元素。
    // pop(): 动态数组拥有pop()成员，可以移除数组最后一个元素。

    function arrayPush() public returns(uint[] memory){
        uint[2] memory a = [uint(1),2];
        array4 = a;
        array4.push(3);
        return array4;
    }
}

pragma solidity ^0.8.21;
contract StructType{
    // 结构体
    struct Student {
        uint256 id;
        uint256 score;
    }
    Student student; // 初始一个student结构体
    // 给结构体赋值
    // 方法1: 在函数中创建一个storage 的struct 引用
    function initStudent1 () external  {
        Student storage _student = student;
        _student.id = 11;
        _student.score = 100;
    }

    // 方法2: 直接饮用状态变量的struct
    function initStudent2() external {
        student.id = 1;
        student.score = 80;
    }

    // 方法3: 构造函数式
    function initStudent3() external {
        student = Student(3, 90);
    }

    // 方法4: key value
    function initStudent4() external {
        student = Student({id: 4, score: 60});
    }
}


pragma solidity ^0.8.21;
contract EnumType {
    // 将uint0， 1， 2表示为Buy， Hold， Sell
    enum ActionSet {Buy, Hold, Sell}
    // 创建enum 变量action
    ActionSet action = ActionSet.Buy;

    // enum 可以和uint 显示的转换
    function enumToUint() external view returns(uint) {
        return uint(action);
    }
}