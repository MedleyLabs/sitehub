import os
import requests


def main():

    base_path = 'http://localhost:9000/static/models'
    model_name = 'model_1.txt'

    url = os.path.join(base_path, model_name)

    r = requests.get(url)
    print(r.content)


if __name__ == '__main__':
    main()