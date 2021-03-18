package gopeg_mv4

import (
	"io/ioutil"
	"path/filepath"
)

// 結合処理(結合ディレクトリパス, 結果保存ディレクトリパス)


// 結合ディレクトリパスから結合ファイルのパスを一覧取得
// reference: https://qiita.com/tanksuzuki/items/7866768c36e13f09eedb
func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths
}