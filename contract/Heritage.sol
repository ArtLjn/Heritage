// SPDX-License-Identifier: UNLICENSED

/*
@Time : 2024/5/23 上午9:49
@Author : ljn
@File : Heritage
@Software: GoLand
*/
pragma solidity ^0.6.0;
pragma experimental ABIEncoderV2;

import "./LibStrings.sol";

contract Heritage {
    using LibStrings for *;

    address public owner;

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
        bool isActive; // 标记是否有效
    }

    // 非遗传承项目
    struct HeritageProject {
        string number; // 编号
        string name; // 名称
        string category; // 类别
        string rank; // 级别
        string locate; // 区域
        string details; // 遗产详情
        bool isActive; // 标记是否有效
    }

    enum Level {
        provincial, // 省级
        city, // 市级
        national, //国家级
        human // 人类非物质文化遗产
    }

    enum CateGory {
        folk_literature, // 民间文学
        traditional_music,// 传统音乐
        traditional_dance, //传统舞蹈
        traditional_drama, // 传统戏剧
        QuYi, // 曲艺
        traditional_sports, //传统体育
        traditional_art, // 传统美术
        traditional_skill, //传统技艺
        traditional_medicine, //传统医药
        folk_custom //民俗
    }

    string[] public categoryList;
    string constant category = "民间文学,传统音乐,传统舞蹈,传统戏剧,曲艺,传统体育,传统美术,传统技艺,传统医药,民俗";
    string constant rankStr = "省级,市级,国家级,人类非物质文化遗产";
    string[] public rankList;

    HeritageInheritor[] public heritageInheritorList;
    HeritageProject[] public heritageProjectList;

    mapping(string => HeritageInheritor) public queryHeritageInheritor;
    mapping(string => HeritageProject) public queryHeritageProject;

    event InheritorCreated(string number, string Inheritor);
    event ProjectCreated(string number, string name);
    event InheritorUpdated(string number, string details, string locate);
    event ProjectUpdated(string number, string details, string locate);
    event InheritorDeleted(string number);
    event ProjectDeleted(string number);
    event InheritorUpdatedRank(string number,string rank);
    event ProjectUpdatedRank(string number,string rank);

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can execute this function.");
        _;
    }

    constructor() public {
        owner = msg.sender;
        LibStrings.slice memory s = category.toSlice();
        LibStrings.slice memory delim = ",".toSlice();
        uint params_total = s.count(delim) + 1;
        for (uint i = 0; i < params_total; i++) {
            categoryList.push(s.split(delim).toString());
        }
        LibStrings.slice memory r = rankStr.toSlice();
        uint rank_total =  r.count(delim) + 1;
        for (uint j = 0; j < rank_total; j++) {
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
            seed += uint256(keccak256(abi.encodePacked(seed, block.difficulty)));
        }
        return string(result);
    }

    function createHeritageInheritor(string memory Inheritor,string memory InheritorImg,
        string memory project,CateGory category,Level rank,string memory details,
        string memory locate
    ) public returns(string memory)  {
        string memory strNumber = generateRandomString(7);
        HeritageInheritor memory heritageInheritor = HeritageInheritor(
            strNumber,
            Inheritor,
            InheritorImg,
            project,
            block.timestamp,
            categoryList[uint256(category)],
            rankList[uint256(rank)],
            details,
            locate,
            true
        );
        heritageInheritorList.push(heritageInheritor);
        queryHeritageInheritor[strNumber] = heritageInheritor;
        emit InheritorCreated(strNumber, Inheritor);
        return strNumber;
    }

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
            details,
            true
        );
        heritageProjectList.push(heritageProject);
        queryHeritageProject[strNumber] = heritageProject;
        emit ProjectCreated(strNumber, name);
        return strNumber;
    }

    function updateHeritageInheritor(string memory number, string memory details, string memory locate) public onlyOwner {
        require(queryHeritageInheritor[number].isActive, "Inheritor does not exist or is inactive.");

        // 更新 map 数据
        queryHeritageInheritor[number].details = details;
        queryHeritageInheritor[number].locate = locate;

        // 同步更新列表数据
        for (uint i = 0; i < heritageInheritorList.length; i++) {
            if (keccak256(abi.encodePacked(heritageInheritorList[i].number)) == keccak256(abi.encodePacked(number))) {
                heritageInheritorList[i].details = details;
                heritageInheritorList[i].locate = locate;
                break;
            }
        }

        emit InheritorUpdated(number, details, locate);
    }

    function updateHeritageProject(string memory number, string memory details, string memory locate) public onlyOwner {
        require(queryHeritageProject[number].isActive, "Project does not exist or is inactive.");

        // 更新 map 数据
        queryHeritageProject[number].details = details;
        queryHeritageProject[number].locate = locate;

        // 同步更新列表数据
        for (uint i = 0; i < heritageProjectList.length; i++) {
            if (keccak256(abi.encodePacked(heritageProjectList[i].number)) == keccak256(abi.encodePacked(number))) {
                heritageProjectList[i].details = details;
                heritageProjectList[i].locate = locate;
                break;
            }
        }

        emit ProjectUpdated(number, details, locate);
    }

    function updateHeritageInheritorLevel(string memory number, Level rank) public onlyOwner {
        require(queryHeritageInheritor[number].isActive, "Inheritor does not exist or is inactive.");

        // 更新 map 数据
        queryHeritageInheritor[number].rank = rankList[uint256(rank)];

        // 同步更新列表数据
        for (uint i = 0; i < heritageInheritorList.length; i++) {
            if (keccak256(abi.encodePacked(heritageInheritorList[i].number)) == keccak256(abi.encodePacked(number))) {
                heritageInheritorList[i].rank = rankList[uint256(rank)];
                break;
            }
        }

        emit InheritorUpdatedRank(number, rankList[uint256(rank)]);
    }

    function updateHeritageProjectLevel(string memory number, Level rank) public onlyOwner {
        require(queryHeritageProject[number].isActive, "Project does not exist or is inactive.");

        // 更新 map 数据
        queryHeritageProject[number].rank = rankList[uint256(rank)];

        // 同步更新列表数据
        for (uint i = 0; i < heritageProjectList.length; i++) {
            if (keccak256(abi.encodePacked(heritageProjectList[i].number)) == keccak256(abi.encodePacked(number))) {
                heritageProjectList[i].rank = rankList[uint256(rank)];
                break;
            }
        }

        emit ProjectUpdatedRank(number, rankList[uint256(rank)]);
    }
    function deleteHeritageInheritor(string memory number) public onlyOwner {
        require(queryHeritageInheritor[number].isActive, "Inheritor already inactive.");

        // 将 map 中的继承人标记为无效
        queryHeritageInheritor[number].isActive = false;

        // 从列表中删除继承人
        for (uint i = 0; i < heritageInheritorList.length; i++) {
            if (keccak256(abi.encodePacked(heritageInheritorList[i].number)) == keccak256(abi.encodePacked(number))) {
                // 移除继承人
                removeInheritorAtIndex(i);
                break;
            }
        }

        emit InheritorDeleted(number);
    }

    function deleteHeritageProject(string memory number) public onlyOwner {
        require(queryHeritageProject[number].isActive, "Project already inactive.");

        // 将 map 中的项目标记为无效
        queryHeritageProject[number].isActive = false;

        // 从列表中删除项目
        for (uint i = 0; i < heritageProjectList.length; i++) {
            if (keccak256(abi.encodePacked(heritageProjectList[i].number)) == keccak256(abi.encodePacked(number))) {
                // 移除项目
                removeProjectAtIndex(i);
                break;
            }
        }

        emit ProjectDeleted(number);
    }

    // 辅助函数：从 heritageInheritorList 中移除指定索引的继承人
    function removeInheritorAtIndex(uint index) internal {
        require(index < heritageInheritorList.length, "Index out of bounds");

        for (uint i = index; i < heritageInheritorList.length - 1; i++) {
            heritageInheritorList[i] = heritageInheritorList[i + 1];
        }
        heritageInheritorList.pop(); // 删除最后一个元素
    }

    // 辅助函数：从 heritageProjectList 中移除指定索引的项目
    function removeProjectAtIndex(uint index) internal {
        require(index < heritageProjectList.length, "Index out of bounds");

        for (uint i = index; i < heritageProjectList.length - 1; i++) {
            heritageProjectList[i] = heritageProjectList[i + 1];
        }
        heritageProjectList.pop(); // 删除最后一个元素
    }


    function filterHeritageInheritorsByCategory(CateGory category) public view returns (HeritageInheritor[] memory) {
        uint count = 0;
        for (uint i = 0; i < heritageInheritorList.length; i++) {
            if (keccak256(abi.encodePacked(heritageInheritorList[i].category)) == keccak256(abi.encodePacked(categoryList[uint256(category)])) && heritageInheritorList[i].isActive) {
                count++;
            }
        }

        HeritageInheritor[] memory filtered = new HeritageInheritor[](count);
        uint index = 0;
        for (uint i = 0; i < heritageInheritorList.length; i++) {
            if (keccak256(abi.encodePacked(heritageInheritorList[i].category)) == keccak256(abi.encodePacked(categoryList[uint256(category)])) && heritageInheritorList[i].isActive) {
                filtered[index] = heritageInheritorList[i];
                index++;
            }
        }
        return filtered;
    }

    function filterHeritageProjectsByLocation(string memory locate) public view returns (HeritageProject[] memory) {
        uint count = 0;
        for (uint i = 0; i < heritageProjectList.length; i++) {
            if (keccak256(abi.encodePacked(heritageProjectList[i].locate)) == keccak256(abi.encodePacked(locate)) && heritageProjectList[i].isActive) {
                count++;
            }
        }

        HeritageProject[] memory filtered = new HeritageProject[](count);
        uint index = 0;
        for (uint i = 0; i < heritageProjectList.length; i++) {
            if (keccak256(abi.encodePacked(heritageProjectList[i].locate)) == keccak256(abi.encodePacked(locate)) && heritageProjectList[i].isActive) {
                filtered[index] = heritageProjectList[i];
                index++;
            }
        }
        return filtered;
    }
}
