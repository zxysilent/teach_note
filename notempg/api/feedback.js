import fetch from "./fetch"

// 添加反馈
export const apiFeedbackAdd = (data) => {
  return fetch.post("/api/feedback/add", data)
}