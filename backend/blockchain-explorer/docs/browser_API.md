[toc]

## API接口文档

#### 1. 整体请求格式

查询请求全部采用 **GET** 请求方式，订阅，修改，取消订阅等操作采用 **POST** 请求方式

URL格式：http://localhost/chainmaker?cmb={OpHandle}

contentType采用 **application/form-data** 

请求的数据通过body体中以 **form-data** 方式发送

URL参数说明如下：

| 字段       | 描述                |
|:-------- |:----------------- |
| OpHandle | 固定字符串，用于描述本次操作的类型 |

#### 2. 整体返回数据格式

默认的格式会包括Response一个字段，使用下面的格式：

```json
{
    "Response": {
        "Data": {}
    }
}
```

若返回的数据是一个数组的话，和上面的一样，可参考下面的格式：

```json
{
    "Response": {
        "TotalCount": 2,
        "GroupList": [
            {},
            {}
        ]
    }

}
```

若返回的数据是一个标识异常的结果，则可参考下面的格式：

```json
{
    "Response": {
        "Error": {
            "Code": "AuthFailure",
            "Message": "pamam is nil"
        }
    }

}
```

说明如下：

| 字段         | 描述        | 取值范围   | 备注              |
|:---------- |:--------- |:------ |:--------------- |
| Response   | 应答关键字     | -----  | 固定格式            |
| Data       | 返回对象      | json格式 | 返回结果是对象（非集合）时使用 |
| TotalCount | 总量，用于计算分页 | 数字     | 返回结果是集合时使用      |
| GroupList  | 结果列表      | 集合     | 返回结果是集合时使用      |

### 1. 查询类接口

#### 1. 首页数据统计

请求URL： http://localhost/chainmaker?cmb=Decimal

请求Handle：Decimal

请求数据：

```json
{
    "ChainId": "chain1"
}
```

请求参数说明：

| 字段      | 描述  | 取值范围 | 备注               |
| ------- | --- | ---- | ---------------- |
| ChainId | 链id | 字符串  | 是否取得指定链的数据，为非必填项 |

返回数据：

```json
{
    "Response": {
        "Data": {
            "BlockHeight": 26,
            "TransactionNum":100,
            "ContractNum": 12,
            "OrgNum": 2,
            "NodeNum": 26,
              "UserNum": 12
        }
    }
}
```

应答结果说明如下：

| 字段             | 描述     | 取值范围 | 备注         |
| -------------- | ------ | ---- | ---------- |
| BlockHeight    | 区块高度   | 数字   | 区块链中最近区块高度 |
| TransactionNum | 今日交易数量 | 数字   | 区块链中交易的数量和 |
| NodeNum        | 节点数量   | 数字   | 区块链中节点的数量  |
| ContractNum    | 合约数量   | 数字   |            |
| OrgNum         | 组织数量   | 数字   |            |
| UserNum        | 用户数量   | 数字   |            |

#### 2. 交易统计曲线

请求URL： http://localhost/chainmaker?cmb=GetTransactionNumByTime

请求Handle：GetTransactionNumByTime

请求数据：

```json
{
    "ChainId": "chain1"
}
```

请求参数说明：

| 字段      | 描述  | 取值范围 | 备注               |
| ------- | --- | ---- | ---------------- |
| ChainId | 链id | 字符串  | 是否取得指定链的数据，为非必填项 |

返回数据：

```json
{
    "Response": {
        "TotalCount": 1,
        "GroupList": [{
            "TxNum": 4,
            "Timestamp": 1606669261
            }
        ]
    }
}
```

应答结果说明：

| 字段        | 描述   | 取值范围 | 备注                  |
| --------- | ---- | ---- | ------------------- |
| TxNum     | 交易数量 | 数字   | 当前链中所有的交易总和         |
| Timestamp | 时间戳  | 数字   | 该链创建时间，**秒** 级UTC时间 |

#### 3. 组织列表

请求URL： http://localhost/chainmaker?cmb=GetOrgList

请求Handle：GetOrgList

请求数据：

```json
{
    "ChainId": "chain1",
      "OrgId": "a",
    "Offset": 0,
    "Limit": 5
}
```

