KURINPLATFORM=kurin64
KURINDIR=$(USERPROFILE)\Desktop\kurin

GOOS=windows
GOARCH=amd64
GOTAGS=$(KURINPLATFORM),kurindebug

.PHONY: buildgame build install rungame run clean
buildgame:
	@set GOOS=$(GOOS)
	@set GOARCH=$(GOARCH)
	@go build -o build/$(KURINPLATFORM)/kurinstation.exe -tags $(GOTAGS) cmd/game/main.go

build: buildgame

install: build
	@if exist "$(KURINDIR)\bin\$(KURINPLATFORM)" rmdir /S /Q "$(KURINDIR)\bin\$(KURINPLATFORM)"
	@xcopy "build\$(KURINPLATFORM)" "$(KURINDIR)\bin\$(KURINPLATFORM)\" /E /C /I >nul
	@if exist"$(KURINDIR)\resources" rmdir /S /Q "$(KURINDIR)\resources"
	@xcopy "resources" "$(KURINDIR)\resources\" /E /C /I >nul

rungame: buildgame
	@if not exist "$(KURINDIR)\bin\dev" mkdir "$(KURINDIR)\bin\dev"
	@copy "build\$(KURINPLATFORM)\kurin.exe" "$(KURINDIR)\bin\dev\kurinstation.exe" >nul
	@cd "build\$(KURINPLATFORM)" && .\kurinstation.exe

run: build
	@if exist "$(KURINDIR)\bin\dev" rmdir /S /Q "$(KURINDIR)\bin\dev"
	@xcopy "build\$(KURINPLATFORM)" "$(KURINDIR)\bin\dev\" /E /C /I >nul
	@cd "build\$(KURINPLATFORM)" && .\kurinstation.exe

clean:
	@if exist "build" rmdir /S /Q build