package localizations

const (
	En = "en"
	Th = "th"
)

type Localization map[string]string

type Localizations map[string]Localization

const (
	CommonInvalidData                = "common_invalid_data"
	CommonNotFoundData               = "common_not_found"
	CommonFieldIsRequiredAndNotEmpty = "common_field_is_required_and_not_empty"
	CommonCannotDeletePleaseTryAgain = "common_cannot_delete_please_try_again"
	CommonDeleteSuccess              = "common_delete_success"
	CommonDataIsDuplicated           = "common_data_is_duplicated"
	CommonCannotAddData              = "common_cannot_add_data"
	CommonCannotUpdateData           = "common_cannot_update_data"
	CommonCannotDeleteData           = "common_cannot_delete_data"
	CommonThereIsNoDataUpdate        = "common_there_is_no_data_update"
	CommonPagingInvalid              = "common_paging_invalid"
	CommonDataDuplicated             = "common_data_duplicated"
)

var Localizes = Localizations{
	En: Localization{
		CommonInvalidData:                "Invalid data",
		CommonNotFoundData:               "Not found data",
		CommonFieldIsRequiredAndNotEmpty: "%s is required and not empty",
		CommonCannotDeletePleaseTryAgain: "Can't delete, Please try again",
		CommonDeleteSuccess:              "Delete success",
		CommonDataIsDuplicated:           "Data is duplicated",
		CommonCannotAddData:              "Cannot add data",
		CommonCannotUpdateData:           "Cannot update data",
		CommonCannotDeleteData:           "Cannot delete data",
		CommonThereIsNoDataUpdate:        "There is no data to update",
		CommonPagingInvalid:              "Page or Limit data invalid",
		CommonDataDuplicated:             "Some of the data is duplicated",
	},
	Th: Localization{
		CommonInvalidData:                "ข้อมูลไม่ถูกต้อง",
		CommonNotFoundData:               "ไม่พบข้อมูล",
		CommonFieldIsRequiredAndNotEmpty: "กรุณากรอกข้อมูล %s และไม่ให้เป็นค่าว่าง",
		CommonCannotDeletePleaseTryAgain: "ไม่สามารถลบได้ กรุณาลองใหม่อีกครั้ง",
		CommonDeleteSuccess:              "ลบสำเร็จ",
		CommonDataIsDuplicated:           "ข้อมูลถูกทำซ้ำ",
		CommonCannotAddData:              "ไม่สามารถเพิ่มข้อมูลได้",
		CommonCannotUpdateData:           "ไม่สามารถอัปเดตข้อมูลได้",
		CommonCannotDeleteData:           "ไม่สามารถลบข้อมูลได้",
		CommonThereIsNoDataUpdate:        "ไม่มีข้อมูลที่จะอัปเดต",
		CommonPagingInvalid:              "จำนวนหน้าและการจำกัดข้อมูลไม่ถูกต้อง",
		CommonDataDuplicated:             "ข้อมูลบางส่วนซ้ำกัน",
	},
}

func Translate(locale string, key string) string {
	return Localizes[locale][key]
}
