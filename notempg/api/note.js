import fetch from "./fetch"

// 笔记
export const apiNoteGet = (id) => {
	return fetch.get("/api/note/get", { id: id })
}
export const apiNoteAll = () => {
	return fetch.get("/api/note/all")
}
export const apiNoteAdd = (data) => {
	return fetch.post("/api/note/add", data)
}
export const apiNoteEdit = (data) => {
	return fetch.post("/api/note/edit", data)
}
export const apiNoteDrop = (id) => {
	return fetch.post("/api/note/drop", { id: id })
}