请求参数说明：

| 字段      | 描述  | 取值范围 | 备注                |
| ------- | --- | ---- | ----------------- |
| chainId | 链id | 字符串  | 返回指定区块id的节点信息     |
| OrgId   | 组织名 | 字符串  | 选填，过滤使用           |
| Offset  | 偏移量 | 数字   | 列表分页显示            |
| Limit   | 限制  | 数字   | 当前查询的数量，默认为每页显示数量 |

返回数据：

```json
{
    "Response": {
        "TotalCount": 2,
        "GroupList": [{
            "Id": 1,
            "OrgId": "org1.chainmaker.com",
            "NodeCount": 2,
            "UserCount": 3
            },
            {
            "Id": 1,
            "OrgId": "org1.chainmaker.com",
            "NodeCount": 1,
            "UserCount": 3
            }
        ]
    }
}
```

应答结果说明：

| 字段        | 描述      | 取值范围 | 备注                |
| --------- | ------- | ---- | ----------------- |
| Id        | 逻辑id    | 数字   | 区块逻辑id            |
| Org       | 组织id    | 字符串  |                   |
| NodeCount | 组织内的节点数 | 字符串  | 区块哈希              |
| UserCount | 组织内的用户数 | 数字   | 区块内打包的交易数量        |
| Timestamp | 时间戳     | 数字   | 创建时间，**秒** 级UTC时间 |

#### 5. 用户列表

请求URL： http://localhost/chainmaker?cmb=GetUserList

请求Handle：GetUserList

请求数据：

```json
{ 
    "ChainId": "chain1",
    "UserId": "a",
    "OrgId": "",
    "Offset": 0,
    "Limit": 5
}
```

请求参数说明：

| 字段      | 描述   | 取值范围 | 备注                |
| ------- | ---- | ---- | ----------------- |
| chainId | 链id  | 字符串  | 返回指定区块id的节点信息     |
| UserId  | 用户id | 字符串  | 选填，过滤使用           |
| OrgId   | 组织id | 字符串  | 选填，过滤使用           |
| Offset  | 偏移量  | 数字   | 列表分页显示            |
| Limit   | 限制   | 数字   | 当前查询的数量，默认为每页显示数量 |

返回数据：

```json
{
    "Response": {
        "TotalCount": 2,
        "GroupList": [{
            "Id": 1,
            "UserId": "sdasdb12d",
            "Role": "admin",
            "OrgId": "xxx",
            "Timestamp": 1606669261
            },
            {
            "Id": 2,
            "UserId": "sdasdb12d",
            "Role": "user",
            "OrgId": "SUCCESS",
            "Timestamp": 1606669261
            }
        ]
    }
}
```

应答结果说明：

| 字段        | 描述     | 取值范围 | 备注                   |
| --------- | ------ | ---- | -------------------- |
| Id        | 交易逻辑id | 数字   | 交易逻辑id               |
| UserId    | 用户名称   | 字符串  |                      |
| Role      | 用户角色   | 字符串  |                      |
| OrgId     | 所属组织   | 字符串  |                      |
| Timestamp | 时间戳    | 数字   | 该合约创建时间，**秒** 级UTC时间 |

#### 6. 节点列表

请求URL： http://localhost/chainmaker?cmb=GetNodeList

请求Handle：GetNodeList

是否需要Token：否

请求数据：

```json
{
    "ChainId": "chain1",
    "NodeName": "node1",
    "OrgId": "xxx",
    "Offset": 0,
    "Limit": 5
}
```

请求参数说明：

| 字段       | 描述   | 取值范围 | 备注                |
| -------- | ---- | ---- | ----------------- |
| chainId  | 链id  | 字符串  | 返回制定区块id的节点信息     |
| NodeName | 节点名称 | 字符串  | 选填，过滤使用           |
| OrgId    | 组织id | 字符串  | 选填，过滤使用           |
| Offset   | 偏移量  | 数字   | 列表分页显示            |
| Limit    | 限制   | 数字   | 当前查询的数量，默认为每页显示数量 |

返回数据：

