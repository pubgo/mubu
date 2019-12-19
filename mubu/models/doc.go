package models

type GetDoc struct {
	Code int `json:"code"`
	Data struct {
		BaseVersion int    `json:"baseVersion"`
		Definition  string `json:"definition"`
		Name        string `json:"name"`
		Role        int    `json:"role"`
	} `json:"data"`
	Msg string `json:"msg"`
}

type ListDoc struct {
	Code int `json:"code"`
	Data struct {
		Dir []struct {
			CommitClient string `json:"commitClient"`
			CreateTime   int    `json:"createTime"`
			DeleteTime   int    `json:"deleteTime"`
			Deleted      int    `json:"deleted"`
			Encrypted    int    `json:"encrypted"`
			FolderID     string `json:"folderId"`
			ID           string `json:"id"`
			MetaVersion  int    `json:"metaVersion"`
			Name         string `json:"name"`
			StarIndex    int    `json:"starIndex"`
			Stared       int    `json:"stared"`
			UpdateTime   int    `json:"updateTime"`
			UserID       int    `json:"userId"`
			Version      int    `json:"version"`
		} `json:"dir"`
		Documents []struct {
			Deleted    int    `json:"deleted"`
			Encrypted  int    `json:"encrypted"`
			FolderID   string `json:"folderId"`
			ID         string `json:"id"`
			Name       string `json:"name"`
			Role       int    `json:"role"`
			ShareID    string `json:"shareId"`
			Stared     int    `json:"stared"`
			UpdateTime int    `json:"updateTime"`
			UserID     int    `json:"userId"`
			UserName   string `json:"userName"`
		} `json:"documents"`
		FolderID string `json:"folderId"`
		Folders  []struct {
			CommitClient string `json:"commitClient"`
			CreateTime   int    `json:"createTime"`
			DeleteTime   int    `json:"deleteTime"`
			Deleted      int    `json:"deleted"`
			Encrypted    int    `json:"encrypted"`
			FolderID     string `json:"folderId"`
			ID           string `json:"id"`
			MetaVersion  int    `json:"metaVersion"`
			Name         string `json:"name"`
			StarIndex    int    `json:"starIndex"`
			Stared       int    `json:"stared"`
			UpdateTime   int    `json:"updateTime"`
			UserID       int    `json:"userId"`
			Version      int    `json:"version"`
		} `json:"folders"`
	} `json:"data"`
	Msg interface{} `json:"msg"`
}
