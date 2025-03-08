@echo off
chcp 936
setlocal EnableDelayedExpansion

:: 设置服务目录列表
set services=auth biz connector gateway im push

:: 设置 app 目录路径
set app_dir=%~dp0app

echo 工作目录: %app_dir%
echo 正在启动所有服务...
echo =====================================

:: 遍历所有服务
for %%s in (%services%) do (
    if exist "%app_dir%\%%s\cmd" (
        cd "%app_dir%\%%s\cmd"
        echo 启动 %%s 服务...
        start /b cmd /c "go run . 2>&1"
        cd %~dp0
        timeout /t 2 /nobreak > nul
    ) else (
        echo 警告: app\%%s\cmd 目录不存在
    )
)

echo =====================================
echo 所有服务已启动完成
echo 按 Ctrl+C 可以终止所有服务
echo =====================================

:: 使用无限循环保持窗口打开并显示输出
:loop
timeout /t 1 /nobreak > nul
goto loop