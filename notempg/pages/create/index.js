
var app = getApp();
import { apiNoteAdd } from "../../api/note"
Page({
    data: {
        focus: true,
        form: {
            title: "",
            content: ""
        }
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
            form: event.detail.value
        })
        apiNoteAdd(this.data.form).then(resp => {
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
});