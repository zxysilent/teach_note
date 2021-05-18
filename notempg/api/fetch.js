// 服务地址
import Utils from "../utils/utils";
const baseUrl = "http://127.0.0.1:8085" //本地开发
// const baseUrl = "http://127.0.0.1:8085" //测试开发
// const baseUrl = "http://127.0.0.1:8085"; //正式服务
const request = (method, url, data) => {
    return new Promise((resolve, reject) => {
        method = method || "GET";
        wx.showNavigationBarLoading();
        wx.request({
            url: baseUrl + url,
            data: data,
            method: method,
            header: {
                // "content-type": "application/json; charset=UTF-8" //默认
                Authorization: Utils.getToken(),
            },
            timeout: 10000, //10s,
            complete: (resp) => {
                wx.hideNavigationBarLoading();
                if (resp.statusCode >= 200 && resp.statusCode < 300) {
                    if (resp.data.code == 340) {
                        Utils.removeToken();
                    }
                    resolve(resp.data);
                } else {
                    reject(resp);
                }
            },
            fail: (err) => {
                wx.showToast({
                    title: "网络异常,请重试",
                    icon: "none",
                    duration: 5000,
                });
                console.log(err);
            },
        });
    });
};
const fetch = {
    // get
    get: (url, data) => {
        return request("GET", url, data);
    },
    // post
    post: (url, data) => {
        return request("POST", url, data);
    },
};

export default fetch;
