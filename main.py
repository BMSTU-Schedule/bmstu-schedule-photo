import imgkit
from PIL import Image
import time
import requests

option = {
    "format": "png",
    "crop-h": "1370",
    "crop-w": "1500",
    "crop-x": "20",
    "crop-y": "150",
    "encoding": "UTF-8"
    }


def get_photo_by_url(url, groupID):
    start = time.time()
    photo_path = f"{groupID}.png"
    imgkit.from_url(url, photo_path, options=option)

    end = time.time()
    print("{}s".format(end - start))

get_photo_by_url("https://students.bmstu.ru/schedule/62eff510-a264-11e5-bdae-005056960017", "ИУ6-54Б")

