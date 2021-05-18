var app = getApp();
import {
    apiNoteAll
} from "../../api/note"
import Utils from "../../utils/utils"
Page({
    data: {
        notes: [],
    },
    onShow: function () {
        this.init()
    },
    onAdd: function (event) {
        wx.navigateTo({
            url: "../create/index",
        });
    },
    // 加载数据
    init() {
        apiNoteAll().then(resp => {
            if (resp.code == 200) {
                resp.data = resp.data.map(item => {
                    item.ctime = Utils.fmtDate(item.ctime * 1000, 'yyyy-MM-dd')
                    return item
                })
                this.setData({
                    notes: resp.data
                })
            }
        })
    },
    onEdit: function (event) {
        console.log("onEdit", event)
        wx.navigateTo({
            url: "../edit/index?id=" + event.currentTarget.dataset.id,
        });
    },
    onPullDownRefresh: function () {
        wx.showToast({
            title: "正在加载数据",
            icon: "loading",
        });
        this.init()
        setTimeout(() => {
            wx.hideToast({})
            wx.stopPullDownRefresh();
        }, 1500)
    },
});