```json
{
    "Response": {
        "TotalCount": 2,
        "GroupList": [{
            "Id": 1,
            "NodeId": "xxxx",
            "NodeName": "node1",
            "NodeAddress": "/ip4/127.0.0.1/tcp/11303",
            "Role": "common",
            "OrgId": "org1"
            },{
            "Id": 2,
            "NodeId": "xxxx",
            "NodeName": "node1",
            "NodeAddress": "/ip4/127.0.0.1/tcp/11303",
            "Role": "common",
            "OrgId": "org1"
            }
        ]
    }
}
```

应答结果说明：

| 字段          | 描述     | 取值范围 | 备注        |
| ----------- | ------ | ---- | --------- |
| Id          | 节点逻辑id | 数字   | 节点逻辑id    |
| NodeId      | 节点Id   | 字符串  | 节点唯一标识    |
| NodeName    | 节点名称   | 字符串  |           |
| NodeAddress | 节点地址   | 字符串  | 节点所在地址    |
| OrgId       | 组织名    | 字符串  | 节点所属组织    |
| Status      | 节点状态   | 数字   | 1：正常；0：异常 |

#### 7. 查看区块列表

请求URL： http://localhost/chainmaker?cmb=GetBlockList

请求Handle：GetBlockList

是否需要Token：是

请求数据：

```json
{
    "ChainId": "chain1",
    "Block": "da13ddad12",
    "StartTime": 1606669261,
    "EndTime": 1606669261,
    "Offset": 0,
    "Limit": 10
}
```

请求参数说明：

| 字段        | 描述           | 取值范围 | 备注                |
| --------- | ------------ | ---- | ----------------- |
| ChainId   | 链id          | 字符串  | 是否取得指定链的数据，为非必填项  |
| Block     | 区块hash值或区块高度 | 字符串  | 非必填用于筛选           |
| StartTime | 开始时间         | 数字   | 非必填用于筛选           |
| EndTime   | 结束时间         | 数字   | 非必填用于筛选           |
| Offset    | 偏移量          | 数字   | 列表分页显示            |
| Limit     | 限制           | 数字   | 当前查询的数量，默认为每页显示数量 |

返回数据：

```json
{
    "Response": {
        "TotalCount": 2,
        "GroupList": [{
            "Id": 1,
            "BlockHeight": 102,
            "BlockHash": "asasd12",
            "TxCount": 4,
            "ProposalNodeId": "dasd",
            "Timestamp": 1606669261
            },
            {
            "Id": 2,
            "BlockHeight": 102,
            "BlockHash": "asddwerwe",
            "TxCount": 3,
            "ProposalNodeId": "ghffw",
            "Timestamp": 1606669261
            }
        ]
    }
}
```

应答结果说明：

| 字段             | 描述     | 取值范围 | 备注                   |
| -------------- | ------ | ---- | -------------------- |
| Id             | 区块逻辑ID | 数字   | 该值用户分页时作为from使用      |
| BlockHash      | 区块hash | 字符串  | 当前区块的hash值           |
| TxCount        | 区块交易数量 | 数字   | 区块交易数量               |
| ProposalNodeId | 提案节点id | 字符串  | 提案节点id               |
| Timestamp      | 时间戳    | 数字   | 该区块创建时间，**秒** 级UTC时间 |

#### 8. 查看交易列表

请求URL： http://localhost/chainmaker?cmb=GetTxList

请求Handle：GetTxList

是否需要Token：是

请求数据：

```json
{
    "ChainId": "chain1",
    "TxId": "da13ddad12",
    "ContractName": "increase",
    "BlockHash": "kbddbjd",
    "StartTime": 1606669261,
    "EndTime": 1606669261,
    "Offset": 0,
    "Limit": 10
}
```

请求参数说明：

