package model

// Feedback 文章
type Feedback struct {
	Id      int    `json:"id"`      //主键
	Content string `json:"content"` //内容
	Ctime   int64  `json:"ctime"`   //时间
}

// FeedbackAll 查询所有反馈
func FeedbackAll() ([]Feedback, error) {
	mods := make([]Feedback, 0, 4)
	err := Db.Select(&mods, "SELECT * FROM Feedback")
	return mods, err
}

// FeedbackAdd 添加反馈
func FeedbackAdd(mod *Feedback) error {
	Tx, err := Db.Beginx()
	if err != nil {
		return err
	}
	_, err = Tx.NamedExec("INSERT INTO Feedback(content,ctime) VALUES(:content,:ctime)", mod)
	if err != nil {
		Tx.Rollback()
		return err
	}
	Tx.Commit()
	return nil
}
