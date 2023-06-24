package main

import (
	"io/ioutil"
	"log"

	"github.com/tidwall/gjson"
	"github.com/xeipuuv/gojsonschema"
)

func main() {
	// 診療情報提供書のHL7 FHIRのサンプルデータBundle-BundleReferralExample01.jsonを読み込み
	fileData, err := ioutil.ReadFile("Bundle-BundleReferralExample01.json")
	if err != nil {
		log.Fatal(err) //終了
	}
	// HL7 FHIRのJSONスキーマfhir.schema.json
	schemaData, err := ioutil.ReadFile("fhir.schema.json")
	if err != nil {
		log.Fatal(err) //終了
	}
	schemaLoader := gojsonschema.NewBytesLoader(schemaData)
	// 診療情報提供書のHL7 FHIRのサンプルデータBundle-BundleReferralExample01.json
	docummentLoader := gojsonschema.NewBytesLoader(fileData)
	// JSONスキーマチェック
	result, err := gojsonschema.Validate(schemaLoader, docummentLoader)
	if err != nil {
		log.Fatal(err)
	}
	if result.Valid() {
		log.Println("JSON Schema Check: ドキュメントは有効です")
	} else {
		//検証エラー
		log.Println("JSON Schema Check: ドキュメントに不備があります")
		for _, desc := range result.Errors() {
			log.Printf("- %s\n", desc)
		}
	}

	// 最初のEntryであるCompositionリソースの解析する例
	composition := gjson.GetBytes(fileData, "entry.0.resource")

	log.Printf("文書名: %s", composition.Get("title"))
	subject := composition.Get("subject")
	subjefctRefId := subject.Get("reference")
	log.Printf("subject reference Id: %s", subjefctRefId)
	log.Printf("subject display: %s", subject.Get("display"))
	log.Printf("subject reference type: %s", subject.Get("type"))

	// Compostion.subjectが参照するPatientの解析する例
	patientEntry := gjson.GetBytes(fileData, "entry.1")
	patientFullUrl := patientEntry.Get("fullUrl")
	if patientFullUrl.String() != subjefctRefId.String() {
		log.Fatal("患者情報のUUIDが不一致です")
	}
	patient := patientEntry.Get("resource")
	// 患者番号の取得
	log.Printf("患者番号: %s", patient.Get("identifier.0.value"))
	// 患者氏名の取得
	patientNames := patient.Get("name")
	patientNames.ForEach(func(key, value gjson.Result) bool {
		if value.Get("extension.0.valueCode").String() == "IDE" {
			log.Printf("患者氏名: %s", value.Get("text"))
		} else {
			log.Printf("患者カナ氏名: %s", value.Get("text"))
		}
		return true
	})

}
