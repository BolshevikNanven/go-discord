@echo off
chcp 936
setlocal EnableDelayedExpansion

:: ���÷���Ŀ¼�б�
set services=auth biz connector gateway im push

:: ���� app Ŀ¼·��
set app_dir=%~dp0app

echo ����Ŀ¼: %app_dir%
echo �����������з���...
echo =====================================

:: �������з���
for %%s in (%services%) do (
    if exist "%app_dir%\%%s\cmd" (
        cd "%app_dir%\%%s\cmd"
        echo ���� %%s ����...
        start /b cmd /c "go run . 2>&1"
        cd %~dp0
        timeout /t 2 /nobreak > nul
    ) else (
        echo ����: app\%%s\cmd Ŀ¼������
    )
)

echo =====================================
echo ���з������������
echo �� Ctrl+C ������ֹ���з���
echo =====================================

:: ʹ������ѭ�����ִ��ڴ򿪲���ʾ���
:loop
timeout /t 1 /nobreak > nul
goto loop