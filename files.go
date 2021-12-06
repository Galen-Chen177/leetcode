package leetcode

import "os"

// 将str保存到文件中
func SaveFileByStr(str, filename string) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	//清空文件
	if err := f.Truncate(0); err != nil {
		return err
	}

	//设置偏移量
	if _, err := f.Seek(0, 0); err != nil {
		return err
	}

	if _, err := f.WriteString(str); err != nil {
		return err
	}
	return nil
}

//判断文件是否存在
func FileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
