package model

// Note 笔记
type Note struct {
	Id      int    `json:"id"`      //主键
	Title   string `json:"title"`   //标题
	Content string `json:"content"` //详细
	Utime   int64  `json:"utime"`   //修改时间
	Ctime   int64  `json:"ctime"`   //创建时间
}

// NoteGet 查询一条笔记
func NoteGet(id int) (*Note, error) {
	mod := &Note{}
	err := Db.Get(mod, "SELECT * FROM Note WHERE id = ?", id)
	return mod, err
}

// NoteAll 查询所有笔记
func NoteAll() ([]Note, error) {
	mods := make([]Note, 0, 4)
	err := Db.Select(&mods, "SELECT * FROM Note")
	return mods, err
}

// NoteAdd 添加笔记
func NoteAdd(mod *Note) error {
	Tx, err := Db.Beginx()
	if err != nil {
		return err
	}
	_, err = Tx.NamedExec("INSERT INTO Note(title,content,utime,ctime) VALUES(:title,:content,:utime,:ctime)", mod)
	if err != nil {
		Tx.Rollback()
		return err
	}
	Tx.Commit()
	return nil
}

// NoteEdit 编辑笔记
func NoteEdit(mod *Note) error {
	Tx, err := Db.Beginx()
	if err != nil {
		return err
	}
	_, err = Tx.NamedExec("UPDATE Note SET title=:title,content=:content,utime=:utime WHERE id=:id", mod)
	if err != nil {
		Tx.Rollback()
		return err
	}
	Tx.Commit()
	return nil
}

// NoteDrop 删除笔记
func NoteDrop(id int) error {
	Tx, err := Db.Beginx()
	if err != nil {
		return err
	}
	_, err = Tx.Exec(`DELETE FROM Note WHERE id=?`, id)
	if err != nil {
		Tx.Rollback()
		return err
	}
	Tx.Commit()
	return nil
}
