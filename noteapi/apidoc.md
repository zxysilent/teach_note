# noteapi’s Api文档
凭证传递方式包括 get、post、header、cookie
数据传递方式包括 json、formData
/api/*公共访问
/adm/* 必须传入token

## Version: 1.0

### /adm/note/add

#### POST
##### Summary

添加笔记

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| token | query | 凭证 | Yes | string |
| body | body | 请求数据 | Yes | [model.Note](#modelnote) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /adm/note/drop

#### POST
##### Summary

删除笔记

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body | 请求数据 | Yes | [model.IptId](#modeliptid) |
| token | query | 凭证 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /adm/note/edit

#### POST
##### Summary

修改笔记

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| token | query | 凭证 | Yes | string |
| body | body | 请求数据 | Yes | [model.Note](#modelnote) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /api/feedback/add

#### POST
##### Summary

添加反馈

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body | 请求数据 | Yes | [model.Feedback](#modelfeedback) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /api/feedback/all

#### GET
##### Summary

所有反馈

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) & object |

### /api/note/all

#### GET
##### Summary

所有笔记

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) & object |

### /api/note/get

#### GET
##### Summary

通过id获取笔记

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | query | id | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) & object |

### Models

#### model.Feedback

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| content | string | 内容 | No |
| ctime | integer | 时间 | No |
| id | integer | 主键 | No |

#### model.IptId

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | integer | 仅包含Id的请求 | No |

#### model.Note

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| content | string | 详细 | No |
| ctime | integer | 创建时间 | No |
| id | integer | 主键 | No |
| title | string | 标题 | No |
| utime | integer | 修改时间 | No |

#### model.Reply

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| code | integer | _Example:_ `200` | No |
| msg | string | _Example:_ `"提示信息"` | No |
