package errors

// アプリで発生したエラーをハンドルするための独自エラー
type ApplicationError struct {
	Code  int
	Debug string
	Front string
	Cause error
}

// ApplicationErrorの生成
func New(Code int, Debug string, Front string, Cause error) ApplicationError {
	return ApplicationError{
		Code:  Code,
		Debug: Debug,
		Front: Front,
		Cause: Cause,
	}
}

// デバック用のエラーの説明を返す
func (appError ApplicationError) Error() string {
	return appError.Debug
}

// フロント用のエラーコードを返す
func (appError ApplicationError) GetCode() int {
	return appError.Code
}

// フロント用のエラーの説明を返す
func (appError ApplicationError) GetFrontMessage() string {
	return appError.Debug
}
