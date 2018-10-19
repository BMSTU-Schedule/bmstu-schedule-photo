import imgkit
from PIL import Image
import time


def get_photo_by_url(url, groupID):
    start = time.time()
    photo_path = f'{groupID}.jpg'
    imgkit.from_url(url, photo_path)

    img = Image.open(photo_path)
    area = (20, 150, 1000, 1500)
    cropped_img = img.crop(area)
    cropped_img.save(photo_path)
    end = time.time()
    print("{}s".format(end - start))

get_photo_by_url("https://students.bmstu.ru/schedule/62ec2976-a264-11e5-96b1-005056960017", "ИУ6-54Б")