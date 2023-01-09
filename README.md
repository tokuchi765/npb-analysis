# npb-analysis

日本野球機構公式サイトからデータを抽出し分析します。

# 推奨環境
- windows
- golang 1.18.3
- node.js 16.14.0 LTS
  - https://nodejs.org/ja/

# 使い方

### 起動
1. フロントのソースコードをビルドする<br>
   - `/npb-analysis/frontend/` 配下で `yarn run-script build` を実行
 
2. DB起動<br>
    - `docker-compose.yml.example` ファイルの `volumes` をインストールしたディレクトリに書きかえる
    - ファイル名を `docker-compose.yml` に書き換える
    - `/npb-analysis/` 配下で `docker compose up -d` を実行<br>
 ※dockerが使える環境前提です<br>
  
1. アプリケーション起動<br>
    - `/npb-analysis/` 配下で `go run .\main.go` を実行する
 
4. `http://localhost:8081/`  をブラウザで開く
![image](https://user-images.githubusercontent.com/55987154/156882493-b333037b-a9ea-4740-b0ca-4fff1334262a.png)
 
5. テストコード実行
   - `test\testUtil.go` の `/home/runner/work/` をインストールしたディレクトリに書き換える<br>
     ※windowsの場合は `file:C:/home/xxx` の形式で書き換える

 # 表示しているデータについて
   分析に使用しているデータは日本プロ野球機構の公式サイトからスクレイピングしたデータを使用してます<br>
 <a href="https://npb.jp/bis/2020/stats/" target="_blank">日本プロ野球機構</a>
 
 # 今後の拡張予定（目標）
  - セイバーメトリクスの数値を算出
  - 機械学習による順位予想機能
