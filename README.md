# CSV2DDL

このツールは、CSVファイルを読み込み、MySQLのテーブル定義（DDL）を生成するためのコマンドラインツールです。

## ファイル説明
- `main.go`: エントリーポイント。コマンドライン引数を解析し、CSVファイルを読み込み、DDLを生成してファイルに出力します。
- `utils.go`: ユーティリティ関数を提供します。
- `models.go`: テーブルやカラムなどのデータ構造を定義します。
- `csv_parser.go`: CSVファイルを解析し、テーブル情報を抽出します。
- `ddl_generator.go`: 抽出されたテーブル情報をもとにMySQLのDDLを生成します。

## CSVヘッダーの説明
CSVファイルのフォーマットは以下の通りです。

```csv
table_name,column_name,type,is_primary_key,is_not_null,is_unique,foreign_key_table,foreign_key_column,check,comment
```

- `table_name`: テーブル名
- `column_name`: カラム名
- `type`: カラムのデータ型
- `is_primary_key`: 主キーである場合は**TRUE**、そうでない場合は**FALSE**
- `is_not_null`: NOT NULL制約がある場合は**TRUE**、そうでない場合は**FALSE**
- `is_unique`: UNIQUE制約がある場合は**TRUE**、そうでない場合は**FALSE**
- `foreign_key_table`: 外部キー制約がある場合、参照先のテーブル名
- `foreign_key_column`: 外部キー制約がある場合、参照先のカラム名
- `check`: CHECK制約の条件式
- `comment`: カラムの説明（コメント

## 実行方法
ビルド済みのバイナリを実行するには、以下のコマンドを使用します。

```bash
./csv-to-mysql-ddl -csv <input_csv_file> -output <output_ddl_file>
```

- -csv: 入力となるCSVファイルのパス
- -output: 生成されたDDLを保存するファイルのパス。省略した場合は、入力CSVファイルと同じ名前になります。

例:

```bash
./csv-to-mysql-ddl -csv input.csv -output output.sql
```

これにより、input.csvファイルを読み込み、output.sqlファイルにMySQLのDDLが生成されます。