KITSUNEPLATFORM=kitsune64
KITSUNEDIR=$(USERPROFILE)\Desktop\kitsune

GOOS=windows
GOARCH=amd64
GOTAGS=$(KITSUNEPLATFORM),kitsunedebug

.PHONY: buildbrowser build install runbrowser run clean
buildbrowser:
	@set GOOS=$(GOOS)
	@set GOARCH=$(GOARCH)
	@go build -o build/$(KITSUNEPLATFORM)/kitsune_browser.exe -tags $(GOTAGS) cmd/browser/main.go

build: buildbrowser

install: build
	@if exist "$(KITSUNEDIR)\bin\$(KITSUNEPLATFORM)" rmdir /S /Q "$(KITSUNEDIR)\bin\$(KITSUNEPLATFORM)"
	@xcopy "build\$(KITSUNEPLATFORM)" "$(KITSUNEDIR)\bin\$(KITSUNEPLATFORM)\" /E /C /I >nul
	@if exist"$(KITSUNEDIR)\resources" rmdir /S /Q "$(KITSUNEDIR)\resources"
	@xcopy "resources" "$(KITSUNEDIR)\resources\" /E /C /I >nul

runbrowser: buildbrowser
	@if not exist "$(KITSUNEDIR)\bin\dev" mkdir "$(KITSUNEDIR)\bin\dev"
	@copy "build\$(KITSUNEPLATFORM)\kitsune_browser.exe" "$(KITSUNEDIR)\bin\dev\kitsune_browser.exe" >nul
	@cd "build\$(KITSUNEPLATFORM)" && .\kitsune_browser.exe

run: build
	@if exist "$(KITSUNEDIR)\bin\dev" rmdir /S /Q "$(KITSUNEDIR)\bin\dev"
	@xcopy "build\$(KITSUNEPLATFORM)" "$(KITSUNEDIR)\bin\dev\" /E /C /I >nul
	@cd "build\$(KITSUNEPLATFORM)" && .\kitsune_browser.exe

clean:
	@if exist "build" rmdir /S /Q build