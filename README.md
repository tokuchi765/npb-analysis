# npb-analysis

日本野球機構公式サイトからデータを抽出し分析します。

# 使い方

### 起動
 1. フロントのソースコードをビルドする<br>
 /npb-analysis/frontend/ 配下で `npm run-script build` を実行
 
 2. DB起動<br>
 /npb-analysis/ 配下で `docker-compose up -d` を実行<br>
 ※dockerが使える環境前提です
  
 3. アプリケーション起動<br>
 /npb-analysis/ 配下で `go run .\main.go` を実行する
 
 4. `http://localhost:8081/`  をブラウザで開く
 ![image](https://user-images.githubusercontent.com/55987154/106752847-3f567b00-666e-11eb-9306-097133579c05.png)
 
 # 表示しているデータについて
   分析に使用しているデータは日本プロ野球機構の公式サイトからスクレイピングしたデータを使用してます<br>
 <a href="https://npb.jp/bis/2020/stats/" target="_blank">日本プロ野球機構</a>
 
 # 今後の拡張予定（目標）
  - チーム打撃成績を表示
  - チーム投手成績を表示
  - セイバーメトリクスの数値を算出
  - 機械学習による順位予想機能
