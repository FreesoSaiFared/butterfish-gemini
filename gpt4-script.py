import requests
import sys
import json

if len(sys.argv) < 2:
    print("Error: No prompt provided.", file=sys.stderr)
    sys.exit(1)
prompt_content = " ".join(sys.argv[1:])

url = "https://www.chatwithmono.xyz/api/chat"
data = {
    "messages": [
        {"role": "system", "content": "You are a helpful assistant. You are an AI integrated into a shell called butterfish. The user will ask you to perform some task. You should respond with a shell command that will accomplish the task. Only respond with the shell command and nothing else. Do not use a code block. Do not explain the command."},
        {"role": "user", "content": prompt_content}
    ],
    "model": "o3"
}
try:
    response = requests.post(url, json=data)
    response.raise_for_status()
    raw_response = response.text
    response_parts = [part.split(':"')[1].rstrip('"') for part in raw_response.splitlines() if part.startswith('0:"')]
    full_response = ''.join(response_parts)
    print(full_response.strip())
except Exception as e:
    print(f"Error: {e}", file=sys.stderr)
    sys.exit(1)
