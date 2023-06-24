# Go言語での FHIRのサンプル2

- [診療情報提供書HL7FHIR記述仕様](https://std.jpfhir.jp/)に基づくサンプルデータ（Bundle-BundleReferralExample01.json）に対して、汎用的なライブラリのみで、FHIRプロファイルでの検証、パースするサンプルプログラムです。

## プロファイルの検証（バリデーション）とパース
- FHIRプロファイルでの検証
    - Goの場合、[HL7のConfluenceのページ「Open Source Implementations」](https://confluence.hl7.org/display/FHIR/Open+Source+Implementations)で紹介されている、FHIRのリファレンス実装がありません。このため、[FHIR package仕様](https://registry.fhir.org/learn)に従ったnpmパッケージ形式での検証方法は、難しそうです。
    - ですが、[HL7 FHIR v4.0.1:R4のダウンロードページ](https://hl7.org/fhir/R4/downloads.html)に、[JSON Schema形式のファイル](https://hl7.org/fhir/R4/fhir.schema.json.zip)が提供されています。そこで、JSONスキーマによる検証ができる[gojsonschema](https://github.com/xeipuuv/gojsonschema)というライブラリを使って、JSON Schemaによる検証をしてみました。
    - なお、HL FHIRでのバリデーションは複数の方法が提供されており、JSON Schemaもその1つです。方法によって検証可能な内容が若干異なり、公式Validator等に比べるとJSON Schemaで検証できる内容は限定されるようです。
        - [HL7 FHIR:R4 Validating Resources](http://hl7.org/fhir/R4/validation.html)

- 【未実施】JPCoreプロファイル、文書情報プロファイルでの検証
    - JSON Schemaでの提供がされていないことから、上記と同様の理由で、 [JPCoreプロファイル](https://jpfhir.jp/fhir/core/)、[https://std.jpfhir.jp/](https://std.jpfhir.jp/)にある[診療情報提供書の文書情報プロファイル（IGpackage2023.4.27 snapshot形式: jp-ereferral-0.9.6-snap.tgz）](https://jpfhir.jp/fhir/eReferral/jp-ereferral-0.9.7-snap.tgz)レベルの検証を実施するのが難しそうです。

- FHIRデータのパース
    - 前述の通り、Goの場合、FHIRのリファレンス実装がありません。
    - [別のリポジトリのgoのサンプルAP](https://github.com/mysd33/gofhirsample)では、[Golang FHIR Models](https://github.com/samply/golang-fhir-models)というライブラリを使用してみましたが、HAPIのようなリファレンス実装と比較しての信頼性、今後のR5等のFHIRバージョンアップ対応等の将来性が保証がされないことから、ここでは別の手段として、汎用的なJSONライブラリのみでFHIRのパースを実現できないかを検討しています。
    - 最初に、[Go標準のJSONライブラリ(encoding/json)](https://pkg.go.dev/encoding/json)を使用する方法が考えられますが、通常は、FHIRのモデルに対応した各構造体を定義する必要がありますし、interface{}型で動的に参照することもできますが複雑なJSON構造ではデータの抽出の実装がしにくいです。FHIRの特性上、構造体の定義する作業は煩雑なのと、バージョンアップに追従することも考えると、現実的な方法とは言えません。
    - そこで、このサンプルAPでは、[gjson](https://github.com/tidwall/gjson)というサードパーティのライブラリを使って、構造体定義せず、かつJSONのパスの指定で簡単にデータ抽出できる方法で実装しました。

## 実行方法
- 検証・パースするサンプルAPの使い方
    - ビルド後、生成されたexeファイルを実行してください。
```sh
# parsingフォルダへ移動
cd parsing
# ビルド
go build
# 実行
parsing-example.exe
```


## 検証・パースの実行結果の例

```sh
# JSONスキーマチェック結果
2023/06/24 22:13:41 JSON Schema Check: ドキュメントは有効です
# テストデータのパース結果
2023/06/24 22:13:42 文書名: 診療情報提供書
2023/06/24 22:13:42 subject reference Id: urn:uuid:0a48a4bf-0d87-4efb-aafd-d45e0842a4dd
2023/06/24 22:13:42 subject display: 患者リソースPatient
2023/06/24 22:13:42 subject reference type: Patient
2023/06/24 22:13:42 患者番号: 12345
2023/06/24 22:13:42 患者氏名: 田中 太郎
2023/06/24 22:13:42 患者カナ氏名: タナカ タロウ
```