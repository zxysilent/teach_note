const utils = {};
const tokenKey = "mpg-app-token-key";
//保存数据
utils.setToken = (token) => {
    wx.setStorageSync(tokenKey, token);
};
utils.getToken = () => {
    return wx.getStorageSync(tokenKey);
};
utils.removeToken = () => {
    wx.removeStorageSync(tokenKey);
};

// yyyy-MM-dd hh:mm:ss
utils.fmtDate = (date, fmt) => {
    date = new Date(date);
    var o = {
        "M+": date.getMonth() + 1, //月份
        "d+": date.getDate(), //日
        "h+": date.getHours(), //小时
        "m+": date.getMinutes(), //分
        "s+": date.getSeconds(), //秒
    };
    if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (date.getFullYear() + "").substr(4 - RegExp.$1.length));
    for (var k in o) if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, RegExp.$1.length == 1 ? o[k] : ("00" + o[k]).substr(("" + o[k]).length));
    return fmt;
};
utils.redirect = (rawUrl) => {
    // 带问号的，用 redirectTo，不带问号的，用switchTab
    if (rawUrl.indexOf("?") > -1) {
        wx.redirectTo({
            url: rawUrl,
        });
    } else {
        wx.switchTab({
            url: rawUrl,
        });
    }
};
export default utils;
