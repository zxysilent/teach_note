
var app = getApp();
import { apiNoteGet, apiNoteEdit, apiNoteDrop } from "../../api/note"
Page({
    data: {
        form: {
            id: 0,
            title: "",
            content: ""
        },
        focus: true
    },
    onLoad: function (options) {
        console.log(options)
        this.init(options.id)
    },
    onSubmit: function (event) {
        console.log(event, this.data.form)
        if (!event.detail.value.title) {
            wx.showToast({
                icon: "error",
                title: "请填写标题"
            });
            return;
        }
        if (!event.detail.value.content) {
            wx.showToast({
                icon: "error",
                title: "请填写内容"
            });
            return;
        }
        this.setData({
            "form.title": event.detail.value.title,
            "form.content": event.detail.value.content
        })
        apiNoteEdit(this.data.form).then(resp => {
            if (resp.code == 200) {
                wx.showToast({
                    title: "保存成功",
                    duration: 3000,
                    complete: () => {
                        setTimeout(() => {
                            wx.navigateBack();
                        }, 3000)
                    }
                });
            } else {
                wx.showToast({
                    icon: "error",
                    title: "保存失败"
                });
            }
        })
    },
    onDrop(event) {
        console.log(event)
        apiNoteDrop(this.data.form.id).then(resp => {
            if (resp.code == 200) {
                wx.showToast({
                    title: "删除成功",
                    duration: 3000,
                    complete: () => {
                        setTimeout(() => {
                            wx.switchTab({
                                url: "../index/index"
                            })
                        }, 3000)
                    }
                });
            } else {
                wx.showToast({
                    icon: "error",
                    title: "删除失败"
                });
            }
        });
    },
    init(id) {
        apiNoteGet(id).then(resp => {
            if (resp.code == 200) {
                this.setData({
                    form: resp.data
                })
            }
        })
    }
});