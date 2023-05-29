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
table_name,column_name,column_type,is_primary_key,is_unique,is_not_null,foreign_key_table,foreign_key_column,check_constraint,column_comment
```

- `table_name`: テーブル名
- `column_name`: カラム名
- `column_type`: カラムのデータ型
- `is_primary_key`: 主キーである場合は**TRUE**、そうでない場合は**FALSE**
- `is_unique`: UNIQUE制約がある場合は**TRUE**、そうでない場合は**FALSE**
- `is_not_null`: NOT NULL制約がある場合は**TRUE**、そうでない場合は**FALSE**
- `foreign_key_table`: 外部キー制約がある場合、参照先のテーブル名
- `foreign_key_column`: 外部キー制約がある場合、参照先のカラム名
- `check_constraint`: CHECK制約の条件式
- `column_comment`: カラムの説明（コメント

## 実行方法
ビルド済みのバイナリを実行するには、以下のコマンドを使用します。

```bash
./csv2ddl -csv <input_csv_file> -output <output_ddl_file> -dir <dirctory_path> -single
```

- -csv: 入力となるCSVファイルのパス
- -output: 生成されたDDLを保存するファイルのパス。省略した場合は、入力CSVファイルと同じ名前になります。
- -dir: ディレクトリパスを指定し、そのディレクトリ内にある全CSVに対してcsv2ddlを実行します。出力されるDDLファイル名はCSVと同名になります。(-dirオプションをつけた場合-csvオプションは不要)
- -single: -dirと同時に使用。つけると1つのDDLファイルに出力します。

例:

```bash
./csv2ddl -csv input.csv -output output.sql
```

これにより、input.csvファイルを読み込み、output.sqlファイルにMySQLのDDLが生成されます。

```bash
./csv2ddl -dir ./csv/
```

これにより、csvディレクトリ以下の全CSVファイルからDDLファイルを生成します。