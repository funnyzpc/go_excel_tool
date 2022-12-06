# exce 批量解密工具 go_excel_tool

## 简要说明
以下文件均为go_excel_tool不同平台的编译结果,具体平台可见文件夹名可知
go_excel_tool_linux
go_excel_tool_macos_intel
go_excel_tool_window.exe

## 使用前提
加密的xlsx文件与密码文件必须在同一个目录，密码文件必须包含“密码”或“password” 字样

### window下使用
0.windiw下使用需要用到exec.bat
1.调整exec.bat内DIR(目录位置)
2.双击执行exec.bat
3.DEC_FILE目录下文件即为解密数据

### linux or macOs使用
0.需要给对应可执行应用赋予可执行权限 例: chmod a+x go_excel_tool_linux
1.直接执行文件命令或编写shell脚本文件
2.主要执行命令 go_excel_tool_window -dir="数据文件目录位置"
3.DEC_FILE目录下文件即为解密数据
