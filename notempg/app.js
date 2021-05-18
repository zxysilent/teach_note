// import fetch from "./api/fetch"
import fetch from "./api/fetch.js";
App({
    onLaunch: function () {
        console.log("onLaunch");
    },
    Get(url, data) {
        return fetch.get(url, data);
    },
    Post(url, data) {
        return fetch.post(url, data);
    },
    /**
     * 全局变量
     */
    globalData: {
        //  这里作为一个全局变量, 方便其它页面调用
    },
});