| 字段           | 描述      | 取值范围 | 备注                |
| ------------ | ------- | ---- | ----------------- |
| ChainId      | 链id     | 字符串  | 是否取得指定链的数据，为非必填项  |
| TxId         | 交易id值   | 字符串  | 非必填用于筛选           |
| BlockHash    | 区块hash值 | 字符串  | 非必填用于筛选           |
| ContractName | 合约名     | 字符串  |                   |
| StartTime    | 开始时间    | 数字   | 非必填用于筛选           |
| EndTime      | 结束时间    | 数字   | 非必填用于筛选           |
| Offset       | 偏移量     | 数字   | 列表分页显示            |
| Limit        | 限制      | 数字   | 当前查询的数量，默认为每页显示数量 |

返回数据：

```json
{
    "Response": {
        "TotalCount": 2,
        "GroupList": [{
            "Id": 1,
              "BlockHeight": 1,
            "TxId": "asasd12",
            "Sender": "com",
              "SenderOrg": "org1.com",
            "ContractName": "ChainConfig",
            "Status": "",
            "BlockHash": "dasd",
            "Timestamp": 1606669261
            },{
            "Id": 2,
            "TxId": "asasd13",
            "Sender": "com",
            "BlockHash": "dasdd",
            "Timestamp": 1606669261
            }
        ]
    }
}
```

应答结果说明：

| 字段                 | 描述     | 取值范围 | 备注                   |
| ------------------ | ------ | ---- | -------------------- |
| Id                 | 区块逻辑ID | 数字   | 该值用户分页时作为from使用      |
| BlockHeight        | 区块高度   | 数字   |                      |
| TxId               | 交易Id   | 字符串  | 交易唯一标识               |
| Sender             | 交易发送者  | 字符串  | 交易发送者                |
| SenderOrg          | 交易发送组织 | 字符串  |                      |
| BlockHash          | 区块hash | 字符串  | 当前区块的hash值           |
| ContractName       | 合约名    | 字符串  |                      |
| ContractMethod     | 合约调用方法 | 字符串  |                      |
| ContractParameters | 合约参数列表 | 字符串  |                      |
| Status             | 交易状态   | 字符串  | "SUCCESS" 或其他错误      |
| Timestamp          | 时间戳    | 数字   | 该区块创建时间，**秒** 级UTC时间 |

#### 9. 合约列表

请求URL： http://localhost/chainmaker?cmb=GetContractList

请求Handle：GetContractList

请求数据：

```json
{
    "ChainId": "chain1",
    "ContractName": "incre",
    "Offset": 0,
    "Limit": 5,
    "Order": "desc"
}
```

请求参数说明：

| 字段           | 描述  | 取值范围 | 备注                |
| ------------ | --- | ---- | ----------------- |
| ChainId      | 链id | 字符串  | 是否取得指定链的数据        |
| ContractName | 合约名 | 字符串  | 选填，过滤使用           |
| Offset       | 偏移量 | 数字   | 列表分页显示            |
| Limit        | 限制  | 数字   | 当前查询的数量，默认为每页显示数量 |
| Order        | 限制  | 字符串  | "desc" or "asc"   |

返回数据：

```json
{
    "Response": {
        "TotalCount": 2,
        "GroupList": [{
            "Id": 1,
            "Name": "token",
              "Version": "v1.1.0",
              "TxCount": 1000,
            "Creator": "a.b.com",
            "CreateTimestamp": 1606669263,
            "UpgradeUser": "a.b.com",
              "UpgradeTimestamp": 1606669261
            },
            {
            "Id": 2,
            "Name": "token",
              "Version": "v1.1.0",
              "TxCount": 1000,
            "Creator": "a.b.com",
            "CreateTimestamp": 1606669263,
            "UpgradeUser": "a.b.com",
              "UpgradeTimestamp": 1606669261
            }
        ]
    }
}
```

应答结果说明：

| 字段               | 描述      | 取值范围 | 备注                   |
| ---------------- | ------- | ---- | -------------------- |
| Id               | 逻辑ID    | 数字   | 该值用户分页时作为from使用      |
| Name             | 合约名     | 字符串  |                      |
| Version          | 当前版本号   | 字符串  |                      |
| TxCount          | 合约交易数量  | 整型   |                      |
| Creator          | 创建者     | 字符串  |                      |
| CreateTimestamp  | 创建时间戳   | 数字   | 该合约创建时间，**秒** 级UTC时间 |
| UpgradeUser      | 升级合约操作人 | 字符串  |                      |
| UpgradeTimestamp | 升级时间戳   | 数字   | 该合约创建时间，**秒** 级UTC时间 |

