[Setup]
AppName=dryLang
AppVersion=1.0.0
AppPublisher=zakyislm
AppPublisherURL=https://github.com/zakyislm/drylang
DefaultDirName={pf}\dryLang
DefaultGroupName=dryLang
UninstallDisplayIcon={app}\y.exe
Compression=lzma2
SolidCompression=yes
OutputDir=.\bin
OutputBaseFilename=drylang-installer-x64
ArchitecturesInstallIn64BitMode=x64
ChangesEnvironment=yes

[Files]
; The actual binary is y.exe. We duplicate it as aliases so users can run 'y', 'dry', or 'drylang'.
Source: "..\..\y.exe"; DestDir: "{app}"; DestName: "y.exe"; Flags: ignoreversion
Source: "..\..\y.exe"; DestDir: "{app}"; DestName: "dry.exe"; Flags: ignoreversion
Source: "..\..\y.exe"; DestDir: "{app}"; DestName: "drylang.exe"; Flags: ignoreversion

[Icons]
Name: "{group}\Uninstall dryLang"; Filename: "{uninstallexe}"

[Registry]
; Add the installation directory to the system PATH so 'y' works from anywhere
Root: HKLM; Subkey: "SYSTEM\CurrentControlSet\Control\Session Manager\Environment"; \
    ValueType: expandsz; ValueName: "Path"; ValueData: "{olddata};{app}"; \
    Check: NeedsAddPath(ExpandConstant('{app}'))

[Code]
function NeedsAddPath(Param: string): boolean;
var
  OrigPath: string;
begin
  if not RegQueryStringValue(HKEY_LOCAL_MACHINE,
    'SYSTEM\CurrentControlSet\Control\Session Manager\Environment',
    'Path', OrigPath)
  then begin
    Result := True;
    exit;
  end;
  // look for the path with leading and trailing semicolon
  // Pos() returns 0 if not found
  Result := Pos(';' + Param + ';', ';' + OrigPath + ';') = 0;
end;
