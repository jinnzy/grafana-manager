package manager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/grafana-tools/sdk"
	"grafana-manager/pkg/logger"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func (m *Manager)ImportDashboard(dir string)  {
	var (
		filesInDir []os.FileInfo
		rawBoard   []byte
		err        error
		folder    sdk.Folder
		folderID = 0
	)
	baseDir := filepath.Base(dir)
	ctx := context.Background()
	c := sdk.NewClient(m.URL, m.BasicAuthOrToken, sdk.DefaultHTTPClient)
	filesInDir, err = ioutil.ReadDir(dir)
	if err != nil {
		logger.Fatal(err)
	}
	searchParams := sdk.SearchParam(func(values *url.Values) {
		values.Add("query",baseDir)
	})
	data, err := c.Search(ctx, searchParams)
	if err != nil {
		logger.Errorf("get folder err: ", err)
	}
	if len(data) == 0 {
		if baseDir == "" || baseDir == "." {
			// 目录为空或者.则使用默认目录
			folderID = 0
		}else {
			folder = sdk.Folder{
				Title: baseDir,
			}
			folder, err = c.CreateFolder(ctx, folder)
			if err != nil {
				log.Fatal(err)
			}
			folderID = folder.ID
		}
	}else {
		folderID = int(data[0].ID)
	}


	for _, file := range filesInDir {
		if strings.HasSuffix(file.Name(), ".json") {
			// 遍历得到的只是文件名，加上目录组合成相对路径。
			if rawBoard, err = ioutil.ReadFile(filepath.Join(dir,file.Name())); err != nil {
				logger.Errorf("read file %s err: %s", filepath.Join(dir,file.Name()), err)
				continue
			}
			var board sdk.Board
			if err = json.Unmarshal(rawBoard, &board); err != nil {
				logger.Errorf("unmarshal %s err: %s", filepath.Join(dir,file.Name()),err)
				continue
			}
			if _, err = c.DeleteDashboard(ctx, board.UpdateSlug()); err != nil {
				logger.Errorf( "error on deleting dashboard  %s with %s", filepath.Join(dir,file.Name()), err)
				continue
			}
			params := sdk.SetDashboardParams{
				FolderID:  folderID,
				Overwrite: false,
			}
			_, err := c.SetDashboard(ctx, board, params)
			if err != nil {
				logger.Errorf("error on importing dashboard %s. err: %s", board.Title, err.Error())
				continue
			}
		}
	}
}

func (m *Manager)ImportDatasource(dir string)  {
	var (
		datasources []sdk.Datasource
		filesInDir  []os.FileInfo
		rawDS       []byte
		status      sdk.StatusMessage
		err         error
	)
	fmt.Println(dir)
	ctx := context.Background()
	c := sdk.NewClient(m.URL, m.BasicAuthOrToken, sdk.DefaultHTTPClient)
	if datasources, err = c.GetAllDatasources(ctx); err != nil {
		logger.Errorf("import data source err: %s", err)
		os.Exit(1)
	}
	filesInDir, err = ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	for _, file := range filesInDir {
		if strings.HasSuffix(file.Name(), ".json") {
			if rawDS, err = ioutil.ReadFile(filepath.Join(dir,file.Name())); err != nil {
				logger.Errorf("read file %s err: %s", filepath.Join(dir,file.Name()),err)
				continue
			}
			var newDS sdk.Datasource
			if err = json.Unmarshal(rawDS, &newDS); err != nil {
				logger.Errorf("unmarshal %s err: %s", filepath.Join(dir,file.Name()),err)
				continue
			}
			for _, existingDS := range datasources {
				if existingDS.Name == newDS.Name {
					if status, err = c.DeleteDatasource(ctx, existingDS.ID); err != nil {
						logger.Errorf( "error on deleting datasource %s with %s", newDS.Name, err)
					}
					break
				}
			}
			if status, err = c.CreateDatasource(ctx, newDS); err != nil {
				logger.Errorf("error on importing datasource %s with %s (%s)", newDS.Name, err, *status.Message)
			}
		}
	}
}