#### 10. 查询区块详情

请求URL： http://localhost/chainmaker?cmb=GetBlockDetail

请求Handle：GetBlockDetail

请求数据：

```json
{
    "ChainId": "chain1",
    "BlockHash": "aaaaaaaaaaaaaaa",
    "BlockHeight": 1
}
```

请求参数说明：

| 字段          | 描述     | 取值范围 | 备注             |
| ----------- | ------ | ---- | -------------- |
| ChainId     | 链ID    | 字符串  | 必填             |
| BlockHash   | 区块hash | 字符串  | 选填（与区块高度二选一）   |
| BlockHeight | 区块高度   | int  | 选填（与区块hash二选一） |

返回数据：

```json
{
    "Response": {
        "Data": {
            "BlockHash": "asdffk2",
            "RwSetHash": "xxxxx"
            "PreBlockHash": "asdh123123",
            "ProposalNodeId": "common",
            "TxRootHash": "asdkkf",
            "TxCount": 2,
            "BlockHeight": 17,
            "OrgId": "org1",
            "Dag": "asdkkf",
            "Timestamp": 1606669261
            }
    }
}
```

应答结果说明：

| 字段             | 描述         | 取值范围 | 备注                  |
| -------------- | ---------- | ---- | ------------------- |
| BlockHash      | 区块hash     | 字符串  | 区块hash              |
| RwSetHash      | 读写集hash    | 字符串  |                     |
| PreBlockHash   | 前一个区块hash  | 字符串  | 前一个区块hash           |
| ProposalNodeId | 提案节点       | 字符串  | 发起交易的节点             |
| TxRootHash     | 交易merkle哈希 | 字符串  | 交易merkle哈希          |
| TxCount        | 交易数量       | 数字   | 区块内的交易数量            |
| BlockHeight    | 区块高度       | 数字   | 区块所在高度              |
| OrgId          | 提案组织id     | 字符串  | 提案组织id              |
| Dag            | Dag特征摘要    | 字符串  | Dag特征摘要             |
| Timestamp      | 时间戳        | 数字   | 区块创建时间，**秒** 级UTC时间 |

#### 11. 查询交易详情

请求URL： http://localhost/chainmaker?cmb=GetTxDetail

请求Handle：GetTxDetail

请求数据：

```json
{
    "Id": 10,
    "ChainId": "chain1",
    "TxId": "aaaaaaaaaaaaaaa"
}
```

请求参数说明：

| 字段      | 描述     | 取值范围 | 备注  |
| ------- | ------ | ---- | --- |
| Id      | 交易逻辑ID | 数字   | 选填  |
| ChainId | 链ID    | 字符串  | 选填  |
| TxId    | 交易ID   | 字符串  | 选填  |

返回数据：

```json
{
    "Response": {
        "Data": {
            "TxId": "chain1",
            "TxHash": "asdh123123",
            "BlockHash": "asdh123123",
            "BlockHeight": 5,
            "Sender": "common",
            "OrgId": "xx.org",
            "ContractName": "asset",
            "TxStatusCode": "success",
            "ContractResultCode": "OK",
            "ContractResult": "OK",
            "RwSetHash": "fa11a5f1",
            "ContractMethod": "add_evidence",
            "ContractParameters": "[{\"key\":\"app_id\",\"value\":\"20051500011\"},{\"key\":\"evi_id\",\"value\":\"5c7100e1d62745d3ab8a91184be3fcca\"}]",
            "Timestamp": 1606669261
            }
    }
}
```

应答结果说明：

