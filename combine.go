package gopeg_mv4

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// 結合処理(結合ディレクトリパス, 結果保存ディレクトリパス)
func CombineMp4(fileName, inputPath, outputPath string) error {
	if err := WriteTxt(dirwalk(inputPath)); err != nil {
		panic(err)
	}
	log.Println("処理開始")
	cmd := exec.Command("ffmpeg", "-f", "concat", "-i", "vr_memo.txt", fmt.Sprintf("%s/%s", outputPath, fileName))
	cmd.Start()
	log.Println("処理実行中")
	cmd.Wait()
	log.Println("処理終了")
	return nil
}

func WriteTxt(paths []string) error {
	file, err := os.Create("vr_memo.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var content string
	for i := range paths {
		content += fmt.Sprintf("file '%s'\n", paths[i])
	}

	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}

	return err
}

func RemoveTxt() error {
	err := os.Remove("vr_memo.txt")
	if err != nil {
		panic(err)
	}
	return err
}

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