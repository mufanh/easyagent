package model

type ScriptUploadRequest struct {
	Token   string `json:"token"`
	Content string `json:"content"`
	// 脚本分组名（脚本上一级目录名）
	GroupDir string `json:"group_dir"`
	// 脚本名
	Name string `json:"name"`
	// 脚本对应的操作系统
	OS string `json:"os"`
}

type ScriptUploadResponse struct {
	BaseResponse
}

type ShowScriptRequest struct {
	Token string `json:"token"`
	// 脚本分组名（脚本上一级目录名）
	GroupDir string `json:"group_dir"`
	// 脚本名
	Name string `json:"name"`
}

type ShowScriptResponse struct {
	BaseResponse
	Content string `json:"content"`
}

type DeleteScriptRequest struct {
	Token    string `json:"token"`
	GroupDir string `json:"group_dir"`
	Name     string `json:"name"`
}

type DeleteScriptResponse struct {
	BaseResponse
}

type ScriptUpdateRequest struct {
	Token   string `json:"token"`
	Content string `json:"content"`
	// 脚本分组名（脚本上一级目录名）
	GroupDir string `json:"group_dir"`
	// 脚本名
	Name string `json:"name"`
}

type ScriptUpdateResponse struct {
	BaseResponse
}

type DeleteScriptGroupDirRequest struct {
	Token    string `json:"token"`
	GroupDir string `json:"group_dir"`
}

type DeleteScriptGroupDirResponse struct {
	BaseResponse
}

type ScriptExecRequest struct {
	Token string `json:"token"`
	// 脚本分组名（脚本上一级目录名）
	GroupDir string `json:"group_dir"`
	// 脚本名
	Name string `json:"name"`

	Async   bool   `json:"async"`
	Logfile string `json:"logfile"`
}

type ScriptExecResponse struct {
	BaseResponse
	// 若async=false，那么日志会直接记录到该字段返回
	Log string `json:"log,omitempty"`
}

type ScriptShowGroupDirsRequest struct {
	Token string `json:"token"`
}

type ScriptShowGroupDirsResponse struct {
	BaseResponse
	GroupDirs []string `json:"group_dirs"`
}

type ScriptShowFilesRequest struct {
	Token    string `json:"token"`
	GroupDir string `json:"group_dir"`
}

type ScriptShowFilesResponse struct {
	BaseResponse
	ScriptFiles []string `json:"script_files"`
}