| 字段                 | 描述       | 取值范围 | 备注                  |
| ------------------ | -------- | ---- | ------------------- |
| TxId               | 交易id     | 字符串  | 交易id                |
| TxHash             | 交易hash   | 字符串  | 交易hash              |
| Sender             | 交易发送者    | 字符串  | 交易发送者               |
| OrgId              | 交易发送所属组织 | 字符串  |                     |
| ContractName       | 合约名      | 字符串  | 合约名                 |
| BlockHash          | 区块hash   | 字符串  | 区块hash              |
| BlockHeight        | 区块高度     | 数字   | 区块高度                |
| ContractMethod     | 方法名      | 字符串  | 合约方法                |
| ContractParameters | 参数       | 字符串  | 合约参数                |
| TxStatusCode       | 交易状态     | 字符串  | 交易状态                |
| ContractResultCode | 合约执行结果码  | 字符串  | 合约执行结果码             |
| ContractResult     | 合约执行结果   | 字符串  | 合约执行结果              |
| RwSetHash          | 交易信息     | 字符串  | 交易信息                |
| Timestamp          | 时间戳      | 数字   | 交易创建时间，**秒** 级UTC时间 |

#### 12. 首页查询

请求URL： http://localhost/chainmaker?cmb=Search

请求Handle：Search

请求数据：

```json
{
    "Id": "",
    "ChainId": "chain1"
}
```

请求参数说明：

| 字段      | 描述    | 取值范围 | 备注                       |
| ------- | ----- | ---- | ------------------------ |
| Id      | 查询输入值 | 字符串  | 可以为交易id，区块hash, 区块高度，合约名 |
| ChainId | 链id   | 字符串  | 是否取得指定链的数据，必填            |

返回数据：

```json
{
    "Response": {
        "Data": {
            "Type": 0,
            "Id": 4,
            "Data": "ddffww",
            "ChainId": "chain1"
            }
    }
}
```

应答结果说明：

| 字段      | 描述        | 取值范围 | 备注                            |
| ------- | --------- | ---- | ----------------------------- |
| Type    | 返回数据类型    | 数字   | 0：区块详情；1：交易详情；2：合约详情； -1:未知结果 |
| Id      | 查询详情逻辑ID  | 数字   | 根据不同类型，返回不同的类型的id             |
| Data    | 不同类型的不同标识 | 字符串  | 区块：区块hash 交易：交易id，合约：合约名      |
| ChainId | 链id       | 字符串  |                               |

#### 13. 合约详情

请求URL： http://localhost/chainmaker?cmb=GetContractDetail

请求Handle：GetContractDetail

请求数据：

```json
{
    "ChainId": "chain1",
    "ContractName": "increase"
}
```

请求参数说明：

| 字段           | 描述      | 取值范围 | 备注  |
| ------------ | ------- | ---- | --- |
| ChainId      | chainid | 字符串  |     |
| ContractName | 合约名     | 字符串  |     |

返回数据：

```json
{
    "Response": {
        "Data": {
            "Name": "increase",
              "Version": "v1.1.0",
              "TxCount": 1000,
            "Creator": "a.b.com",
            "CreateTxId": "xxxxxxxxxxx",
            "CreateTimestamp": 1606669263,
            "UpgradeUser": "a.b.com",
              "UpgradeTimestamp": 1606669261
            }
        ]
    }
}
```

应答结果说明：

| 字段               | 描述      | 取值范围 | 备注                   |
| ---------------- | ------- | ---- | -------------------- |
| Id               | 逻辑ID    | 数字   | 该值用户分页时作为from使用      |
| Name             | 合约名     | 字符串  |                      |
| Version          | 当前版本号   | 字符串  |                      |
| TxCount          | 合约交易数量  | 整型   |                      |
| Creator          | 创建者     | 字符串  |                      |
| CreateTxId       | 创建交易id  |      |                      |
| CreateTimestamp  | 创建时间戳   | 数字   | 该合约创建时间，**秒** 级UTC时间 |
| UpgradeUser      | 升级合约操作人 | 字符串  |                      |
| UpgradeTimestamp | 升级时间戳   | 数字   | 该合约创建时间，**秒** 级UTC时间 |

#### 14. 事件列表

请求URL： http://localhost/chainmaker?cmb=GetEventList

请求Handle：GetEventList

请求数据：

```json
{
  "ChainId": "chain2",
    "ContractName": "increase",
   "Offset": 0,
    "Limit": 10
}
```

请求参数说明：

