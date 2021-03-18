# dev-nrt-splitter

NRT-Splitter is a command-line utility that does post-processes for csv reports. It walks through a designated folder to deal with each valid csv files, and ignores any other file type.

## basic usage

Download a pre-built binary distribution for your platform (Windows, Linux, Mac), from the releases area of this repository:

https://github.com/nsip/dev-nrt-splitter/releases/latest

Unpack the downloaded zip file, and you should see a folder structure like this:
```
/NRT-Splitter-Linux-v0_0_1
    report_splitter(.exe)
    config.toml
    /data
        └──system_reports.zip
```

In the folder:
-  the report_splitter executable (report_splitter.exe if on windows)
-  a configuration file (config.toml)
-  a subfolder called /data which has a sample package file (system_reports.zip)

report_splitter ignores any command-line parameters or flags except its designated configuration file path.

If running report_splitter without designating configuration file path, default `config.toml` under same directory would applies to executable.

### configuration

```
InFolder (string) -> in which folder splitter processes report csv file. 
WalkSubFolders (bool) -> if true, splitter process every file even if in sub-folder; otherwise, it ignores sub-folder files.

[Trim]
Enabled (bool) -> turn on/off Trim function.
Columns (string array) -> which columns want to be removed from original csv file.
OutFolder (string) -> in which folder trimmed csv files should be output.

[Splitting]
Enabled (bool) -> turn on/off Splitting function.
OutFolder (string) -> in which folder split results should be output.
Schema (string array) -> header sequence for splitting. Each header creates its split category folder. 
```

### play with sample

1. Under `/NRT-Splitter-Linux-v0_0_1`, unpack sample package, `unzip ./data/system_reports.zip -d ./in/`.
2. Modify `config.toml`, set `InFolder` value to `"./in/"`.
3. Make sure `config.toml` is in the same directory of report_splitter.
4. Results should be in `./in/` folder after running `./report_splitter(.exe)`.
