import json

def recover(log_path):
    with open(log_path, 'r', encoding='utf-8') as f:
        lines = f.readlines()
    
    for line in reversed(lines):
        if "file:///c:/Users/user/Desktop/zaky/drylang/core/types.go" in line:
            obj = json.loads(line)
            if "output" in obj.get("content", ""):
                print(obj["content"])
            if "tool_calls" in obj:
                for tc in obj["tool_calls"]:
                    if "output" in tc:
                        print(tc["output"])

recover(r"C:\Users\user\.gemini\antigravity-ide\brain\a7efbd08-afaa-4f77-807e-8617010d316f\.system_generated\logs\transcript_full.jsonl")
