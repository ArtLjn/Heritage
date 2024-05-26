// SPDX-License-Identifier: UNLICENSED

/*
@Time : 2024/5/23 上午9:49 
@Author : ljn
@File : Heritage
@Software: GoLand
*/
pragma solidity ^0.6.0;
pragma experimental ABIEncoderV2;

import "./data.sol";
// 遗产信息一旦写入就不可以被修改，确保遗产的安全性
contract Heritage is data{

    // createHeritageInheritor
    // 创建非遗人项目
    // exam: 张三,zhang_san_img.jpg,糖画,6,1,糖画......., 江苏省
    function createHeritageInheritor(string memory Inheritor,string memory InheritorImg,
        string memory project,CateGory category,Level rank,string memory details,
        string memory locate
    ) public returns(string memory)  {
        string memory strNumber = generateRandomString(7);
        HeritageInheritor memory heritageInheritor = HeritageInheritor(
            strNumber
            ,Inheritor
            ,InheritorImg
            ,project
            ,block.timestamp
            ,categoryList[uint256(category)]
            ,rankList[uint256(rank)]
            ,details,locate
        );
        heritageInheritorList.push(heritageInheritor);
        // 存入map
        queryHeritageInheritor[strNumber] = heritageInheritor;
        return strNumber;
    }

    // createHeritageProject
    // 创建非物质遗产项目
    // exam: 中国剪纸,6,0,江苏省,details
    function createHeritageProject(string memory name,
        CateGory category,Level rank,string memory locate,string memory details
    ) public returns(string memory) {
        string memory strNumber = generateRandomString(8);
        HeritageProject memory heritageProject = HeritageProject(
            strNumber,
            name,
            categoryList[uint256(category)],
            rankList[uint256(rank)],
            locate,
            details
        );
        heritageProjectList.push(heritageProject);
        //存入map
        queryHeritageProject[strNumber] = heritageProject;
        return strNumber;
    }
}
