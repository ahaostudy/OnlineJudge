import json
import requests

# 设置为⾃⼰的API_KEY，API_KEY由⾃⼰的账号⽣成
api_key = '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'


def chatgpt():
    # 官⽅接⼝（国内⽆法直接访问）
    # url = 'https://api.openai.com/v1/chat/completions'
    # 代理接⼝
    url = 'https://openai.ahao.ink/v1/chat/completions'
    headers = {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer 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',
    }
    data = {
        'model': 'gpt-3.5-turbo',
        'messages': [
            {'role': 'user', 'content': 'hello'}
        ],
    }
    response = requests.post(url=url, headers=headers, data=json.dumps(data))
    print(response.text)


if __name__ == '__main__':
    chatgpt()
