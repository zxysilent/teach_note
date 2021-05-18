var app = getApp();
import { apiFeedbackAdd } from "../../api/feedback"
Page({
  data: {
    content: ""
  },
  onSubmit(event) {
    console.log(event)
    if (!event.detail.value.content) {
      wx.showToast({
        icon: "error",
        title: "请填写反馈内容"
      });
      return;
    }
    this.setData({
      content: event.detail.value.content
    })
    apiFeedbackAdd({ content: event.detail.value.content }).then(resp => {
      // app.Post("/api/feedback/add", { content: event.detail.value.content }).then(resp => {
      if (resp.code == 200) {
        wx.showToast({
          title: "提交成功",
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
          title: "提交失败"
        });
      }
    })
  },
  onReset(event) {
    console.log(event)
    this.setData({
      content: ""
    })
  }
})
