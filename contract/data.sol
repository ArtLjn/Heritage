// SPDX-License-Identifier: UNLICENSED

/*
@Time : 2024/5/22 下午10:09 
@Author : ljn
@File : data
@Software: GoLand
*/
pragma solidity ^0.6.0;
pragma experimental ABIEncoderV2;
import "./LibStrings.sol";
contract data {

    using LibStrings for *;
    // 非遗传承人
    struct HeritageInheritor {
        string number;  // 编号
        string Inheritor; //传承人
        string InheritorImg; // 传承人证件照
        string project; // 传承项目
        uint256 creationTime; //创建时间
        string category; // 所属类别
        string rank; //级别
        string details; //遗产详情
        string locate; // 申请地区或单位
    }

    // 非遗传承项目
    struct HeritageProject {
        string number; // 编号
        string name; // 名称
        string category; // 类别
        string rank; // 级别
        string locate; // 区域
        string details; // 遗产详情
    }

    enum Level {
        national, //国家级
        provincial, // 省级
        city, // 市级
        human // 人类非物质文化遗产
    }


    enum CateGory {
        folk_literature, // 民间文学
        traditional_music,// 传统音乐
        traditional_dance, //传统舞蹈
        traditional_drama, // 传统戏剧
        QuYi, // 曲艺
        traditional_sports, //传统体育
        traditional_skill, //传统技艺
        traditional_medicine, //传统技艺
        folk_custom //民俗
    }

    // 遗产类别列表
    string[] public categoryList;
    // 遗产类型构造
    string constant category = "民间文学,传统音乐,传统舞蹈,传统戏剧,曲艺,传统体育,传统美术,传统技艺,传统医药,民俗";
    // 级别类型构造
    string constant rankStr = "国家级,省级,市级,列入人类非物质文化遗产代表作名录的项目";
    // 遗产级别列表
    string[] public rankList;
    // 非遗传承人列表
    HeritageInheritor[] public heritageInheritorList;
    // 非遗传承项目列表
    HeritageProject[] public heritageProjectList;
    // 根据编号查询非遗传承人项目
    mapping(string => HeritageInheritor) public queryHeritageInheritor;
    // 根据编号查询非遗传承项目
    mapping(string => HeritageProject) public queryHeritageProject;

    constructor() public {
        LibStrings.slice memory s = category.toSlice();
        LibStrings.slice memory delim = ",".toSlice();
        uint params_total = s.count(delim) + 1;
        for(uint i = 0; i < params_total; i++) {
            categoryList.push(s.split(delim).toString());
        }
        LibStrings.slice memory r = rankStr.toSlice();
        uint rank_total =  r.count(delim) + 1;
        for (uint j = 0; j < rank_total; j ++) {
            rankList.push(r.split(delim).toString());
        }
    }

    function getCateGoryList() public view returns(string[] memory)  {
        return categoryList;
    }
    function getHeritageInheritorList() public view returns(HeritageInheritor[] memory) {
        return heritageInheritorList;
    }
    function getHeritageProject() public view returns(HeritageProject[] memory)  {
        return heritageProjectList;
    }
    function getRankList() public view returns(string[] memory) {
        return rankList;
    }
    function generateRandomString(uint length) internal view returns (string memory) {
        uint256 seed = uint256(keccak256(abi.encodePacked(block.timestamp, block.difficulty)));
        bytes memory alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
        bytes memory result = new bytes(length);
        for (uint256 i = 0; i < length; i++) {
            result[i] = alphabet[uint256(seed) % alphabet.length];
            // 为了增加随机性，可以在每次迭代后更新seed
            seed += uint256(keccak256(abi.encodePacked(seed, block.difficulty)));
        }
        return string(result);
    }
}

