# Go言語での FHIRのサンプル

- [診療情報提供書HL7FHIR記述仕様](https://std.jpfhir.jp/)に基づくサンプルデータ（Bundle-BundleReferralExample01.json）に対して、FHIRプロファイルでの検証、パースするサンプルプログラムです。

## プロファイルの検証（バリデーション）とパース
- FHIRプロファイルでの検証
    - Goの場合、[HL7のConfluenceのページ「Open Source Implementations」](https://confluence.hl7.org/display/FHIR/Open+Source+Implementations)で紹介されている、FHIRのリファレンス実装がありません。
    - また、その他、検索しても、Goでは、JavaのHAPI等と違い、FHIRの構造定義ファイルでの検証を行うライブラリがなさそうです。
    - ですが、[HL7 FHIR v4.0.1:R4のダウンロードページ](https://hl7.org/fhir/R4/downloads.html)に[JSON Schema形式のファイル](https://hl7.org/fhir/R4/fhir.schema.json.zip)が提供されています。        
    - そこで、JSONスキーマによる検証ができる[gojsonschema](https://github.com/xeipuuv/gojsonschema)というライブラリを使って、JSONスキーマの検証をしています。
- 【未実施】JPCoreプロファイル、文書情報プロファイルでの検証
    - [別のrepositoryにあるサンプルAP](https://github.com/mysd33/gofhirsample)に記載の通り、 [JPCoreプロファイル](https://jpfhir.jp/fhir/core/)、[診療情報提供書の文書情報プロファイル（IGpackage2023.4.27 snapshot形式: jp-ereferral-0.9.6-snap.tgz）](https://jpfhir.jp/fhir/eReferral/jp-ereferral-0.9.7-snap.tgz)レベルの検証を実施するのが難しそうです。
- FHIRデータのパース
    - 前述の通り、Goの場合、[HL7のConfluenceのページ「Open Source Implementations」](https://confluence.hl7.org/display/FHIR/Open+Source+Implementations)で紹介されている、FHIRのリファレンス実装がありません。
    - [別のリポジトリのgoのサンプルAP](https://github.com/mysd33/gofhirsample)では、[Golang FHIR Models](https://github.com/samply/golang-fhir-models)というライブラリを使用していますが、ライブラリの信頼性、今後のR5等のFHIRバージョンアップ対応等の将来性が期待できないことから、そういったライブラリを使用せずに、JSONパースを実現できないかを検討します。ただ、[Go標準のJSONライブラリ(encoding/json)](https://pkg.go.dev/encoding/json)を使用するには、FHIRのモデルに対応した各構造体を定義する必要がありますが、FHIRの特性上、構造体の定義する作業は煩雑なのと、バージョンアップに追従する負担もあり、現実的ではありません。
    - そこで、このサンプルAPでは、[gjson](https://github.com/tidwall/gjson)というサードパーティのライブラリを使って、JSONのパスをしているすることで、構造体定義せずにパースできる方法で、実装します。

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