| 字段           | 描述  | 取值范围 | 备注                |
| ------------ | --- | ---- | ----------------- |
| ChainId      |     | 字符串  | 根据参数返回数据数量        |
| ContractName | 合约名 |      |                   |
| Offset       | 偏移量 | 数字   | 列表分页显示            |
| Limit        | 限制  | 数字   | 当前查询的数量，默认为每页显示数量 |

返回数据：

```json
{
    "Response": {
        "TotalCount": 2,
        "GroupList": [{
            "Topic": "chain1",
            "EventInfo": "[xxxxxxxx]",
            "Timestamp": 1606669261
            },
            {
            "Topic": "chain1",
            "EventInfo": "[xxxxxxxx]",
            "Timestamp": 1606669261
            }
        ]
    }
}
```

应答结果说明：

| 字段        | 描述   | 取值范围 | 备注  |
| --------- | ---- | ---- | --- |
| Id        | 逻辑ID | 数字   |     |
| Topic     | 事件主题 | 字符串  |     |
| EventInfo | 事件信息 | 字符串  |     |
| Timestamp | 时间戳  | 数字   |     |

### 2. 链管理类接口

#### 15 查看链列表【拿不到链的启动时间】

请求URL： http://localhost/chainmaker?cmb=GetChainList

请求Handle：GetChainList

是否需要Token：是

请求数据：

```json
{
    "ChainId": "chain1",
    "Offset": 0,
    "Limit": 10
}
```

请求参数说明：

| 字段      | 描述  | 取值范围 | 备注                |
| ------- | --- | ---- | ----------------- |
| ChainId | 链id | 字符串  | 非必填用于筛选           |
| Offset  | 偏移量 | 数字   | 列表分页显示            |
| Limit   | 限制  | 数字   | 当前查询的数量，默认为每页显示数量 |

返回数据：

```json
{
    "Response": {
        "TotalCount": 2,
        "GroupList": [{
            "Id": 1,
            "ChainId": "chain1",
            "ChainVersion": "v2.1.0",
            "Consensus": "PBFT",
            "Timestamp": 1606669261,
          "Status": 1
            },{
            "Id": 2,
            "ChainId": "chain2",
              "ChainVersion": "v2.0.0",
            "Consensus": "PBFT",
            "Timestamp": 1606669261,
              "Status": 0
            }
        ]
    }
}
```

应答结果说明：

| 字段           | 描述    | 取值范围 | 备注                  |
| ------------ | ----- | ---- | ------------------- |
| Id           | 链逻辑ID | 数字   | 该值用户分页时作为from使用     |
| ChainId      | 链ID   | 字符串  | 链的唯一ID              |
| ChainVersion | 链版本   | 字符串  |                     |
| Consensus    | 共识算法  | 字符串  | 共识算法                |
| Timestamp    | 时间戳   | 数字   | 该链创建时间，**秒** 级UTC时间 |
| Status       | 状态    | int  | 0: 正常 1：异常 2.取消订阅   |

#### 16 订阅链（绑定链）

请求URL： http://localhost/chainmaker?cmb=SubscribeChain

请求Handle：SubscribeChain

是否需要Token：否

请求数据：

```json
{
    "ChainId": "chain1",
    "Addr": "127.0.0.1:12301",
    "OrgId": "wx-org1.org",
    "Tls": true,
    "OrgCA": "",
    "UserCert": "",
    "UserKey": "",
    "TLSHostName": "chainmaker.org",
}
```

请求参数说明：

| 字段          | 描述          | 取值范围 | 备注        |
| ----------- | ----------- | ---- | --------- |
| ChainId     | 链id         | 字符串  | 必填        |
| Addr        | 地址          | 字符串  | 必填        |
| OrgId       | 证书所属组织id    | 字符串  | tls开启为必填  |
| OrgCA       | node节点的CA证书 | 字符串  | pem格式的字符串 |
| UserCert    | 用户证书        | 字符串  | Pem 格式    |
| UserKey     | 用户key       | 字符串  | Pem 格式    |
| TLSHostName | tls连接使用的域名  | 字符串  | 选填        |

返回数据：

