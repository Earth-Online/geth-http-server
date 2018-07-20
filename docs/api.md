~~~[api]
post:/api/update
*Address#账号地址
*Password#原密码
*NewPassword#新密码
<<<
success
{
    "error": 0,
    "info": "更新密码成功",
}
<<<
error 
500
{
    "error": 1,
    "info": "具体错误信息",
}
~~~

~~~[api]
post:/api/new_account
*Password#账号密码
<<<
success
{
    "error": 0,
    "PrivateKey": "账号私钥",
    "Address":"账号地址"
}
<<<
error 
500
{
    "error": 1,
    "info": "具体错误信息",
}
~~~

~~~[api]
post:/api/import
*Keydata#私钥
*Password#加密密码
<<<
success
{
    "error": 0,
    "info": "账号地址",
}
<<<
error 
500
{
    "error": 1,
    "info": "具体错误信息",
}
~~~

~~~[api]
post:/api/query
*Address#查询地址
<<<
success
{
    "error": 0,
    "info": [{"from":"发送地址","to":"接受地址","gasUsed":"矿工费","hash":"交易号","value":"金额"}],
}
<<<
error 
500
{
    "error": 1,
    "info": "具体错误信息",
}
~~~

~~~[api]
post:/api/check
*address#接受地址
*hash#交易号
*info#备注信息
*value#金额
<<<
success
{
    "error": 0,
    "info": "确认成功"
}
<<<
error 
500
{
    "error": 1,
    "info": "具体错误信息",
}

~~~

~~~[api]
post:/api/get_balance
*address#地址
<<<
success
{
    "error": 0,
    "info": "余额wei为单位"
}
<<<
error 
500
{
    "error": 1,
    "info": "具体错误信息",
}

~~~

~~~[api]
get:/api/Price
<<<
success
{
    "error": 0,
    "info": "现在以太币美元价格"
}
<<<
error 
500
{
    "error": 1,
    "info": "具体错误信息",
}

~~~

~~~[api]
post:/api/send
*from#发送账号
*to#接受账号
*value#金额
*password#发送账号密码
gas#准备的矿工费
input#备注信息
gasPrice#每种付费gas的价格
<<<
success
{
    "error": 0,
    "info": "交易号"
}
<<<
error 
500
{
    "error": 1,
    "info": "具体错误信息",
}

~~~









