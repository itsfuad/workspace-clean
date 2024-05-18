package main
import (
    "strings"
    "path/filepath"
)
// Analyze summarizes a stream of input without modifying user files.
func Analyze(input string) map[string]interface{} {

targets:=[]string{};total:=0
patterns:=map[string]bool{"node_modules":true,"__pycache__":true,".pytest_cache":true,"dist":true,"target":true}
for _,line:=range strings.Split(input,"\n"){p:=strings.TrimSpace(line);if p==""{continue};total++
if patterns[filepath.Base(p)]{targets=append(targets,p)}}
return map[string]interface{}{"scanned":total,"candidates":targets,"dryRun":true}
}