```json
{
    "Response": {
      "Data": {
            "Id": 2,
            "ChainId": "chain2",
              "ChainVersion": "v2.0.0",
            "Consensus": "PBFT",
            "Timestamp": 1606669261,
              "Status": 0
     }
}
```

应答结果说明：

| 字段           | 描述    | 取值范围 | 备注                  |
| ------------ | ----- | ---- | ------------------- |
| Id           | 链逻辑ID | 数字   | 该值用户分页时作为from使用     |
| ChainId      | 链ID   | 字符串  | 链的唯一ID              |
| ChainVersion | 链版本   | 字符串  |                     |
| Consensus    | 共识算法  | 字符串  | 共识算法                |
| Timestamp    | 时间戳   | 数字   | 该链创建时间，**秒** 级UTC时间 |
| Status       | 状态    | int  | 0: 正常 1：异常 2.取消订阅   |

#### 17 取消订阅信息

请求URL： http://localhost/chainmaker?cmb=CancelSubscribe

请求Handle：CancelSubscribe

是否需要Token：否

请求数据：

```json
{
    "ChainId": "chain1",
}
```

请求参数说明：

| 字段      | 描述  | 取值范围 | 备注     |
| ------- | --- | ---- | ------ |
| ChainId | 链id | 字符串  | 必填用于筛选 |

返回数据：

```json
{
    "Response": {
      "Data": {
            "Id": 2,
            "ChainId": "chain2",
              "ChainVersion": "v2.0.0",
            "Consensus": "PBFT",
            "Timestamp": 1606669261,
              "Status": 1
     }
}
```

应答结果说明：

| 字段           | 描述    | 取值范围 | 备注                  |
| ------------ | ----- | ---- | ------------------- |
| Id           | 链逻辑ID | 数字   | 该值用户分页时作为from使用     |
| ChainId      | 链ID   | 字符串  | 链的唯一ID              |
| ChainVersion | 链版本   | 字符串  |                     |
| Consensus    | 共识算法  | 字符串  | 共识算法                |
| Timestamp    | 时间戳   | 数字   | 该链创建时间，**秒** 级UTC时间 |
| Status       | 状态    | int  | 0: 正常 1：异常 2.取消订阅   |

#### 18 修改订阅信息

请求URL： http://localhost/chainmaker?cmb=ModifySubscribe

请求Handle：ModifySubscribe

是否需要Token：否

请求数据：

```json
{
    "ChainId": "chain1",
    "Addr": "127.0.0.1:12301",
    "Tls": true,
    "OrgId": "wx-org1.org",
    "OrgCA": "",
    "UserCert": "",
    "UserKey": "",
    "TLSHostName": "chainmaker.org",
}
```

请求参数说明：

| 字段          | 描述          | 取值范围 | 备注            |
| ----------- | ----------- | ---- | ------------- |
| ChainId     | 链id         | 字符串  | 必填用于筛选        |
| Addr        | 地址          | 字符串  | 选填            |
| OrgCA       | node节点的CA证书 | 字符串  | pem格式的字符串  选填 |
| UserCert    | 用户证书        | 字符串  | 选填            |
| UserKey     | 用户key       | 字符串  | 选填            |
| TLSHostName | tls连接使用的域名  | 字符串  | 选填            |

返回数据：

```json
{
    "Response": {
      "Data": {
            "Id": 2,
            "ChainId": "chain2",
              "ChainVersion": "v2.0.0",
            "Consensus": "PBFT",
            "Timestamp": 1606669261,
              "Status": 0
     }
}
```

应答结果说明：

| 字段           | 描述    | 取值范围 | 备注                  |
| ------------ | ----- | ---- | ------------------- |
| Id           | 链逻辑ID | 数字   | 该值用户分页时作为from使用     |
| ChainId      | 链ID   | 字符串  | 链的唯一ID              |
| ChainVersion | 链版本   | 字符串  |                     |
| Consensus    | 共识算法  | 字符串  | 共识算法                |
| Timestamp    | 时间戳   | 数字   | 该链创建时间，**秒** 级UTC时间 |
| Status       | 状态    | int  | 0: 正常 1：异常  2.取消订阅  |
