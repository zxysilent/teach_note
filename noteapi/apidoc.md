# noteapi’s Api文档
凭证传递方式包括 get、post、header、cookie
数据传递方式包括 json、formData
/api/*公共访问
/adm/* 必须传入token

## Version: 1.0

### /adm/article/add

#### POST
##### Summary

添加article信息

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| token | query | 凭证 | Yes | string |
| body | body | 请求数据 | Yes | [model.Article](#modelarticle) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /adm/article/drop

#### POST
##### Summary

删除article信息

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body | 请求数据 | Yes | [model.IptId](#modeliptid) |
| token | query | 凭证 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /adm/article/edit

#### POST
##### Summary

修改article信息

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| token | query | 凭证 | Yes | string |
| body | body | 请求数据 | Yes | [model.Article](#modelarticle) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /adm/auth/edit/info

#### POST
##### Summary

修改自己的信息

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| token | query | 凭证 | No | string |
| Authorization | header | 凭证 | No | string |
| body | body | request | Yes | [model.User](#modeluser) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /adm/auth/edit/passwd

#### POST
##### Summary

修改自己的密码

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| token | query | 凭证 | No | string |
| Authorization | header | 凭证 | No | string |
| opass | formData | 旧密码 | Yes | string |
| npass | formData | 新密码 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /adm/auth/get

#### GET
##### Summary

登录信息

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| token | query | 凭证 | No | string |
| Authorization | header | 凭证 | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) & object |

### /adm/class/add

#### POST
##### Summary

添加class信息

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| token | query | 凭证 | Yes | string |
| body | body | 请求数据 | Yes | [model.Class](#modelclass) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /adm/class/drop

#### POST
##### Summary

删除class信息

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body | 请求数据 | Yes | [model.IptId](#modeliptid) |
| token | query | 凭证 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /adm/class/edit

#### POST
##### Summary

修改class信息

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| token | query | 凭证 | Yes | string |
| body | body | 请求数据 | Yes | [model.Class](#modelclass) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /adm/user/add

#### POST
##### Summary

添加user信息

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| token | query | 凭证 | Yes | string |
| body | body | 请求数据 | Yes | [model.User](#modeluser) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /adm/user/drop

#### POST
##### Summary

删除user信息

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body | 请求数据 | Yes | [model.IptId](#modeliptid) |
| token | query | 凭证 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /adm/user/edit

#### POST
##### Summary

修改user信息

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| token | query | 凭证 | Yes | string |
| body | body | 请求数据 | Yes | [model.User](#modeluser) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /api/article/edit/hits

#### GET
##### Summary

修改article信息hits

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | query | id | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) & object |

### /api/article/get

#### GET
##### Summary

通过id获取article信息

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | query | id | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) & object |

### /api/article/page

#### GET
##### Summary

分页数据

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| pi | query | 分页数 | Yes | integer |
| ps | query | 每页条数[5,30] | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) & object |

### /api/auth/login

#### POST
##### Summary

登陆

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| num | formData | 账号 | Yes | string |
| passwd | formData | 密码 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) & object |

### /api/banner/all

#### GET
##### Summary

所有轮播

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) & object |

### /api/class/all

#### GET
##### Summary

所有分类

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) & object |

### /api/class/get

#### GET
##### Summary

通过id获取class信息

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | query | id | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) & object |

### /api/class/page

#### GET
##### Summary

分页数据

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| pi | query | 分页数 | Yes | integer |
| ps | query | 每页条数[5,30] | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) & object |

### /api/sys/info

#### GET
##### Summary

系统信息

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) & object |

### /api/upload/file

#### POST
##### Summary

上传文件

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| file | formData | file | Yes | file |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /api/upload/image

#### POST
##### Summary

上传图片

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| file | formData | file | Yes | file |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) |

### /api/user/get

#### GET
##### Summary

通过id获取user信息

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | query | id | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) & object |

### /api/user/page

#### GET
##### Summary

分页数据

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| pi | query | 分页数 | Yes | integer |
| ps | query | 每页条数[5,30] | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | 返回数据 | [model.Reply](#modelreply) & object |

### Models

#### model.Article

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| author | string | 作者 | No |
| cid | integer | cls id | No |
| content | string | 详细 | No |
| cunix | integer | 创建时间 | No |
| hits | integer | 点击数 | No |
| id | integer | 主键 | No |
| title | string | 标题 | No |
| uunix | integer | 修改时间 | No |

#### model.Banner

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| cunix | integer | 创建时间 | No |
| id | integer | 主键 | No |
| title | string | 标题 | No |
| url | string | 图片 | No |

#### model.Class

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | integer | 主键 | No |
| name | string | 类别名称 | No |
| url | string | 图片url | No |

#### model.IptId

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | integer | 仅包含Id的请求 | No |

#### model.Reply

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| code | integer | _Example:_ `200` | No |
| msg | string | _Example:_ `"提示信息"` | No |

#### model.SysInfo

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| arch | string |  | No |
| num_cpu | integer |  | No |
| os | string |  | No |
| version | string |  | No |

#### model.User

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| cunix | integer | 创建时间 | No |
| email | string | 邮箱 | No |
| id | integer | 主键 | No |
| name | string | 名称 | No |
| num | string | 账号 | No |
| phone | string | 手机 | No |
