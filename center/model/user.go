package model

import (
	"fmt"
	"time"
)

const (
	//TABLE 用户表
	TABLE = "user"
)


//User 用户结构体
type User struct {
	UID        int32
	Username   string
	Password   string
	AppID      int32
	CreateTime int32
	UpdateTime int32
}

// GetUser 获取用户信息
func (m *Model) GetUser(uid int32, username string) (*User, error) {
	var user User
	var whereStr string
	if uid > 0 {
		whereStr = fmt.Sprintf("uid = %d", uid)
	} else if len(username) > 0 {
		whereStr = fmt.Sprintf("username = %s", username)
	}

	sqlStr := fmt.Sprintf("select uid,username,password,salt,app_id,create_time from %s where %s limit 1", TABLE, whereStr)
	rows, err := m.db.Query(sqlStr, uid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&user.UID, &user.Username, &user.Password, &user.AppID, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser 创建用户
func (m *Model) CreateUser(appID int32, username string, password string, salt string) (int64, error) {
	createTime := time.Now().Unix()
	updateTime := time.Now().Unix()
	sqlStr := "insert into user (app_id, username, password, create_time, update_time, salt) values (?, ?, ?, ?, ?, ?)"
	rs, err := m.db.Exec(sqlStr, appID, username, password, createTime, updateTime, salt)
	if err != nil {
		return 0, err
	}
	uid, err := rs.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uid, nil
}

// UpdateUser 更新用户
func (m *Model) UpdateUser(uid int32, username string, password string) error {
	set := ""
	if len(username) > 0 {
		set += fmt.Sprintf(" username = %s", username)
	}
	if len(password) > 0 {
		set += fmt.Sprintf(" password = %s", password)
	}
	updateTime := time.Now().Unix()
	set += fmt.Sprintf(" update_time = %d", updateTime)
	sqlStr := fmt.Sprintf("update user set %s where uid = %d", set, uid)
	_, err := m.db.Exec(sqlStr)
	if err != nil {
		return err
	}
	return nil
}
